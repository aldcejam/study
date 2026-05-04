# 📚 Study Assistant Pipeline Manager

Sistema automatizado de 3 estágios para extrair, processar e exportar metadados de documentos de estudo.

---

## 🎯 Objetivo

Criar um pipeline que:
1. **Extrai** tabelas marcadas de arquivos markdown
2. **Processa** metadados e identifica revisões pendentes
3. **Exporta** dados estruturados em JSON para integração com outras ferramentas

---

## 📂 Conteúdo

```
__MANAGER__/
├── 📄 extract.sh              # Estágio 1: Extrai tabelas por tag
├── 📄 process-metadata.sh     # Estágio 2: Processa e classifica datas
├── 📄 export-json.sh          # Estágio 3: Exporta para JSON
├── 📄 pipeline.sh             # Wrapper: Executa tudo em sequência
├── 📂 output/                 # Pasta com arquivos gerados
├── 📄 PIPELINE.md             # Documentação detalhada do pipeline
├── 📄 README.md               # Este arquivo
└── 📄 .gitignore              # Ignora arquivos gerados
```

---

## 🚀 Quick Start

### 1. Extrair Metadados

```bash
./extract.sh .. METADATA
# Gera: output/tabelas_METADATA.csv
```

### 2. Processar Datas

```bash
./process-metadata.sh tabelas_METADATA.csv
# Mostra: Revisões pendentes e atrasadas
```

### 3. Exportar JSON

```bash
./export-json.sh tabelas_METADATA.csv
# Gera: output/revisoes.json
```

### Tudo em Uma Vez

```bash
./pipeline.sh ..
# Executa os 3 estágios automaticamente
```

---

## 📋 Scripts Disponíveis

### `extract.sh` - Estágio 1: Extração

Extrai tabelas markdown marcadas com tags HTML.

**Uso:**
```bash
./extract.sh <diretório> <tag> [nome_saída]
```

**Tags Suportadas:**
- `METADATA` - Metadados do documento
- `LEARNING` - Processo de aprendizado
- `SPACED-REVIEW` - Revisões espaçadas
- `RECALL` - Perguntas Active Recall
- `CUSTOM` - Tags personalizadas

**Exemplos:**
```bash
./extract.sh .. METADATA
./extract.sh ../Programação\ distribuida LEARNING
./extract.sh .. SPACED-REVIEW reviews.csv
```

---

### `process-metadata.sh` - Estágio 2: Processamento

Processa CSV e identifica revisões pendentes/atrasadas.

**Uso:**
```bash
./process-metadata.sh <arquivo_csv>
```

**Lógica:**
- Extrai datas do campo `Dias de Resumo`
- Calcula dias até revisão
- Classifica como "prevista" ou "atrasada"
- Gera mensagens formatadas

**Classificação:**
- ✅ **Prevista**: -2 a 0 dias (dentro do próximo 2 dias ou atrasada até 2 dias)
- ❌ **Atrasada**: < -2 dias (mais de 2 dias no passado)

**Exemplo:**
```bash
./process-metadata.sh tabelas_METADATA.csv
```

---

### `export-json.sh` - Estágio 3: Exportação

Converte CSV para JSON estruturado.

**Uso:**
```bash
./export-json.sh <arquivo_csv> [nome_saída]
```

**Formato de Saída:**
```json
{
  "data_geracao": "2026-03-15T21:33:45.123456",
  "data_atual": "15/03/26",
  "total_topicos": 1,
  "revisoes": [
    {
      "tema": "Arquitetura",
      "subtema": "Client-Server",
      "revisoes": {
        "prevista": [...],
        "atrasada": [],
        "futura": [...]
      }
    }
  ]
}
```

**Exemplo:**
```bash
./export-json.sh tabelas_METADATA.csv
./export-json.sh tabelas_METADATA.csv meusdados.json
```

---

### `pipeline.sh` - Wrapper Automatizado

Executa os 3 estágios em sequência.

**Uso:**
```bash
./pipeline.sh <diretório> [tag] [nome_saída]
```

**O que faz:**
1. ✅ Extrai tabelas do diretório
2. ✅ Processa metadados
3. ✅ Mostra revisões pendentes

**Exemplo:**
```bash
./pipeline.sh ..
./pipeline.sh .. LEARNING
```

---

## 📊 Exemplos de Fluxo

### Cenário 1: Revisar estudos de Programação Distribuída

```bash
# Ir para o diretório
cd estudos/__MANAGER__

# Executar pipeline completo
./pipeline.sh ..

# Resultado: Mostra todos os tópicos com revisões pendentes
```

### Cenário 2: Exportar dados para integração

```bash
# 1. Extrair
./extract.sh .. METADATA

# 2. Exportar para JSON
./export-json.sh tabelas_METADATA.csv revisoes_todas.json

# 3. Usar o JSON em outra aplicação
cat output/revisoes_todas.json | jq '.revisoes[0]'
```

### Cenário 3: Revisar apenas Learning

```bash
# Extrair LEARNING (aprendizado)
./extract.sh .. LEARNING dados_learning.csv

# Processar
./process-metadata.sh dados_learning.csv

# Exportar
./export-json.sh dados_learning.csv learning.json
```

---

## 🏷️ Sistema de Tags

Os scripts usam tags HTML para identificar tabelas no markdown:

```markdown
<!-- TABLE:METADATA -->
| Campo | Valor |
| Tema | Meu Tópico |
<!-- /TABLE:METADATA -->
```

**Marcação automática:**
- Tags são **case-insensitive**
- Espaços são ignorados
- Suporta qualquer nome de tag

---

## 📈 Fluxo de Dados

```
Arquivos Markdown
       ↓
[Estágio 1: extract.sh]
       ↓
CSV com tabelas
       ↓
[Estágio 2: process-metadata.sh]
       ↓
Mensagens formatadas (Terminal)
       ↓
[Estágio 3: export-json.sh]
       ↓
JSON estruturado (output/)
```

---

## 🎓 Formato de Metadados

Os scripts esperan campos específicos:

```csv
Campo,Valor
Disciplina,Programação Distribuída
Tema,Arquitetura Client-Server
Subtema,Stateful vs Stateless
Data de criação,15/03/26
Tags,#tag1 #tag2
Dias de Resumo,[x] 16/03/26, [ ] 18/03/26, [ ] 22/03/26, [ ] 29/03/26
```

**Campo Chave:** `Dias de Resumo`
- Formato: `[status] DD/MM/YY`
- Status: `x` = completada, ` ` = pendente
- Múltiplas datas separadas por vírgula

---

## ⚙️ Configuração

### Permissões

```bash
chmod +x *.sh
```

### Dependências

- `bash` - Estágio 1
- `awk` - Estágio 1
- `python3` - Estágios 2-3
- Módulos Python (built-in): `csv`, `re`, `datetime`, `sys`, `json`, `os`

---

## 📁 Arquivos Gerados

Todos os arquivos gerados ficam em `output/`:

```bash
output/
├── tabelas_METADATA.csv        # Dados extraídos (Estágio 1)
├── tabelas_LEARNING.csv
├── tabelas_SPACED-REVIEW.csv
├── revisoes.json               # Exportado (Estágio 3)
└── ...
```

Para limpar:
```bash
rm -rf output/*
```

---

## 🔍 Troubleshooting

### Erro: "Arquivo não encontrado"

Os scripts procuram automaticamente em `output/` se você passar apenas o nome:
```bash
./process-metadata.sh tabelas_METADATA.csv
# Procura em: output/tabelas_METADATA.csv
```

### Erro: "Nenhuma tabela encontrada"

Verifique que o markdown tem as tags corretas:
```markdown
<!-- TABLE:METADATA -->
... dados ...
<!-- /TABLE:METADATA -->
```

### Saída vazia no JSON

Verifique se o CSV tem o campo `Dias de Resumo` preenchido com datas.

---

## 📚 Referências

- [PIPELINE.md](./PIPELINE.md) - Documentação técnica detalhada
- [study-assistant-skill.md](../Programação\ distribuida/study-assistant-skill.md) - Framework de estudo

---

## 🎯 Próximas Funcionalidades

- [ ] Integração com calendário (Google Calendar, iCal)
- [ ] Notificações por email
- [ ] Dashboard web
- [ ] Slack/Discord integration
- [ ] Exportar para Anki
- [ ] Sincronização com banco de dados

---

## 📝 Notas

- Todos os scripts são independentes e podem ser executados isoladamente
- Os arquivos gerados são idempotentes (rodar novamente não duplica dados)
- O sistema foi projetado para funcionar com markdown do Obsidian

---

## 📄 Licença

Sistema criado para estudo pessoal | 2026

