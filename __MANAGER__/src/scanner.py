#!/usr/bin/env python3
import os
import re
import json
import sys
import yaml
import subprocess
from datetime import datetime, date
from typing import List
from models import NoteMetadata

def extract_yaml(content: str) -> dict:
    """
    Extrai o YAML Frontmatter do conteúdo de uma nota Obsidian.
    Garante que o retorno seja sempre um dicionário.
    """
    yaml_match = re.search(r'\A---\s*\n(.*?)\n---', content, re.DOTALL)
    if not yaml_match:
        return {}
    
    try:
        data = yaml.safe_load(yaml_match.group(1))
        if isinstance(data, dict):
            return data
        return {}
    except Exception as e:
        # Logs de erro silenciosos para o stderr para não quebrar o pipe JSON
        print(f"Erro ao parsear YAML: {e}", file=sys.stderr)
        return {}

def get_git_updated_at(filepath: str) -> str:
    """
    Recupera a data da última alteração do arquivo registrada no Git.
    """
    try:
        cmd = ["git", "log", "-1", "--format=%cI", "--", filepath]
        result = subprocess.check_output(cmd, stderr=subprocess.DEVNULL).decode().strip()
        if result:
            return result
    except Exception:
        # Se não estiver em um repo git ou git falhar, não usamos datetime.now()
        # pois em CI o checkout reseta o mtime, invalidando o cache.
        pass
    
    # Fallback para mtime apenas se necessário, mas com cautela
    try:
        mtime = os.path.getmtime(filepath)
        return datetime.fromtimestamp(mtime).isoformat()
    except Exception:
        # Se tudo falhar, retorna uma data muito antiga (epoch)
        # Isso garante que a regra 'updated_at > last_notified' não dispare sem motivo
        return "1970-01-01T00:00:00Z"

def scan_notes(base_dir: str) -> List[NoteMetadata]:
    """
    Varre o diretório base em busca de arquivos Markdown e extrai seus metadados.
    """
    notes = []
    for root, dirs, files in os.walk(base_dir):
        if any(ignored in root for ignored in ["__MANAGER__", ".obsidian", ".git", ".agents", ".gemini"]):
            continue
        
        for file in files:
            if file.endswith(".md"):
                full_path = os.path.join(root, file)
                rel_path = os.path.relpath(root, base_dir)
                
                try:
                    with open(full_path, 'r', encoding='utf-8') as f:
                        content = f.read()
                except Exception:
                    continue
                    
                metadata = extract_yaml(content)
                if not isinstance(metadata, dict):
                    metadata = {}
                
                revisoes = []
                
                # Suporte apenas ao formato de propriedades dinâmicas: revision_DD-MM-YYYY
                for key, value in metadata.items():
                    if key.startswith("revision_"):
                        date_part = key.replace("revision_", "").replace("-", "/")
                        # Status: 'x' se for verdadeiro (concluído), ' ' se for falso (pendente)
                        status = 'x' if value is True else ' '
                        revisoes.append({"status": status, "data": date_part})

                note: NoteMetadata = {
                    "filename": file,
                    "relative_path": rel_path,
                    "tema": str(metadata.get('tema', metadata.get('title', file.replace('.md', '')))),
                    "subtema": metadata.get('subtema'),
                    "revisoes": revisoes,
                    "tags": metadata.get('tags', []),
                    "references": metadata.get('references', []),
                    "activity": metadata.get('activity'),
                    "updatedAt": get_git_updated_at(full_path)
                }
                notes.append(note)
    return notes

if __name__ == "__main__":
    manager_dir = os.path.dirname(os.path.abspath(__file__))
    estudos_dir = os.path.abspath(os.path.join(manager_dir, "../.."))
    
    results = scan_notes(estudos_dir)
    
    def json_serial(obj):
        if isinstance(obj, (datetime, date)):
            return obj.isoformat()
        raise TypeError ("Type %s not serializable" % type(obj))

    print(json.dumps(results, ensure_ascii=False, default=json_serial))
