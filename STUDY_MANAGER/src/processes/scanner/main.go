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

	"study_manager/src/utils/frontmatter"
	"study_manager/src/utils/models"
)

// ignoredDirs são pastas que devem ser ignoradas ao varrer o vault.
var ignoredDirs = []string{"STUDY_MANAGER", "__MANAGER__", ".obsidian", ".git", ".agents", ".gemini"}

// shouldIgnore verifica se o caminho informado pertence a uma lista de diretórios
// que devem ser ignorados durante a varredura do vault (ex: .git, .obsidian).
func shouldIgnore(path string) bool {
	for _, ignored := range ignoredDirs {
		if strings.Contains(path, ignored) {
			return true
		}
	}
	return false
}

// getGitUpdatedAt recupera a data da última alteração do arquivo registrada no Git.
//
// O comando tenta obter o timestamp via 'git log' no formato ISO 8601.
// Em caso de falha (arquivo não versionado ou erro), utiliza o mtime do sistema de arquivos.
// Como fallback final, retorna a época Unix (1970-01-01).
func getGitUpdatedAt(filePath string) string {
	// Prioridade: mtime do sistema (mais fiel a edições locais)
	info, err := os.Stat(filePath)
	if err == nil {
		return info.ModTime().Format(time.RFC3339)
	}

	// Fallback: Git log
	cmd := exec.Command("git", "log", "-1", "--format=%cI", "--", filePath)
	out, err := cmd.Output()
	if err == nil {
		result := strings.TrimSpace(string(out))
		if result != "" {
			return result
		}
	}

	return "1970-01-01T00:00:00Z"
}

// parseRevisions processa o map do frontmatter em busca de chaves no formato
// 'revision_DD-MM-YYYY' e as converte para uma lista de modelos Revision.
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

// strOrNil é uma função utilitária que converte uma interface para um ponteiro de string.
// Retorna nil caso o valor seja nulo, não seja uma string ou seja uma string vazia.
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

// generateID cria um identificador único de 12 caracteres (MD5 truncado)
// baseado no caminho do arquivo, nome e data de revisão para rastreamento.
func generateID(path, filename, revDate string) string {
	h := md5.Sum([]byte(fmt.Sprintf("%s/%s@%s", path, filename, revDate)))
	return fmt.Sprintf("%x", h)[:12]
}

// scanNotes executa a varredura recursiva do diretório base, filtrando arquivos Markdown
// e extraindo seus metadados apenas se seguirem o padrão esperado (primeiro atributo 'tema').
func scanNotes(baseDir string) ([]models.ScannerOutput, error) {
	var notes []models.ScannerOutput

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
		if raw == "" {
			return nil
		}

		// Validação: o primeiro atributo deve ser "tema"
		lines := strings.Split(raw, "\n")
		isFirstTema := false
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" || strings.HasPrefix(trimmed, "#") {
				continue
			}
			if strings.HasPrefix(trimmed, "tema:") {
				isFirstTema = true
			}
			break
		}

		if !isFirstTema {
			return nil
		}

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

		var hw interface{} = []interface{}{}
		if h, ok := fm["homework"]; ok && h != nil {
			hw = h
		}

		note := models.ScannerOutput{
			Filename:     d.Name(),
			RelativePath: relPath,
			Tema:         tema,
			Subtema:      strOrNil(fm["subtema"]),
			Revisoes:     parseRevisions(fm),
			Tags:         tags,
			References:   refs,
			Homework:     hw,
			Activity:     strOrNil(fm["activity"]),
			UpdatedAt:    getGitUpdatedAt(path),
		}
		notes = append(notes, note)
		return nil
	})

	return notes, err
}

func main() {
	selfDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, "scanner: erro ao obter diretório:", err)
		os.Exit(1)
	}

	// Prioridade para variável de ambiente (útil no Docker)
	vaultDir := os.Getenv("VAULT_PATH")
	if vaultDir == "" {
		// Fallback: Sobe dois níveis: STUDY_MANAGER/ -> estudos/
		vaultDir = filepath.Join(selfDir, "..", "..")
	}
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
