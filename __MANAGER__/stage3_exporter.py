#!/usr/bin/env python3
import sys
import json
import os
from typing import List
from models import ProcessedNote

def export_and_format(notes: List[ProcessedNote]):
    # 1. Salvar JSON estruturado no diretório de output
    script_dir = os.path.dirname(os.path.abspath(__file__))
    output_dir = os.path.join(script_dir, "output")
    os.makedirs(output_dir, exist_ok=True)
    
    output_file = os.path.join(output_dir, "revisoes_METADATA.json")
    with open(output_file, "w", encoding='utf-8') as f:
        json.dump(notes, f, indent=2, ensure_ascii=False)

    # 2. Gerar mensagem para WhatsApp apenas para o que for urgente
    urgentes = [n for n in notes if n['status_revisao'] in ["ATRASADA", "HOJE"]]
    
    if not urgentes:
        print("SEM_REVISOES")
        return

    # Ordenar por atraso (mais atrasadas primeiro)
    urgentes.sort(key=lambda x: x['dias_atraso'])

    msg = "📚 *RESUMO DE ESTUDOS*\n"
    msg += "──────────────────\n\n"
    
    for note in urgentes:
        icon = "🚨" if note['status_revisao'] == "ATRASADA" else "📅"
        atraso_str = f"{abs(note['dias_atraso'])}d atrás" if note['status_revisao'] == "ATRASADA" else "HOJE"
        
        # Pega a última pasta do caminho relativo
        contexto = note['relative_path'].split(os.sep)[-1]
        
        msg += f"{icon} *{note['tema']}*\n"
        msg += f"   └ 📂 {contexto} | ⏳ {atraso_str}\n"
        
        if note.get('references'):
            for ref in note['references']:
                desc = ref.get('description', 'Ref')
                source = ref.get('source', '')
                msg += f"   🔗 _{desc}: {source}_\n"
            msg += "\n"
        else:
            msg += "\n"
    
    msg += "──────────────────\n"
    msg += "_Abra seu Obsidian para revisar!_"
    
    # Output da mensagem para o estágio de notificação
    print(msg)

if __name__ == "__main__":
    try:
        input_data = json.load(sys.stdin)
    except Exception as e:
        print(f"Erro ao ler stdin no Estágio 3: {e}", file=sys.stderr)
        sys.exit(1)
        
    export_and_format(input_data)
