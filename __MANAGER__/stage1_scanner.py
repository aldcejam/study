#!/usr/bin/env python3
import os
import re
import json
import sys
import yaml
from typing import List
from models import NoteMetadata, Revision, Reference

def extract_yaml(content: str) -> dict:
    """Extrai YAML Frontmatter usando a biblioteca PyYAML."""
    yaml_match = re.search(r'^---\s*\n(.*?)\n---', content, re.DOTALL | re.MULTILINE)
    if not yaml_match:
        return {}
    
    try:
        return yaml.safe_load(yaml_match.group(1))
    except Exception as e:
        print(f"Erro ao parsear YAML: {e}", file=sys.stderr)
        return {}

def scan_notes(base_dir: str) -> List[NoteMetadata]:
    notes = []
    for root, dirs, files in os.walk(base_dir):
        # Ignora pastas de gerenciamento e ocultas
        if "__MANAGER__" in root or ".obsidian" in root or ".git" in root or ".agents" in root:
            continue
        
        for file in files:
            if file.endswith(".md"):
                full_path = os.path.join(root, file)
                # Calcula a pasta relativa a partir de 'estudos'
                rel_path = os.path.relpath(root, base_dir)
                
                try:
                    with open(full_path, 'r', encoding='utf-8') as f:
                        content = f.read()
                except Exception:
                    continue
                    
                metadata = extract_yaml(content)
                
                # Pega as revisões do YAML
                revisoes = metadata.get('revisoes', [])
                
                # Se revisões vieram como string no YAML (separadas por vírgula)
                if isinstance(revisoes, str):
                    revisoes = [{"status": " ", "data": d.strip()} for d in revisoes.split(',')]

                note: NoteMetadata = {
                    "filename": file,
                    "relative_path": rel_path,
                    "tema": metadata.get('tema', metadata.get('title', file.replace('.md', ''))),
                    "subtema": metadata.get('subtema'),
                    "revisoes": revisoes,
                    "tags": metadata.get('tags', []),
                    "references": metadata.get('references', [])
                }
                notes.append(note)
    return notes

if __name__ == "__main__":
    # O diretório base de estudos é um nível acima do __MANAGER__
    manager_dir = os.path.dirname(os.path.abspath(__file__))
    estudos_dir = os.path.abspath(os.path.join(manager_dir, ".."))
    
    results = scan_notes(estudos_dir)
    # Output JSON para o próximo estágio
    print(json.dumps(results, ensure_ascii=False))
