package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"study_manager/internal/models"
)

const bigFuture = 9999

var dateFormats = []string{"02/01/06", "02/01/2006", "2006-01-02"}

func parseDate(s string) time.Time {
	s = strings.TrimSpace(s)
	for _, fmt := range dateFormats {
		t, err := time.Parse(fmt, s)
		if err == nil {
			return t
		}
	}
	// Data muito no futuro — não dispara alerta
	return time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
}

func generateID(path, filename, revDate string) string {
	h := md5.Sum([]byte(fmt.Sprintf("%s/%s@%s", path, filename, revDate)))
	return fmt.Sprintf("%x", h)[:12]
}

func processNotes(notes []models.NoteMetadata) []models.ProcessedNote {
	today := time.Now().Truncate(24 * time.Hour)
	var result []models.ProcessedNote

	for _, note := range notes {
		diasAtraso := bigFuture
		status := "FUTURA"
		dataRef := "00/00/00"

		// Filtra apenas revisões pendentes (status == " ")
		var pendentes []models.Revision
		for _, r := range note.Revisoes {
			if r.Status == " " {
				pendentes = append(pendentes, r)
			}
		}

		if len(pendentes) == 0 {
			status = "EM_DIA"
			diasAtraso = 0
		} else {
			// Encontra a revisão mais urgente
			for _, rev := range pendentes {
				d := parseDate(rev.Data)
				diff := int(d.Truncate(24 * time.Hour).Sub(today).Hours() / 24)
				if diff < diasAtraso {
					diasAtraso = diff
					dataRef = rev.Data
				}
			}
			switch {
			case diasAtraso < 0:
				status = "ATRASADA"
			case diasAtraso == 0:
				status = "HOJE"
			default:
				status = "FUTURA"
			}
		}

		pn := models.ProcessedNote{
			NoteMetadata:   note,
			ID:             generateID(note.RelativePath, note.Filename, dataRef),
			DiasAtraso:     diasAtraso,
			StatusRevisao:  status,
			LastNotifiedAt: nil,
		}
		result = append(result, pn)
	}
	return result
}

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil || len(strings.TrimSpace(string(raw))) == 0 {
		fmt.Fprintln(os.Stderr, "processor: nenhum dado recebido via stdin")
		os.Exit(1)
	}

	var notes []models.NoteMetadata
	if err := json.Unmarshal(raw, &notes); err != nil {
		fmt.Fprintln(os.Stderr, "processor: erro ao parsear JSON:", err)
		os.Exit(1)
	}

	processed := processNotes(notes)

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(processed); err != nil {
		fmt.Fprintln(os.Stderr, "processor: erro ao serializar JSON:", err)
		os.Exit(1)
	}
}
