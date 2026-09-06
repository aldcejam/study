---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, cloud, arquitetura, certificacao, solutions-architect, devops]
---
# 🗂️ AWS - Amazon Web Services

## 🎯 Escopo da Área
A AWS é a nuvem pública de referência do mercado e o campo de batalha onde a teoria de sistemas distribuídos (consenso, particionamento, falha parcial, latência de rede) vira produto gerenciado. Dominar a AWS em nível de arquiteto significa não decorar nomes de serviços, mas entender **os trade-offs por trás de cada serviço**: quando um serviço gerenciado remove um gargalo operacional e quando ele impõe lock-in, limites de escala ou custo oculto. O objetivo desta trilha é formar um arquiteto capaz de desenhar sistemas multi-região, resilientes a falhas de AZ inteiras, seguros por padrão e otimizados em custo.

Esta trilha é estruturada para levar o estudo do zero absoluto até o nível das **certificações mais altas da AWS**: `Solutions Architect – Professional (SAP-C02)` e `DevOps Engineer – Professional (DOP-C02)`, além das *Specialties* (Advanced Networking, Security, Machine Learning, Data Engineer, Database). Cada categoria é um domínio de serviços; cada UE é uma etapa de aprendizado com Core Problem, grade atômica e sandbox prática.

A ordem sugerida respeita dependências: primeiro os fundamentos (infra global, IAM, Well-Architected), depois os pilares técnicos (rede, computação, armazenamento, dados), depois as camadas transversais (segurança, integração, observabilidade, DevOps) e, por fim, a síntese arquitetural (alta disponibilidade, DR, custo, certificação).

## 🗺️ Mapa de Exploração
- [[./01 - Fundamentos e Well-Architected/main|01 - Fundamentos e Well-Architected]] -> Infra global, conta, IAM base, Well-Architected e billing.
- [[./02 - Redes e Content Delivery/main|02 - Redes e Content Delivery]] -> VPC, DNS, CDN, conectividade híbrida (base para Advanced Networking).
- [[./03 - Computacao/main|03 - Computacao]] -> EC2, Auto Scaling, Load Balancing, containers e serverless.
- [[./04 - Armazenamento/main|04 - Armazenamento]] -> S3, EBS, EFS, FSx, backup e transferência.
- [[./05 - Bancos de Dados/main|05 - Bancos de Dados]] -> RDS/Aurora, DynamoDB, caching e bancos especializados (base para Database).
- [[./06 - Seguranca Identidade e Compliance/main|06 - Seguranca, Identidade e Compliance]] -> IAM avançado, KMS, detecção e proteção (base para Security).
- [[./07 - Integracao de Aplicacoes e Mensageria/main|07 - Integracao de Aplicacoes e Mensageria]] -> SQS, SNS, EventBridge, Step Functions, API Gateway.
- [[./08 - Observabilidade e Governanca/main|08 - Observabilidade e Governanca]] -> CloudWatch, X-Ray, CloudTrail, Config, Organizations.
- [[./09 - DevOps e IaC/main|09 - DevOps e IaC]] -> CloudFormation, CDK, CodePipeline e estratégias de deploy (base para DevOps Pro).
- [[./10 - Analytics e Big Data/main|10 - Analytics e Big Data]] -> Kinesis, Glue, Athena, Redshift, EMR (base para Data Engineer).
- [[./11 - Migracao e Transferencia/main|11 - Migracao e Transferencia]] -> 7 Rs, DMS, MGN, Snow Family, DataSync.
- [[./12 - Arquitetura Avancada e Certificacoes/main|12 - Arquitetura Avancada e Certificacoes]] -> Alta disponibilidade, DR, FinOps e roadmap de provas.

## 🎓 Roadmap de Certificações (Ordem de Ataque)
1. **Foundational:** `Cloud Practitioner (CLF-C02)` — opcional; cubra lendo a categoria 01.
2. **Associate (base sólida):** `Solutions Architect – Associate (SAA-C03)` — exige categorias 01 a 08.
3. **Professional (alvo principal):**
   - `Solutions Architect – Professional (SAP-C02)` — síntese de 01–12, foco em 12.
   - `DevOps Engineer – Professional (DOP-C02)` — foco em 09, 08 e 03.
4. **Specialties (aprofundamento):**
   - `Advanced Networking – Specialty (ANS-C01)` — aprofunda 02.
   - `Security – Specialty (SCS-C02)` — aprofunda 06.
   - `Machine Learning / Data Engineer – Associate (DEA-C01)` — aprofunda 10.
   - `Database – Specialty` (descontinuada, conteúdo migrou p/ 05).

## 🧭 Como usar esta trilha
- O `main.md` de cada categoria é um **overview** (o que existe e por quê).
- Cada `UE - ...` é uma **etapa de aprendizado** com blueprint técnico e sandbox.
- As notas heutagógicas atômicas dentro de cada UE devem ser criadas conforme você avança (Active Recall + revisão espaçada).
