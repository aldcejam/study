#!/bin/bash

# 🔄 Pipeline de Gerenciamento de Estudos
# Executa todos os estágios em sequência usando pipes

# Garante que estamos no diretório correto
cd "$(dirname "$0")"

# Cores
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE} iniciando Pipeline de Estudos...${NC}"
echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

# Execução em Pipe
python3 stage1_scanner.py | \
python3 stage2_processor.py | \
python3 stage3_exporter.py | \
python3 stage4_notifier.py

echo ""
echo "✨ Processo concluído!"
