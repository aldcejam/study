package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	ColorBlue  = "\033[0;34m"
	ColorGreen = "\033[0;32m"
	ColorRed   = "\033[0;31m"
	ColorReset = "\033[0m"
)

func main() {
	header()

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

	// No modo imutável, usamos os binários que estão na pasta /app/bin/
	scannerCmd := exec.Command("./bin/scanner")
	summaryCmd := exec.Command("./bin/summaryNotification")
	notifierCmd := exec.Command("./bin/notifier")

	// Pipeline: Scanner -> SummaryNotification
	pr1, pw1 := io.Pipe()
	scannerCmd.Stdout = pw1
	summaryCmd.Stdin = pr1

	// Pipeline: SummaryNotification -> Notifier
	pr2, pw2 := io.Pipe()
	summaryCmd.Stdout = pw2
	notifierCmd.Stdin = pr2

	notifierCmd.Stdout = os.Stdout
	scannerCmd.Stderr = os.Stderr
	summaryCmd.Stderr = os.Stderr
	notifierCmd.Stderr = os.Stderr

	if err := scannerCmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar scanner: %v", err)
	}
	if err := summaryCmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar summary: %v", err)
	}
	if err := notifierCmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar notifier: %v", err)
	}

	errScanner := scannerCmd.Wait()
	pw1.Close()
	if errScanner != nil {
		return fmt.Errorf("scanner error: %v", errScanner)
	}

	errSummary := summaryCmd.Wait()
	pw2.Close()
	if errSummary != nil {
		return fmt.Errorf("summary error: %v", errSummary)
	}

	if err := notifierCmd.Wait(); err != nil {
		return fmt.Errorf("notifier error: %v", err)
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
