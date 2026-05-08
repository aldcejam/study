package main

import (
	"fmt"
	"log"
	"study_manager/src/infra/database"
	"github.com/ostafen/clover/v2/query"
)

func main() {
	dbPath := "./output/clover_db"
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Erro ao abrir banco: %v", err)
	}
	defer db.Close()

	docs, _ := db.FindAll(query.NewQuery("notes"))
	fmt.Printf("Total de notas no banco: %d\n", len(docs))
	for _, doc := range docs {
		fmt.Printf("- %s: %s (ShortID: %s)\n", doc.Get("filename"), doc.Get("relative_path"), doc.Get("short_id"))
	}
}
