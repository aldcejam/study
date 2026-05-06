import pytest
import os
import json
import subprocess
import shutil
from datetime import datetime

def test_full_pipeline_flow(tmp_path):
    """
    Testa a integração entre scanner, processor e exporter usando arquivos temporários.
    """
    # 1. Preparar ambiente temporário (Simulando a estrutura real)
    test_vault = tmp_path / "vault"
    test_vault.mkdir()
    manager_dir = test_vault / "__MANAGER__"
    manager_dir.mkdir()
    engine_dir = manager_dir / "study-metadata-generator"
    engine_dir.mkdir()
    output_dir = manager_dir / "output"
    output_dir.mkdir()
    
    # Copia os scripts da pasta geradora real para a pasta de teste
    src_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), '../study-metadata-generator'))
    for script in ["scanner.py", "processor.py", "exporter.py", "models.py"]:
        shutil.copy(os.path.join(src_dir, script), engine_dir)
    
    # Cria uma nota de teste com data bem antiga para garantir status ATRASADA
    note_content = """---
tema: Teste Integração
revision_01-01-2020: false
---
Conteúdo da nota."""
    note_file = test_vault / "nota_teste.md"
    note_file.write_text(note_content, encoding='utf-8')
    
    # 2. Executar Pipeline via Shell (Chamando a pasta correta)
    cmd = (
        f"python3 {engine_dir}/scanner.py | "
        f"python3 {engine_dir}/processor.py | "
        f"python3 {engine_dir}/exporter.py"
    )
    
    result = subprocess.run(cmd, shell=True, capture_output=True, text=True, cwd=engine_dir)
    
    # Debug em caso de erro
    if result.returncode != 0 or "RESUMO DE ESTUDOS" not in result.stdout:
        print(f"Stdout: {result.stdout}")
        print(f"Stderr: {result.stderr}")
    
    # 3. Validar resultados
    assert result.returncode == 0
    assert "RESUMO DE ESTUDOS" in result.stdout
    assert "Teste Integração" in result.stdout
    
    # Verifica se o arquivo de metadados foi criado na pasta output (um nível acima da engine)
    meta_file = output_dir / "revisoes_METADATA.json"
    assert meta_file.exists()
    
    with open(meta_file, 'r', encoding='utf-8') as f:
        meta_data = json.load(f)
        assert len(meta_data) == 1
        assert meta_data[0]['tema'] == "Teste Integração"
        assert meta_data[0]['last_notified_at'] is not None
