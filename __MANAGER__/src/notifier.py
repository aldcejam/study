#!/usr/bin/env python3
import sys
import os
import subprocess

def load_env():
    """
    Carrega variáveis do arquivo .env manualmente para evitar dependências extras.
    Busca na raiz do projeto ou no diretório do script.
    """
    script_dir = os.path.dirname(os.path.abspath(__file__))
    # Tenta encontrar o .env na raiz (estudos/) ou no __MANAGER__/
    env_paths = [
        os.path.join(script_dir, "../../.env"),
        os.path.join(script_dir, "../.env"),
        ".env"
    ]
    
    for path in env_paths:
        if os.path.exists(path):
            try:
                with open(path, 'r') as f:
                    for line in f:
                        line = line.strip()
                        if line and not line.startswith("#"):
                            key, value = line.split("=", 1)
                            os.environ[key.strip()] = value.strip()
                break
            except Exception:
                continue

# Carrega o .env antes de definir as constantes
load_env()

# Agora usamos as chaves do Telegram como padrão
TOKEN = os.getenv("TELEGRAM_TOKEN")
CHAT_ID = os.getenv("TELEGRAM_CHAT_ID")

def send():
    """
    Lê a mensagem do stdin e envia para o Telegram via script shell.
    """
    mensagem = sys.stdin.read().strip()
    
    if mensagem == "SEM_REVISOES" or not mensagem:
        print("✅ Nenhuma revisão pendente. Notificação não enviada.")
        return

    if mensagem == "SEM_ALTERACOES":
        print("ℹ️ Cache detectado: Nenhuma alteração relevante desde o último alerta.")
        return

    if not TOKEN or not CHAT_ID:
        print(f"⚠️ Erro: Credenciais do Telegram não encontradas no .env")
        return

    script_dir = os.path.dirname(os.path.abspath(__file__))
    send_script = os.path.join(script_dir, "../scripts/send-message.sh")
    
    if os.path.exists(send_script):
        print(f"🚀 Enviando alerta para o Telegram...")
        # Note que agora passamos TOKEN e CHAT_ID para o script
        result = subprocess.run(["bash", send_script, TOKEN, CHAT_ID, mensagem])
        if result.returncode == 0:
            print("✨ Sucesso!")
        else:
            print("❌ Falha no envio via send-message.sh.")
    else:
        print(f"❌ Erro crítico: Script '{send_script}' not found.")

if __name__ == "__main__":
    send()
