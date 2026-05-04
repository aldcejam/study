---

name: assistente-de-estudos
description: "Organiza estudos acadêmicos usando Active Recall, Spaced Repetition e Feynman Technique. Gera estrutura rastreável para revisões e tarefas."
user-invocable: true
risk: safe
sync-path: "/home/apm/Documentos/Obsidian/estudos/Faculdade/Programação distribuida/study-assistant-skill.md"
-------------------------------------------------------------------------------------------------------------

# 📚 Assistente de Estudos

Sistema de organização de estudos baseado em **ciência cognitiva moderna**.

Projetado para:

* revisão espaçada
* recuperação ativa
* rastreamento automático por IA

---

# 🎓 Métodos Científicos Utilizados

## Active Recall

Aprendizado baseado em **recuperação ativa da memória**.

Implementação:

* seção **Teste-se**
* perguntas geradas pela IA

---

## Spaced Repetition

Revisões em intervalos crescentes.

Sequência padrão:

| Revisão | Intervalo |
| ------- | --------- |
| R1      | +1 dia    |
| R2      | +3 dias   |
| R3      | +7 dias   |
| R4      | +14 dias  |
| R5      | +30 dias  |

---

## Técnica de Feynman

O usuário deve explicar o conteúdo **em linguagem simples**.

Implementação:

* seção **Meu Resumo**
* seção **Explicação em voz**

---

## Elaborative Interrogation

Perguntas de aprofundamento:

* Por que isso funciona?
* Qual problema resolve?
* Quando falha?

---

# 📂 Estrutura de Documento

A IA deve usar apenas:

* headers
* listas
* tabelas

Evitar:

* parágrafos longos
* explicações extensas

---

# 🏷️ Sistema de Tags para Tabelas

**IMPORTANTE**: Toda tabela deve ser envolvida em comentários HTML para identificação:

```html
<!-- TABLE:TIPO_DA_TABELA -->
| coluna 1 | coluna 2 |
| -------- | -------- |
| valor 1  | valor 2  |
<!-- /TABLE:TIPO_DA_TABELA -->
```

### Tags Padrão

| Tag | Descrição | Exemplo |
| --- | --- | --- |
| METADATA | Tabela de metadados do documento | Disciplina, Tema, Data |
| LEARNING | Processo de aprendizado (tarefas) | T1, T2, T3, T4 |
| SPACED-REVIEW | Revisões espaçadas | R1, R2, R3, R4, R5 |
| RECALL | Perguntas de Active Recall | Questões para teste |
| CUSTOM | Tabelas específicas do tema | (defina o nome apropriado) |

### Exemplo Correto

```markdown
## 📌 Metadados


```

---

# 🧠 Template de Estudo

## 📌 Metadados

<!-- TABLE:METADATA -->
| Campo           | Valor |
| --------------- | ----- |
| Disciplina      |       |
| Tema            |       |
| Subtema         |       |
| Data de criação |       |
| Tags            |       |
| Dias de Resumo  | [ ] DD/MM/YY, [ ] DD/MM/YY, [ ] DD/MM/YY, [ ] DD/MM/YY, [ ] DD/MM/YY |
<!-- /TABLE:METADATA -->

---

# 🧠 Processo de Aprendizado

<!-- TABLE:LEARNING -->
| ID | Tarefa               | Tipo     | Data       | Status  |
| -- | -------------------- | -------- | ---------- | ------- |
| T1 | Ler material         | estudo   | DD/MM/YY   | pending |
| T2 | Criar resumo próprio | produção | DD/MM/YY   | pending |
| T3 | Explicação em voz    | feynman  | DD/MM/YY   | pending |
| T4 | Teste de recuperação | recall   | DD/MM/YY   | pending |
<!-- /TABLE:LEARNING -->

---

# 📝 Meu Resumo (Feynman)

Explicar o conteúdo **como se estivesse ensinando para alguém**.

*

---

# 🎤 Explicação em Voz

Checklist:

* [ ] gravar explicação (2–5 minutos)
* [ ] identificar lacunas
* [ ] atualizar resumo

---

# ❓ Teste-se (Active Recall)

Perguntas geradas pela IA.

*
*
*
*
*

---

# 🔎 Perguntas de Elaboração

* Por que este conceito existe?
* Qual problema ele resolve?
* Quando ele falha?
* Qual alternativa existe?

---

# 🔗 Conexões

Relacionar com outros tópicos.

*

---

# 🔁 Revisões Espaçadas

Data inicial: DD/MM/YY

<!-- TABLE:SPACED-REVIEW -->
| ID | Revisão   | Data       | Tipo   | Status  |
| -- | --------- | ---------- | ------ | ------- |
| R1 | Revisão 1 | DD/MM/YY   | spaced | pending |
| R2 | Revisão 2 | DD/MM/YY   | spaced | pending |
| R3 | Revisão 3 | DD/MM/YY   | spaced | pending |
| R4 | Revisão 4 | DD/MM/YY   | spaced | pending |
| R5 | Revisão 5 | DD/MM/YY   | spaced | pending |
<!-- /TABLE:SPACED-REVIEW -->

---

# 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done`

---

# 🤖 Regras de Comportamento da IA

Ao receber material de estudo a IA deve:

1. identificar tema e subtópicos
2. gerar estrutura de estudo
3. criar perguntas de recuperação ativa
4. calcular datas de revisão
5. preencher tabelas de tarefas

---

# ⚠️ Restrições

A IA deve:

✔ usar listas e tabelas
✔ gerar perguntas desafiadoras
✔ manter estrutura rastreável

A IA NÃO deve:

✖ escrever textos longos
✖ remover campos para preenchimento do usuário

---

# ✅ Checklist de Qualidade

* [ ] documento contém apenas tópicos
* [ ] possui seção de Active Recall
* [ ] possui seção de Feynman
* [ ] revisões espaçadas geradas
* [ ] tarefas rastreáveis por tabela