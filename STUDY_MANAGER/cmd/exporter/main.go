package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"study_manager/internal/models"
)

// loadHistory carrega o cache persistente de notificações.
func loadHistory(outputDir string) map[string]models.ProcessedNote {
	path := filepath.Join(outputDir, "revisoes_METADATA.json")
	hist := make(map[string]models.ProcessedNote)

	data, err := os.ReadFile(path)
	if err != nil {
		return hist
	}

	var items []models.ProcessedNote
	if err := json.Unmarshal(data, &items); err != nil {
		return hist
	}
	for _, item := range items {
		hist[item.ID] = item
	}
	return hist
}

// saveHistory persiste o estado atual das notas processadas.
func saveHistory(outputDir string, notes []models.ProcessedNote) error {
	os.MkdirAll(outputDir, 0755)
	path := filepath.Join(outputDir, "revisoes_METADATA.json")
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// shouldNotify decide se um item deve ser notificado agora.
func shouldNotify(item models.ProcessedNote, hist map[string]models.ProcessedNote) bool {
	prev, ok := hist[item.ID]
	if !ok || prev.LastNotifiedAt == nil {
		return true
	}

	lastNotified, err1 := time.Parse(time.RFC3339, *prev.LastNotifiedAt)
	updatedAt, err2 := time.Parse(time.RFC3339, item.UpdatedAt)

	if err1 != nil || err2 != nil {
		return true
	}

	// Notifica se a nota foi modificada após o último alerta
	if updatedAt.After(lastNotified) {
		return true
	}

	// Notifica se faz mais de 2 dias desde o último alerta
	if time.Since(lastNotified).Hours() >= 48 {
		return true
	}

	return false
}

func exportAndFormat(notes []models.ProcessedNote, outputDir string) string {
	hist := loadHistory(outputDir)
	now := time.Now().Format(time.RFC3339)

	var urgentes []models.ProcessedNote
	for _, n := range notes {
		if n.StatusRevisao == "ATRASADA" || n.StatusRevisao == "HOJE" {
			urgentes = append(urgentes, n)
		}
	}

	if len(urgentes) == 0 {
		saveHistory(outputDir, notes)
		return "SEM_REVISOES"
	}

	var aNotificar []models.ProcessedNote
	for i := range notes {
		n := &notes[i]
		if n.StatusRevisao != "ATRASADA" && n.StatusRevisao != "HOJE" {
			continue
		}
		if shouldNotify(*n, hist) {
			n.LastNotifiedAt = &now
			aNotificar = append(aNotificar, *n)
		} else if prev, ok := hist[n.ID]; ok {
			n.LastNotifiedAt = prev.LastNotifiedAt
		}
	}

	saveHistory(outputDir, notes)

	if len(aNotificar) == 0 {
		return "SEM_ALTERACOES"
	}

	// Ordena por urgência (dias_atraso crescente)
	sort.Slice(aNotificar, func(i, j int) bool {
		return aNotificar[i].DiasAtraso < aNotificar[j].DiasAtraso
	})

	var sb strings.Builder
	sb.WriteString("📚 *RESUMO DE ESTUDOS*\n")
	sb.WriteString("──────────────────\n\n")

	for _, note := range aNotificar {
		icon := "📅"
		atrasoStr := "HOJE"
		if note.StatusRevisao == "ATRASADA" {
			icon = "🚨"
			atrasoStr = fmt.Sprintf("%dd atrás", abs(note.DiasAtraso))
		}

		// Pega o último segmento do caminho como contexto
		parts := strings.Split(note.RelativePath, string(filepath.Separator))
		contexto := parts[len(parts)-1]

		sb.WriteString(fmt.Sprintf("%s *%s*\n", icon, note.Tema))
		sb.WriteString(fmt.Sprintf("   └ 📂 %s | ⏳ %s\n", contexto, atrasoStr))

		if note.Activity != nil && *note.Activity != "" {
			sb.WriteString(fmt.Sprintf("   📝 *Atividade:* %s\n", *note.Activity))
		}

		if note.References != nil {
			switch refs := note.References.(type) {
			case []interface{}:
				for _, ref := range refs {
					switch r := ref.(type) {
					case map[string]interface{}:
						desc, _ := r["description"].(string)
						src, _ := r["source"].(string)
						sb.WriteString(fmt.Sprintf("   🔗 _%s: %s_\n", desc, src))
					case string:
						sb.WriteString(fmt.Sprintf("   🔗 _%s_\n", r))
					}
				}
			}
			sb.WriteString("\n")
		} else {
			sb.WriteString("\n")
		}
	}

	sb.WriteString("──────────────────\n")
	sb.WriteString("_Abra seu Obsidian para revisar!_")

	return sb.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil || len(strings.TrimSpace(string(raw))) == 0 {
		fmt.Fprintln(os.Stderr, "exporter: nenhum dado recebido via stdin")
		os.Exit(1)
	}

	var notes []models.ProcessedNote
	if err := json.Unmarshal(raw, &notes); err != nil {
		fmt.Fprintln(os.Stderr, "exporter: erro ao parsear JSON:", err)
		os.Exit(1)
	}

	// output/ fica ao lado do binário (dentro de STUDY_MANAGER/)
	selfDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	outputDir := filepath.Join(selfDir, "output")

	msg := exportAndFormat(notes, outputDir)
	fmt.Print(msg)
}
