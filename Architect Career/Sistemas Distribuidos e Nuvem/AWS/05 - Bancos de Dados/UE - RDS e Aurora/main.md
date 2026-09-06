---
tema: Amazon RDS e Aurora
tipo: unidade-estudo
tags: [aws, rds, aurora, relacional, alta-disponibilidade]
---
# 🧪 UE - RDS e Aurora

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Bancos relacionais são o coração transacional de quase todo sistema, mas escalá-los e mantê-los disponíveis é notoriamente difícil (o write master é um gargalo e um ponto único de falha). RDS remove o fardo operacional (patching, backup, failover), e o Aurora reinventa a arquitetura separando compute de storage para atingir HA e escala de leitura que o RDS clássico não alcança. Entender Multi-AZ (HA) vs read replicas (escala de leitura) é crítico — são coisas diferentes e frequentemente confundidas.

## 🧬 Grade Atômica de Tópicos
1. **RDS engines e operação:** MySQL, PostgreSQL, MariaDB, Oracle, SQL Server; parameter/option groups, automated backups vs snapshots, maintenance windows.
2. **Alta disponibilidade (Multi-AZ):** Standby síncrono em outra AZ, failover automático (mudança de DNS), diferença entre HA e escalabilidade.
3. **Escala de leitura (Read Replicas):** Réplicas assíncronas (lag eventual), promoção, cross-Region replicas para DR e latência.
4. **Aurora (arquitetura):** Storage distribuído em 6 cópias/3 AZs, log-structured, até 15 réplicas com baixo lag, Aurora Serverless v2, Global Database, replicação por quorum.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html).
- **Paper clássico:** [Amazon Aurora: Design Considerations for High Throughput Cloud-Native Relational Databases](https://www.allthingsdistributed.com/files/p1041-verbitski.pdf) (SIGMOD 2017) — leitura obrigatória para entender "the log is the database". Relacione com [[../../../Consenso e Concorrencia/UE - Raft Consensus - Leader Election e Log Replication/main|Raft]].

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Comparar HA e escala de leitura.
- [ ] Criar RDS PostgreSQL Multi-AZ e forçar um failover; medir o downtime.
- [ ] Adicionar uma read replica e observar o lag sob carga de escrita.
- [ ] Subir um cluster Aurora e comparar o comportamento de failover e leitura.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - RDS - Engines, Backup e Operacao]]
- [[./02 - Multi-AZ vs Read Replicas]]
- [[./03 - Arquitetura Interna do Aurora]]
