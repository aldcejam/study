---
tema: Bancos de Dados Especializados AWS
tipo: unidade-estudo
tags: [aws, redshift, neptune, documentdb, timestream, purpose-built]
---
# 🧪 UE - Bancos Especializados

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Usar um banco relacional para tudo é o antipadrão que a AWS combate com sua filosofia de "purpose-built databases". Cada formato de dado e padrão de query tem um motor otimizado: relações de grafo, séries temporais, documentos, colunar analítico, ledger imutável. Forçar o modelo errado gera queries lentas, custo alto e complexidade evitável. O arquiteto Professional precisa reconhecer o sinal no enunciado que aponta para o banco certo.

## 🧬 Grade Atômica de Tópicos
1. **Analítico colunar (Redshift):** Data warehouse MPP, armazenamento colunar, distribution/sort keys, Redshift Spectrum (query no S3), Serverless. (Aprofundado na categoria 10.)
2. **Grafos e documentos (Neptune, DocumentDB):** Neptune (Gremlin/SPARQL/openCypher) para relacionamentos densos; DocumentDB (compatível com MongoDB) para documentos.
3. **Séries temporais e chave-valor amplo (Timestream, Keyspaces):** Timestream para métricas/IoT; Keyspaces (Cassandra gerenciado) para escala wide-column.
4. **Ledger e casos de nicho (QLDB, MemoryDB):** QLDB (ledger imutável e verificável); MemoryDB (Redis durável como banco primário).

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** Guias de [Amazon Redshift](https://docs.aws.amazon.com/redshift/), [Neptune](https://docs.aws.amazon.com/neptune/), [Timestream](https://docs.aws.amazon.com/timestream/).
- **System Blueprint:** Arquitetura polyglot persistence — cada microserviço usa o banco adequado ao seu domínio.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Mapear casos de uso a bancos.
- [ ] Criar uma tabela de decisão: dado o padrão de acesso X, qual banco AWS e por quê.
- [ ] Subir um Timestream ou Neptune e rodar uma query nativa do modelo.
- [ ] Comparar o custo/latência de uma agregação em Redshift vs RDS.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Redshift e Analitico Colunar]]
- [[./02 - Neptune e DocumentDB]]
- [[./03 - Timestream, Keyspaces, QLDB e MemoryDB]]
