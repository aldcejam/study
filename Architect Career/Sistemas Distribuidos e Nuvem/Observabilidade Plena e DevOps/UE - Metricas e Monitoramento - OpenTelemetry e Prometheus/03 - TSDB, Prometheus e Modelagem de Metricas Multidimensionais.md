---
tema: TSDB, Prometheus e Modelagem de Metricas Multidimensionais
_revision_14-06-2026: false
_revision_15-06-2026: false
_revision_20-06-2026: false
_revision_13-07-2026: false
_revision_11-09-2026: false
_revision_10-11-2026: false
references:
  - https://www.vldb.org/pvldb/vol8/p1816-teller.pdf
homework:
  - [[homeworks/TSDB, Prometheus e Modelagem de Metricas Multidimensionais-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar TSDB, Prometheus e Modelagem de Metricas Multidimensionais em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* Como funciona a técnica de compressão Gorilla (Double-Delta para timestamps e XOR para float64 values) adotada por motores modernos de TSDB?
* O que é o problema de alta cardinalidade em métricas e de que forma ele afeta a indexação de rótulos (labels/attributes) e o uso de memória no Prometheus?
* Qual a diferença prática e de modelagem matemática entre um Counter, um UpDownCounter (Gauge no Prometheus) e um Histogram no OpenTelemetry?
* Como o Prometheus lida com a raspagem de métricas (pull-based model) e quais estratégias existem (como pushgateways ou OTLP Push) para jobs efêmeros?

---
## 🔎 Perguntas de Elaboração

* Por que armazenar métricas com alta granularidade temporal e alta dimensionalidade diretamente em bancos de dados relacionais (SQL) tradicionais é inviável em grande escala?
* Sob quais cenários de falha de rede o modelo push do OpenTelemetry/OTLP é superior ao modelo pull tradicional do Prometheus (e vice-versa)?
* Como a representação de Histograms cumulativos no Prometheus se diferencia dos Exponential Histograms do OpenTelemetry em termos de precisão e tamanho de armazenamento?

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
