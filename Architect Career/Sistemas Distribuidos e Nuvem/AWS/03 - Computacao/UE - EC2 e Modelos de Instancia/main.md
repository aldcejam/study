---
tema: EC2 e Modelos de Instancia
tipo: unidade-estudo
tags: [aws, ec2, computacao, instancias]
---
# 🧪 UE - EC2 e Modelos de Instancia

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> EC2 é o serviço fundacional de computação e a maior fonte de desperdício de custo quando mal dimensionado. Escolher a família errada (CPU vs memória vs rede), ignorar Spot para workloads tolerantes a interrupção, ou não entender placement groups para cargas de HPC/baixa latência resulta em performance ruim e faturas infladas. Também é a base conceitual de containers e serverless (que rodam sobre a mesma infraestrutura).

## 🧬 Grade Atômica de Tópicos
1. **Famílias e dimensionamento:** General purpose, Compute/Memory/Storage optimized, Accelerated (GPU); vCPU/memória/rede; burstable (T-series) e CPU credits.
2. **Ciclo de vida e imagens:** AMI, User Data, Instance Metadata Service (IMDSv2), estados de instância, hibernação.
3. **Modelos de compra e Spot:** On-Demand, Reserved, Savings Plans, Spot (interrupção com aviso de 2 min), Capacity Reservations. Estratégia de mix.
4. **Placement e tenancy:** Placement groups (cluster/spread/partition), dedicated hosts/instances, Elastic Network Adapter, Nitro System.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon EC2 User Guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts.html) e whitepaper do [AWS Nitro System](https://aws.amazon.com/ec2/nitro/).
- **System Blueprint:** Frota mista On-Demand + Spot gerenciada por Auto Scaling Group para reduzir custo mantendo baseline de disponibilidade.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Explorar tipos de instância e Spot.
- [ ] Lançar instâncias de 2 famílias diferentes e comparar métricas sob carga (`stress`).
- [ ] Solicitar uma instância Spot e observar o comportamento/aviso de interrupção.
- [ ] Configurar IMDSv2 e ler metadados/User Data de dentro da instância.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Familias de Instancia e Dimensionamento]]
- [[./02 - AMIs, User Data e IMDS]]
- [[./03 - Spot, Placement Groups e Nitro]]
