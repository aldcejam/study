package topic

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"study_manager/src/infra/database"
)

//go:embed system_prompt.txt
var systemPrompt string

// HandleStartTopic carrega a nota, cria um tópico e salva a sessão de estudo.
func HandleStartTopic(connStr string, shortID string, chatID int64, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) (int, error)) {
	pool, err := database.InitDB(connStr)
	if err != nil {
		log.Printf("DB Error: %v", err)
		return
	}
	defer pool.Close()

	var tema string
	err = pool.QueryRow(context.Background(), "SELECT tema FROM notes WHERE short_id = $1", shortID).Scan(&tema)
	if err != nil {
		log.Printf("Nota não encontrada para o shortID: %s", shortID)
		sendTopicFunc(token, chatID, 0, "❌ Nota não encontrada.", "")
		return
	}

	// Limita o nome do tópico a 128 caracteres (limite do Telegram)
	topicName := "Estudo: " + tema
	if len(topicName) > 128 {
		topicName = topicName[:125] + "..."
	}

	threadID, err := createTopicFunc(token, chatID, topicName)
	if err != nil {
		log.Printf("Erro ao criar tópico: %v", err)
		sendTopicFunc(token, chatID, 0, "❌ Erro ao criar tópico. Verifique se o recurso de 'Tópicos' (Fórum) está habilitado no grupo e se o bot tem permissão 'Gerenciar Tópicos'.", "")
		return
	}

	// Salva a sessão de estudo no banco de dados (upsert)
	_, err = pool.Exec(context.Background(), `
		INSERT INTO study_sessions (chat_id, thread_id, short_id, history) 
		VALUES ($1, $2, $3, '[]'::jsonb)
		ON CONFLICT (chat_id, thread_id) DO NOTHING
	`, chatID, threadID, shortID)

	if err != nil {
		log.Printf("Erro ao salvar sessão: %v", err)
	}

	// Envia mensagem de boas vindas dentro do tópico
	welcomeMsg := fmt.Sprintf("🤖 **Tópico de Estudo Iniciado!**\n\nEstou pronto para te ajudar a estudar a nota **%s**.\n\nPode mandar suas dúvidas ou me pedir para testar seus conhecimentos sobre o assunto!", tema)
	sendTopicFunc(token, chatID, threadID, welcomeMsg, "Markdown")
}

// HandleTopicMessage processa mensagens enviadas em tópicos.
func HandleTopicMessage(connStr string, chatID int64, threadID int, text string, token string, sendTopicFunc func(string, int64, int, string, string) (int, error), editTopicMessageFunc func(string, int64, int, string, string) error, sendChatActionFunc func(string, int64, int, string) error) {
	log.Printf("HandleTopicMessage called: chatID=%d, threadID=%d", chatID, threadID)
	pool, err := database.InitDB(connStr)
	if err != nil {
		log.Printf("DB error: %v", err)
		return
	}
	defer pool.Close()

	var shortID string
	var historyBytes []byte
	err = pool.QueryRow(context.Background(), "SELECT short_id, history FROM study_sessions WHERE chat_id = $1 AND thread_id = $2", chatID, threadID).Scan(&shortID, &historyBytes)
	if err != nil {
		log.Printf("Active session not found: %v", err)
		return
	}

	var relativePath, filename string
	err = pool.QueryRow(context.Background(), "SELECT relative_path, filename FROM notes WHERE short_id = $1", shortID).Scan(&relativePath, &filename)
	if err != nil {
		log.Printf("Note info not found: %v", err)
		return
	}

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, relativePath, filename)
	contentBytes, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Failed to read note file: %v", err)
		return
	}
	contentStr := string(contentBytes)

	var history []string
	if len(historyBytes) > 0 {
		json.Unmarshal(historyBytes, &history)
	}

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
	
	newHistoryBytes, _ := json.Marshal(history)
	_, err = pool.Exec(context.Background(), "UPDATE study_sessions SET history = $1 WHERE chat_id = $2 AND thread_id = $3", newHistoryBytes, chatID, threadID)
	if err != nil {
		log.Printf("Failed to update history: %v", err)
	}

	if msgID != 0 {
		editTopicMessageFunc(token, chatID, msgID, responseStr, "Markdown")
	} else {
		sendTopicFunc(token, chatID, threadID, responseStr, "Markdown")
	}
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

