# 📚 Study Metadata Generator

Sistema automatizado projetado em Python para extrair metadados, gerenciar o estado das notificações e alertar sobre revisões de estudo pendentes diretamente das suas notas no Obsidian.

---

## 🎯 Objetivo

Criar uma pipeline que:
1. **Varre** os arquivos Markdown no seu vault em busca de propriedades de revisão.
2. **Processa** o status das revisões (pendente vs. concluída) cruzando dados temporais usando o cache e o histórico do Git.
3. **Persiste** os dados via JSON para controle de repetição de notificações.
4. **Notifica** ativamente o usuário (ex: WhatsApp via integração API) sobre as revisões que estão previstas ou atrasadas, respeitando uma política de intervalo de 2 dias (cache).

---

## 📂 Arquitetura (Engine V2)

```
__MANAGER__/
├── study-metadata-generator/ # O coração da Pipeline (Scripts em Python)
│   ├── scanner.py            # Varre notas e extrai YAML Frontmatter (propriedades dinâmicas)
│   ├── processor.py          # Analisa pendências, atrasos e datas
│   ├── exporter.py           # Analisa cache, Git e escreve no revisoes_METADATA.json
│   ├── notifier.py           # Aciona scripts externos de comunicação
│   └── models.py             # Tipagens de dados usadas pelo sistema
├── output/                   # Contém o arquivo JSON de persistência de status (cache)
├── tests/                    # Suíte de testes automatizados via Pytest
├── pipeline.sh               # Script orquestrador que roda todos os scripts do generator
├── send-message.sh           # Script de integração usado pelo notifier.py (ex: WhatsApp)
├── EXAMPLES.md               # Guia sobre como estruturar os metadados
└── README.md                 # Este arquivo
```

---

## 🚀 Como Funciona

### Formato Obrigatório nas Notas (YAML Frontmatter)

O sistema escaneia o topo das suas notas Obsidian (`---`) e procura por propriedades que começam exatamente com `revision_` acompanhadas de uma data no formato `DD-MM-YYYY`.

O valor (`true` ou `false`) define o status da revisão:

```yaml
---
tema: Arquitetura Client-Server
subtema: Orquestração de Container
revision_18-04-2026: false
revision_20-04-2026: false
references:
  - "tal_coisa: link"
  - "tal_coisa_la:relative_path"
---
```

*   `false` = Revisão pendente (Ativa o alerta).
*   `true` = Revisão concluída (Ignora o alerta).

### A Lógica do Cache e Notificação

O sistema não fica mandando spam de notificações todos os dias. Ele segue regras estritas implementadas no `exporter.py`:

1.  **Edição Recente (Git):** Se o arquivo sofreu alteração no Git recentemente (a data de `updatedAt` é superior à data de notificação registrada no JSON), ele força um envio **imediato** da notificação, assumindo que algo novo foi adicionado que o usuário deva ver.
2.  **Intervalo de Respiro:** Se a nota continuar com itens atrasados/pendentes e não for editada, o sistema aguardará um período de **2 dias** desde a última vez que notificou sobre este arquivo antes de enviar outra notificação.
3.  **Conclusão:** Quando todas as revisões pendentes (aquelas que já passaram ou estão na data de hoje) receberem o status `true`, o sistema para de enviar notificações referentes a essa nota.

---

## ⚙️ Executando

### Execução Completa (Manual ou via Cron/GitHub Actions)

Apenas rode o script da pipeline:

```bash
cd estudos/__MANAGER__
./pipeline.sh
```

A saída exibirá todas as revisões pendentes agrupadas e formatadas, acionando o `send-message.sh` no final.

### Validando os Testes

Para garantir que toda a lógica de negócio esteja funcionando, instale o `pytest` e rode na raiz do gerenciador:

```bash
cd estudos/__MANAGER__
python3 -m venv venv
source venv/bin/activate
pip install pytest pyyaml
export PYTHONPATH=$PYTHONPATH:$(pwd)/study-metadata-generator
pytest tests/
```

---

## 📝 Notas
- O sistema foi projetado para funcionar perfeitamente com propriedades nativas do Obsidian.
- Não é mais necessário usar campos longos como `Dias de Resumo` ou listas YAML; cada data é sua própria chave (`revision_DD-MM-YYYY`).
