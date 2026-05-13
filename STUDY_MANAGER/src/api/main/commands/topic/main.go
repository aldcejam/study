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
	"strconv"
	"strings"
	"time"

	"study_manager/src/infra/database"
	"study_manager/src/infra/formatter"
	"github.com/jackc/pgx/v5"
	"sync"
)

var creationLocks sync.Map

//go:embed system_prompt.txt
var systemPrompt string

// HandleStartTopic carrega a nota, cria um tópico e salva a sessão de estudo.
func HandleStartTopic(repo *database.Repository, shortID string, chatID int64, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) (int, error)) {
	lockKey := fmt.Sprintf("%d:%s", chatID, shortID)
	if _, loaded := creationLocks.LoadOrStore(lockKey, true); loaded {
		return
	}
	defer creationLocks.Delete(lockKey)

	note, err := repo.GetNoteByShortID(context.Background(), shortID)
	if err != nil {
		log.Printf("Nota não encontrada para o shortID: %s", shortID)
		sendTopicFunc(token, chatID, 0, "❌ Nota não encontrada.", "")
		return
	}

	// Verifica se já existe um tópico para esta nota
	existingThreadID, err := repo.GetThreadIDForNote(context.Background(), chatID, shortID)
	if err != nil && err != pgx.ErrNoRows {
		log.Printf("Erro ao verificar tópico existente para chat %d, nota %s: %v", chatID, shortID, err)
		return
	}

	if err == nil && existingThreadID != 0 {
		// Tenta enviar o "bump" para ver se o tópico ainda existe
		bumpMsg := "🔔 **Retornando aos Estudos**\n\nEste é o tópico ativo para esta nota. O que vamos estudar agora?"
		conv := formatter.NewMarkdownConverter()
		formattedMsg := conv.Convert(bumpMsg)
		_, sendErr := sendTopicFunc(token, chatID, existingThreadID, formattedMsg, "HTML")
		
		if sendErr == nil {
			// Sucesso: Tópico existe
			msg := fmt.Sprintf("⚠️ Já existe um tópico de estudo ativo para a nota <b>%s</b>. Acesse-o na lista de tópicos do grupo.", note.Tema)
			sendTopicFunc(token, chatID, 0, msg, "HTML")
			return
		}
		
		// Falha: O tópico provavelmente foi deletado no Telegram
		log.Printf("Erro ao enviar bump para o tópico %d (provavelmente deletado). Recriando... (Erro: %v)", existingThreadID, sendErr)
		repo.DeleteStudySession(context.Background(), chatID, existingThreadID)
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
	welcomeMsg := fmt.Sprintf("🤖 <b>Tópico de Estudo Iniciado!</b>\n\nEstou pronto para te ajudar a estudar a nota <b>%s</b>.\n\nPode mandar suas dúvidas ou me pedir para testar seus conhecimentos sobre o assunto!", note.Tema)
	sendTopicFunc(token, chatID, threadID, welcomeMsg, "HTML")
}

func HandleTopicMessage(repo *database.Repository, chatID int64, threadID int, text string, token string, createTopicFunc func(string, int64, string) (int, error), sendTopicFunc func(string, int64, int, string, string) (int, error), editTopicMessageFunc func(string, int64, int, string, string) error, sendChatActionFunc func(string, int64, int, string) error) {
	log.Printf("HandleTopicMessage called: chatID=%d, threadID=%d, text=%q", chatID, threadID, text)

	shortID, historyBytes, err := repo.GetStudySession(context.Background(), chatID, threadID)
	if err != nil {
		log.Printf("Active session not found for thread %d: %v", threadID, err)
		return
	}

	note, err := repo.GetNoteByShortID(context.Background(), shortID)
	if err != nil {
		log.Printf("Note info not found: %v", err)
		return
	}

	// --- Tratamento de Comandos Especiais ---
	cleanText := text
	if strings.HasPrefix(cleanText, "/") {
		parts := strings.Split(cleanText, " ")
		cmd := parts[0]
		if idx := strings.Index(cmd, "@"); idx != -1 {
			cmd = cmd[:idx]
			parts[0] = cmd
			cleanText = strings.Join(parts, " ")
		}
	}
	trimmedText := strings.TrimSpace(cleanText)
	
	// Comando /fork
	if trimmedText == "/fork" {
		log.Printf("Processing /fork command for thread %d", threadID)
		newThreadID, err := createTopicFunc(token, chatID, "[Fork] "+note.Tema)
		if err != nil {
			log.Printf("Error creating fork topic: %v", err)
			sendTopicFunc(token, chatID, threadID, "❌ Falha ao criar novo tópico (Fork).", "HTML")
			return
		}
		
		err = repo.CreateForkSession(context.Background(), chatID, threadID, newThreadID, historyBytes)
		if err != nil {
			log.Printf("Error saving fork session: %v", err)
			return
		}

		sendTopicFunc(token, chatID, threadID, "✅ <b>Tópico clonado com sucesso!</b> Acesse a nova aba do fórum.", "HTML")
		
		welcomeMsg := fmt.Sprintf("🤖 <b>Tópico Clonado!</b>\n\nTodo o contexto da nossa conversa sobre <b>%s</b> foi preservado aqui.\n\nComo deseja continuar?", note.Tema)
		sendTopicFunc(token, chatID, newThreadID, welcomeMsg, "HTML")
		return
	}

	// Comando /apagar N
	if strings.HasPrefix(trimmedText, "/apagar") {
		log.Printf("Processing /apagar command for thread %d: %s", threadID, trimmedText)
		parts := strings.Split(trimmedText, " ")
		if len(parts) >= 2 {
			n, err := strconv.Atoi(parts[1])
			if err == nil && n > 0 {
				var history []string
				if len(historyBytes) > 0 {
					json.Unmarshal(historyBytes, &history)
				}
				
				itemsToRemove := n * 2
				if itemsToRemove > len(history) {
					itemsToRemove = len(history)
				}
				
				if itemsToRemove > 0 {
					history = history[:len(history)-itemsToRemove]
					newHistoryBytes, _ := json.Marshal(history)
					repo.UpdateStudySessionHistory(context.Background(), chatID, threadID, newHistoryBytes)
					sendTopicFunc(token, chatID, threadID, fmt.Sprintf("🧹 <b>As últimas %d interações foram apagadas do contexto da IA.</b>", itemsToRemove/2), "HTML")
				} else {
					sendTopicFunc(token, chatID, threadID, "⚠️ O histórico já está vazio.", "HTML")
				}
				return
			}
		}
		sendTopicFunc(token, chatID, threadID, "⚠️ Formato inválido. Use <code>/apagar N</code> onde N é o número de interações.", "HTML")
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
	msgID, err := sendTopicFunc(token, chatID, threadID, "🤔 *Pensando...*", "Markdown")
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

	conv := formatter.NewMarkdownConverter()
	formattedResponse := conv.Convert(responseStr)

	if msgID != 0 {
		editTopicMessageFunc(token, chatID, msgID, formattedResponse, "HTML")
	} else {
		sendTopicFunc(token, chatID, threadID, formattedResponse, "HTML")
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

