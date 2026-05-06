#!/bin/bash
# 📱 Script de envio de mensagens via WhatsApp (CallMeBot)

PHONE=$1
API_KEY=$2
MESSAGE=$3

if [ -z "$PHONE" ] || [ -z "$API_KEY" ] || [ -z "$MESSAGE" ]; then
    echo "Uso: $0 <telefone> <api_key> <mensagem>"
    exit 1
fi

# URL Encode da mensagem usando Python para garantir compatibilidade
ENCODED_MSG=$(python3 -c "import urllib.parse, sys; print(urllib.parse.quote(sys.stdin.read()))" <<< "$MESSAGE")

# Chamada para a API do CallMeBot
# https://www.callmebot.com/blog/free-api-whatsapp-messages/
RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "https://api.callmebot.com/whatsapp.php?phone=$PHONE&text=$ENCODED_MSG&apikey=$API_KEY")

if [ "$RESPONSE" == "200" ]; then
    exit 0
else
    echo "Erro na API: Código $RESPONSE"
    exit 1
fi
