---
tema: VPC Avancado e Interconexao
tipo: unidade-estudo
tags: [aws, vpc, transit-gateway, privatelink, peering]
---
# 🧪 UE - VPC Avancado e Interconexao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em ambientes corporativos com dezenas ou centenas de VPCs, a interconexão vira um problema de topologia em escala: peering full-mesh explode combinatorialmente (N²), rotas transitivas não funcionam por padrão, e expor serviços entre contas sem passar pela internet exige entender PrivateLink. Errar aqui gera arquiteturas de rede caras, inseguras e impossíveis de manter.

## 🧬 Grade Atômica de Tópicos
1. **VPC Peering e seus limites:** Conexão 1:1, não-transitiva, sem CIDR sobreposto. Por que não escala para muitas VPCs.
2. **Transit Gateway (hub-and-spoke):** Roteamento centralizado e transitivo, route tables do TGW, segmentação, conexões inter-Region (peering de TGW).
3. **VPC Endpoints e PrivateLink:** Gateway Endpoints (S3/DynamoDB) vs Interface Endpoints (ENI + PrivateLink) para acessar serviços sem sair para a internet; expor serviços próprios via NLB + Endpoint Service.
4. **IPv6, Flow Logs e observabilidade de rede:** VPC Flow Logs para auditoria/debug, Traffic Mirroring, Reachability Analyzer; endereçamento IPv6.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Transit Gateway](https://docs.aws.amazon.com/vpc/latest/tgw/what-is-transit-gateway.html) e [AWS PrivateLink](https://docs.aws.amazon.com/vpc/latest/privatelink/what-is-privatelink.html).
- **System Blueprint:** Arquitetura de rede corporativa com TGW central + VPC de inspeção (firewall) + Direct Connect.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Conectar múltiplas VPCs via Transit Gateway.
- [ ] Criar 3 VPCs e conectá-las via Transit Gateway com roteamento transitivo.
- [ ] Criar um Gateway Endpoint para S3 e validar acesso sem NAT.
- [ ] Ativar Flow Logs e usar Reachability Analyzer para diagnosticar um caminho bloqueado.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - VPC Peering e Transit Gateway]]
- [[./02 - VPC Endpoints e PrivateLink]]
- [[./03 - IPv6, Flow Logs e Diagnostico de Rede]]
