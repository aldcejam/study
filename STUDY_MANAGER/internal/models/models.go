package models

// Revision representa uma entrada de revisão em uma nota.
type Revision struct {
	Data   string `json:"data"`   // "DD/MM/YYYY"
	Status string `json:"status"` // "x" = feito, " " = pendente
}

// Reference representa uma referência bibliográfica ou link.
type Reference struct {
	Description string `json:"description"`
	Source      string `json:"source"`
}

// NoteMetadata é a saída do scanner: metadados brutos de uma nota.
type NoteMetadata struct {
	Filename     string      `json:"filename"`
	RelativePath string      `json:"relative_path"`
	Tema         string      `json:"tema"`
	Subtema      *string     `json:"subtema"`
	Revisoes     []Revision  `json:"revisoes"`
	Tags         []string    `json:"tags"`
	References   interface{} `json:"references"` // string ou {description, source}
	Activity     *string     `json:"activity"`
	UpdatedAt    string      `json:"updatedAt"` // ISO 8601
}

// ProcessedNote é a saída do processor: nota com status calculado.
type ProcessedNote struct {
	NoteMetadata
	ID             string  `json:"id"`
	DiasAtraso     int     `json:"dias_atraso"`
	StatusRevisao  string  `json:"status_revisao"` // ATRASADA, HOJE, FUTURA, EM_DIA
	LastNotifiedAt *string `json:"last_notified_at"`
}
