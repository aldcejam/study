package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"study_manager/src/api/main/commands"
	"study_manager/src/api/main/commands/topic"
)

// --- Models ---

type CallbackQuery struct {
	ID      string   `json:"id"`
	From    Chat     `json:"from"`
	Message *Message `json:"message"`
	Data    string   `json:"data"`
}

type Update struct {
	UpdateID      int            `json:"update_id"`
	Message       *Message       `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type Message struct {
	MessageID       int    `json:"message_id"`
	Chat            Chat   `json:"chat"`
	Text            string `json:"text"`
	MessageThreadID int    `json:"message_thread_id,omitempty"`
	IsTopicMessage  bool   `json:"is_topic_message,omitempty"`
}

type Chat struct {
	ID int64 `json:"id"`
}

// handleWebhook atua como um Proxy/Roteador, despachando comandos para o pacote commands
func handleWebhook(dbPath string, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var update Update
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			log.Printf("Decode error: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Função anônima para injetar as dependências no helper de envio
		sender := func(id int64, msg string, mode string) error {
			return sendTelegramMessage(token, id, msg, mode)
		}
		


		// --- TRATAMENTO DE CALLBACK QUERIES (Cliques em botões) ---
		if update.CallbackQuery != nil {
			cq := update.CallbackQuery
			var chatID int64
			if cq.Message != nil {
				chatID = int64(cq.Message.Chat.ID)
			} else {
				chatID = int64(cq.From.ID)
			}
			data := cq.Data
			log.Printf("Callback received: %s from %d", data, chatID)

			answerCallbackQuery(token, cq.ID)

			// O formato será: acao:shortID  ex: "view_note:1ece4e"
			parts := strings.Split(data, ":")
			if len(parts) >= 2 {
				action := parts[0]
				shortID := parts[1]

				switch action {
				case "view_note":
					commands.HandleViewContent(dbPath, shortID, sender, chatID)
				case "start_topic":
					go topic.HandleStartTopic(dbPath, shortID, chatID, token, createForumTopic, sendTelegramMessageToTopic)
				case "view_hw":
					menuSenderFunc := func(id int64, msg string, mode string, markupStr string) error {
						return sendTelegramMarkup(token, id, msg, mode, markupStr)
					}
					commands.HandleViewList(dbPath, shortID, "homework", menuSenderFunc, chatID)
				case "view_ref":
					menuSenderFunc := func(id int64, msg string, mode string, markupStr string) error {
						return sendTelegramMarkup(token, id, msg, mode, markupStr)
					}
					commands.HandleViewList(dbPath, shortID, "references", menuSenderFunc, chatID)
				case "open_hw", "open_ref":
					// Callback format: type:shortID:index
					parts := strings.Split(data, ":")
					if len(parts) == 3 {
						index := 0
						fmt.Sscanf(parts[2], "%d", &index)
						listType := "homework"
						if parts[0] == "open_ref" {
							listType = "references"
						}
						commands.HandleOpenFileByIndex(dbPath, parts[1], index, listType, sender, chatID)
					}
				case "url":
					// Callback format: url:link
					link := strings.TrimPrefix(data, "url:")
					sender(chatID, "🌐 Link para abrir: "+link, "")
				}

				answerCallbackQuery(token, update.CallbackQuery.ID)
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		// --- TRATAMENTO DE MENSAGENS NORMAIS ---
		if update.Message == nil || update.Message.Text == "" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if update.Message.MessageThreadID != 0 && update.Message.IsTopicMessage {
			go topic.HandleTopicMessage(dbPath, update.Message.Chat.ID, update.Message.MessageThreadID, update.Message.Text, token, sendTelegramMessageToTopic, editTelegramMessageText)
			w.WriteHeader(http.StatusOK)
			return
		}

		text := update.Message.Text
		rawCmd := strings.Split(text, " ")[0]
		chatID := update.Message.Chat.ID

		log.Printf("Command received: %s from %d", rawCmd, chatID)

		// Remove o sufixo @MeusEstudosBot (ou qualquer outro bot) em grupos
		cmd := rawCmd
		if idx := strings.Index(cmd, "@"); idx != -1 {
			cmd = cmd[:idx]
		}

		// Roteamento dos comandos
		switch {
		case cmd == "/meus_estudos":
			commands.HandleMeusEstudos(dbPath, syncDatabase, sender, escapeHTML, chatID)

		case strings.HasPrefix(cmd, "/ver_"):
			shortID := strings.TrimPrefix(cmd, "/ver_")
			// Mudou: agora enviar o menu em vez do conteúdo direto
			menuSenderFunc := func(id int64, msg string, mode string, markupStr string) error {
				return sendTelegramMarkup(token, id, msg, mode, markupStr)
			}
			commands.HandleVerMenu(dbPath, shortID, sender, menuSenderFunc, chatID)

		case strings.HasPrefix(cmd, "/abrir_"):
			// ler_arquivo decodificado em hex
			hexPath := strings.TrimPrefix(cmd, "/abrir_")
			commands.HandleAbrirArquivo(hexPath, sender, chatID)

		default:
			log.Printf("Unknown command: %s", cmd)
		}

		w.WriteHeader(http.StatusOK)
	}
}