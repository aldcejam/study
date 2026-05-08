package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"study_manager/src/utils/dotenv"
)

func main() {
	dotenv.Load(".env", "../.env", "../../.env")
	
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN not found")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./output/clover_db"
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

	// Telegram Webhook Route - Passamos o dbPath para que o handler abra o banco sob demanda
	http.HandleFunc("/webhook", handleWebhook(dbPath, token))

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
