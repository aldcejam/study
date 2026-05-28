---
não iniciado-tema: Reconexão em Massa Thundering Herd e Stateful Drops
revision_29-05-2026: false
revision_30-05-2026: false
revision_04-06-2026: false
revision_27-06-2026: false
revision_26-08-2026: false
revision_25-10-2026: false
references:
  - 
homework:
  - [[homeworks/Reconexão em Massa Thundering Herd e Stateful Drops-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar Reconexão em Massa Thundering Herd e Stateful Drops em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* O que caracteriza exatamente o problema de "Thundering Herd" em um sistema distribuído de WebSockets?
* O que é Exponential Backoff e como ele atua em reconexões?
* Qual o impacto direto de perder estado de sessão (Stateful Drop) em um nó WebSocket que subitamente cai?
* O que é "Jitter" e qual o seu papel vital junto ao Exponential Backoff?

---
## 🔎 Perguntas de Elaboração

* Por que gerenciar balanceamento de carga para WebSockets (Stateful) é significativamente mais frágil e complexo do que o roteamento de requisições REST (Stateless)?
* Em um evento em que um servidor cai desconectando 100.000 usuários simultaneamente, como você projeta os clientes para que a tentativa de reconexão deles não cause DDoS e derrube os servidores restantes?
* Como soluções como Redis Pub/Sub ou Kafka ajudam a distribuir o estado das sessões WebSockets em um cluster horizontal?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[Padrões de Resiliência e Backoff Strategy]]
* [[Sistemas Distribuídos e Load Balancing L7]]
* [[Redis Pub/Sub e Sincronização de Estado]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
