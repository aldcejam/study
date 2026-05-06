#!/usr/bin/env python3
import sys
import json
import hashlib
from datetime import datetime
from typing import List
from models import NoteMetadata, ProcessedNote

def parse_date(date_str: str) -> datetime:
    """
    Converte strings de data (vários formatos) para objeto datetime.
    Formatos suportados: DD/MM/YY, DD/MM/YYYY, YYYY-MM-DD.
    
    Args:
        date_str: String representando a data.
        
    Returns:
        Objeto datetime (hora zerada) ou data longínqua em caso de erro.
    """
    date_str = date_str.strip()
    formats = ["%d/%m/%y", "%d/%m/%Y", "%Y-%m-%d"]
    
    for fmt in formats:
        try:
            return datetime.strptime(date_str, fmt)
        except ValueError:
            continue
            
    # Se falhar em todos, retorna data muito no futuro para não disparar alerta
    return datetime(2099, 1, 1)

def generate_id(path: str, filename: str, revision_date: str) -> str:
    """
    Gera um ID único estável para um item de revisão.
    
    Args:
        path: Caminho relativo do arquivo.
        filename: Nome do arquivo.
        revision_date: Data da revisão que está sendo processada.
        
    Returns:
        Hash MD5 curto representando a identidade única do item.
    """
    unique_str = f"{path}/{filename}@{revision_date}"
    return hashlib.md5(unique_str.encode()).hexdigest()[:12]

def process_notes(notes: List[NoteMetadata]) -> List[ProcessedNote]:
    """
    Analisa as notas e calcula o status de revisão e atraso para cada uma.
    Identifica a revisão pendente mais urgente.
    
    Args:
        notes: Lista de metadados brutos das notas.
        
    Returns:
        Lista de notas processadas com informações de status e IDs de rastreamento.
    """
    today = datetime.now().replace(hour=0, minute=0, second=0, microsecond=0)
    processed = []
    
    for note in notes:
        dias_atraso = 9999
        status = "FUTURA"
        data_referencia = "00/00/00"
        
        # Filtra apenas revisões pendentes
        pendentes = [r for r in note['revisoes'] if r['status'] == ' ']
        
        if not pendentes:
            status = "EM_DIA"
            dias_atraso = 0
        else:
            # Encontra a revisão mais urgente (menor diff em relação a hoje)
            for rev in pendentes:
                date_obj = parse_date(rev['data'])
                diff = (date_obj - today).days
                
                if diff < dias_atraso:
                    dias_atraso = diff
                    data_referencia = rev['data']
            
            if dias_atraso < 0:
                status = "ATRASADA"
            elif dias_atraso == 0:
                status = "HOJE"
            else:
                status = "FUTURA"
            
        processed_note: ProcessedNote = {
            **note,
            "id": generate_id(note['relative_path'], note['filename'], data_referencia),
            "dias_atraso": dias_atraso,
            "status_revisao": status,
            "last_notified_at": None
        }
        processed.append(processed_note)
    return processed

if __name__ == "__main__":
    try:
        input_data = json.load(sys.stdin)
    except Exception as e:
        print(f"Erro ao ler stdin no Estágio 2: {e}", file=sys.stderr)
        sys.exit(1)
        
    results = process_notes(input_data)
    print(json.dumps(results, ensure_ascii=False))
