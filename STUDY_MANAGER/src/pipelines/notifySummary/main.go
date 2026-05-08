package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	ColorBlue  = "\033[0;34m"
	ColorGreen = "\033[0;32m"
	ColorRed   = "\033[0;31m"
	ColorReset = "\033[0m"
)

func main() {
	header()

	// 0. Ensure 'go' is in PATH
	if err := ensureGoInPath(); err != nil {
		fatal("Go não encontrado: %v", err)
	}

	// 1. Build stages
	stages := []string{"scanner", "summaryNotification", "notifier"}
	if err := buildStages(stages); err != nil {
		fatal("Erro na compilação: %v", err)
	}

	// 2. Execute pipeline
	if err := executePipeline(); err != nil {
		fatal("Erro na execução do pipeline: %v", err)
	}

	success()
}

func header() {
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s STUDY_MANAGER — Pipeline de Estudos (Go)       %s\n", ColorBlue, ColorReset)
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", ColorBlue, ColorReset)
}

func buildStages(stages []string) error {
	fmt.Printf("%s⚙️  Compilando binários...%s\n", ColorBlue, ColorReset)

	if err := os.MkdirAll("./bin", 0755); err != nil {
		return err
	}

	for _, stage := range stages {
		cmd := exec.Command("go", "build", "-o", filepath.Join("bin", stage), "./src/processes/"+stage)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("falha ao compilar %s: %v\n%s", stage, err, string(output))
		}
	}

	fmt.Printf("%s✔  Compilação concluída.%s\n", ColorGreen, ColorReset)
	return nil
}

func executePipeline() error {
	fmt.Printf("%s▶  Executando pipeline...%s\n", ColorBlue, ColorReset)

	// Define the commands
	scanner := exec.Command("./bin/scanner")
	summary := exec.Command("./bin/summaryNotification")
	notifier := exec.Command("./bin/notifier")

	// Chain: scanner -> summaryNotification -> notifier

	// Scanner Out -> Summary In
	pr1, pw1 := io.Pipe()
	scanner.Stdout = pw1
	summary.Stdin = pr1

	// Summary Out -> Notifier In
	pr2, pw2 := io.Pipe()
	summary.Stdout = pw2
	notifier.Stdin = pr2

	// Final output to stdout
	notifier.Stdout = os.Stdout
	notifier.Stderr = os.Stderr

	// Collect errors
	scanner.Stderr = os.Stderr
	summary.Stderr = os.Stderr

	// Start all
	if err := scanner.Start(); err != nil {
		return err
	}
	if err := summary.Start(); err != nil {
		return err
	}
	if err := notifier.Start(); err != nil {
		return err
	}

	// Wait in order
	if err := scanner.Wait(); err != nil {
		pw1.Close()
		return fmt.Errorf("scanner: %v", err)
	}
	pw1.Close()

	if err := summary.Wait(); err != nil {
		pw2.Close()
		return fmt.Errorf("summaryNotification: %v", err)
	}
	pw2.Close()

	if err := notifier.Wait(); err != nil {
		return fmt.Errorf("notifier: %v", err)
	}

	return nil
}

func ensureGoInPath() error {
	_, err := exec.LookPath("go")
	if err == nil {
		return nil
	}

	// Try common paths
	tryPaths := []string{
		"/home/apm/go/bin",
		"/usr/local/go/bin",
		"/snap/go/current/bin",
	}

	for _, path := range tryPaths {
		goPath := filepath.Join(path, "go")
		if _, err := os.Stat(goPath); err == nil {
			newPath := path + string(os.PathListSeparator) + os.Getenv("PATH")
			os.Setenv("PATH", newPath)
			return nil
		}
	}

	return fmt.Errorf("não foi possível encontrar o executável 'go' no PATH ou em locais comuns")
}

func success() {
	fmt.Println("")
	fmt.Printf("%s✨ Processo concluído!%s\n", ColorGreen, ColorReset)
}

func fatal(format string, a ...any) {
	fmt.Printf("\n%s❌ %s%s\n", ColorRed, fmt.Sprintf(format, a...), ColorReset)
	os.Exit(1)
}
