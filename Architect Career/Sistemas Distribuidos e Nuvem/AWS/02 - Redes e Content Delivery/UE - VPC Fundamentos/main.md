---
tema: VPC Fundamentos
tipo: unidade-estudo
tags: [aws, vpc, rede, subnet, security-group]
---
# 🧪 UE - VPC Fundamentos

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Uma VPC mal planejada é uma dívida técnica quase impossível de pagar: CIDRs sobrepostos impedem peering futuro, subnets públicas por engano expõem bancos de dados, e a confusão entre Security Group (stateful) e NACL (stateless) gera falhas de conectividade intermitentes que consomem dias de debugging. Este é o alicerce de toda segurança e conectividade na AWS.

## 🧬 Grade Atômica de Tópicos
1. **CIDR, subnets e AZs:** Planejamento de blocos de endereço não-sobrepostos, subnets por AZ, reserva de IPs pela AWS (5 por subnet). Público vs privado é definido pela route table, não pela subnet.
2. **Roteamento e gateways:** Route tables, Internet Gateway (tráfego bidirecional público), NAT Gateway/Instance (saída para privadas), Egress-only IGW (IPv6).
3. **Firewall em camadas (Security Group vs NACL):** SG é stateful e opera na ENI; NACL é stateless e opera na subnet (regras numeradas, allow/deny). Por que ambos existem e ordem de avaliação.
4. **Interfaces de rede (ENI, IP secundário, ENA):** Como o tráfego entra/sai de uma instância; Elastic IP; impacto de placement e bandwidth.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon VPC User Guide](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
- **System Blueprint:** Topologia clássica de 3 tiers (subnets públicas para ALB, privadas para app, isoladas para banco) multi-AZ.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Construir uma VPC 3-tier do zero manualmente.
- [ ] Criar VPC /16, subnets pública/privada em 2 AZs, IGW e NAT Gateway.
- [ ] Colocar uma instância pública (bastion) e uma privada; acessar a privada só via bastion.
- [ ] Provocar uma falha propositalmente com NACL bloqueando a porta de retorno efêmera e observar o comportamento stateless.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CIDR, Subnets e Route Tables]]
- [[./02 - Gateways - IGW, NAT e Egress-only]]
- [[./03 - Security Groups vs NACLs]]
