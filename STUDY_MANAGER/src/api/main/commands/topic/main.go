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
	"github.com/jackc/pgx/v5"
)

//go:embed system_prompt.txt
var systemPrompt string

// HandleStartTopic carrega a nota, cria um tópico e salva a sessão de estudo.
func HandleStartTopic(repo *database.Repository, shortID string, chatID int64, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) (int, error)) {
	note, err := repo.GetNoteByShortID(context.Background(), shortID)
	if err != nil {
		log.Printf("Nota não encontrada para o shortID: %s", shortID)
		sendTopicFunc(token, chatID, 0, "❌ Nota não encontrada.", "")
		return
	}

	// Verifica se já existe um tópico para esta nota
	existingThreadID, err := repo.GetThreadIDForNote(context.Background(), chatID, shortID)
	if err == nil && existingThreadID != 0 {
		sendTopicFunc(token, chatID, 0, fmt.Sprintf("⚠️ Já existe um tópico de estudo ativo para a nota **%s**. Acesse-o na lista de tópicos do grupo.", note.Tema), "Markdown")
		
		// Opcional: Reenviar uma mensagem no tópico existente para dar um "bump" nele
		bumpMsg := "🔔 **Retornando aos Estudos**\n\nEste é o tópico ativo para esta nota. O que vamos estudar agora?"
		sendTopicFunc(token, chatID, existingThreadID, bumpMsg, "Markdown")
		return
	}
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Erro ao verificar tópico existente: %v", err)
	}

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

	// Salva a sessão de estudo no banco de dados (upsert)
	err = repo.UpsertStudySession(context.Background(), chatID, threadID, shortID)

	if err != nil {
		log.Printf("Erro ao salvar sessão: %v", err)
	}

	// Envia mensagem de boas vindas dentro do tópico
	welcomeMsg := fmt.Sprintf("🤖 **Tópico de Estudo Iniciado!**\n\nEstou pronto para te ajudar a estudar a nota **%s**.\n\nPode mandar suas dúvidas ou me pedir para testar seus conhecimentos sobre o assunto!", note.Tema)
	sendTopicFunc(token, chatID, threadID, welcomeMsg, "Markdown")
}

// HandleTopicMessage processa mensagens enviadas em tópicos.
func HandleTopicMessage(repo *database.Repository, chatID int64, threadID int, text string, token string, sendTopicFunc func(string, int64, int, string, string) (int, error), editTopicMessageFunc func(string, int64, int, string, string) error, sendChatActionFunc func(string, int64, int, string) error) {
	log.Printf("HandleTopicMessage called: chatID=%d, threadID=%d", chatID, threadID)

	shortID, historyBytes, err := repo.GetStudySession(context.Background(), chatID, threadID)
	if err != nil {
		log.Printf("Active session not found: %v", err)
		return
	}

	note, err := repo.GetNoteByShortID(context.Background(), shortID)
	if err != nil {
		log.Printf("Note info not found: %v", err)
		return
	}

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, note.RelativePath, note.Filename)
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
	err = repo.UpdateStudySessionHistory(context.Background(), chatID, threadID, newHistoryBytes)
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

