#!/bin/bash

# ============================================================================
# Script: extract-tables.sh
# Descrição: Extrai todas as tabelas de arquivos markdown em um diretório
# Uso: ./extract-tables.sh <diretório> [arquivo_saída]
# Exemplo: ./extract-tables.sh . tables.txt
# ============================================================================

set -euo pipefail

# Validar argumentos
if [[ $# -lt 1 ]]; then
    echo "Uso: $0 <diretório> [arquivo_saída]"
    echo ""
    echo "Argumentos:"
    echo "  <diretório>      Diretório contendo arquivos .md"
    echo "  [arquivo_saída]  Arquivo de saída (padrão: tables.txt)"
    echo ""
    echo "Exemplo:"
    echo "  $0 . tables.txt"
    echo "  $0 ./estudos"
    exit 1
fi

DIR="${1:-.}"
OUTPUT="${2:-tables.txt}"

# Validar se diretório existe
if [[ ! -d "$DIR" ]]; then
    echo "❌ Erro: Diretório '$DIR' não encontrado"
    exit 1
fi

# Inicializar arquivo de saída
: > "$OUTPUT"

echo "🔍 Extraindo tabelas de markdown..."
echo "📁 Diretório: $DIR"
echo "📄 Saída: $OUTPUT"
echo ""

# Contador de tabelas encontradas
table_count=0
file_count=0

# Iterar sobre todos os arquivos .md
while IFS= read -r -d '' file; do
    file_count=$((file_count + 1))
    
    # Caminho relativo para display
    rel_path="${file#$DIR/}"
    
    # Extrair tabelas do arquivo
    # Tabelas markdown começam com | e contêm pelo menos um ---
    in_table=false
    current_table=""
    
    while IFS= read -r line; do
        # Detectar início de tabela (linha com |)
        if [[ "$line" =~ \|.*\| ]]; then
            if ! $in_table; then
                in_table=true
                current_table=""
            fi
            current_table+="$line"$'\n'
        else
            # Se estava em tabela e encontrou linha sem |, finalizar
            if $in_table; then
                in_table=false
                # Validar se tem ao menos uma linha de separação
                if [[ "$current_table" =~ \-\-\- ]]; then
                    echo "## 📋 Arquivo: $rel_path" >> "$OUTPUT"
                    echo "$current_table" >> "$OUTPUT"
                    echo "" >> "$OUTPUT"
                    table_count=$((table_count + 1))
                fi
                current_table=""
            fi
        fi
    done < "$file"
    
    # Finalizar tabela se arquivo terminou enquanto estava em tabela
    if $in_table && [[ "$current_table" =~ \-\-\- ]]; then
        echo "## 📋 Arquivo: $rel_path" >> "$OUTPUT"
        echo "$current_table" >> "$OUTPUT"
        echo "" >> "$OUTPUT"
        table_count=$((table_count + 1))
    fi
    
    echo "✅ $rel_path"
    
done < <(find "$DIR" -type f -name "*.md" -print0)

# Relatório final
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📊 RELATÓRIO DE EXTRAÇÃO"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "📁 Arquivos processados: $file_count"
echo "📋 Tabelas extraídas: $table_count"
echo "💾 Arquivo de saída: $OUTPUT"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [[ $table_count -eq 0 ]]; then
    echo "⚠️  Nenhuma tabela encontrada"
    exit 1
fi

echo "✨ Extração concluída com sucesso!"
exit 0
