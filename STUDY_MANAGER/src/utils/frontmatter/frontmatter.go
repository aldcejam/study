// Package frontmatter extrai YAML frontmatter de arquivos Markdown.
// Usa apenas regexp e stdlib, sem dependências externas.
package frontmatter

import (
	"regexp"
	"strings"
)

var fmRegex = regexp.MustCompile(`(?s)\A---\s*\n(.*?)\n---`)

// Extract retorna o bloco YAML bruto entre os delimitadores --- de um conteúdo Markdown.
// Retorna string vazia se não encontrar frontmatter.
func Extract(content string) string {
	m := fmRegex.FindStringSubmatch(content)
	if len(m) < 2 {
		return ""
	}
	return m[1]
}

// ParseSimple faz um parsing linha-a-linha do YAML simples (chave: valor).
// Suporta apenas o subconjunto necessário: strings, booleans e listas simples.
// Retorna um map[string]interface{}.
func ParseSimple(yamlStr string) map[string]interface{} {
	result := make(map[string]interface{})
	lines := strings.Split(yamlStr, "\n")

	var currentKey string
	var listItems []interface{}
	inList := false

	flush := func() {
		if currentKey != "" && inList {
			result[currentKey] = listItems
		}
		currentKey = ""
		listItems = nil
		inList = false
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Item de lista
		if strings.HasPrefix(trimmed, "- ") {
			val := strings.TrimPrefix(trimmed, "- ")
			val = unquote(val)

			// Se o item da lista contiver ":", tentamos tratar como um mapa simples (ex: - link_1: url)
			if idx := strings.Index(val, ":"); idx > 0 {
				k := strings.TrimSpace(val[:idx])
				v := strings.TrimSpace(val[idx+1:])
				listItems = append(listItems, map[string]interface{}{k: unquote(v)})
			} else {
				listItems = append(listItems, val)
			}

			inList = true
			continue
		}

		// Chave: valor
		idx := strings.Index(trimmed, ":")
		if idx < 0 {
			continue
		}
		flush()

		key := strings.TrimSpace(trimmed[:idx])
		val := strings.TrimSpace(trimmed[idx+1:])
		currentKey = key

		if val == "" {
			// Pode ser início de lista — aguarda próximas linhas
			continue
		}

		// Boolean
		lower := strings.ToLower(val)
		if lower == "true" {
			result[key] = true
			currentKey = ""
			continue
		}
		if lower == "false" {
			result[key] = false
			currentKey = ""
			continue
		}
		if lower == "null" || lower == "~" {
			result[key] = nil
			currentKey = ""
			continue
		}

		result[key] = unquote(val)
		currentKey = ""
	}
	flush()
	return result
}

func unquote(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
