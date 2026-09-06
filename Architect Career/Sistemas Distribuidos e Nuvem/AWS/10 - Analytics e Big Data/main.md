---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, analytics, big-data, kinesis, redshift]
---
# 🗂️ 10 - Analytics e Big Data

## 🎯 Escopo da Área
Dados são o ativo mais valioso e o mais difícil de mover em escala. Este domínio cobre o ciclo completo de dados na AWS: ingestão em tempo real (Kinesis/MSK), armazenamento em data lake (S3 + Lake Formation), transformação (Glue/EMR), consulta (Athena/Redshift) e visualização (QuickSight). É a base da `Data Engineer – Associate`. O arquiteto precisa desenhar pipelines que equilibram latência (batch vs streaming), custo (serverless vs cluster) e governança de acesso a dados sensíveis.

## 🗺️ Mapa de Exploração
- [[./UE - Ingestao e Streaming/main|UE - Ingestao e Streaming]] -> Kinesis Data Streams/Firehose, MSK (Kafka gerenciado).
- [[./UE - Data Lake e Processamento/main|UE - Data Lake e Processamento]] -> S3, Glue, Lake Formation, EMR, Athena.
- [[./UE - Data Warehouse e BI/main|UE - Data Warehouse e BI]] -> Redshift, QuickSight, OpenSearch.
