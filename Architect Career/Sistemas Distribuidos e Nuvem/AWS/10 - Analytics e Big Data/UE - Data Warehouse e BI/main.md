---
tema: Data Warehouse e BI - Redshift e QuickSight
tipo: unidade-estudo
tags: [aws, redshift, quicksight, opensearch, data-warehouse]
---
# 🧪 UE - Data Warehouse e BI

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Consultas analíticas complexas (agregações sobre bilhões de linhas, joins pesados) mataram um banco OLTP: ele é otimizado para transações, não para varreduras analíticas. Data warehouses colunar MPP (Redshift) resolvem isso, mas exigem entender distribution/sort keys — errá-las causa data skew e shuffles de rede que arruínam a performance. Além disso, é preciso levar o insight ao usuário de negócio (QuickSight) e habilitar busca full-text/observabilidade (OpenSearch).

## 🧬 Grade Atômica de Tópicos
1. **Redshift (MPP colunar):** Arquitetura leader/compute nodes, armazenamento colunar, distribution styles (KEY/EVEN/ALL), sort keys, compressão, workload management (WLM).
2. **Redshift avançado:** Spectrum (query no S3 sem carga), Concurrency Scaling, Redshift Serverless, materialized views, data sharing entre clusters.
3. **BI e visualização (QuickSight):** SPICE (in-memory), dashboards, embedding, ML Insights; modelo de custo por sessão.
4. **Busca e observabilidade (OpenSearch):** Full-text search, log analytics, index/shard design, comparação com casos de uso do Redshift/Athena.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/welcome.html) e [Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/welcome.html).
- **System Blueprint:** Warehouse moderno: ingestão do data lake (S3) via Redshift Spectrum + modelagem dimensional + dashboards em QuickSight.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Modelar e consultar um mini-warehouse.
- [ ] Criar um Redshift Serverless, carregar dados e testar distribution/sort keys diferentes medindo a query.
- [ ] Consultar dados do S3 via Redshift Spectrum sem carregá-los.
- [ ] Construir um dashboard simples no QuickSight sobre esses dados.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Redshift - MPP, Distribution e Sort Keys]]
- [[./02 - Spectrum, Serverless e Data Sharing]]
- [[./03 - QuickSight e OpenSearch]]
