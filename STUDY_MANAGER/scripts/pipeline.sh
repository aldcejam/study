#!/bin/bash

# 🔄 Pipeline de Gerenciamento de Estudos (Go Edition)
# Compila os binários e executa o pipeline com inspeção de saída por estágio.

set -e

# Garante que estamos no diretório STUDY_MANAGER/
cd "$(dirname "$0")/.."

# Detecta Go em locais comuns se não estiver no PATH
if ! command -v go &>/dev/null; then
    for try_path in /home/apm/go/bin /usr/local/go/bin /snap/go/current/bin; do
        if [ -x "$try_path/go" ]; then
            export PATH="$PATH:$try_path"
            break
        fi
    done
fi

if ! command -v go &>/dev/null; then
    echo "❌ Go não encontrado. Instale em: https://go.dev/dl/"
    exit 1
fi

# Evita o aviso "GOPATH set to GOROOT"
export GOPATH="${GOPATH:-$HOME/go_workspace}"


BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE} STUDY_MANAGER — Pipeline de Estudos (Go)       ${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# 1. Compila todos os binários
echo -e "${BLUE}⚙️  Compilando binários...${NC}"
go build -o bin/scanner  ./cmd/scanner
go build -o bin/processor ./cmd/processor
go build -o bin/exporter  ./cmd/exporter
go build -o bin/notifier  ./cmd/notifier
echo -e "${GREEN}✔  Compilação concluída.${NC}"

# 2. Pasta de rastreio: limpa e recria para novos arquivos de saída
TRACE_DIR="./trace"
rm -rf "$TRACE_DIR"
mkdir -p "$TRACE_DIR"

# 3. Execução da pipeline com salvamento dos outputs intermediários
echo -e "${BLUE}▶  Executando pipeline...${NC}"
./bin/scanner \
    | tee "$TRACE_DIR/01_scanner.json" \
    | ./bin/processor \
    | tee "$TRACE_DIR/02_processor.json" \
    | ./bin/exporter \
    | tee "$TRACE_DIR/03_exporter.txt" \
    | ./bin/notifier

echo ""
echo -e "${GREEN}✨ Processo concluído!${NC}"
echo -e "${BLUE}📄 Arquivos de rastreio gerados em: ${TRACE_DIR}/${NC}"
