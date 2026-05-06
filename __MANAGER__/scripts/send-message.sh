#!/bin/bash
# 🤖 Script de envio de mensagens via Telegram (Migrado do WhatsApp)

TOKEN=$1
CHAT_ID=$2
MESSAGE=$3

if [ -z "$TOKEN" ] || [ -z "$CHAT_ID" ] || [ -z "$MESSAGE" ]; then
    echo "Uso: $0 <token> <chat_id> <mensagem>"
    exit 1
fi

# Chamada para a API do Telegram
RESPONSE=$(curl -s -X POST "https://api.telegram.org/bot$TOKEN/sendMessage" \
    -d "chat_id=$CHAT_ID" \
    -d "text=$MESSAGE" \
    -d "parse_mode=Markdown")

if echo "$RESPONSE" | grep -q '"ok":true'; then
    exit 0
else
    echo "Erro na API Telegram: $RESPONSE"
    exit 1
fi
