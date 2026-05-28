---
não iniciado-tema: Diferenciação entre SSE, Long Polling e WS - Teoria e Fundamentos
revision_29-05-2026: false
revision_30-05-2026: false
revision_04-06-2026: false
revision_27-06-2026: false
revision_26-08-2026: false
revision_25-10-2026: false
references:
  - 
homework:
  - [[homeworks/Diferenciação entre SSE, Long Polling e WS - Teoria e Fundamentos-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)
Explicar Diferenciação entre SSE, Long Polling e WS - Teoria e Fundamentos em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* Qual a diferença fundamental na direcionalidade entre Server-Sent Events (SSE) e WebSockets?
* Como o Long Polling simula tempo real e qual é o seu principal gargalo de recursos no servidor?
* Como funciona o handshake de "Upgrade" no início de uma conexão WebSocket?
* Por que SSE pode sofrer gargalos ao ser utilizado sobre HTTP/1.1 em navegadores modernos?

---
## 🔎 Perguntas de Elaboração

* O que é polling, long-polling, qual a diferença entre as duas e em quais contextos faz sentido aplica-las
* Por que você escolheria SSE ao invés de WebSockets para implementar um feed ou dashboard financeiro?
* Em que cenários extremos ou restritivos o Long Polling ainda seria justificável no desenvolvimento web atual?
* Explique detalhadamente o custo oculto de manter milhares de threads bloqueadas (I/O bloqueante) aguardando respostas no modelo de Long Polling.
---
## 🔗 Conexões

Relacionar com outros tópicos:
* [[Protocolos TCP e UDP]]
* [[Modelos de I/O Bloqueante e Não-Bloqueante]]
* [[HTTP/2 e Evolução Protocolar]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
