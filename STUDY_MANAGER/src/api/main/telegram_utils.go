package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"log"
	"net/http"
	"net/url"
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
