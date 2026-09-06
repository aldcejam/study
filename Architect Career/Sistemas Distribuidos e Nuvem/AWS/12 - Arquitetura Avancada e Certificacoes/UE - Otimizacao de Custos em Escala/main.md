---
tema: Otimizacao de Custos em Escala (FinOps)
tipo: unidade-estudo
tags: [aws, finops, custo, otimizacao, arquitetura]
---
# 🧪 UE - Otimizacao de Custos em Escala

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em escala, pequenas ineficiências arquiteturais viram milhões em desperdício: uma escolha de NAT Gateway vs VPC Endpoint, um storage class errado, computação superdimensionada rodando 24/7. FinOps é a disciplina que trata custo como um requisito de engenharia de primeira classe, não como uma surpresa no fim do mês. As provas Professional testam a habilidade de escolher, entre várias soluções que funcionam, a mais econômica que ainda atende aos requisitos.

## 🧬 Grade Atômica de Tópicos
1. **Alavancas de custo de computação:** Rightsizing (Compute Optimizer), Savings Plans/RI, Spot para cargas tolerantes, serverless para tráfego variável, Graviton (ARM).
2. **Custo de dados e rede:** Escolha de storage class/lifecycle, evitar NAT desnecessário via endpoints, minimizar cross-AZ/cross-Region, compressão/formatos colunares.
3. **Visibilidade e accountability:** CUR + Athena/QuickSight, cost allocation tags, showback/chargeback por conta/OU, anomaly detection de custo.
4. **Trade-offs conscientes:** Custo vs resiliência vs performance; quando pagar mais por HA é justificado pelo negócio.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Cost Optimization Pillar](https://docs.aws.amazon.com/wellarchitected/latest/cost-optimization-pillar/welcome.html) e AWS Compute Optimizer.
- **System Blueprint:** Migração de EC2 x86 para Graviton + Savings Plans + Spot em workers assíncronos, com dashboard de custo por tag. Conecta com [[../../01 - Fundamentos e Well-Architected/UE - Billing FinOps e Precificacao/main|Billing e FinOps]].

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Reduzir custo de uma arquitetura mantendo requisitos.
- [ ] Analisar uma workload no Compute Optimizer e aplicar rightsizing.
- [ ] Substituir um NAT Gateway por VPC Gateway Endpoint para S3 e medir a economia.
- [ ] Comparar o custo mensal de uma app em EC2 On-Demand vs Fargate vs Lambda.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Alavancas de Custo de Computacao]]
- [[./02 - Custo de Dados e Rede]]
- [[./03 - Visibilidade, Tags e Trade-offs]]
