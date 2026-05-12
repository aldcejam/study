package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"os/exec"
	"time"

	"study_manager/src/utils/dotenv"
)



func main() {
	dotenv.Load(".env", "../.env", "../../../.env")


	// Ticker para rodar o pipeline de resumo a cada 5s (Teste de Otimização)
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			log.Println("🔄 [Background] Rodando pipeline de notificação...")
			cmd := exec.Command("./bin/pipeline")
			// Define o CWD para a raiz do projeto
			cmd.Dir = "." 
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("❌ [Background] Erro no pipeline: %v\n%s", err, string(output))
			} else {
				log.Println("✅ [Background] Pipeline finalizado.")
			}
		}
	}()
	
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN not found")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://study_user:study_password@localhost:5432/study_db?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8443"
	}

	// Health Check Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "🚀 Study Manager API is up and running!")
	})

	// Telegram Webhook Route - Passamos o dbUrl para que o handler abra o pool sob demanda
	http.HandleFunc("/webhook", handleWebhook(dbUrl, token))

	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile != "" && keyFile != "" {
		log.Printf("Server HTTPS active on port %s", port)
		log.Fatal(http.ListenAndServeTLS(":"+port, certFile, keyFile, nil))
	} else {
		log.Printf("Server HTTP active on port %s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
}
