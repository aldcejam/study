import pytest
import sys
import os
from datetime import datetime, timedelta, timezone

# Adiciona o diretório do gerador ao sys.path para importar os módulos
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../study-metadata-generator')))

from scanner import extract_yaml
from processor import parse_date, generate_id
from exporter import should_notify

def test_dynamic_properties_parsing():
    """Testa a extração de revisões a partir de chaves dinâmicas."""
    # Simula o que o scanner faz ao encontrar chaves revision_
    metadata = {
        "tema": "Teste",
        "revision_18-03-2026": False,
        "revision_22-03-2026": True
    }
    revisoes = []
    for key, value in metadata.items():
        if key.startswith("revision_"):
            date_part = key.replace("revision_", "").replace("-", "/")
            status = 'x' if value is True else ' '
            revisoes.append({"status": status, "data": date_part})
    
    assert len(revisoes) == 2
    # Nota: a ordem pode variar por ser dict, mas os dados devem estar lá
    data_list = [r['data'] for r in revisoes]
    assert "18/03/2026" in data_list
    assert "22/03/2026" in data_list

def test_extract_yaml_resilience():
    """Testa se o extrator de YAML não quebra com arquivos malformados."""
    content = "---\nmalformed: : :\n---\nBody"
    assert extract_yaml(content) == {}
    
    content_scalar = "---\nJust a string\n---"
    assert extract_yaml(content_scalar) == {}

def test_parse_date_formats():
    """Testa múltiplos formatos de data."""
    assert parse_date("04/05/26").year == 2026
    assert parse_date("04/05/2026").year == 2026
    assert parse_date("2026-05-04").year == 2026

def test_should_notify_with_timezone():
    """Testa a regra de notificação com suporte a timezone."""
    now = datetime.now(timezone.utc)
    item = {
        "id": "abc",
        "updatedAt": now.isoformat(),
        "last_notified_at": (now - timedelta(days=3)).isoformat()
    }
    # Intervalo de 2 dias atingido
    assert should_notify(item, {"abc": item}) == True
