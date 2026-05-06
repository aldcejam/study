#!/bin/bash
# 🤖 Script de teste para envio de mensagens via Telegram

# Carrega variáveis do arquivo .env
if [ -f "$(dirname "$0")/../.env" ]; then
    export $(grep -v '^#' "$(dirname "$0")/../.env" | xargs)
fi

TOKEN=$TELEGRAM_TOKEN
MESSAGE="Olá Mundo! 📚 Este é um teste da Pipeline de Estudos."

if [ -z "$TOKEN" ]; then
    echo "❌ Erro: TELEGRAM_TOKEN não encontrado no arquivo .env"
    exit 1
fi

# Se o CHAT_ID não for passado como argumento, tenta usar o do .env ou buscar no bot
CHAT_ID=$1
if [ -z "$CHAT_ID" ]; then
    CHAT_ID=$TELEGRAM_CHAT_ID
fi

if [ -z "$CHAT_ID" ]; then
    echo "🔍 Buscando Chat ID (certifique-se de ter enviado um 'oi' para o bot)..."
    CHAT_ID=$(curl -s "https://api.telegram.org/bot$TOKEN/getUpdates" | python3 -c "import sys, json; data=json.load(sys.stdin); print(data['result'][-1]['message']['chat']['id']) if data['result'] else print('')")
fi

if [ -z "$CHAT_ID" ]; then
    echo "❌ Erro: Não foi possível encontrar um Chat ID. Envie uma mensagem para o bot e tente novamente."
    exit 1
fi

echo "🚀 Enviando mensagem para o Chat ID: $CHAT_ID..."

RESPONSE=$(curl -s -X POST "https://api.telegram.org/bot$TOKEN/sendMessage" \
    -d "chat_id=$CHAT_ID" \
    -d "text=$MESSAGE" \
    -d "parse_mode=Markdown")

if echo "$RESPONSE" | grep -q '"ok":true'; then
    echo "✨ Mensagem enviada com sucesso!"
    exit 0
else
    echo "❌ Erro ao enviar mensagem:"
    echo "$RESPONSE"
    exit 1
fi
