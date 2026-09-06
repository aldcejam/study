---
tema: Infraestrutura como Codigo - CloudFormation e CDK
tipo: unidade-estudo
tags: [aws, iac, cloudformation, cdk, sam]
---
# 🧪 UE - Infraestrutura como Codigo

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Infraestrutura provisionada manualmente é irreproduzível, não-versionada e propensa a "snowflake servers" e drift silencioso entre ambientes. IaC torna a infra declarativa, versionada e auditável, mas introduz seus próprios problemas: gestão de estado, dependências entre recursos, rollback de stacks falhas e organização de código em escala. Dominar IaC é pré-requisito absoluto para qualquer prática DevOps séria na AWS.

## 🧬 Grade Atômica de Tópicos
1. **CloudFormation (declarativo):** Templates, resources, parameters, mappings, outputs, intrinsic functions; change sets, stack policies, rollback e comportamento em falha.
2. **Modularização e escala:** Nested stacks, cross-stack references (exports/imports), StackSets (multi-conta/multi-Region), organização de stacks por ciclo de vida.
3. **CDK e SAM (imperativo/abstração):** AWS CDK (definir infra em linguagem de programação, constructs, síntese para CloudFormation); SAM para serverless.
4. **Drift, estado e boas práticas:** Detecção de drift, imutabilidade, gestão de segredos em templates, testes de infra, comparação com Terraform.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html) e [AWS CDK](https://docs.aws.amazon.com/cdk/v2/guide/home.html).
- **System Blueprint:** Ambiente completo (VPC + ECS + RDS) versionado em CDK, com StackSets replicando guardrails em múltiplas contas.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Provisionar e evoluir infra declarativamente.
- [ ] Escrever um template CloudFormation (VPC + EC2) e aplicar via change set.
- [ ] Alterar um recurso manualmente e detectar o drift.
- [ ] Recriar a mesma stack em CDK e comparar a experiência.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CloudFormation - Templates e Change Sets]]
- [[./02 - Nested Stacks e StackSets]]
- [[./03 - CDK, SAM e Deteccao de Drift]]
