package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

// escapeHTML escapa caracteres especiais para envio em modo HTML
func escapeHTML(text string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(text)
}

// sendTelegramMessage envia uma mensagem simples para o Telegram
func sendTelegramMessage(token string, chatID int64, text string, parseMode string) error {
	if len(text) > 4000 {
		text = text[:3997] + "..."
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatID)},
		"text":    {text},
	}
	if parseMode != "" {
		payload.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram error: %s", string(body))
	}
	return nil
}

// sendTelegramMarkup envia uma mensagem com botões ou teclados embutidos passando o JSON do reply_markup
func sendTelegramMarkup(token string, chatID int64, text string, parseMode string, replyMarkup string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	
	payload := url.Values{
		"chat_id":      {fmt.Sprintf("%d", chatID)},
		"text":         {text},
		"reply_markup": {replyMarkup},
	}
	if parseMode != "" {
		payload.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram menu error: %s", string(body))
	}
	return nil
}

// answerCallbackQuery envia uma resposta vazia para o telegram saber que o click foi processado
func answerCallbackQuery(token string, callbackQueryID string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", token)
	payload := url.Values{
		"callback_query_id": {callbackQueryID},
	}
	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}


// syncDatabase executa o pipeline de sincronização (scanner + summary) com reset total
func syncDatabase() error {
	log.Println("🔄 Resetando e Sincronizando banco de dados...")

	// Limpeza total para evitar otimizações que pulam atualizações
	os.RemoveAll("./output/clover_db")
	os.Remove("./output/last_scan.json")

	scannerCmd := exec.Command("./bin/scanner")
	var scannerOut bytes.Buffer
	scannerCmd.Stdout = &scannerOut
	if err := scannerCmd.Run(); err != nil {
		return fmt.Errorf("falha no scanner: %v", err)
	}

	summaryCmd := exec.Command("./bin/summaryNotification")
	summaryCmd.Stdin = &scannerOut
	if err := summaryCmd.Run(); err != nil {
		return fmt.Errorf("falha no summaryNotification: %v", err)
	}

	log.Println("✅ Sincronização completa (banco recriado).")
	return nil
}

// createForumTopic cria um novo tópico em um supergrupo e retorna o message_thread_id
func createForumTopic(token string, chatID int64, name string) (int, error) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/createForumTopic", token)
	payload := url.Values{
		"chat_id": {fmt.Sprintf("%d", chatID)},
		"name":    {name},
	}

	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("create topic error: %s", string(body))
	}

	var result struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageThreadID int `json:"message_thread_id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result.MessageThreadID, nil
}

// sendTelegramMessageToTopic envia uma mensagem para um tópico específico
func sendTelegramMessageToTopic(token string, chatID int64, threadID int, text string, parseMode string) (int, error) {
	if len(text) > 4000 {
		text = text[:3997] + "..."
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	payload := url.Values{
		"chat_id":           {fmt.Sprintf("%d", chatID)},
		"message_thread_id": {fmt.Sprintf("%d", threadID)},
		"text":              {text},
	}
	if parseMode != "" {
		payload.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("telegram topic error: %s", string(body))
	}
	
	var result struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageID int `json:"message_id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Result.MessageID, nil
}

// editTelegramMessageText edita o texto de uma mensagem já enviada
func editTelegramMessageText(token string, chatID int64, messageID int, text string, parseMode string) error {
	if len(text) > 4000 {
		text = text[:3997] + "..."
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/editMessageText", token)
	payload := url.Values{
		"chat_id":    {fmt.Sprintf("%d", chatID)},
		"message_id": {fmt.Sprintf("%d", messageID)},
		"text":       {text},
	}
	if parseMode != "" {
		payload.Set("parse_mode", parseMode)
	}

	resp, err := http.PostForm(apiURL, payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram edit error: %s", string(body))
	}
	return nil
}

