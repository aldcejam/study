#!/usr/bin/env python3
import sys
import json
from datetime import datetime
from typing import List
from models import NoteMetadata, ProcessedNote

def parse_date(date_str: str) -> datetime:
    """Converte DD/MM/YY para datetime."""
    try:
        # Tenta formato DD/MM/YY
        return datetime.strptime(date_str.strip(), "%d/%m/%y")
    except ValueError:
        try:
            # Tenta formato YYYY-MM-DD
            return datetime.strptime(date_str.strip(), "%Y-%m-%d")
        except:
            # Se falhar, retorna data muito no futuro para não disparar alerta
            return datetime(2099, 1, 1)

def process_notes(notes: List[NoteMetadata]) -> List[ProcessedNote]:
    today = datetime.now().replace(hour=0, minute=0, second=0, microsecond=0)
    processed = []
    
    for note in notes:
        dias_atraso = 9999
        status = "FUTURA"
        
        # Só processamos notas que tenham revisões pendentes
        pendentes = [r for r in note['revisoes'] if r['status'] == ' ']
        
        if not pendentes:
            status = "EM_DIA"
            dias_atraso = 0
        else:
            for rev in pendentes:
                date_obj = parse_date(rev['data'])
                diff = (date_obj - today).days
                
                # Queremos a data pendente mais próxima ou mais atrasada
                if diff < dias_atraso:
                    dias_atraso = diff
            
            if dias_atraso < 0:
                status = "ATRASADA"
            elif dias_atraso == 0:
                status = "HOJE"
            else:
                status = "FUTURA"
            
        processed_note: ProcessedNote = {
            **note,
            "dias_atraso": dias_atraso,
            "status_revisao": status
        }
        processed.append(processed_note)
    return processed

if __name__ == "__main__":
    # Recebe o JSON do estágio anterior via stdin
    try:
        input_data = json.load(sys.stdin)
    except Exception as e:
        print(f"Erro ao ler stdin no Estágio 2: {e}", file=sys.stderr)
        sys.exit(1)
        
    results = process_notes(input_data)
    # Passa para o próximo estágio
    print(json.dumps(results, ensure_ascii=False))
