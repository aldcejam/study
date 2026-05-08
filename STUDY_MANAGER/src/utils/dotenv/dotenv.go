// Package dotenv carrega variáveis de um arquivo .env sem dependências externas.
package dotenv

import (
	"bufio"
	"os"
	"strings"
)

// Load busca um arquivo .env nos caminhos fornecidos (em ordem) e define
// as variáveis de ambiente no processo atual. Para em no primeiro arquivo encontrado.
func Load(paths ...string) {
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			if os.Getenv(key) == "" { // não sobrescreve variáveis já definidas
				os.Setenv(key, val)
			}
		}
		return // para no primeiro arquivo encontrado
	}
}
