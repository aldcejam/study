---
tema: OpenTelemetry Collector - Arquitetura e Pipelines
_revision_14-06-2026: false
_revision_15-06-2026: false
_revision_20-06-2026: false
_revision_13-07-2026: false
_revision_11-09-2026: false
_revision_10-11-2026: false
references:
  - https://opentelemetry.io/docs/collector/
homework:
  - [[homeworks/OpenTelemetry Collector - Arquitetura e Pipelines-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar OpenTelemetry Collector - Arquitetura e Pipelines em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* Como funciona o fluxo de dados em um pipeline do OpenTelemetry Collector entre os componentes Receivers, Processors e Exporters?
* Qual é a função crítica do processor `memory_limiter` e como ele previne que o Collector sofra OOM Kill sob picos extremos de tráfego?
* O que é tail-based sampling (amostragem baseada em cauda), onde ela deve ser implementada na topologia, e qual sua vantagem sobre a head-based sampling?
* Explique a diferença operacional, vantagens e desvantagens entre a implantação do Collector como Agent (Sidecar/DaemonSet) versus Gateway.

---
## 🔎 Perguntas de Elaboração

* Por que é recomendado utilizar um buffering de memória/disco nos Exporters do Collector para mitigar falhas temporárias nos sistemas de armazenamento de destino?
* De que forma o pipeline do OpenTelemetry Collector reduz custos de infraestrutura e tráfego de saída ao lidar com múltiplos destinos (ex: rotear logs para Elasticsearch e traces para Tempo)?
* Qual o impacto de processadores pesados de transformação (ex: processadores regex para parsing de logs) na latência de entrega de traces e métricas pelo Collector?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[../Observabilidade Plena e DevOps/main]]
* [[../../Resiliencia e Confiabilidade/main]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
