---
tema: Infraestrutura Global e Modelos de Nuvem AWS
tipo: unidade-estudo
tags: [aws, fundamentos, regioes, disponibilidade, edge]
---
# 🧪 UE - Infraestrutura Global e Modelos de Nuvem

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Latência é física e falhas são inevitáveis. A escolha errada de Region ou o uso de uma única Availability Zone transforma a queda de um datacenter em indisponibilidade total do produto. O arquiteto precisa entender os limites de isolamento (Region > AZ > data center) para desenhar sistemas que sobrevivem à perda de uma AZ inteira sem perda de dados, e para posicionar dados/computação perto do usuário respeitando soberania de dados e requisitos de latência.

## 🧬 Grade Atômica de Tópicos
1. **Hierarquia física (Region, AZ, Data Center):** AZs são isoladas em energia/rede mas conectadas por links de baixa latência (<~1-2ms). Region é o limite de isolamento e de soberania de dados. Impacto no design multi-AZ vs multi-região.
2. **Rede de borda (Edge Locations, Regional Edge Caches, PoPs):** Onde CloudFront, Route 53 e Global Accelerator terminam conexões perto do usuário. Diferença entre borda e Region.
3. **Extensões da nuvem (Local Zones, Wavelength, Outposts):** Quando a computação precisa ficar fisicamente dentro de uma cidade, numa rede 5G ou no datacenter do cliente (híbrido). Trade-off entre latência ultrabaixa e conjunto reduzido de serviços.
4. **Modelos de serviço e responsabilidade compartilhada:** IaaS vs PaaS vs SaaS aplicados à AWS; a linha "segurança DA nuvem" (AWS) vs "segurança NA nuvem" (cliente) muda conforme o serviço é gerenciado ou não.

> [!TIP] Material aprofundado
> Explicação completa de todos os tópicos (baseada na doc oficial da AWS): [[1-Detalhamento|📚 Detalhamento Técnico]]

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Global Infrastructure](https://aws.amazon.com/about-aws/global-infrastructure/) e o whitepaper [AWS Well-Architected — Reliability Pillar](https://docs.aws.amazon.com/wellarchitected/latest/reliability-pillar/welcome.html).
- **System Blueprint:** Arquiteturas Active-Active multi-região da Netflix/Amazon.com e o conceito de "cell-based architecture" para conter blast radius.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Mapear domínios de falha na prática e medir latência entre Regions.
- [ ] Listar Regions/AZs via `aws ec2 describe-availability-zones` e desenhar o diagrama de isolamento.
- [ ] Subir a mesma instância em 2 AZs e simular a perda de uma (parar/terminar) observando o comportamento.
- [ ] Medir RTT entre 3 Regions distintas e correlacionar com distância geográfica.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Region, AZ e Dominios de Falha]]
- [[./02 - Edge, Local Zones, Wavelength e Outposts]]
- [[./03 - Modelo de Responsabilidade Compartilhada]]
