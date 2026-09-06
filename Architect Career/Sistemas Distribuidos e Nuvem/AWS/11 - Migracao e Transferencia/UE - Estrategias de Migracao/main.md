---
tema: Estrategias de Migracao para AWS
tipo: unidade-estudo
tags: [aws, migracao, 7rs, dms, mgn]
---
# 🧪 UE - Estrategias de Migracao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Migrar "lift-and-shift" tudo cegamente desperdiça a nuvem; refatorar tudo antes de migrar atrasa anos e quebra o business case. A decisão certa é caso a caso, guiada por um framework (7 Rs) e por dados de descoberta. Migrações de banco são as mais arriscadas: heterogêneas (Oracle→PostgreSQL) exigem conversão de esquema, e minimizar downtime exige replicação contínua (CDC). Errar aqui gera cutover fracassado, perda de dados e rollback caótico.

## 🧬 Grade Atômica de Tópicos
1. **As 7 Rs:** Retire, Retain, Rehost (lift-and-shift), Relocate, Repurchase, Replatform, Refactor; critérios de escolha e trade-offs de esforço vs benefício.
2. **Descoberta e planejamento (Migration Hub / ADS):** Application Discovery Service, dependências, waves de migração, portfolio assessment.
3. **Migração de servidores (MGN):** Application Migration Service (replicação em bloco, cutover com mínimo downtime), test instances.
4. **Migração de bancos (DMS + SCT):** DMS (homogênea e heterogênea, full load + CDC), Schema Conversion Tool, estratégias de cutover e validação de dados.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Migration Hub](https://docs.aws.amazon.com/migrationhub/latest/ug/whatishub.html) e [AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/Welcome.html).
- **System Blueprint:** Migração de Oracle on-premises para Aurora PostgreSQL usando SCT (esquema) + DMS (full load + CDC) com cutover de baixo downtime.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Migrar um banco com downtime mínimo.
- [ ] Usar DMS para migrar dados de um banco fonte para RDS com full load + CDC.
- [ ] Validar consistência dos dados pós-migração.
- [ ] Classificar 5 aplicações fictícias em uma das 7 Rs justificando cada escolha.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - As 7 Rs e Planejamento]]
- [[./02 - Migration Hub, ADS e MGN]]
- [[./03 - DMS e SCT - Migracao de Bancos]]
