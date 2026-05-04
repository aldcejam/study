#!/usr/bin/env python3
import sys
import os
import subprocess

# Carregar variáveis de ambiente ou usar placeholders
PHONE = os.getenv("WHATSAPP_PHONE", "SEU_NUMERO_AQUI")
API_KEY = os.getenv("WHATSAPP_API_KEY", "SUA_API_KEY_AQUI")

def send():
    # Lê a mensagem formatada do stdin
    mensagem = sys.stdin.read().strip()
    
    if mensagem == "SEM_REVISOES" or not mensagem:
        print("✅ Nenhuma revisão pendente. Notificação não enviada.")
        return

    if PHONE == "SEU_NUMERO_AQUI" or API_KEY == "SUA_API_KEY_AQUI":
        print("⚠️ Erro: Telefone ou API Key não configurados no script stage4_notifier.py.")
        return

    script_dir = os.path.dirname(os.path.abspath(__file__))
    send_script = os.path.join(script_dir, "send-message.sh")
    
    if os.path.exists(send_script):
        print(f"🚀 Enviando alerta para o WhatsApp ({PHONE})...")
        result = subprocess.run(["bash", send_script, PHONE, API_KEY, mensagem])
        if result.returncode == 0:
            print("✨ Sucesso!")
        else:
            print("❌ Falha no envio via send-message.sh.")
    else:
        print(f"❌ Erro crítico: Script '{send_script}' não encontrado.")

if __name__ == "__main__":
    send()
