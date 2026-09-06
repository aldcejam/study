---
tema: Estrategias de Deploy e Automacao
tipo: unidade-estudo
tags: [aws, deploy, blue-green, canary, rollback]
---
# 🧪 UE - Estrategias de Deploy e Automacao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Todo deploy é uma mudança de estado em produção — e a estratégia define se um bug afeta 100% dos usuários instantaneamente ou 1% de forma reversível. In-place, blue/green e canary trocam velocidade, custo e risco de forma diferente. Não ter rollback automático baseado em métricas transforma um deploy ruim em incidente prolongado. Este é o núcleo do domínio de maior peso na prova DevOps Pro. Conecta com [[../../../Observabilidade Plena e DevOps/UE - Estrategias de Deploy e GitOps/main|Estrategias de Deploy e GitOps]].

## 🧬 Grade Atômica de Tópicos
1. **Padrões de deploy:** In-place vs rolling vs blue/green vs canary/linear; trade-offs de downtime, custo de infra duplicada e blast radius.
2. **Implementação na AWS:** CodeDeploy (blue/green EC2/ECS/Lambda), Lambda aliases + weighted routing, ECS deployment controllers, ALB weighted target groups.
3. **Rollback e safety:** Rollback automático em CloudWatch alarms, deployment health, feature flags (AppConfig), pre/post-traffic hooks.
4. **Plataformas gerenciadas:** Elastic Beanstalk (deploys gerenciados, environments), AWS Amplify para frontend; quando abstrair vs controlar.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [CodeDeploy Deployment Types](https://docs.aws.amazon.com/codedeploy/latest/userguide/welcome.html) e [AWS AppConfig](https://docs.aws.amazon.com/appconfig/latest/userguide/what-is-appconfig.html).
- **System Blueprint:** Canary de Lambda com deslocamento de tráfego linear (10% a cada 5 min) e rollback automático se o alarme de erro disparar.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Deploy canary com rollback automático.
- [ ] Configurar um deploy canary de uma Lambda (alias + weighted) via CodeDeploy.
- [ ] Ligar rollback automático a um alarme de erro do CloudWatch.
- [ ] Publicar uma versão com bug e confirmar o rollback.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Padroes de Deploy - Blue-Green e Canary]]
- [[./02 - Rollback Automatico e Feature Flags]]
- [[./03 - Elastic Beanstalk e Plataformas Gerenciadas]]
