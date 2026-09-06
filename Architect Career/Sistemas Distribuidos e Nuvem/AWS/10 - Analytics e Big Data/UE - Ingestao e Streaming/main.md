---
tema: Ingestao e Streaming - Kinesis e MSK
tipo: unidade-estudo
tags: [aws, kinesis, msk, kafka, streaming]
---
# 🧪 UE - Ingestao e Streaming

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Dados de eventos (clicks, logs, telemetria IoT) chegam em alta velocidade e volume, e processá-los em batch introduz latência inaceitável para casos como detecção de fraude ou dashboards em tempo real. Streaming resolve isso, mas traz desafios de particionamento (shards), ordenação, reprocessamento e backpressure. Escolher errado entre Kinesis (gerenciado, integrado) e MSK (Kafka, portável) ou dimensionar mal os shards causa throttling e perda de dados. Conecta com [[../../../Comunicacao e Mensageria/UE - Message Brokers vs Event Streaming - RabbitMQ e Kafka/main|Message Brokers vs Event Streaming]].

## 🧬 Grade Atômica de Tópicos
1. **Kinesis Data Streams:** Shards, partition key, ordenação por shard, retenção, consumidores (fan-out enhanced), reprocessamento; capacidade provisioned vs on-demand.
2. **Kinesis Data Firehose:** Entrega gerenciada near-real-time para S3/Redshift/OpenSearch, transformação via Lambda, buffering, conversão de formato (Parquet).
3. **Amazon MSK (Kafka gerenciado):** Quando escolher Kafka (ecossistema, portabilidade, semântica) vs Kinesis; MSK Serverless, consumidores.
4. **Padrões de streaming:** Processamento com Kinesis Data Analytics/Flink, exactly-once, janelas, integração com data lake.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon Kinesis](https://docs.aws.amazon.com/streams/latest/dev/introduction.html) e [Amazon MSK](https://docs.aws.amazon.com/msk/latest/developerguide/what-is-msk.html).
- **System Blueprint:** Pipeline de telemetria: produtores→Kinesis Data Streams→Firehose→S3 (Parquet)→Athena, com análise em tempo real via Flink.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Ingerir e entregar dados de streaming.
- [ ] Criar um Kinesis Data Stream e produzir/consumir eventos, observando ordenação por shard.
- [ ] Configurar Firehose para entregar no S3 convertendo para Parquet.
- [ ] Provocar throttling reduzindo shards e observar o comportamento.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Kinesis Data Streams e Shards]]
- [[./02 - Firehose e Entrega Gerenciada]]
- [[./03 - MSK e Padroes de Streaming]]
