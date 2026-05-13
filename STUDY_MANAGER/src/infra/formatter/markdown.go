package formatter

import (
	"regexp"
	"strings"
)

// MarkdownConverter é responsável por transformar Markdown do Obsidian em HTML compatível com o Telegram.
type MarkdownConverter struct{}

// NewMarkdownConverter cria uma nova instância do conversor.
func NewMarkdownConverter() *MarkdownConverter {
	return &MarkdownConverter{}
}

/**
 * Convert é o método principal que orquestra a conversão de Markdown para HTML.
 * Ele segue uma ordem específica de processamento para garantir que as tags HTML
 * geradas não sejam interferidas por outras regras.
 */
func (c *MarkdownConverter) Convert(text string) string {
	// 1. Escapar caracteres HTML básicos para evitar conflitos
	text = c.EscapeHTML(text)

	// 2. Formatar blocos de código (deve ser um dos primeiros para "proteger" o conteúdo)
	text = c.FormatCodeBlocks(text)

	// 3. Formatar código inline
	text = c.FormatInlineCode(text)

	// 4. Formatar cabeçalhos (Headers)
	text = c.FormatHeaders(text)

	// 5. Formatar negrito e itálico
	text = c.FormatBold(text)
	text = c.FormatItalic(text)

	// 6. Formatar checklists e listas
	text = c.FormatCheckboxes(text)
	text = c.FormatLists(text)

	// 7. Formatar links (padrão e Obsidian)
	text = c.FormatLinks(text)
	text = c.FormatObsidianLinks(text)

	return text
}

/**
 * EscapeHTML escapa os caracteres que possuem significado especial em HTML (&, <, >).
 * Este método deve ser chamado ANTES de qualquer outro que adicione tags HTML.
 */
func (c *MarkdownConverter) EscapeHTML(text string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(text)
}

/**
 * FormatHeaders converte cabeçalhos Markdown (#, ##, ###) em texto em negrito.
 * Como o Telegram não suporta tamanhos de fonte diferentes, usamos negrito para destaque.
 */
func (c *MarkdownConverter) FormatHeaders(text string) string {
	re := regexp.MustCompile(`(?m)^#{1,6}\s+(.*)$`)
	return re.ReplaceAllString(text, "<b>$1</b>")
}

/**
 * FormatBold converte a sintaxe de negrito (**texto** ou __texto__) para tags <b>.
 */
func (c *MarkdownConverter) FormatBold(text string) string {
	// Go regexp não suporta backreferences (\1), então fazemos em duas etapas
	reStar := regexp.MustCompile(`\*\*(.*?)\*\*`)
	text = reStar.ReplaceAllString(text, "<b>$1</b>")
	reUnder := regexp.MustCompile(`__(.*?)__`)
	return reUnder.ReplaceAllString(text, "<b>$1</b>")
}

/**
 * FormatItalic converte a sintaxe de itálico (*texto* ou _texto_) para tags <i>.
 */
func (c *MarkdownConverter) FormatItalic(text string) string {
	// Go regexp não suporta backreferences (\1), então fazemos em duas etapas
	reStar := regexp.MustCompile(`\*([^*]+)\*`)
	text = reStar.ReplaceAllString(text, "<i>$1</i>")
	reUnder := regexp.MustCompile(`_([^_]+)_`)
	return reUnder.ReplaceAllString(text, "<i>$1</i>")
}

/**
 * FormatCodeBlocks converte blocos de código (```...```) para tags <pre>.
 */
func (c *MarkdownConverter) FormatCodeBlocks(text string) string {
	re := regexp.MustCompile("(?s)```(.*?)\n?(.*?)```")
	return re.ReplaceAllString(text, "<pre>$2</pre>")
}

/**
 * FormatInlineCode converte código inline (`texto`) para tags <code>.
 */
func (c *MarkdownConverter) FormatInlineCode(text string) string {
	re := regexp.MustCompile("`(.*?)`")
	return re.ReplaceAllString(text, "<code>$1</code>")
}

/**
 * FormatCheckboxes converte a sintaxe de checklist do Obsidian (- [ ] e - [x]) em emojis.
 */
func (c *MarkdownConverter) FormatCheckboxes(text string) string {
	text = strings.ReplaceAll(text, "- [ ]", "⬜")
	text = strings.ReplaceAll(text, "- [x]", "✅")
	text = strings.ReplaceAll(text, "- [X]", "✅")
	return text
}

/**
 * FormatLists converte marcadores de lista simples (- ou *) em um caractere de bullet (•).
 */
func (c *MarkdownConverter) FormatLists(text string) string {
	re := regexp.MustCompile(`(?m)^(\s*)[\-\*]\s+(.*)$`)
	return re.ReplaceAllString(text, "$1• $2")
}

/**
 * FormatLinks converte links Markdown padrão [texto](url) em tags <a>.
 */
func (c *MarkdownConverter) FormatLinks(text string) string {
	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	return re.ReplaceAllString(text, `<a href="$2">$1</a>`)
}

/**
 * FormatObsidianLinks trata links internos do Obsidian [[Link]]. 
 * Como o Telegram não sabe resolver esses links, apenas removemos os colchetes.
 */
func (c *MarkdownConverter) FormatObsidianLinks(text string) string {
	re := regexp.MustCompile(`\[\[(.*?)\]\]`)
	return re.ReplaceAllString(text, "$1")
}
