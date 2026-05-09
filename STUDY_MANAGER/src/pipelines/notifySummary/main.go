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

	// 1. Execute pipeline
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



func success() {
	fmt.Println("")
	fmt.Printf("%s✨ Processo concluído!%s\n", ColorGreen, ColorReset)
}

func fatal(format string, a ...any) {
	fmt.Printf("\n%s❌ %s%s\n", ColorRed, fmt.Sprintf(format, a...), ColorReset)
	os.Exit(1)
}
