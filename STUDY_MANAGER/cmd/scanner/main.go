package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"study_manager/internal/frontmatter"
	"study_manager/internal/models"
)

// ignoredDirs são pastas que devem ser ignoradas ao varrer o vault.
var ignoredDirs = []string{"STUDY_MANAGER", "__MANAGER__", ".obsidian", ".git", ".agents", ".gemini"}

func shouldIgnore(path string) bool {
	for _, ignored := range ignoredDirs {
		if strings.Contains(path, ignored) {
			return true
		}
	}
	return false
}

// getGitUpdatedAt retorna a data do último commit que tocou o arquivo.
// Faz fallback para mtime e depois para epoch se tudo falhar.
func getGitUpdatedAt(filepath string) string {
	cmd := exec.Command("git", "log", "-1", "--format=%cI", "--", filepath)
	out, err := cmd.Output()
	if err == nil {
		result := strings.TrimSpace(string(out))
		if result != "" {
			return result
		}
	}

	// Fallback: mtime
	info, err := os.Stat(filepath)
	if err == nil {
		return info.ModTime().Format(time.RFC3339)
	}

	return "1970-01-01T00:00:00Z"
}

// parseRevisions extrai pares revision_DD-MM-YYYY: true/false do frontmatter.
func parseRevisions(fm map[string]interface{}) []models.Revision {
	var revs []models.Revision
	for key, val := range fm {
		if !strings.HasPrefix(key, "revision_") {
			continue
		}
		datePart := strings.ReplaceAll(strings.TrimPrefix(key, "revision_"), "-", "/")
		status := " "
		if b, ok := val.(bool); ok && b {
			status = "x"
		}
		revs = append(revs, models.Revision{Data: datePart, Status: status})
	}
	return revs
}

// strOrNil converte interface{} para *string, retornando nil se vazio ou nil.
func strOrNil(v interface{}) *string {
	if v == nil {
		return nil
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return nil
	}
	return &s
}

// generateID gera um hash MD5 de 12 caracteres para identificar unicamente uma nota.
func generateID(path, filename, revDate string) string {
	h := md5.Sum([]byte(fmt.Sprintf("%s/%s@%s", path, filename, revDate)))
	return fmt.Sprintf("%x", h)[:12]
}

func scanNotes(baseDir string) ([]models.NoteMetadata, error) {
	var notes []models.NoteMetadata

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // ignora erros de acesso
		}
		if d.IsDir() && shouldIgnore(path) {
			return filepath.SkipDir
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		raw := frontmatter.Extract(string(content))
		fm := frontmatter.ParseSimple(raw)

		relPath, _ := filepath.Rel(baseDir, filepath.Dir(path))

		// Tema: tenta campo 'tema', depois 'title', depois nome do arquivo
		tema := d.Name()[:len(d.Name())-3] // remove .md
		if t, ok := fm["tema"].(string); ok && t != "" {
			tema = t
		} else if t, ok := fm["title"].(string); ok && t != "" {
			tema = t
		}

		// Tags e references como interface{} para preservar formato original
		var tags []string
		if rawTags, ok := fm["tags"]; ok {
			if tagList, ok := rawTags.([]interface{}); ok {
				for _, t := range tagList {
					if s, ok := t.(string); ok {
						tags = append(tags, s)
					}
				}
			}
		}

		var refs interface{} = []interface{}{}
		if r, ok := fm["references"]; ok && r != nil {
			refs = r
		}

		note := models.NoteMetadata{
			Filename:     d.Name(),
			RelativePath: relPath,
			Tema:         tema,
			Subtema:      strOrNil(fm["subtema"]),
			Revisoes:     parseRevisions(fm),
			Tags:         tags,
			References:   refs,
			Activity:     strOrNil(fm["activity"]),
			UpdatedAt:    getGitUpdatedAt(path),
		}
		notes = append(notes, note)
		return nil
	})

	return notes, err
}

func main() {
	// O binário é executado de dentro de STUDY_MANAGER/,
	// o vault está dois níveis acima (estudos/)
	selfDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, "scanner: erro ao obter diretório:", err)
		os.Exit(1)
	}
	// Sobe dois níveis: STUDY_MANAGER/ -> estudos/
	vaultDir := filepath.Join(selfDir, "..", "..")
	vaultDir, _ = filepath.Abs(vaultDir)

	notes, err := scanNotes(vaultDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "scanner: erro ao varrer vault:", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(notes); err != nil {
		fmt.Fprintln(os.Stderr, "scanner: erro ao serializar JSON:", err)
		os.Exit(1)
	}
}
