---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, banco-de-dados, rds, dynamodb, aurora]
---
# 🗂️ 05 - Bancos de Dados

## 🎯 Escopo da Área
A escolha do banco de dados é a decisão arquitetural mais difícil de reverter. A AWS oferece "purpose-built databases": relacional gerenciado (RDS/Aurora), NoSQL de chave-valor em escala (DynamoDB), cache em memória (ElastiCache), grafos (Neptune), colunar analítico (Redshift) e mais. O arquiteto deve mapear o padrão de acesso e os requisitos de consistência/latência/escala ao banco certo — aplicando na prática a teoria de CAP/PACELC. Este domínio sustentava a `Database – Specialty`.

## 🗺️ Mapa de Exploração
- [[./UE - RDS e Aurora/main|UE - RDS e Aurora]] -> Relacional gerenciado, Multi-AZ, read replicas, arquitetura do Aurora.
- [[./UE - DynamoDB/main|UE - DynamoDB]] -> Modelagem NoSQL, partition keys, índices, streams, capacidade.
- [[./UE - Caching e In-Memory/main|UE - Caching e In-Memory]] -> ElastiCache (Redis/Memcached), DAX, estratégias de cache.
- [[./UE - Bancos Especializados/main|UE - Bancos Especializados]] -> Redshift, Neptune, DocumentDB, Timestream, Keyspaces, QLDB.
