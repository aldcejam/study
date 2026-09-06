---
tema: Alta Disponibilidade e Disaster Recovery
tipo: unidade-estudo
tags: [aws, alta-disponibilidade, disaster-recovery, rto, rpo]
---
# 🧪 UE - Alta Disponibilidade e Disaster Recovery

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> "Fica online" não é um requisito; "99,99% com RTO de 5 min e RPO de 1 min" é. Alta disponibilidade (tolerar falha de componente/AZ) e Disaster Recovery (sobreviver à perda de uma Region inteira) são objetivos diferentes com custos radicalmente diferentes. Escolher uma estratégia de DR sem casar RTO/RPO com o custo aceitável leva a gastar demais (multi-região ativo-ativo desnecessário) ou de menos (backup-restore quando o negócio exige segundos). Este é o assunto mais cobrado no SA Pro.

## 🧬 Grade Atômica de Tópicos
1. **Métricas de resiliência:** RTO (tempo para recuperar) vs RPO (perda de dados aceitável); SLAs, nines de disponibilidade e o custo marginal de cada nine.
2. **Estratégias de DR (do barato ao caro):** Backup & Restore, Pilot Light, Warm Standby, Multi-site Active/Active; trade-off custo vs RTO/RPO.
3. **Alta disponibilidade dentro da Region:** Multi-AZ em cada camada, health checks, auto-healing, remoção de pontos únicos de falha, quorum.
4. **DR multi-região:** Replicação de dados (Aurora Global, DynamoDB Global Tables, S3 CRR), Route 53 failover, roteamento global, teste de DR (game days). Conecta com [[../../../Resiliencia e Confiabilidade/main|Resiliencia e Confiabilidade]].

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Disaster Recovery whitepaper](https://docs.aws.amazon.com/whitepapers/latest/disaster-recovery-workloads-on-aws/disaster-recovery-workloads-on-aws.html) e [Reliability Pillar](https://docs.aws.amazon.com/wellarchitected/latest/reliability-pillar/welcome.html).
- **System Blueprint:** Warm Standby multi-região com Aurora Global Database + Route 53 health-check failover; game day de failover regional.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar e testar um DR.
- [ ] Montar um Pilot Light: dados replicados cross-Region + infra mínima pronta para escalar.
- [ ] Executar um failover controlado via Route 53 e medir RTO real.
- [ ] Desenhar uma tabela mapeando cada estratégia de DR ao seu RTO/RPO/custo.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - RTO, RPO e Nines de Disponibilidade]]
- [[./02 - Estrategias de DR - Pilot Light a Active-Active]]
- [[./03 - DR Multi-regiao e Game Days]]
