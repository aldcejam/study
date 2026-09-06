---
tema: DNS e Amazon Route 53
tipo: unidade-estudo
tags: [aws, dns, route53, roteamento, alta-disponibilidade]
---
# 🧪 UE - DNS e Route 53

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> DNS é o primeiro ponto de decisão de qualquer requisição e uma das alavancas mais poderosas — e mais mal usadas — de alta disponibilidade e roteamento global. Políticas de roteamento erradas causam failover que não acontece, tráfego mandado para a região errada e TTLs que impedem recuperação rápida. Route 53 é peça central em arquiteturas multi-região e DR.

## 🧬 Grade Atômica de Tópicos
1. **Fundamentos DNS e registros:** A, AAAA, CNAME, Alias (específico AWS, resolve para recursos sem custo de query), NS, TTL e cache.
2. **Políticas de roteamento:** Simple, Weighted, Latency-based, Failover, Geolocation, Geoproximity, Multivalue. Quando usar cada e como combinam.
3. **Health checks e failover:** Monitoramento de endpoints, failover ativo-passivo e ativo-ativo, integração com CloudWatch alarms.
4. **Zonas privadas e resolver híbrido:** Private Hosted Zones associadas a VPCs; Route 53 Resolver para DNS híbrido on-premises↔VPC.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/Welcome.html).
- **System Blueprint:** Arquitetura multi-região ativa-ativa usando latency-based routing + health checks para failover automático de Region.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar failover DNS entre duas regiões.
- [ ] Registrar/usar um domínio e criar registros Alias para dois ALBs em Regions distintas.
- [ ] Configurar política Failover com health check e derrubar o primário para observar o failover.
- [ ] Comparar o comportamento com política Latency-based.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Registros DNS e Alias]]
- [[./02 - Politicas de Roteamento]]
- [[./03 - Health Checks, Failover e Zonas Privadas]]
