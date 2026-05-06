#!/usr/bin/env python3
import sys
import json
import os
from datetime import datetime, timezone
from typing import List, Dict
from models import ProcessedNote

def load_history(output_dir: str) -> Dict[str, dict]:
    """
    Carrega o histórico de notificações do arquivo JSON.
    """
    history_path = os.path.join(output_dir, "revisoes_METADATA.json")
    if os.path.exists(history_path):
        try:
            with open(history_path, 'r', encoding='utf-8') as f:
                data = json.load(f)
                return {item['id']: item for item in data}
        except Exception:
            return {}
    return {}

def should_notify(item: ProcessedNote, history: Dict[str, dict]) -> bool:
    """
    Decide se um item deve ser notificado com base no histórico e updatedAt.
    """
    hist = history.get(item['id'])
    if not hist or not hist.get('last_notified_at'):
        return True
    
    try:
        last_notified = datetime.fromisoformat(hist['last_notified_at'])
        updated_at = datetime.fromisoformat(item['updatedAt'])
        
        if last_notified.tzinfo is None:
            last_notified = last_notified.replace(tzinfo=timezone.utc)
        if updated_at.tzinfo is None:
            updated_at = updated_at.replace(tzinfo=timezone.utc)
        
        if updated_at > last_notified:
            return True
            
        now_aware = datetime.now(timezone.utc)
        delta = now_aware - last_notified
        if delta.days >= 2:
            return True
    except Exception:
        return True
        
    return False

def export_and_format(notes: List[ProcessedNote]):
    """
    Filtra as notas, gera a mensagem e atualiza o histórico.
    Distingue entre 'SEM_REVISOES' (nada urgente) e 'SEM_ALTERACOES' (tudo já notificado).
    """
    script_dir = os.path.dirname(os.path.abspath(__file__))
    output_dir = os.path.join(script_dir, "../output")
    os.makedirs(output_dir, exist_ok=True)
    
    history = load_history(output_dir)
    today_iso = datetime.now(timezone.utc).astimezone().isoformat()
    
    urgentes_candidatas = [n for n in notes if n['status_revisao'] in ["ATRASADA", "HOJE"]]
    
    if not urgentes_candidatas:
        # Nada para revisar hoje
        print("SEM_REVISOES")
        # Mesmo assim, salvamos o estado atual das notas (pode haver mudanças em notas futuras)
        output_file = os.path.join(output_dir, "revisoes_METADATA.json")
        with open(output_file, "w", encoding='utf-8') as f:
            json.dump(notes, f, indent=2, ensure_ascii=False)
        return

    a_notificar = []
    for note in urgentes_candidatas:
        if should_notify(note, history):
            note['last_notified_at'] = today_iso
            a_notificar.append(note)
        else:
            if note['id'] in history:
                note['last_notified_at'] = history[note['id']].get('last_notified_at')

    output_file = os.path.join(output_dir, "revisoes_METADATA.json")
    with open(output_file, "w", encoding='utf-8') as f:
        json.dump(notes, f, indent=2, ensure_ascii=False)

    if not a_notificar:
        # Existem coisas urgentes, mas todas já foram notificadas recentemente
        print("SEM_ALTERACOES")
        return

    a_notificar.sort(key=lambda x: x['dias_atraso'])

    msg = "📚 *RESUMO DE ESTUDOS*\n"
    msg += "──────────────────\n\n"
    
    for note in a_notificar:
        icon = "🚨" if note['status_revisao'] == "ATRASADA" else "📅"
        atraso_str = f"{abs(note['dias_atraso'])}d atrás" if note['status_revisao'] == "ATRASADA" else "HOJE"
        contexto = note['relative_path'].split(os.sep)[-1]
        
        msg += f"{icon} *{note['tema']}*\n"
        msg += f"   └ 📂 {contexto} | ⏳ {atraso_str}\n"
        
        if note.get('activity'):
            msg += f"   📝 *Atividade:* {note['activity']}\n"

        if note.get('references'):
            for ref in note['references']:
                if isinstance(ref, dict):
                    desc = ref.get('description', 'Ref')
                    source = ref.get('source', '')
                    msg += f"   🔗 _{desc}: {source}_\n"
                else:
                    # Se for apenas uma string (formato novo)
                    msg += f"   🔗 _{ref}_\n"
            msg += "\n"
        else:
            msg += "\n"
    
    msg += "──────────────────\n"
    msg += "_Abra seu Obsidian para revisar!_"
    
    print(msg)

if __name__ == "__main__":
    try:
        input_data = json.load(sys.stdin)
    except Exception:
        sys.exit(1)
        
    export_and_format(input_data)
