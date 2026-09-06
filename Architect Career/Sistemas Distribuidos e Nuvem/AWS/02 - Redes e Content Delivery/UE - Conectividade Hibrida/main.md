---
tema: Conectividade Hibrida AWS
tipo: unidade-estudo
tags: [aws, vpn, direct-connect, hibrido, rede]
---
# 🧪 UE - Conectividade Hibrida

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Poucas empresas migram 100% de uma vez; a maioria vive anos em modelo híbrido. Conectar datacenter e nuvem com latência previsível, banda garantida e failover exige entender as diferenças entre VPN (rápida de subir, sobre a internet) e Direct Connect (dedicado, caro, previsível). Um design híbrido frágil derruba sistemas de produção quando o link cai sem redundância.

## 🧬 Grade Atômica de Tópicos
1. **Site-to-Site VPN:** Túneis IPsec redundantes, roteamento estático vs dinâmico (BGP), Virtual Private Gateway vs Transit Gateway como terminação.
2. **Direct Connect (DX):** Conexão física dedicada, Virtual Interfaces (Private/Public/Transit), Direct Connect Gateway para múltiplas Regions/VPCs.
3. **Resiliência híbrida:** VPN como backup do DX, múltiplas locações DX, BGP para failover automático, MACsec para criptografia.
4. **DNS híbrido e integração:** Route 53 Resolver endpoints (inbound/outbound) para resolução de nomes entre on-premises e VPC.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Direct Connect](https://docs.aws.amazon.com/directconnect/latest/UserGuide/Welcome.html) e [Site-to-Site VPN](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html).
- **System Blueprint:** Padrão de conectividade resiliente DX+VPN backup recomendado no Well-Architected (Reliability).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular conectividade híbrida (sem hardware real).
- [ ] Subir uma Site-to-Site VPN entre duas VPCs (uma simulando on-premises com software VPN).
- [ ] Configurar BGP e observar propagação de rotas.
- [ ] Desenhar o failover DX->VPN e descrever o comportamento do BGP.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Site-to-Site VPN e BGP]]
- [[./02 - Direct Connect e Virtual Interfaces]]
- [[./03 - Resiliencia Hibrida e DNS Resolver]]
