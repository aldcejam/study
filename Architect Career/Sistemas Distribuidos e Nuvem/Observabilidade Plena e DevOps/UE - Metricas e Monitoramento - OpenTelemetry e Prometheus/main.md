---
tema: Métricas, Monitoramento e OpenTelemetry
tipo: unidade-estudo
tags: [sistemas-distribuidos, observabilidade, opentelemetry, prometheus, devops]
---
# 🧪 UE - Metricas e Monitoramento - OpenTelemetry e Prometheus

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas distribuídos modernos e microserviços de alta escala, a falta de padronização na instrumentação de código gera lock-in de fornecedores (vendor lock-in) e degrada o desempenho devido a agentes proprietários pesados. Ignorar o ecossistema OpenTelemetry impede a correlação nativa tridimensional (Métricas, Logs e Traces) com propagação de contexto, tornando impossível debugar gargalos complexos ou degradações amortecidas que ocorrem sob estresse concorrencial e falhas parciais de rede.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **Padrão OpenTelemetry (API vs SDK vs OTLP):** Diferença arquitetural entre a API de instrumentação (acoplamento fraco e no-op por padrão) e o SDK de implementação (processadores de span, exportadores, gerenciamento de memória). Protocolo de transporte OTLP (OpenTelemetry Protocol) sobre gRPC/HTTP2 e serialização via Protobuf.
2. **Arquitetura do OpenTelemetry Collector:** O pipeline tripartido (Receivers, Processors, Exporters) do coletor. Configuração de técnicas críticas de produção como batching, filtering, amostragem baseada em cauda (tail-based sampling) e o processor `memory_limiter` para evitar OOM kills. Topologias de implantação: Agent (sidecar/daemonset) vs Gateway (cluster centralizado).
3. **Modelagem de Métricas Multidimensionais e Motores TSDB:** Tipos de instrumentos do OpenTelemetry (Counter, Asynchronous Counter, UpDownCounter, Histogram, Gauge) e como eles são traduzidos em séries temporais. Funcionamento interno de TSDBs (Time Series Databases) como Prometheus e VictoriaMetrics, incluindo compressão Gorilla/Double-Delta, indexação baseada em índices invertidos de labels e impacto na cardinalidade.
4. **Instrumentação Avançada e Correlação de Sinais (Exemplars e Context):** Mecanismo de correlação de sinais através de Context Propagation (W3C Trace Context, Baggage). Utilização de *Exemplars* para associar diretamente uma métrica agregada de latência (ex: percentil no Prometheus) ao ID de um trace específico no backend de tracing (ex: Jaeger/Tempo) sem custo computacional excessivo.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **O Padrão da Indústria:** [OpenTelemetry Specification](https://opentelemetry.io/docs/specs/otel/) - Especificações formais de modelos de dados de métricas, traces e logs.
- **Bancos de Dados Temporais:** [Gorilla: A Fast, Scalable, In-Memory Time Series Database](https://www.vldb.org/pvldb/vol8/p1816-teller.pdf) (VLDB 2015) - O paper clássico do Facebook que inspirou o motor de armazenamento de séries temporais do Prometheus.
- **System Blueprint:** Arquiteturas corporativas modernas combinando o OpenTelemetry Collector como gateway unificador, Prometheus/VictoriaMetrics para armazenamento de métricas de alta cardinalidade, e Grafana para visualização operacional.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Desenvolver um ecossistema local multi-serviço instrumentado nativamente com a API do OpenTelemetry (em Go, Rust ou Python), roteando todos os sinais para um OpenTelemetry Collector central que os distribui para o Prometheus (métricas) e Tempo/Jaeger (traces).
- [ ] Configurar um pipeline com `docker-compose` contendo: App Instrumentada, OpenTelemetry Collector, Prometheus (com suporte a exemplars habilitado) e Jaeger.
- [ ] Implementar instrumentação manual usando a API do OpenTelemetry: criar um Counter customizado de requisições e um Histogram para latência.
- [ ] Configurar propagação de contexto HTTP (W3C) entre dois serviços da Sandbox.
- [ ] Visualizar métricas no Prometheus contendo exemplars e navegar diretamente para o trace correspondente no Jaeger com um clique.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./01 - OpenTelemetry API vs SDK e OTLP]]
- [[./02 - OpenTelemetry Collector - Arquitetura e Pipelines]]
- [[./03 - TSDB, Prometheus e Modelagem de Metricas Multidimensionais]]
- [[./04 - Correlacao de Sinais, Context Propagation e Exemplars]]
