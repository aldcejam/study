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

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq("1ece4e")))
	if err != nil {
		log.Fatalf("Erro ao buscar: %v", err)
	}

	if doc == nil {
		fmt.Println("Nota não encontrada no banco.")
		return
	}

	fmt.Printf("Campos do documento: %+v\n", doc.AsMap())
}
