---
tema: Diferenciação entre SSE, Long Polling e WS - Teoria e Fundamentos
revision_31-05-2026: false
revision_01-06-2026: false
revision_04-06-2026: false
revision_11-06-2026: false
revision_25-06-2026: false
revision_25-07-2026: false
references:
  - 
homework:
  - [[homeworks/Diferenciação entre SSE, Long Polling e WS - Teoria e Fundamentos-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)
- **WebSockets (WS)**:
  - **Direcionalidade**: Bidirecional (full-duplex) simultâneo. Tanto cliente quanto servidor transmitem dados a qualquer momento.
  - **Protocolo**: Inicia via handshake HTTP (`Upgrade: websocket`) e transiciona para o protocolo WebSocket independente sobre TCP (HTTP status `101 Switching Protocols`).
- **Server-Sent Events (SSE)**:
  - **Direcionalidade**: Unidirecional (downlink). Apenas o servidor envia dados após a conexão inicial do cliente.
  - **Protocolo**: Mantém-se sob o protocolo HTTP padrão usando o cabeçalho `Content-Type: text/event-stream` e conexão persistente (`keep-alive`).

---
## 🎤 Explicação em Voz

Checklist:
* [x] Gravar explicação (2–5 minutos)
* [x] Identificar lacunas
* [x] Atualizar resumo

### 🔍 Lacunas e Refinamentos Identificados:
* **WebSocket Handshake**: Detalhar a transição via código HTTP `101 Switching Protocols` e o uso dos cabeçalhos `Upgrade: websocket` e `Connection: Upgrade`.
* **SSE Headers**: Adicionar a necessidade de cabeçalhos como `Content-Type: text/event-stream` e `Cache-Control: no-cache` para evitar buffering ou cache indesejado.
* **HTTP/2 Multiplexing**: Entender que o SSE aproveita nativamente o HTTP/2 (evitando o limite de 6 conexões do HTTP/1.1 no navegador), enquanto o WebSocket exige extensões adicionais (como RFC 8441) para rodar sobre HTTP/2.

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
