---
tema: Padroes de Arquitetura e Resiliencia AWS
tipo: unidade-estudo
tags: [aws, arquitetura, resiliencia, cell-based, padroes]
---
# 🧪 UE - Padroes de Arquitetura e Resiliencia

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Serviços gerenciados não tornam um sistema resiliente por mágica — a resiliência emerge de padrões de design que contêm falhas. Sem bulkheads, um cliente ruidoso derruba todos; sem backpressure, um pico afunda o sistema; sem contenção de blast radius, um bug se propaga para 100% dos usuários. Esta UE conecta a teoria de sistemas distribuídos (CAP, particionamento, circuit breaker) à sua materialização em serviços AWS, que é o raciocínio central do arquiteto Professional.

## 🧬 Grade Atômica de Tópicos
1. **Contenção de blast radius:** Cell-based architecture, shuffle sharding, isolamento por AZ/Region/conta; por que reduzir o raio de uma falha importa mais que evitá-la.
2. **Padrões de resiliência:** Circuit breaker, retry com backoff + jitter, timeout, bulkhead, graceful degradation — e onde a AWS os implementa (SDK, ALB, App Mesh). Conecta com [[../../../Resiliencia e Confiabilidade/main|Resiliencia e Confiabilidade]].
3. **Desacoplamento e elasticidade:** Filas para absorver picos, autoscaling, idempotência, statelessness; padrões event-driven e CQRS na AWS.
4. **Aplicação do Well-Architected sob ambiguidade:** Ler um enunciado, identificar o pilar priorizado e escolher a arquitetura que otimiza aquele pilar sem violar os outros — a habilidade central das provas Pro.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon Builders' Library](https://aws.amazon.com/builders-library/) — artigos de engenheiros da AWS sobre timeouts, retries, shuffle sharding e mais.
- **System Blueprint:** Aplicação de cell-based architecture para limitar o impacto de uma falha a uma célula de usuários.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Injetar falhas e observar contenção.
- [ ] Implementar retry com backoff+jitter e circuit breaker numa chamada entre serviços.
- [ ] Usar SQS como buffer e demonstrar que um pico não derruba o consumidor.
- [ ] Fazer um "game day": derrubar uma AZ e verificar que o sistema degrada sem cair.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Blast Radius, Cells e Shuffle Sharding]]
- [[./02 - Circuit Breaker, Retries e Bulkhead na AWS]]
- [[./03 - Well-Architected Aplicado a Cenarios]]
