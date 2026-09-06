---
tema: Caching e In-Memory - ElastiCache e DAX
tipo: unidade-estudo
tags: [aws, cache, elasticache, redis, memcached]
---
# 🧪 UE - Caching e In-Memory

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O caminho mais barato para uma requisição é aquele que nunca chega ao banco. Cache reduz latência e protege bancos de sobrecarga, mas introduz o problema mais difícil da computação: invalidação. Estratégias erradas de cache (sem TTL, sem invalidação, thundering herd no cache miss) causam dados obsoletos servidos aos usuários ou colapso do banco quando o cache expira em massa.

## 🧬 Grade Atômica de Tópicos
1. **ElastiCache Redis vs Memcached:** Redis (estruturas ricas, persistência, réplicas, cluster mode, pub/sub) vs Memcached (multi-thread, simples, sharding). Quando usar cada.
2. **Estratégias de cache:** Cache-aside (lazy loading), write-through, write-behind; TTL, eviction policies (LRU/LFU); mitigação de thundering herd/cache stampede.
3. **Alta disponibilidade no Redis:** Cluster mode enabled/disabled, replicas, Multi-AZ com failover automático, Global Datastore.
4. **DAX (DynamoDB Accelerator):** Cache in-memory específico para DynamoDB (microssegundos), write-through, quando substitui ElastiCache.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/WhatIs.html).
- **System Blueprint:** Padrão cache-aside com Redis à frente de RDS/Aurora, com TTL e jitter para evitar expiração em massa.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar cache-aside e medir ganho.
- [ ] Subir um ElastiCache Redis e implementar cache-aside numa app com fallback ao RDS.
- [ ] Medir latência e taxa de hit/miss; simular expiração em massa e observar o efeito no banco.
- [ ] Adicionar jitter aos TTLs e comparar.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Redis vs Memcached]]
- [[./02 - Estrategias de Cache e Invalidacao]]
- [[./03 - HA no Redis e DAX]]
