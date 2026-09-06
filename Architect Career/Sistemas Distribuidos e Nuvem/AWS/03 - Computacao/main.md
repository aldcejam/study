---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, computacao, ec2, containers, serverless]
---
# 🗂️ 03 - Computacao

## 🎯 Escopo da Área
Computação é onde o código roda, e a AWS oferece um espectro que vai de máquinas virtuais tradicionais (EC2) até funções efêmeras que escalam a zero (Lambda). O arquiteto precisa escolher o modelo certo por workload avaliando controle vs esforço operacional, custo por padrão de tráfego, latência de inicialização e limites de escala. Escolher errado aqui gera desperdício massivo ou gargalos de escalabilidade. Este domínio é central para SA Pro e DevOps Pro.

## 🗺️ Mapa de Exploração
- [[./UE - EC2 e Modelos de Instancia/main|UE - EC2 e Modelos de Instancia]] -> Famílias, AMIs, placement, tenancy, spot/reserved.
- [[./UE - Escalabilidade - Auto Scaling e Load Balancing/main|UE - Escalabilidade - Auto Scaling e Load Balancing]] -> ASG, políticas de scaling, ELB (ALB/NLB/GWLB).
- [[./UE - Containers - ECS EKS e Fargate/main|UE - Containers - ECS, EKS e Fargate]] -> Orquestração gerenciada, ECR, modelos de execução.
- [[./UE - Serverless Compute - Lambda/main|UE - Serverless Compute - Lambda]] -> Modelo de execução, cold starts, concorrência, limites.
