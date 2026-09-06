---
tema: Infraestrutura Global e Modelos de Nuvem AWS
tipo: detalhamento-tecnico
contexto_pai: [[./main|Voltar para a UE]]
tags: [aws, fundamentos, regioes, disponibilidade, edge, outposts]
fontes:
  - AWS Global Infrastructure (aws.amazon.com/about-aws/global-infrastructure)
  - AWS Well-Architected Reliability Pillar (whitepaper)
  - AWS Fault Isolation Boundaries (whitepaper)
  - AWS Shared Responsibility Model (docs oficiais)
---
# 📚 Detalhamento — Infraestrutura Global e Modelos de Nuvem

> [!ABSTRACT] Como usar esta nota
> Este arquivo é o material de referência denso. Leia-o em paralelo com a `main.md` da UE. Cada seção corresponde a um item da **Grade Atômica**. Ao final há um glossário, uma tabela de decisão e perguntas de fixação (active recall).

---

## 1. Hierarquia física: Region, Availability Zone e Data Center

### 1.1 O modelo em três níveis

A AWS organiza sua infraestrutura em uma hierarquia de **domínios de isolamento de falha** (fault isolation boundaries). Do maior para o menor:

```
Partition (aws | aws-cn | aws-us-gov)
   └── Region (ex.: us-east-1, sa-east-1)
          └── Availability Zone (ex.: us-east-1a, us-east-1b)
                 └── Data Center (1..N por AZ, opaco ao cliente)
```

- **Region:** Área geográfica isolada (ex.: `sa-east-1` = São Paulo). Contém **no mínimo 3 AZs** (a maioria das novas). É o **maior limite de isolamento**: falhas, dados e a maioria dos serviços **não cruzam Regions** por padrão. É também a fronteira de **soberania de dados** — seus dados ficam na Region escolhida até você mover explicitamente.
- **Availability Zone (AZ):** Um ou mais data centers com **energia, refrigeração e rede física independentes**, separados por distância significativa (tipicamente quilômetros, mas < 100 km) para evitar que um único desastre — enchente, incêndio, falha elétrica — atinja duas AZs ao mesmo tempo.
- **Data Center:** Unidade física real (prédio). **Opaco** ao cliente: você nunca escolhe um data center, apenas a AZ. Uma AZ pode conter vários.

### 1.2 A propriedade-chave: latência vs isolamento

As AZs de uma Region são conectadas por **links de fibra dedicados, redundantes e de altíssima taxa**, com latência **single-digit millisecond** (tipicamente **< 1–2 ms** round-trip). Isso cria o trade-off central:

| Propriedade | Dentro de 1 AZ | Multi-AZ (mesma Region) | Multi-Region |
|---|---|---|---|
| Latência entre nós | µs | ~1–2 ms | 10–300+ ms |
| Isolamento de falha | Nenhum | Alto (energia/rede independentes) | Máximo (geografia + regulatório) |
| Custo de transferência | Grátis (mesma AZ)* | Pago (cross-AZ) | Pago (cross-Region, mais caro) |
| Complexidade | Baixa | Média (replicação síncrona viável) | Alta (replicação assíncrona, conflitos) |

\* Ver detalhes de billing na UE de FinOps — cross-AZ tem custo em ambos os sentidos.

> [!IMPORTANT] Regra mental de design
> A latência **< 2 ms** entre AZs é baixa o suficiente para **replicação síncrona** (ex.: RDS Multi-AZ, quorum de escrita). Entre Regions, a latência força **replicação assíncrona**, o que introduz **RPO > 0** (possível perda de dados no failover) e problemas de consistência.

### 1.3 Como os serviços se posicionam nessa hierarquia

Entender o **escopo** de cada serviço é decisivo para arquitetura e para provas:

- **Zonais (vivem em 1 AZ):** EC2 instance, EBS volume, subnet. Se a AZ cai, o recurso cai.
- **Regionais (abstraem múltiplas AZs automaticamente):** S3, DynamoDB, SQS, Lambda, ELB, RDS Multi-AZ. A AWS distribui a redundância entre AZs por você.
- **Globais:** IAM, Route 53, CloudFront, WAF (associado ao CloudFront), Organizations. Não pertencem a uma Region específica.

> [!TIP] Blast radius / raio de explosão
> Bom design **contém o blast radius**. A arquitetura **cell-based** (usada por Amazon.com e AWS internamente) divide a workload em células independentes por AZ/Region, de forma que a falha de uma célula afete apenas uma fração dos usuários, nunca 100%.

### 1.4 Conceitos avançados citados em provas

- **AZ IDs vs AZ names:** `us-east-1a` na sua conta pode ser um data center físico diferente do `us-east-1a` de outra conta — a AWS **embaralha** os nomes por conta para balancear carga. O **AZ ID** (ex.: `use1-az1`) é o identificador físico real e estável entre contas. Relevante para arranjos multi-conta (VPC sharing, RAM).
- **Region opt-in:** Regions mais novas (ex.: `me-south-1`, `af-south-1`) vêm **desabilitadas por padrão** e precisam ser explicitamente ativadas.
- **Local Zones e Wavelength** (seção 3) estendem a topologia, mas **não são AZs completas**.

---

## 2. Rede de borda: Edge Locations, Regional Edge Caches e PoPs

### 2.1 Por que existe uma camada de borda

O problema físico: a velocidade da luz na fibra impõe **~5 µs/km**. Um usuário em São Paulo acessando um servidor em `us-east-1` (Virgínia) paga **~120–150 ms** só de ida e volta. Para conteúdo estático e terminação de conexão TLS, isso é inaceitável. A **rede de borda** aproxima o ponto de contato do usuário.

### 2.2 Os componentes

- **Edge Locations (PoPs — Points of Presence):** Centenas de sites espalhados globalmente (muito mais que Regions). Terminam:
  - **Amazon CloudFront** — CDN: cacheia conteúdo estático/dinâmico perto do usuário.
  - **Amazon Route 53** — DNS: resolve nomes a partir do PoP mais próximo (latência-based, geolocation routing).
  - **AWS Global Accelerator** — usa a **rede backbone da AWS** para tráfego TCP/UDP, entrando pela borda mais próxima e trafegando pela rede privada da AWS em vez da internet pública.
  - **AWS WAF / Shield** — filtragem e mitigação de DDoS na borda.
- **Regional Edge Caches:** Camada **intermediária** entre os Edge Locations e a origem. Cacheam objetos menos populares (que expiram do cache do PoP) para reduzir idas à origem. Maiores e em menor número que os PoPs.

### 2.3 Borda vs Region — a distinção que cai em prova

| | Edge Location | Region |
|---|---|---|
| Quantidade | Centenas (600+) | ~30+ |
| Função | Cache, DNS, terminação TLS, DDoS | Computação e armazenamento completos |
| Serviços | CloudFront, Route 53, GA, WAF | EC2, RDS, S3, Lambda, etc. |
| Estado | Efêmero (cache) | Persistente (dados) |

> [!NOTE] Global Accelerator vs CloudFront
> Ambos usam a borda, mas: **CloudFront** é para **HTTP(S) cacheável** (conteúdo). **Global Accelerator** é para **qualquer TCP/UDP** (jogos, IoT, VoIP, APIs não-cacheáveis) e fornece **IPs Anycast estáticos** que roteiam para o endpoint saudável mais próximo.

---

## 3. Extensões da nuvem: Local Zones, Wavelength e Outposts

Essas são respostas da AWS para casos onde a Region "padrão" está **longe demais** ou os dados **não podem sair** de um local físico.

### 3.1 AWS Local Zones

- **O que é:** Uma extensão de uma Region colocada **dentro de uma grande área metropolitana** (ex.: Los Angeles, São Paulo, Bogotá).
- **Problema que resolve:** Latência **single-digit ms** para usuários da cidade, em aplicações sensíveis (edição de vídeo, gaming, AR/VR, ML inference em tempo real).
- **Trade-off:** Oferece um **subconjunto** de serviços (EC2, EBS, VPC subnets, ELB, alguns bancos). Você conecta a Local Zone à sua VPC da Region-pai via subnet estendida.

### 3.2 AWS Wavelength

- **O que é:** Infraestrutura AWS **embarcada dentro da rede 5G** de operadoras de telecom (ex.: no datacenter da operadora, no edge da rede móvel).
- **Problema que resolve:** Latência **ultrabaixa para dispositivos móveis/5G** sem o tráfego sair da rede da operadora para a internet pública.
- **Casos:** Veículos conectados, smart factories, live streaming interativo, AR/VR móvel.

### 3.3 AWS Outposts

- **O que é:** Rack(s) de hardware AWS **físico entregue e instalado no data center do próprio cliente** (on-premises). Gerenciado pela AWS, mesma API/console.
- **Problema que resolve:**
  - **Residência de dados** (regulação que exige dados no local físico).
  - **Latência ultrabaixa a sistemas locais legados** (ex.: chão de fábrica).
  - **Processamento local** que não tolera a viagem até a Region.
- **Variantes:** Outposts **racks** (42U) e Outposts **servers** (1U/2U para espaços pequenos, lojas, filiais).
- **Trade-off:** Você opera **hardware físico** (espaço, energia, rede), subconjunto de serviços, e depende do link com a Region-pai para o plano de controle.

### 3.4 Tabela comparativa das extensões

| | Local Zone | Wavelength | Outposts |
|---|---|---|---|
| Onde fica | Metrópole (AWS) | Rede 5G da telecom | Data center do cliente |
| Dono do local | AWS | Operadora | Cliente |
| Caso principal | Latência urbana | Latência 5G/móvel | Residência de dados + híbrido |
| Serviços | Subconjunto | Subconjunto (menor) | Subconjunto |
| Conectividade | VPC estendida da Region | Carrier gateway | Service link à Region-pai |

> [!TIP] Mnemônico
> **Local Zone** = cidade. **Wavelength** = 5G/telecom. **Outposts** = "AWS no *seu* rack".

---

## 4. Modelos de serviço e Responsabilidade Compartilhada

### 4.1 IaaS, PaaS, SaaS aplicados à AWS

O nível de abstração define **quanto você gerencia vs quanto a AWS gerencia**:

| Modelo | Você gerencia | AWS gerencia | Exemplo AWS |
|---|---|---|---|
| **IaaS** (Infra) | SO, runtime, app, dados, scaling | Virtualização, hardware, rede física | EC2, EBS, VPC |
| **PaaS** (Plataforma) | App, dados, config | SO, runtime, patching, scaling | Elastic Beanstalk, RDS, Lambda* |
| **SaaS** (Software) | Apenas seus dados/uso | Tudo | WorkMail, Chime, QuickSight |

\* Lambda é frequentemente classificado como **FaaS/Serverless**, um subtipo de PaaS onde você gerencia só o código.

> [!NOTE] Quanto mais gerenciado, menos controle e menos responsabilidade operacional
> Subir para PaaS/SaaS reduz **undifferentiated heavy lifting** (patching, scaling manual), mas reduz flexibilidade e pode aumentar lock-in.

### 4.2 O Modelo de Responsabilidade Compartilhada (Shared Responsibility Model)

Princípio central de segurança na AWS. Divide a segurança em duas metades:

- **AWS → "Security OF the Cloud" (segurança DA nuvem):** Hardware, data centers físicos, rede global, virtualização (hypervisor), e a infraestrutura que roda os serviços gerenciados.
- **Cliente → "Security IN the Cloud" (segurança NA nuvem):** Dados, configuração de IAM, criptografia, patching do SO (em IaaS), regras de firewall/Security Groups, gestão de identidades.

### 4.3 A linha se move conforme o serviço

Este é o ponto sutil que **cai muito em prova**: a fronteira **não é fixa**, ela desliza conforme o modelo:

```
EC2 (IaaS):        Cliente patcha o SO, gerencia tudo acima do hypervisor.
RDS (PaaS):        AWS patcha o SO e o engine do banco; cliente cuida de dados, users, SGs.
S3/DynamoDB:       AWS gerencia toda a infra; cliente cuida de policies, criptografia e dados.
Lambda:            AWS gerencia SO e runtime; cliente cuida só do código e permissões.
```

> [!IMPORTANT] Sempre do cliente, em TODOS os modelos
> Independentemente do serviço, **três coisas nunca deixam de ser responsabilidade do cliente**:
> 1. **Os dados** (classificação, criptografia, backup lógico).
> 2. **Gestão de identidade e acesso** (IAM, quem pode fazer o quê).
> 3. **Configuração** (Security Groups, buckets públicos vs privados, políticas de rede).
>
> A maioria dos incidentes ("bucket S3 vazado") é falha do **cliente** na parte "IN the cloud", não da AWS.

---

## 🧭 Tabela de Decisão Rápida (multi-AZ vs multi-Region vs extensões)

| Requisito dominante | Escolha |
|---|---|
| Alta disponibilidade dentro de um país, RPO≈0 | **Multi-AZ** (síncrono) |
| Disaster recovery geográfico, sobreviver à perda de uma Region | **Multi-Region** (assíncrono) |
| Soberania/residência de dados regulatória | **Region específica** ou **Outposts** |
| Latência < 10 ms para uma cidade | **Local Zone** |
| Latência ultrabaixa para 5G/móvel | **Wavelength** |
| Dados devem ficar no data center do cliente | **Outposts** |
| Entregar conteúdo estático global rápido | **CloudFront (Edge)** |
| Acelerar TCP/UDP não-cacheável global | **Global Accelerator (Edge)** |

---

## 📖 Glossário

- **Region:** Área geográfica isolada com ≥3 AZs; limite de isolamento e soberania.
- **Availability Zone (AZ):** Conjunto de 1+ data centers com infra independente; conectadas por links < 2 ms.
- **AZ ID:** Identificador físico estável da AZ entre contas (ex.: `use1-az1`).
- **Edge Location / PoP:** Site de borda para CDN, DNS, DDoS e terminação de conexão.
- **Regional Edge Cache:** Camada intermediária de cache entre PoP e origem.
- **Local Zone:** Extensão da Region em uma metrópole (baixa latência urbana).
- **Wavelength:** AWS embarcada na rede 5G da operadora.
- **Outposts:** Hardware AWS no data center do cliente (híbrido).
- **Blast radius:** Alcance de impacto de uma falha; bom design o minimiza.
- **Shared Responsibility Model:** Divisão de segurança entre AWS (da nuvem) e cliente (na nuvem).
- **RPO / RTO:** Recovery Point Objective (perda de dados aceitável) / Recovery Time Objective (tempo de recuperação aceitável).

---

## ✅ Active Recall (responda sem olhar)

1. Por que a latência < 2 ms entre AZs viabiliza replicação síncrona, e o que muda entre Regions?
2. Qual a diferença entre **AZ name** e **AZ ID**, e por que ela importa em arquiteturas multi-conta?
3. Diferencie CloudFront de Global Accelerator em termos de protocolo e caso de uso.
4. Um cliente precisa manter dados fisicamente dentro de sua própria fábrica por regulação. Qual extensão da nuvem usar e por quê?
5. No RDS, quem é responsável por: (a) patch do SO, (b) patch do engine, (c) criar usuários do banco, (d) definir Security Groups?
6. Explique como o Shared Responsibility Model "desliza" entre EC2, RDS e Lambda.
7. O que é blast radius e como a cell-based architecture o reduz?

---

## 🔗 Fontes e leitura oficial

- [AWS Global Infrastructure](https://aws.amazon.com/about-aws/global-infrastructure/)
- [Regions, AZs e Local Zones — docs EC2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)
- [AWS Well-Architected — Reliability Pillar](https://docs.aws.amazon.com/wellarchitected/latest/reliability-pillar/welcome.html)
- [AWS Fault Isolation Boundaries (whitepaper)](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/abstract-and-introduction.html)
- [Shared Responsibility Model](https://aws.amazon.com/compliance/shared-responsibility-model/)
- [AWS Outposts](https://aws.amazon.com/outposts/) · [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) · [Wavelength](https://aws.amazon.com/wavelength/)
