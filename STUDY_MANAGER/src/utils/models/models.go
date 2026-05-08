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

// ScannerOutput é a saída do scanner: metadados brutos de uma nota.
type ScannerOutput struct {
	Filename     string      `json:"filename"`
	RelativePath string      `json:"relative_path"`
	Tema         string      `json:"tema"`
	Subtema      *string     `json:"subtema"`
	Revisoes     []Revision  `json:"revisoes"`
	Tags         []string    `json:"tags"`
	References   interface{} `json:"references"` // string ou {description, source}
	Homework     interface{} `json:"homework"`   // string, list of strings or list of maps
	Activity     *string     `json:"activity"`
	UpdatedAt    string      `json:"updatedAt"` // ISO 8601
}

// SummaryNotificationOutput é a estrutura interna gerada para os itens a notificar.
type SummaryNotificationOutput struct {
	ScannerOutput
	ID             string  `json:"id"`
	DiasAtraso     int     `json:"dias_atraso"`
	StatusRevisao  string  `json:"status_revisao"` // ATRASADA, HOJE, FUTURA, EM_DIA
	LastNotifiedAt *string `json:"last_notified_at"`
}
