---
não iniciado-tema: Multiplexação C10K via epoll e Estrutura dos Frames Binários
revision_29-05-2026: false
revision_30-05-2026: false
revision_04-06-2026: false
revision_27-06-2026: false
revision_26-08-2026: false
revision_25-10-2026: false
references:
  - 
homework:
  - [[homeworks/Multiplexação C10K via epoll e Estrutura dos Frames Binários-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar Multiplexação C10K via epoll e Estrutura dos Frames Binários em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* O que é exatamente o problema C10K em engenharia de servidores?
* Qual é a vantagem arquitetural e de performance do `epoll` em relação ao `select` em sistemas Linux?
* O que são os "Opcodes" (como Ping, Pong e Close) na estrutura de um frame binário WebSocket?
* Qual o propósito fundamental do mascaramento (Masking) de dados enviados do cliente para o servidor?

---
## 🔎 Perguntas de Elaboração

* Por que a arquitetura tradicional de "Uma Thread por Conexão" (Thread-per-connection) falha catastroficamente ao tentar escalar para 10.000 conexões WebSocket ativas?
* Como o uso de I/O não bloqueante e event loops puros soluciona o esgotamento de File Descriptors e Context Switches?
* O que poderia dar errado em infraestruturas de rede (proxies intermediários) se os payloads do WebSocket não contivessem obrigatoriamente a máscara XOR (Masking)?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[Sistema Operacional e File Descriptors]]
* [[Arquitetura Orientada a Eventos (Event Loops)]]
* [[Vulnerabilidades de Cache Poisoning (Motivo do Masking)]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
