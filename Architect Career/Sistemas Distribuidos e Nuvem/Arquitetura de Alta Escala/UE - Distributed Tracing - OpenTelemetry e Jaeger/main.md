---
tema: Distributed Tracing - OpenTelemetry e Jaeger
tipo: unidade-estudo
tags: [distribuídos, observabilidade, arquitetura, telemetria]
---
# 🧪 UE - Distributed Tracing - OpenTelemetry e Jaeger

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em uma arquitetura de microserviços, uma única requisição de usuário pode passar por 50 serviços diferentes. Se a requisição falha ou fica lenta, como saber qual serviço é o culpado? O problema é a perda de contexto entre saltos de rede. O **Distributed Tracing** resolve isso anexando um `Trace ID` único a cada requisição que flui por todo o sistema. **OpenTelemetry** fornece o padrão de instrumentação, e o **Jaeger** fornece a visualização da "árvore de chamadas". Sem tracing, debugar microserviços é como procurar uma agulha em um palheiro no escuro.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Conceitos de Trace e Span]:** A estrutura de uma requisição (Trace) composta por múltiplas unidades de trabalho (Spans).
2. **[Context Propagation]:** Como passar o Trace ID via cabeçalhos HTTP (W3C Trace Context) ou gRPC.
3. **[Sampling (Amostragem)]:** Por que não podemos traçar 100% das requisições em sistemas de alta escala e os tipos de amostragem (Head-based vs Tail-based).
4. **[Arquitetura do OpenTelemetry]:** O papel do Collector, SDKs e a exportação de dados (OTLP).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Dapper, a Large-Scale Distributed Systems Tracing Infrastructure* (Google Technical Report, 2010).
- **System Blueprint:** Implementação de Observabilidade na Uber (criadores do Jaeger) e na Netflix.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Instrumentar dois microserviços simples em Node.js ou Go usando OpenTelemetry SDK para que uma chamada de um serviço para o outro gere um trace completo no Jaeger.
- [ ] Subir uma instância do Jaeger via Docker.
- [ ] Configurar o exportador OTLP nos serviços.
- [ ] Realizar uma chamada e visualizar a linha do tempo (Gantt chart) no painel do Jaeger.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Propagacao de Contexto e Padrao W3C Trace Context - Teoria e Fundamentos]]
- [[./Arquitetura de Coleta e Exportacao de Telemetria - Funcionamento Interno e Arquitetura]]
- [[./Sobrecarga de Performance e Estratégias de Amostragem - Casos de Falha e Analise Amortizada]]
- [[./Analise de Caminho Critico com Jaeger - Implementacao de Referencia e Benchmarks]]
