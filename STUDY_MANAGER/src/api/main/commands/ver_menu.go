package commands

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"study_manager/src/infra/database"

	"github.com/ostafen/clover/v2/query"
)

// Estruturas locais para o teclado inline (para evitar importação do main)
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
	Url          string `json:"url,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// HandleVerMenu envia as opções de menu (botões inline) quando se clica em /ver_ID
func HandleVerMenu(dbPath string, shortID string, sendFunc func(int64, string, string) error, menuSenderFunc func(int64, string, string, string) error, chatID int64) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		sendFunc(chatID, "❌ Nota não encontrada.", "")
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	text := fmt.Sprintf("📖 <b>%s</b>\nEscolha o que deseja visualizar:", note.Tema)

	keyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: "📖 Ver Nota", CallbackData: "view_note:" + shortID},
			},
			{
				{Text: "📝 Ver Atividades", CallbackData: "view_hw:" + shortID},
				{Text: "🔗 Ver Referências", CallbackData: "view_ref:" + shortID},
			},
		},
	}

	kbBytes, err := json.Marshal(keyboard)
	if err != nil {
		log.Printf("Marshal keyboard error: %v", err)
		return
	}

	if err := menuSenderFunc(chatID, text, "HTML", string(kbBytes)); err != nil {
		log.Printf("Send menu error: %v", err)
	}
}

// HandleViewContent envia o conteúdo bruto da nota (antigo comportamento do ver_nota)
func HandleViewContent(dbPath string, shortID string, sendFunc func(int64, string, string) error, chatID int64) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		sendFunc(chatID, "❌ Nota não encontrada.", "")
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, note.RelativePath, note.Filename)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Read error: %v", err)
		sendFunc(chatID, fmt.Sprintf("❌ Erro ao ler arquivo: %s", note.Filename), "")
		return
	}

	if err := sendFunc(chatID, string(content), ""); err != nil {
		log.Printf("Send content error: %v", err)
	}
}

// HandleViewList resolve os caminhos das atividades ou referências e envia em forma de botões
func HandleViewList(dbPath string, shortID string, listType string, menuSenderFunc func(int64, string, string, string) error, chatID int64) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	var list interface{}
	var title string
	var prefix string
	if listType == "homework" {
		list = note.Homework
		title = "📝 <b>Atividades Disponíveis</b>"
		prefix = "open_hw"
	} else if listType == "references" {
		list = note.References
		title = "🔗 <b>Referências Disponíveis</b>"
		prefix = "open_ref"
	}

	if list == nil {
		return
	}

	itemsArray, ok := list.([]interface{})
	if !ok || len(itemsArray) == 0 {
		return
	}

	var keyboard InlineKeyboardMarkup
	count := 0
	for i, item := range itemsArray {
		if mapItem, isMap := item.(map[string]interface{}); isMap {
			for k, v := range mapItem {
				valStr := fmt.Sprintf("%v", v)
				if strings.HasPrefix(k, "link_") {
					// Links externos abrem direto no navegador
					keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []InlineKeyboardButton{
						{Text: "🌐 Abrir Link (" + k + ")", Url: valStr},
					})
					count++
				} else if strings.HasPrefix(k, "path_") {
					fileName := filepath.Base(valStr)
					// Callback format: prefix:shortID:index
					callback := fmt.Sprintf("%s:%s:%d", prefix, shortID, i)
					keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []InlineKeyboardButton{
						{Text: "📄 " + fileName, CallbackData: callback},
					})
					count++
				}
			}
		}
	}

	if count == 0 {
		return
	}

	kbBytes, _ := json.Marshal(keyboard)
	menuSenderFunc(chatID, title, "HTML", string(kbBytes))
}

// HandleOpenFileByIndex abre um arquivo da lista (homework ou references) usando o índice
func HandleOpenFileByIndex(dbPath string, shortID string, index int, listType string, sendFunc func(int64, string, string) error, chatID int64) {
	db, err := database.InitDB(dbPath)
	if err != nil {
		return
	}
	defer db.Close()

	doc, err := db.FindFirst(query.NewQuery("notes").Where(query.Field("short_id").Eq(shortID)))
	if err != nil || doc == nil {
		sendFunc(chatID, "❌ Nota original não encontrada.", "")
		return
	}

	var note database.NoteDoc
	doc.Unmarshal(&note)

	var list interface{}
	if listType == "homework" {
		list = note.Homework
	} else {
		list = note.References
	}

	itemsArray, ok := list.([]interface{})
	if !ok || index >= len(itemsArray) {
		sendFunc(chatID, "❌ Item não encontrado na lista.", "")
		return
	}

	item := itemsArray[index]
	mapItem, ok := item.(map[string]interface{})
	if !ok {
		sendFunc(chatID, "❌ Formato de item inválido.", "")
		return
	}

	var targetPath string
	for k, v := range mapItem {
		if strings.HasPrefix(k, "path_") {
			targetPath = fmt.Sprintf("%v", v)
			break
		}
	}

	if targetPath == "" {
		sendFunc(chatID, "❌ Caminho não encontrado neste item.", "")
		return
	}

	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	baseNoteDir := filepath.Join(vaultRoot, note.RelativePath)
	fullPath := filepath.Join(baseNoteDir, targetPath)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Read error: %v", err)
		sendFunc(chatID, "❌ Erro ao ler o arquivo: "+filepath.Base(fullPath), "")
		return
	}

	if err := sendFunc(chatID, string(content), ""); err != nil {
		log.Printf("Send content error: %v", err)
	}
}

// HandleAbrirArquivo abre um arquivo a partir de seu caminho codificado em hex
func HandleAbrirArquivo(hexPath string, sendFunc func(int64, string, string) error, chatID int64) {
	decoded, err := hex.DecodeString(hexPath)
	if err != nil {
		sendFunc(chatID, "❌ Comando de arquivo inválido.", "")
		return
	}
	
	relPath := string(decoded)
	vaultRoot := os.Getenv("VAULT_PATH")
	if vaultRoot == "" {
		vaultRoot, _ = filepath.Abs("..")
	}
	fullPath := filepath.Join(vaultRoot, relPath)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		log.Printf("Read error: %v", err)
		sendFunc(chatID, "❌ Erro ao ler o arquivo especificado.", "")
		return
	}

	if err := sendFunc(chatID, string(content), ""); err != nil {
		log.Printf("Send content error: %v", err)
	}
}
