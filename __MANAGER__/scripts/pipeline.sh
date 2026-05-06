#!/bin/bash

# 🔄 Pipeline de Gerenciamento de Estudos
# Executa todos os estágios em sequência usando pipes

# Garante que estamos no diretório correto (__MANAGER__)
cd "$(dirname "$0")"

# Cores
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE} iniciando Pipeline de Estudos (Engine V2)...${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# Execução em Pipe chamando os scripts na pasta 'src'
python3 ../src/scanner.py | \
python3 ../src/processor.py | \
python3 ../src/exporter.py | \
python3 ../src/notifier.py

echo ""
echo "✨ Processo concluído!"
