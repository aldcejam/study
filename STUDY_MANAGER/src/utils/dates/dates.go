package dates

import (
	"strings"
	"time"
)

var dateFormats = []string{"02/01/06", "02/01/2006", "2006-01-02"}

// ParseDate converte uma string de data para o tipo time.Time.
// Caso não seja possível parsear, ele retorna uma data no futuro distante para ignorá-la de revisões iminentes.
func ParseDate(s string) time.Time {
	s = strings.TrimSpace(s)
	for _, fmt := range dateFormats {
		t, err := time.Parse(fmt, s)
		if err == nil {
			return t
		}
	}
	return time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
}
