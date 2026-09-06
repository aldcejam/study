---
tema: Correlacao de Sinais, Context Propagation e Exemplars
_revision_14-06-2026: false
_revision_15-06-2026: false
_revision_20-06-2026: false
_revision_13-07-2026: false
_revision_11-09-2026: false
_revision_10-11-2026: false
references:
  - https://opentelemetry.io/docs/concepts/signals/
homework:
  - [[homeworks/Correlacao de Sinais, Context Propagation e Exemplars-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar Correlacao de Sinais, Context Propagation e Exemplars em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* O que é Context Propagation (Propagação de Contexto) em sistemas distribuídos e quais são os cabeçalhos padrão W3C (traceparent e tracestate)?
* Como funciona o recurso de *Exemplars* no Prometheus e de que forma ele permite conectar métricas de latência agregadas (ex: buckets de histograma) a traces de requisições específicas?
* Qual a diferença prática entre propagação de contexto em-banda (in-band, ex: W3C HTTP headers) e fora-de-banda (out-of-band, ex: logs estruturados correlacionados)?
* O que é a especificação Baggage do OpenTelemetry e quais são os cuidados necessários ao utilizá-la em termos de performance e segurança?

---
## 🔎 Perguntas de Elaboração

* Por que correlacionar métricas e traces dinamicamente em tempo de query (ex: clicar no gráfico do Grafana e ver o trace) é infinitamente superior a apenas ter ambos os sistemas isolados?
* Como a propagação de Baggage com chaves/valores arbitrários afeta o overhead de payloads HTTP entre microserviços em cascata profunda?
* Como configurar o Prometheus e o OpenTelemetry SDK para garantir que Exemplars sejam coletados e associados sem causar vazamento de memória ou aumento absurdo de consumo de RAM do coletor?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[../Observabilidade Plena e DevOps/main]]
* [[../../Arquitetura de Alta Escala/main]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
