---
tema: Containers na AWS - ECS, EKS e Fargate
tipo: unidade-estudo
tags: [aws, containers, ecs, eks, fargate, kubernetes]
---
# 🧪 UE - Containers - ECS, EKS e Fargate

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Containers padronizam o empacotamento, mas orquestrá-los em produção envolve decisões de plataforma difíceis de reverter: ECS (simples, AWS-nativo) vs EKS (Kubernetes, portável, complexo), e EC2 launch type (você gerencia os nós) vs Fargate (serverless, sem gerenciar nós). Escolher errado gera custo operacional desnecessário ou lock-in indesejado, além de falhas de rede/IAM específicas de cada modelo (task roles, ENI por task).

## 🧬 Grade Atômica de Tópicos
1. **ECS (Elastic Container Service):** Cluster, Task Definition, Service, scheduling; task role vs execution role; networking modes (awsvpc).
2. **EKS (Kubernetes gerenciado):** Control plane gerenciado, worker nodes/managed node groups, IRSA (IAM Roles for Service Accounts), add-ons; quando vale a complexidade do k8s.
3. **Fargate (serverless para containers):** Sem gerenciar nós, cobrança por vCPU/memória por task, isolamento; limites vs EC2 launch type.
4. **Registro e integração (ECR + rede):** ECR (imagens privadas, scanning), integração com ALB, service discovery (Cloud Map), autoscaling de tasks.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/Welcome.html) e [Amazon EKS](https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html).
- **System Blueprint:** Microserviço em ECS Fargate + ALB + ECR + service autoscaling; comparação com o mesmo em EKS. Ver também [[../../../Cloud Native/UE - Kubernetes Internals - API Server e Controller Manager/main|Kubernetes Internals]].

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Rodar o mesmo container em ECS Fargate e comparar com EKS.
- [ ] Publicar uma imagem no ECR e rodar como serviço no ECS Fargate atrás de um ALB.
- [ ] Configurar autoscaling de tasks por CPU e gerar carga.
- [ ] (Opcional) Subir a mesma app num cluster EKS e comparar esforço operacional.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - ECS - Tasks, Services e Roles]]
- [[./02 - EKS e IRSA]]
- [[./03 - Fargate e ECR]]
