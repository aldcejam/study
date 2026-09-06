---
tema: Billing, FinOps e Modelos de Precificacao AWS
tipo: unidade-estudo
tags: [aws, custo, finops, billing, otimizacao]
---
# 🧪 UE - Billing, FinOps e Precificacao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Na nuvem, a arquitetura é a fatura. Decisões técnicas (tipo de instância, storage class, transferência entre AZs/Regions) têm custo direto e frequentemente invisível até o fim do mês. Sem governança de custos — tags, budgets, alertas — o custo cresce de forma não-linear e mata o ROI da migração. FinOps é o pilar que conecta engenharia e finanças e é fortemente cobrado nas provas Professional.

## 🧬 Grade Atômica de Tópicos
1. **Modelos de compra de computação:** On-Demand, Reserved Instances, Savings Plans, Spot. Trade-off entre compromisso/preço/flexibilidade e casos de uso de cada.
2. **Custos ocultos de rede e storage:** Data transfer OUT, cross-AZ, cross-Region, NAT Gateway; storage classes e custos de recuperação/requisição no S3.
3. **Governança e visibilidade:** Cost Explorer, AWS Budgets, Cost and Usage Report (CUR), cost allocation tags, consolidated billing via Organizations.
4. **Estratégias de otimização:** Rightsizing, desligamento de recursos ociosos, arquiteturas serverless para custo variável, uso de Trusted Advisor/Compute Optimizer.

> [!TIP] Material aprofundado
> Explicação completa de todos os tópicos (baseada na doc oficial da AWS): [[4-Detalhamento|📚 Detalhamento Técnico]]

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Billing User Guide](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html) e [Cost Optimization Pillar](https://docs.aws.amazon.com/wellarchitected/latest/cost-optimization-pillar/welcome.html).
- **System Blueprint:** Cultura FinOps (FinOps Foundation) aplicada com showback/chargeback por tags.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Instrumentar visibilidade e controle de custos.
- [ ] Criar um AWS Budget com alerta por e-mail em um limite baixo.
- [ ] Aplicar cost allocation tags e visualizar o custo por tag no Cost Explorer.
- [ ] Comparar o custo estimado de uma workload em On-Demand vs Savings Plan vs Spot na calculadora.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Modelos de Compra - On-Demand, RI, Savings Plans e Spot]]
- [[./02 - Custos de Rede e Storage]]
- [[./03 - Governanca de Custos e Estrategias FinOps]]
