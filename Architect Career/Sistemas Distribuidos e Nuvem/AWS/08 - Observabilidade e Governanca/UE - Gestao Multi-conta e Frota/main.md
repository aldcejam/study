---
tema: Gestao Multi-conta e Frota - Organizations e SSM
tipo: unidade-estudo
tags: [aws, organizations, control-tower, systems-manager, governanca]
---
# 🧪 UE - Gestao Multi-conta e Frota

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Uma conta única não escala com segurança: sem isolamento de blast radius, um erro em dev afeta produção, e o billing/segurança viram caos. A estratégia moderna é multi-conta com uma "landing zone" governada. O desafio é impor guardrails consistentes (SCPs), consolidar billing e gerenciar milhares de instâncias sem SSH manual. Errar a estrutura de OUs/contas cedo cria uma dívida organizacional cara de corrigir depois.

## 🧬 Grade Atômica de Tópicos
1. **AWS Organizations:** Root, Organizational Units (OUs), Service Control Policies (guardrails), consolidated billing, delegated administration.
2. **Landing Zone (Control Tower):** Blueprint automatizado de multi-conta, guardrails preventivos/detectivos, Account Factory, ligação com IAM Identity Center.
3. **Systems Manager (gestão de frota):** Session Manager (acesso sem SSH/bastion), Patch Manager, Run Command, State Manager, Parameter Store, Automation runbooks.
4. **Estratégia de contas e resource sharing:** Padrões de separação (por ambiente/time/workload), RAM (Resource Access Manager) para compartilhar sub-redes/TGW.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Organizations](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html), [AWS Control Tower](https://docs.aws.amazon.com/controltower/latest/userguide/what-is-control-tower.html) e [AWS Systems Manager](https://docs.aws.amazon.com/systems-manager/latest/userguide/what-is-systems-manager.html).
- **System Blueprint:** Landing zone com OUs (Security, Infra, Workloads), SCPs de guardrail e acesso administrativo sem bastion via Session Manager.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Governar múltiplas contas e a frota.
- [ ] Criar uma Organization com uma OU e aplicar uma SCP que nega uma Region inteira.
- [ ] Acessar uma EC2 privada via Session Manager (sem chave SSH/bastion).
- [ ] Guardar um parâmetro no Parameter Store e consumi-lo numa automação.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Organizations, OUs e SCPs]]
- [[./02 - Control Tower e Landing Zone]]
- [[./03 - Systems Manager - Frota e Session Manager]]
