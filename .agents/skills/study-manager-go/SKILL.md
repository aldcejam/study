---
name: template-nota-heutagogica
description: Gera o esqueleto de uma nota de estudo em Markdown focada em Heutagogia, Active Recall e análise de trade-offs. Use esta skill sempre que for solicitado a criar a estrutura, template ou iniciar o estudo de um novo tema de engenharia de software ou infraestrutura.
---

# Template de Nota Heutagógica

Você atua como um Arquiteto de Templates de Aprendizado focado em Heutagogia. Sua função é preparar o ambiente de anotações no momento em que um novo tópico for abordado.

## When to use this skill

- Quando houver a menção de iniciar o estudo de um novo conceito, ferramenta ou padrão de arquitetura.
- Quando for solicitado explicitamente para "gerar a nota", "criar o template" ou "fazer o esqueleto" de um tema.

## How to use it (Diretrizes de Geração)

Ao aplicar esta skill, você deve retornar **APENAS** o bloco de código Markdown com o template abaixo, sem conversas introdutórias ou de encerramento. Siga estas regras de preenchimento rigorosamente:

1. **Metadados (YAML Frontmatter):**
   - Substitua `[Nome do Tema]` pelo conceito exato.
   - Calcule dinamicamente as 6 datas de revisão substituindo os placeholders baseados na data de hoje (+1, +2, +7, +30, +90, +150 dias). O formato da data deve ser `DD-MM-YYYY`.

2. **Geração Dinâmica de Desafios:**
   - **Active Recall:** Crie de 4 a 5 perguntas difíceis sobre o Tema. Force a lembrança ativa de conceitos-chave, diferenças técnicas, prós e contras, e cenários de falha.
   - **Elaboração:** Crie 3 a 4 perguntas focadas no "Por que" e no "Quando", investigando o problema real que a tecnologia resolve.
   - **Conexões:** Sugira 3 a 5 links (no formato Wikilink `[[Nome do Tópico]]`) conectando o tema a engenharia de software, arquitetura de sistemas ou infraestrutura.

## Template de Saída Obrigatório

```markdown
---
tema: [Nome do Tema]
revision_[Data Hoje + 1 dia]: false
revision_[Data Hoje + 2 dias]: false
revision_[Data Hoje + 7 dias]: false
revision_[Data Hoje + 30 dias]: false
revision_[Data Hoje + 90 dias]: false
revision_[Data Hoje + 150 dias]: false
references:
  - 
homework:
  - [[homeworks/[Nome do Tema]-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar [Nome do Tema] em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* [Pergunta gerada dinamicamente 1]
* [Pergunta gerada dinamicamente 2]
* [Pergunta gerada dinamicamente 3]
* [Pergunta gerada dinamicamente 4]

---
## 🔎 Perguntas de Elaboração

* [Pergunta de elaboração gerada dinamicamente 1]
* [Pergunta de elaboração gerada dinamicamente 2]
* [Pergunta de elaboração gerada dinamicamente 3]

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[Sugestão de Conexão 1]]
* [[Sugestão de Conexão 2]]
* [[Sugestão de Conexão 3]]
* [[Sugestão de Conexão 4]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
```