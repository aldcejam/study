---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, fundamentos, well-architected, iam, billing]
---
# 🗂️ 01 - Fundamentos e Well-Architected

## 🎯 Escopo da Área
Toda decisão arquitetural na AWS parte de três fundamentos: **onde** a computação roda fisicamente (infraestrutura global e seus domínios de falha), **quem pode fazer o quê** (IAM e o modelo de responsabilidade compartilhada) e **como avaliar se o design é bom** (Well-Architected Framework). Ignorar esses pilares leva a arquiteturas que falham quando uma AZ cai, contas comprometidas por permissões amplas demais e faturas explosivas sem governança.

Esta categoria estabelece o vocabulário e os mapas mentais reutilizados por todas as outras. É a base obrigatória para qualquer prova, do Cloud Practitioner ao Solutions Architect Professional.

## 🗺️ Mapa de Exploração
- [[./UE - Infraestrutura Global e Modelos de Nuvem/main|UE - Infraestrutura Global e Modelos de Nuvem]] -> Regions, AZs, Edge, Local Zones, Outposts e domínios de falha.
- [[./UE - Identidade Base e Estrutura de Conta/main|UE - Identidade Base e Estrutura de Conta]] -> IAM fundamentos, root, MFA, princípio do menor privilégio.
- [[./UE - Well-Architected Framework/main|UE - Well-Architected Framework]] -> Os 6 pilares e o processo de revisão arquitetural.
- [[./UE - Billing FinOps e Precificacao/main|UE - Billing, FinOps e Precificacao]] -> Modelos de compra, tags de custo, budgets e otimização.

## Ciclos pos-diagnostico
Pastas datadas com nivel por UE e ordem de estudo para uma IA dirigir a sessao (nao recomeçar do zero).

- [[./2026-08-13 - Ciclo pos-diagnostico/main|2026-08-13 — Ciclo pos-diagnostico]]
