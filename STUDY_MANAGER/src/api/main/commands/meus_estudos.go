package commands

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"study_manager/src/infra/database"
)

type TreeNode struct {
	Name     string
	IsDir    bool
	ShortID  string
	Children map[string]*TreeNode
}

func NewTreeNode(name string, isDir bool, shortID string) *TreeNode {
	return &TreeNode{
		Name:     name,
		IsDir:    isDir,
		ShortID:  shortID,
		Children: make(map[string]*TreeNode),
	}
}

// RenderTree gera a representação visual da árvore
func RenderTree(node *TreeNode, level int, sb *strings.Builder, escapeFunc func(string) string) {
	keys := make([]string, 0, len(node.Children))
	for k := range node.Children {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		child := node.Children[k]
		indent := strings.Repeat("  ", level)
		icon := "📄"
		cmd := ""
		if child.IsDir {
			icon = "📁"
		} else if child.ShortID != "" {
			cmd = fmt.Sprintf(" — /ver_%s", child.ShortID)
		}

		safeName := escapeFunc(child.Name)
		sb.WriteString(fmt.Sprintf("%s%s %s%s\n", indent, icon, safeName, cmd))
		if child.IsDir {
			RenderTree(child, level+1, sb, escapeFunc)
		}
	}
}

// HandleMeusEstudos processa a árvore de notas
func HandleMeusEstudos(repo *database.Repository, syncFunc func() error, sendFunc func(int64, string, string) error, escapeFunc func(string) string, chatID int64) {
	if err := syncFunc(); err != nil {
		log.Printf("Sync error: %v", err)
	}

	notes, err := repo.GetAllNotesInfo(context.Background())
	if err != nil {
		log.Printf("Query error: %v", err)
		sendFunc(chatID, "❌ Erro ao acessar o banco de dados.", "")
		return
	}

	root := NewTreeNode("Raiz", true, "")

	for _, n := range notes {
		normalizedPath := filepath.ToSlash(n.RelativePath)
		parts := strings.Split(normalizedPath, "/")

		current := root
		for i, part := range parts {
			if part == "" {
				continue
			}
			isLast := i == len(parts)-1
			displayName := part
			nodeShortID := ""
			if isLast {
				displayName = n.Tema
				nodeShortID = n.ShortID
			}

			if _, exists := current.Children[displayName]; !exists {
				current.Children[displayName] = NewTreeNode(displayName, !isLast, nodeShortID)
			}
			current = current.Children[displayName]
		}
	}

	var sb strings.Builder
	sb.WriteString("📚 <b>MAPA DE ESTUDOS</b>\n")
	sb.WriteString("──────────────────\n")
	RenderTree(root, 0, &sb, escapeFunc)

	if err := sendFunc(chatID, sb.String(), "HTML"); err != nil {
		log.Printf("Send error: %v", err)
	}
}
