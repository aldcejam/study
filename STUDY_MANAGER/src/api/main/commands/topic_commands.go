package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"study_manager/src/infra/database"

	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

// HandleStartTopic carrega a nota, cria um tópico e salva a sessão de estudo.
func HandleStartTopic(dbPath string, shortID string, chatID int64, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) error) {
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
func HandleTopicMessage(dbPath string, chatID int64, threadID int, text string, token string, sendTopicFunc func(string, int64, int, string, string) error) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	// Procura a sessão ativa
	doc, err := db.FindFirst(query.NewQuery("study_sessions").Where(query.Field("thread_id").Eq(threadID).And(query.Field("chat_id").Eq(chatID))))
	if err != nil || doc == nil {
		// Não é um tópico de estudo gerenciado pelo bot
		return
	}

	shortID := doc.Get("short_id").(string)
	
	// Pega a nota
	noteDoc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || noteDoc == nil {
		return
	}

	var note database.NoteDoc
	noteDoc.Unmarshal(&note)

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, note.RelativePath, note.Filename)

	content, err := os.ReadFile(fullPath)
	contentStr := ""
	if err == nil {
		contentStr = string(content)
	}

	// Histórico
	historyRaw := doc.Get("history")
	var history []string
	if historyRaw != nil {
		if arr, ok := historyRaw.([]interface{}); ok {
			for _, h := range arr {
				if s, ok := h.(string); ok {
					history = append(history, s)
				}
			}
		}
	}

	// Monta o super prompt
	prompt := fmt.Sprintf("Você é um tutor de estudos. Baseado na seguinte nota:\n\n---\n%s\n---\n\n", contentStr)
	if len(history) > 0 {
		prompt += "Histórico da conversa:\n"
		for _, h := range history {
			prompt += h + "\n"
		}
	}
	prompt += "\nUsuário: " + text + "\nTutor:"

	// Atualiza histórico com a msg do user
	history = append(history, "Usuário: "+text)

	// Chama o gemini CLI
	// Exemplo de comando: gemini ask "prompt"
	cmd := exec.Command("gemini", "ask", prompt)
	out, err := cmd.Output()
	responseStr := ""
	if err != nil {
		log.Printf("Erro no gemini CLI: %v\nOutput: %s", err, string(out))
		responseStr = "Desculpe, ocorreu um erro ao processar sua resposta via Gemini CLI."
	} else {
		responseStr = string(out)
	}

	// Salva a resposta no histórico
	history = append(history, "Tutor: "+responseStr)
	
	// Atualiza o banco
	db.UpdateById("study_sessions", doc.ObjectId(), func(d *document.Document) *document.Document {
		d.Set("history", history)
		return d
	})

	// Envia a resposta de volta ao tópico
	sendTopicFunc(token, chatID, threadID, responseStr, "Markdown")
}

