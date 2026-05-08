---
name: study-manager-go
description: "Gerencia o pipeline de automação de estudos em Go. Responsável por varrer o vault, processar revisões atrasadas e notificar via Telegram."
user-invocable: true
risk: safe
---

# 🚀 Study Manager Go Skill (Local)

Use esta skill para orientar o desenvolvimento e manutenção do projeto **STUDY_MANAGER**.

## 🏗️ Arquitetura do Pipeline

O sistema opera em uma cadeia de pipes Unix, onde cada estágio é um binário Go independente:

1.  **Scanner** (`cmd/scanner`): Varre o vault Obsidian.
    - **Regra Crítica**: Só processa arquivos onde o primeiro atributo do frontmatter é `tema:`.
    - Extrai revisões no formato `revision_DD-MM-YYYY`.
    - Obtém `updatedAt` via Git ou filesystem.
2.  **Processor** (`cmd/processor`): Calcula o status das revisões.
    - Gera IDs únicos (MD5 truncado).
    - Calcula `dias_atraso` e define status (`ATRASADA`, `HOJE`, `EM_DIA`).
3.  **Exporter** (`cmd/exporter`): Gerencia o estado e formata a saída.
    - Mantém o cache em `output/revisoes_METADATA.json`.
    - Implementa lógica de anti-spam (notifica apenas se houver mudanças ou após 2 dias).
4.  **Notifier** (`cmd/notifier`): Despachante final.
    - Lê o `.env` (Telegram Token/Chat ID).
    - Envia a mensagem formatada via HTTP (API do Telegram).

## 🛠️ Comandos Essenciais

- **Execução Total**: `bash scripts/pipeline.sh` (Compila e roda tudo).
- **Rastreio**: Os arquivos gerados em cada estágio ficam em `./trace/` (01_scanner.json, 02_processor.json, 03_exporter.txt). Servem para identificar o que foi gerado em cada etapa do processo.
- **Compilação Manual**: `go build -o bin/ ./cmd/...`

## ⚠️ Regras de Manutenção (LEIA SEMPRE)

1.  **Zero Dependências**: O projeto deve usar APENAS a biblioteca padrão do Go (`stdlib`). Não adicione `go.mod` externo a menos que estritamente necessário.
2.  **Transparência de Dados**: Qualquer alteração no fluxo de dados entre binários deve ser refletida nas structs em `internal/models/models.go`.
3.  **Rastreio de Processo**: Mantenha o salvamento dos arquivos intermediários no `pipeline.sh` (via `tee`) para que seja possível identificar o que foi gerado em cada etapa.
4.  **Auto-Atualização**: **IMPORTANTE!** Sempre que você alterar a lógica do pipeline, a estrutura de pastas ou as regras de negócio (ex: filtros do scanner), você DEVE atualizar este arquivo `SKILL.md` local para refletir a nova realidade.

## ✅ Checklist de Alteração

- [ ] Os binários continuam se comunicando via stdin/stdout?
- [ ] O parser de frontmatter em `internal/frontmatter` suporta a nova mudança?
- [ ] O cache `revisoes_METADATA.json` permanece compatível?
- [ ] **A SKILL local foi atualizada com as novas informações?**
