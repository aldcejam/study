---
tema: Mensageria - SQS e SNS
tipo: unidade-estudo
tags: [aws, sqs, sns, mensageria, desacoplamento]
---
# 🧪 UE - Mensageria - SQS e SNS

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Chamadas síncronas acoplam a disponibilidade de dois serviços: se o consumidor cai, o produtor também falha. Filas e tópicos quebram esse acoplamento temporal, absorvem picos e garantem entrega, mas introduzem semânticas sutis: entrega at-least-once (duplicatas!), ordenação, visibility timeout e o tratamento de mensagens "venenosas" via DLQ. Ignorar essas semânticas gera processamento duplicado, perda de mensagens ou loops infinitos.

## 🧬 Grade Atômica de Tópicos
1. **SQS Standard vs FIFO:** Standard (at-least-once, ordenação best-effort, throughput alto) vs FIFO (exactly-once processing, ordenação estrita, message group id, dedup).
2. **Semântica de consumo:** Visibility timeout, long vs short polling, DLQ e maxReceiveCount, tratamento de idempotência.
3. **SNS (pub/sub) e fan-out:** Tópicos, subscribers (SQS, Lambda, HTTP, e-mail), padrão fan-out SNS→múltiplas SQS, message filtering, FIFO topics.
4. **Padrões de integração:** Fan-out, dead-letter, message replay, comparação com Kafka/EventBridge; ligação com Saga/Outbox.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/welcome.html) e [Amazon SNS](https://docs.aws.amazon.com/sns/latest/dg/welcome.html).
- **System Blueprint:** Fan-out SNS→SQS para desacoplar um evento em múltiplos processadores. Ver [[../../../Comunicacao e Mensageria/UE - Padroes de Mensageria - Saga e Outbox Pattern/main|Saga e Outbox]].

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar fan-out resiliente com DLQ.
- [ ] Criar um tópico SNS com 2 filas SQS inscritas (fan-out) e message filtering.
- [ ] Configurar DLQ e forçar falhas repetidas para ver a mensagem cair na DLQ.
- [ ] Comparar comportamento de ordenação entre SQS Standard e FIFO sob carga.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - SQS Standard vs FIFO]]
- [[./02 - Visibility Timeout, DLQ e Idempotencia]]
- [[./03 - SNS, Fan-out e Filtering]]
