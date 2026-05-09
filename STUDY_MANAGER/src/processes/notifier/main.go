package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"study_manager/src/utils/dotenv"
)

func loadEnv() {
	selfDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	dotenv.Load(
		filepath.Join(selfDir, "../../../.env"), // estudos/.env
		filepath.Join(selfDir, "../../.env"),    // STUDY_MANAGER/.env
		".env",
	)
}

func sendTelegram(token, chatID, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id":    {chatID},
		"text":       {message},
		"parse_mode": {"HTML"},
	})
	if err != nil {
		return fmt.Errorf("erro de rede: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API Telegram retornou %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func main() {
	loadEnv()

	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "notifier: erro ao ler stdin:", err)
		os.Exit(1)
	}

	mensagem := strings.TrimSpace(string(raw))
	if mensagem == "" {
		fmt.Fprintln(os.Stderr, "notifier: nenhum dado recebido via stdin")
		os.Exit(1)
	}

	switch mensagem {
	case "SEM_REVISOES":
		fmt.Println("✅ Nenhuma revisão pendente. Notificação não enviada.")
		return
	case "SEM_ALTERACOES":
		fmt.Println("ℹ️  Cache detectado: Nenhuma alteração relevante desde o último alerta.")
		return
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if token == "" || chatID == "" {
		fmt.Fprintln(os.Stderr, "notifier: ⚠️  TELEGRAM_TOKEN ou TELEGRAM_CHAT_ID não definidos")
		os.Exit(1)
	}

	fmt.Println("🚀 Enviando alerta para o Telegram...")
	if err := sendTelegram(token, chatID, mensagem); err != nil {
		fmt.Fprintln(os.Stderr, "notifier: ❌ Falha no envio:", err)
		os.Exit(1)
	}
	fmt.Println("✨ Sucesso!")
}
