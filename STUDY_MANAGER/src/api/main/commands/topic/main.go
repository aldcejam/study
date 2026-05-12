package topic

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"study_manager/src/infra/database"

	c "github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

//go:embed system_prompt.txt
var systemPrompt string

// HandleStartTopic carrega a nota, cria um tópico e salva a sessão de estudo.
func HandleStartTopic(dbPath string, shortID string, chatID int64, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) (int, error)) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Printf("DB Error: %v", err)
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		log.Printf("Nota não encontrada para o shortID: %s", shortID)
		sendTopicFunc(token, chatID, 0, "❌ Nota não encontrada.", "")
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	// Limita o nome do tópico a 128 caracteres (limite do Telegram)
	topicName := "Estudo: " + note.Tema
	if len(topicName) > 128 {
		topicName = topicName[:125] + "..."
	}

	threadID, err := createTopicFunc(token, chatID, topicName)
	if err != nil {
		log.Printf("Erro ao criar tópico: %v", err)
		sendTopicFunc(token, chatID, 0, "❌ Erro ao criar tópico. Verifique se o recurso de 'Tópicos' (Fórum) está habilitado no grupo e se o bot tem permissão 'Gerenciar Tópicos'.", "")
		return
	}

	// Salva a sessão de estudo no banco de dados
	sessionDoc := document.NewDocument()
	sessionDoc.Set("thread_id", threadID)
	sessionDoc.Set("chat_id", chatID)
	sessionDoc.Set("short_id", shortID)
	sessionDoc.Set("history", []interface{}{}) // Começa com histórico vazio

	_, err = db.InsertOne("study_sessions", sessionDoc)
	if err != nil {
		log.Printf("Erro ao salvar sessão: %v", err)
	}

	// Envia mensagem de boas vindas dentro do tópico
	welcomeMsg := fmt.Sprintf("🤖 **Tópico de Estudo Iniciado!**\n\nEstou pronto para te ajudar a estudar a nota **%s**.\n\nPode mandar suas dúvidas ou me pedir para testar seus conhecimentos sobre o assunto!", note.Tema)
	sendTopicFunc(token, chatID, threadID, welcomeMsg, "Markdown")
}

// HandleTopicMessage processa mensagens enviadas em tópicos.
func HandleTopicMessage(dbPath string, chatID int64, threadID int, text string, token string, sendTopicFunc func(string, int64, int, string, string) (int, error), editTopicMessageFunc func(string, int64, int, string, string) error, sendChatActionFunc func(string, int64, int, string) error) {
	log.Printf("HandleTopicMessage called: chatID=%d, threadID=%d", chatID, threadID)
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Printf("DB error: %v", err)
		return
	}
	defer db.Close()

	doc, shortID, err := getActiveSession(db, chatID, threadID)
	if err != nil {
		log.Printf("Active session not found: %v", err)
		return
	}

	contentStr, err := getNoteContent(db, shortID)
	if err != nil {
		log.Printf("Note content not found: %v", err)
		return
	}

	history := getHistory(doc)
	prompt := buildPrompt(contentStr, history, text)
	history = append(history, "Usuário: "+text)

	// Envia mensagem de loading e pega o ID
	msgID, err := sendTopicFunc(token, chatID, threadID, "🤔 *Pensando...* ⏳", "Markdown")
	if err != nil {
		log.Printf("Error sending loading message: %v", err)
	}

	// Inicia a animação "digitando..." em background
	stopTyping := make(chan struct{})
	go func() {
		ticker := time.NewTicker(4 * time.Second)
		defer ticker.Stop()
		sendChatActionFunc(token, chatID, threadID, "typing")
		for {
			select {
			case <-ticker.C:
				sendChatActionFunc(token, chatID, threadID, "typing")
			case <-stopTyping:
				return
			}
		}
	}()

	responseStr, err := invokeGeminiCLI(prompt)
	
	close(stopTyping) // Para a animação de digitando

	if err != nil {
		log.Printf("Gemini CLI error: %v", err)
		responseStr = "Desculpe, ocorreu um erro ao processar sua resposta via Gemini CLI."
	}

	history = append(history, "Tutor: "+responseStr)
	saveHistory(db, doc, history)

	if msgID != 0 {
		editTopicMessageFunc(token, chatID, msgID, responseStr, "Markdown")
	} else {
		sendTopicFunc(token, chatID, threadID, responseStr, "Markdown")
	}
}

// Helpers

func getActiveSession(db *c.DB, chatID int64, threadID int) (*document.Document, string, error) {
	docs, err := db.FindAll(query.NewQuery("study_sessions"))
	if err != nil {
		return nil, "", err
	}
	for _, doc := range docs {
		var cID int64
		if val, ok := doc.Get("chat_id").(int64); ok {
			cID = val
		} else if val, ok := doc.Get("chat_id").(float64); ok {
			cID = int64(val)
		}

		var tID int
		if val, ok := doc.Get("thread_id").(int); ok {
			tID = val
		} else if val, ok := doc.Get("thread_id").(float64); ok {
			tID = int(val)
		}

		if cID == chatID && tID == threadID {
			return doc, doc.Get("short_id").(string), nil
		}
	}
	return nil, "", fmt.Errorf("session not found")
}

func getNoteContent(db *c.DB, shortID string) (string, error) {
	noteDoc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || noteDoc == nil {
		return "", fmt.Errorf("note not found")
	}

	var note database.NoteDoc
	noteDoc.Unmarshal(&note)

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, note.RelativePath, note.Filename)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getHistory(doc *document.Document) []string {
	var history []string
	historyRaw := doc.Get("history")
	if historyRaw != nil {
		if arr, ok := historyRaw.([]interface{}); ok {
			for _, h := range arr {
				if s, ok := h.(string); ok {
					history = append(history, s)
				}
			}
		}
	}
	return history
}

func buildPrompt(contentStr string, history []string, text string) string {
	prompt := fmt.Sprintf("%s\n\n---\n%s\n---\n\n", systemPrompt, contentStr)
	if len(history) > 0 {
		prompt += "Histórico da conversa:\n"
		for _, h := range history {
			prompt += h + "\n"
		}
	}
	prompt += "\nUsuário: " + text + "\nTutor:"
	return prompt
}

func invokeGeminiCLI(prompt string) (string, error) {
	cmd := exec.Command("gemini", "ask", prompt)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Erro no gemini CLI: %v\nOutput: %s", err, string(out))
		return "", err
	}
	return string(out), nil
}

func saveHistory(db *c.DB, doc *document.Document, history []string) {
	db.UpdateById("study_sessions", doc.ObjectId(), func(d *document.Document) *document.Document {
		d.Set("history", history)
		return d
	})
}
