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

PHONE = os.getenv("WHATSAPP_PHONE", "SEU_NUMERO_AQUI")
API_KEY = os.getenv("WHATSAPP_API_KEY", "SUA_API_KEY_AQUI")

def send():
    """
    Lê a mensagem do stdin e envia para o WhatsApp via script shell.
    """
    mensagem = sys.stdin.read().strip()
    
    if mensagem == "SEM_REVISOES" or not mensagem:
        print("✅ Nenhuma revisão pendente. Notificação não enviada.")
        return

    if mensagem == "SEM_ALTERACOES":
        print("ℹ️ Cache detectado: Nenhuma alteração relevante desde o último alerta.")
        return

    if PHONE == "SEU_NUMERO_AQUI" or API_KEY == "SUA_API_KEY_AQUI":
        print(f"⚠️ Erro: Credenciais não encontradas (PHONE={PHONE}, KEY_SET={'Sim' if API_KEY != 'SUA_API_KEY_AQUI' else 'Não'})")
        return

    script_dir = os.path.dirname(os.path.abspath(__file__))
    send_script = os.path.join(script_dir, "../send-message.sh")
    
    if os.path.exists(send_script):
        print(f"🚀 Enviando alerta para o WhatsApp ({PHONE})...")
        result = subprocess.run(["bash", send_script, PHONE, API_KEY, mensagem])
        if result.returncode == 0:
            print("✨ Sucesso!")
        else:
            print("❌ Falha no envio via send-message.sh.")
    else:
        print(f"❌ Erro crítico: Script '{send_script}' not found.")

if __name__ == "__main__":
    send()
