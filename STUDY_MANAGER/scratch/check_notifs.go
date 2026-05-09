package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"study_manager/src/infra/database"
	"github.com/ostafen/clover/v2/query"
)

func main() {
	cwd, _ := os.Getwd()
	dbPath := filepath.Join(cwd, "output", "clover_db")
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	docs, err := db.FindAll(query.NewQuery("notifications"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("--- Notificações Salvas (%d) ---\n", len(docs))
	for _, doc := range docs {
		var n database.NotificationDoc
		doc.Unmarshal(&n)
		ln := "nunca"
		if n.LastNotifiedAt != nil {
			ln = *n.LastNotifiedAt
		}
		fmt.Printf("Path: %s | Last: %s | Completed: %v\n", n.RelativePath, ln, n.Completed)
	}
}
