---
tema: Auto Scaling e Elastic Load Balancing
tipo: unidade-estudo
tags: [aws, auto-scaling, elb, alb, nlb, escalabilidade]
---
# 🧪 UE - Escalabilidade - Auto Scaling e Load Balancing

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Elasticidade é a promessa central da nuvem, mas ela só funciona se dois mecanismos operarem juntos: um balanceador que distribui carga e detecta instâncias insalubres, e um grupo que adiciona/remove capacidade conforme a demanda. Configurar mal thresholds, cooldowns ou health checks causa flapping (escala pra cima e pra baixo sem parar), latência em picos e derrubada de instâncias saudáveis.

## 🧬 Grade Atômica de Tópicos
1. **Tipos de Load Balancer:** ALB (L7, HTTP, path/host routing, WAF), NLB (L4, milhões de conexões, IP estático, ultra baixa latência), Gateway LB (appliances de rede). Trade-offs de cada.
2. **Target groups e health checks:** Registro de alvos, health checks ativos, draining/deregistration delay, sticky sessions.
3. **Auto Scaling Group (ASG):** Launch templates, min/max/desired, distribuição multi-AZ, health checks (EC2 vs ELB), lifecycle hooks.
4. **Políticas de escala:** Target tracking, step scaling, scheduled scaling, predictive scaling; cooldown e warm-up.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Elastic Load Balancing](https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/what-is-load-balancing.html) e [Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html).
- **System Blueprint:** Web tier com ALB + ASG multi-AZ e target tracking em CPU/RequestCount.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Montar uma camada web elástica e resiliente.
- [ ] Criar ALB + Target Group + ASG multi-AZ com launch template.
- [ ] Configurar target tracking (ex.: 50% CPU) e gerar carga para observar scale-out/in.
- [ ] Matar uma instância e ver o ASG substituí-la via health check do ELB.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - ALB, NLB e Gateway Load Balancer]]
- [[./02 - Target Groups e Health Checks]]
- [[./03 - Auto Scaling Groups e Politicas de Escala]]
