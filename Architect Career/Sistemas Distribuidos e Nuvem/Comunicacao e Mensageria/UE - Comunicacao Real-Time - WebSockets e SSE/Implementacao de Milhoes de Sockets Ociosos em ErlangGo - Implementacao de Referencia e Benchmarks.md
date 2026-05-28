---
não iniciado-tema: Implementação de Milhões de Sockets Ociosos em ErlangGo
revision_29-05-2026: false
revision_30-05-2026: false
revision_04-06-2026: false
revision_27-06-2026: false
revision_26-08-2026: false
revision_25-10-2026: false
references:
  - 
homework:
  - [[homeworks/Implementação de Milhões de Sockets Ociosos em ErlangGo-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar Implementação de Milhões de Sockets Ociosos em ErlangGo em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* Como o Modelo de Atores no Erlang (ou Elixir) difere de threads padrão do sistema operacional?
* Como Go (Golang) através de "goroutines" sustenta milhares de conexões websocket de forma leve?
* Qual a diferença prática de consumo de memória inicial entre uma thread do SO em Java/C++ e um processo Erlang ou Goroutine?
* O que é o framework Phoenix Channels e por que ele é famoso no mundo de WebSockets?

---
## 🔎 Perguntas de Elaboração

* Por que máquinas virtuais (VMs) focadas em concorrência preemptiva como a BEAM (Erlang) são historicamente imbatíveis na sustentação de milhões de WebSockets simultâneos?
* Explique as dores e complexidades de lidar com o Garbage Collector (GC) em linguagens que mantêm milhões de conexões longas ativas em um heap gigante (como em Java ou Go).
* Sob quais cenários críticos você recomendaria que a equipe migrasse os microserviços de WebSockets de Node.js para Go ou Elixir/Erlang?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[Erlang BEAM VM e Actor Model]]
* [[Concorrência via Goroutines em Go]]
* [[Garbage Collection em High-Throughput Systems]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
