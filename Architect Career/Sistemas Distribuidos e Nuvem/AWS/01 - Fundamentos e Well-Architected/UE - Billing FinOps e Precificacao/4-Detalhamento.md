---
tema: Billing, FinOps e Modelos de Precificacao AWS
tipo: detalhamento-tecnico
contexto_pai: [[./main|Voltar para a UE]]
tags: [aws, custo, finops, billing, otimizacao, savings-plans]
fontes:
  - AWS Billing User Guide (docs oficiais)
  - AWS Well-Architected Cost Optimization Pillar (whitepaper)
  - AWS Pricing (aws.amazon.com/pricing)
  - FinOps Foundation Framework
---
# 📚 Detalhamento — Billing, FinOps e Precificação

> [!ABSTRACT] Como usar esta nota
> Material de referência denso para acompanhar a `main.md`. Máxima central: **na nuvem, a arquitetura é a fatura.** Decisões técnicas (tipo de instância, storage class, transferência entre AZs/Regions) têm custo direto e muitas vezes **invisível até o fim do mês**. FinOps conecta engenharia e finanças e cai forte nas provas Professional.

---

## 0. Os 3 fundamentos da precificação AWS

Praticamente todo serviço cobra sobre alguma combinação de:

1. **Compute** — tempo e tamanho da capacidade (vCPU/memória, ou invocações/duração no serverless).
2. **Storage** — GB armazenados por mês, mais requisições e recuperação.
3. **Data transfer (rede)** — o custo mais **oculto**. Regra geral:
   - **IN (entrada) para a AWS:** geralmente **grátis**.
   - **OUT (saída) para a internet:** **pago** (e cresce).
   - **Entre AZs / entre Regions:** **pago** (cross-Region é o mais caro).

> [!IMPORTANT] Free Tier
> A AWS oferece três tipos de free tier: **12 meses grátis** (ex.: 750h/mês de EC2 t2/t3.micro), **always free** (ex.: 1M invocações Lambda/mês, 25 GB DynamoDB) e **trials** (curta duração). Não confie nele para produção; monitore com Budgets.

---

## 1. Modelos de compra de computação (EC2)

O trade-off é sempre **compromisso × preço × flexibilidade**.

### 1.1 On-Demand
- Paga por segundo/hora **sem compromisso**. Preço cheio.
- **Quando:** cargas **imprevisíveis**, curtas, de desenvolvimento/teste, ou pico temporário.

### 1.2 Reserved Instances (RI)
- **Compromisso de 1 ou 3 anos** em troca de **até ~72% de desconto** vs On-Demand.
- Variações: **Standard RI** (maior desconto, menos flexível) vs **Convertible RI** (permite trocar família/SO, desconto menor).
- Pagamento: **All Upfront** (maior desconto) / **Partial** / **No Upfront**.
- **Quando:** carga **estável e previsível** (baseline 24/7).
- *Nota:* RIs são um mecanismo de **billing/capacidade**, mais rígido que Savings Plans.

### 1.3 Savings Plans (recomendado hoje)
- Compromisso de **valor por hora em USD/h** por 1 ou 3 anos, em troca de descontos comparáveis aos RIs (~até 72%).
- Tipos:
  - **Compute Savings Plans:** mais flexível — aplica a EC2 **qualquer família/region/SO**, Fargate e Lambda. Menor desconto.
  - **EC2 Instance Savings Plans:** restrito a uma **família em uma region**; maior desconto.
- **Quando:** baseline previsível, mas querendo **mais flexibilidade** que RIs.

### 1.4 Spot Instances
- Usa **capacidade ociosa** da AWS com **até ~90% de desconto**. Mas a AWS pode **recuperar (interromper)** a instância com **aviso de 2 minutos**.
- **Quando:** cargas **tolerantes a interrupção e stateless**: batch, CI/CD, big data, rendering, workloads com checkpoint.
- **Nunca:** para estado crítico que não tolera término abrupto (a menos que combinado com estratégias de resiliência).

### 1.5 Outros
- **Dedicated Hosts / Dedicated Instances:** hardware físico dedicado (compliance, licenças BYOL amarradas a socket/core).
- **Capacity Reservations:** garante capacidade em uma AZ sem compromisso de prazo (pode combinar com Savings Plans para desconto).

### 1.6 Tabela comparativa

| Modelo             | Desconto    | Compromisso             | Interrupção?    | Caso de uso                         |
| ------------------ | ----------- | ----------------------- | --------------- | ----------------------------------- |
| On-Demand          | 0%          | Nenhum                  | Não             | Imprevisível, curto, dev/test       |
| Savings Plans      | até ~72%    | 1 ou 3 anos ($/h)       | Não             | Baseline previsível + flexível      |
| Reserved Instances | até ~72%    | 1 ou 3 anos (instância) | Não             | Baseline estável e específico       |
| Spot               | até ~90%    | Nenhum                  | **Sim (2 min)** | Batch, stateless, tolerante a falha |
| Dedicated Host     | — (premium) | Opcional                | Não             | Compliance, licença BYOL            |

> [!TIP] Estratégia combinada ("layered")
> Cobrir o **baseline** 24/7 com **Savings Plans/RIs**, absorver **picos** com **On-Demand**, e rodar **trabalho tolerante a falha** em **Spot**. Isso minimiza custo sem sacrificar disponibilidade.

---

## 2. Custos ocultos de rede e storage

### 2.1 Data transfer (o vilão silencioso)

| Fluxo                               | Custo                                       |
| ----------------------------------- | ------------------------------------------- |
| Internet → AWS (IN)                 | Grátis                                      |
| AWS → Internet (OUT)                | **Pago** (escalonado por volume)            |
| Mesma AZ (via IP privado)           | Grátis                                      |
| **Cross-AZ** (mesma Region)         | **Pago** (ambos os sentidos)                |
| **Cross-Region**                    | **Pago** (mais caro que cross-AZ)           |
| Via **NAT Gateway**                 | Custo por hora **+ por GB processado**      |
| Via **VPC Endpoints / PrivateLink** | Reduz saída pela internet (pode economizar) |
| Saída via **CloudFront**            | Mais barata que saída direta do EC2/S3      |
|                                     |                                             |

> [!WARNING] Armadilhas comuns
> - Tráfego **cross-AZ** em arquiteturas "multi-AZ mal pensadas" (chat entre serviços em AZs diferentes) infla a fatura.
> - **NAT Gateway** cobra por GB processado — mover muito tráfego por ele é caro; VPC Endpoints podem substituir para serviços AWS.
> - Servir mídia direto do S3 para a internet é mais caro que via **CloudFront**.

### 2.2 Storage — S3 classes e custos

O S3 cobra por **armazenamento + requisições + recuperação + transição**. Escolher a classe certa é otimização direta:

| Classe                            | Caso                                       | Custo storage                   | Custo recuperação                   |
| --------------------------------- | ------------------------------------------ | ------------------------------- | ----------------------------------- |
| **S3 Standard**                   | Acesso frequente                           | Alto                            | Nenhum                              |
| **S3 Intelligent-Tiering**        | Padrão de acesso **desconhecido/variável** | Médio (+ taxa de monitoramento) | Automático, sem taxa de recuperação |
| **S3 Standard-IA**                | Infrequente, mas acesso rápido             | Menor                           | Paga por GB recuperado              |
| **S3 One Zone-IA**                | Infrequente, **1 só AZ** (recriável)       | Ainda menor                     | Paga recuperação                    |
| **S3 Glacier Instant Retrieval**  | Arquivo com acesso raro instantâneo        | Baixo                           | Ms                                  |
| **S3 Glacier Flexible Retrieval** | Arquivo, recuperação em min–horas          | Muito baixo                     | Min a 12h                           |
| **S3 Glacier Deep Archive**       | Arquivo de longuíssimo prazo (7–10 anos)   | O mais baixo                    | 12–48h                              |

- **S3 Lifecycle policies:** automatizam a **transição** entre classes e a **expiração** de objetos (ex.: Standard → IA após 30d → Glacier após 90d → delete após 365d).
- **EBS:** cobra por GB **provisionado** (não usado) por mês + IOPS/throughput conforme o tipo (gp3, io2, etc.). Snapshots vão para S3 e cobram incremental.

> [!TIP] Intelligent-Tiering
> Quando o padrão de acesso é imprevisível, **Intelligent-Tiering** costuma ser a escolha segura: move objetos entre camadas automaticamente e evita taxas de recuperação surpresa.

---

## 3. Governança e visibilidade de custos

Ferramentas para **enxergar, alocar e controlar** o gasto:

### 3.1 AWS Cost Explorer
- Visualização **interativa e histórica** (até 12+ meses) de custo e uso, com **previsões**.
- Filtra e agrupa por serviço, conta, **tag**, region, tipo de compra. Recomenda **rightsizing** e compra de Savings Plans.

### 3.2 AWS Budgets
- Define **orçamentos** (custo, uso, RI/SP utilization/coverage) e dispara **alertas** (e-mail/SNS) ao ultrapassar limites reais **ou previstos**.
- **Budget Actions:** pode **aplicar ação automática** (ex.: anexar policy restritiva, parar instâncias) ao estourar o limite.

### 3.3 Cost and Usage Report (CUR)
- O relatório **mais detalhado** (nível de linha por hora/recurso), entregue no **S3**. Base para análise avançada (Athena, QuickSight) e FinOps sofisticado.

### 3.4 Cost Allocation Tags
- **Tags** (chave/valor) aplicadas a recursos para **atribuir custo** por projeto, time, ambiente (`env=prod`), centro de custo.
- Precisam ser **ativadas** no console de billing para aparecerem no Cost Explorer/CUR.
- Tipos: **AWS-generated** (`aws:createdBy`) e **user-defined**.
- Base para **showback** (mostrar custo ao time) e **chargeback** (cobrar o time).

### 3.5 AWS Organizations & Consolidated Billing
- Agrupa **múltiplas contas** sob uma **management account** com **fatura única**.
- Benefícios: **volume discounts agregados**, compartilhamento de **RIs/Savings Plans** entre contas, isolamento por conta (blast radius de billing e segurança), aplicação de **SCPs**.
- **Boa prática:** conta separada por ambiente/time; billing centralizado.

### 3.6 Ferramentas de otimização automática
- **AWS Compute Optimizer:** recomenda rightsizing de EC2/EBS/Lambda/Auto Scaling com base em métricas reais.
- **AWS Trusted Advisor:** checks de Cost Optimization (instâncias ociosas, RIs subutilizadas, IPs elásticos ociosos).
- **AWS Cost Anomaly Detection:** ML que detecta **picos anômalos** de gasto e alerta.

---

## 4. Estratégias de otimização (FinOps na prática)

### 4.1 As alavancas principais
1. **Rightsizing:** ajustar o tamanho ao uso real (Compute Optimizer). Não pagar por capacidade ociosa.
2. **Elasticidade / desligamento:** Auto Scaling para acompanhar demanda; **desligar dev/test fora do horário comercial** (scheduler).
3. **Modelo de compra certo:** baseline em Savings Plans/RI, picos em On-Demand, batch em Spot.
4. **Serverless para custo variável:** Lambda/Fargate cobram só pelo uso — ótimo para cargas intermitentes (evita pagar servidor ocioso).
5. **Storage tiering:** lifecycle S3, escolher a classe certa, limpar snapshots/volumes órfãos.
6. **Reduzir data transfer:** VPC Endpoints, CloudFront, manter tráfego na mesma AZ quando possível.
7. **Governança contínua:** tags obrigatórias, budgets, revisões periódicas de custo.

### 4.2 O ciclo FinOps (FinOps Foundation)

```
INFORM  →  OPTIMIZE  →  OPERATE  →  (loop)
(visibilidade,   (rightsizing,     (governança,
 tags, showback)  compras, tiers)   automação, cultura)
```

- **Inform:** dar visibilidade (Cost Explorer, tags, dashboards) — todos veem o custo do que criam.
- **Optimize:** agir (rightsizing, Savings Plans, Spot, lifecycle).
- **Operate:** institucionalizar (budgets, políticas, anomaly detection, responsabilização por time).

> [!IMPORTANT] Princípio FinOps central
> **Custo é responsabilidade de todos (engenharia + finanças + negócio), decisões são orientadas por dados de negócio, e times donos do gasto são responsáveis pelo seu uso.**

---

## 🧭 Tabela de Decisão Rápida

| Situação | Escolha |
|---|---|
| Carga 24/7 previsível | **Savings Plans** (ou RI) |
| Pico curto e imprevisível | **On-Demand** |
| Batch tolerante a interrupção | **Spot** |
| Padrão de acesso a objetos desconhecido | **S3 Intelligent-Tiering** |
| Arquivamento de longuíssimo prazo | **S3 Glacier Deep Archive** |
| Alerta ao estourar orçamento | **AWS Budgets (+ Actions)** |
| Atribuir custo por time/projeto | **Cost Allocation Tags** + Cost Explorer |
| Análise por recurso, nível de linha | **Cost and Usage Report (CUR)** |
| Detectar gasto anômalo automaticamente | **Cost Anomaly Detection** |
| Recomendação de tamanho ideal | **Compute Optimizer** |
| Várias contas, fatura única, descontos agregados | **Organizations + Consolidated Billing** |

---

## 📖 Glossário

- **On-Demand / RI / Savings Plans / Spot:** modelos de compra de compute (sem compromisso → comprometido → oportunístico).
- **Data transfer OUT:** custo de saída de dados; IN geralmente grátis.
- **Cross-AZ / Cross-Region:** transferência entre AZs/Regions, ambas pagas.
- **S3 storage classes:** níveis de armazenamento com trade-off custo × latência de recuperação.
- **Lifecycle policy:** regra de transição/expiração automática de objetos S3.
- **Cost Explorer:** visualização histórica e previsão de custos.
- **AWS Budgets:** orçamentos com alertas e ações automáticas.
- **CUR:** Cost and Usage Report, o relatório mais granular.
- **Cost allocation tags:** tags para atribuir custo por dimensão de negócio.
- **Consolidated Billing:** fatura única para várias contas via Organizations.
- **Showback / Chargeback:** mostrar / cobrar o custo do time.
- **Rightsizing:** ajustar capacidade ao uso real.
- **FinOps:** disciplina que une engenharia e finanças para gestão de custo na nuvem.

---

## ✅ Active Recall (responda sem olhar)

1. Quais os 3 eixos que compõem quase toda precificação AWS? Qual é o mais "oculto"?
2. Diferencie Savings Plans (Compute vs EC2 Instance) e Reserved Instances.
3. Quando Spot é apropriado e quando é proibido? Qual o aviso de interrupção?
4. Data transfer: o que é grátis e o que é pago (IN, OUT, cross-AZ, cross-Region)?
5. Um cliente não sabe o padrão de acesso dos objetos. Qual classe S3 recomendar e por quê?
6. Diferencie Cost Explorer, Budgets e CUR — quando usar cada um?
7. O que são cost allocation tags e como habilitam showback/chargeback?
8. Descreva o ciclo Inform → Optimize → Operate do FinOps com um exemplo de ação em cada fase.
9. Como o Consolidated Billing via Organizations reduz custo?

---

## 🔗 Fontes e leitura oficial

- [AWS Billing User Guide](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html)
- [Well-Architected — Cost Optimization Pillar](https://docs.aws.amazon.com/wellarchitected/latest/cost-optimization-pillar/welcome.html)
- [EC2 Pricing / Purchase Options](https://aws.amazon.com/ec2/pricing/) · [Savings Plans](https://aws.amazon.com/savingsplans/) · [Spot](https://aws.amazon.com/ec2/spot/)
- [S3 Storage Classes](https://aws.amazon.com/s3/storage-classes/)
- [Cost Explorer](https://aws.amazon.com/aws-cost-management/aws-cost-explorer/) · [Budgets](https://aws.amazon.com/aws-cost-management/aws-budgets/) · [CUR](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html)
- [Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
- [FinOps Foundation](https://www.finops.org/framework/)
