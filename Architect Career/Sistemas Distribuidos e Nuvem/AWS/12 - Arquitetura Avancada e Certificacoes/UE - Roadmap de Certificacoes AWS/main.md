---
tema: Roadmap de Certificacoes AWS
tipo: unidade-estudo
tags: [aws, certificacao, roadmap, saa, sap, devops]
---
# 🧪 UE - Roadmap de Certificacoes AWS

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Certificação não é o objetivo — é a prova externa de que o raciocínio de arquiteto foi internalizado. As provas Professional (SAP-C02, DOP-C02) não testam memorização, mas julgamento sob ambiguidade: cenários longos onde várias respostas "funcionam" e apenas uma é ótima para o requisito destacado. Estudar sem estratégia (só decorando serviços) reprova; estudar mapeando cada domínio da prova aos blueprints práticos aprova. Esta UE é o plano de ataque.

## 🧬 Grade Atômica de Tópicos
1. **Trilha e pré-requisitos:** Ordem recomendada — SAA-C03 (base) → SAP-C02 + DOP-C02 (alvos) → Specialties (ANS/SCS/DEA). Peso de cada domínio no blueprint oficial.
2. **Técnica de prova (cenários):** Como ler questões longas, eliminar distratores, identificar a palavra-chave do requisito (custo? RTO? menor esforço operacional?) e o "least X" que define a resposta.
3. **Preparação ativa:** Simulados cronometrados, revisão dos erros por domínio, hands-on obrigatório (free tier), fichamento dos whitepapers-chave (Well-Architected, DR, Security).
4. **Manutenção do conhecimento:** Recertificação, acompanhamento de novidades (What's New/re:Invent), releitura espaçada das UEs desta trilha.

## 📜 Fronteira Acadêmica e Referências
- **Blueprints oficiais:** Exam Guides de [SAP-C02](https://d1.awsstatic.com/training-and-certification/docs-sa-pro/AWS-Certified-Solutions-Architect-Professional_Exam-Guide.pdf) e [DOP-C02](https://d1.awsstatic.com/training-and-certification/docs-devops-pro/AWS-Certified-DevOps-Engineer-Professional_Exam-Guide.pdf).
- **Prática:** AWS Skill Builder (official practice exams), whitepapers e a Builders' Library.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Construir e executar um plano de estudo mensurável.
- [ ] Montar um cronograma mapeando cada domínio da prova alvo às categorias 01–12 desta trilha.
- [ ] Fazer um simulado diagnóstico e registrar os domínios fracos.
- [ ] Ciclo semanal: 1 domínio → hands-on → simulado → revisão de erros → nota heutagógica.

## 🗺️ Plano por Certificação
- **SAA-C03 (Associate):** Categorias 01–08. Foco em escolher o serviço certo e arquiteturas resilientes/econômicas simples.
- **SAP-C02 (SA Professional):** Todas as categorias, ênfase em 12 (DR, custo, migração, multi-conta) e cenários complexos multi-serviço.
- **DOP-C02 (DevOps Professional):** Ênfase em 09 (CI/CD, IaC), 08 (observabilidade/governança), 03 (containers/serverless) e automação de resiliência.
- **ANS-C01 (Advanced Networking):** Aprofundar 02 (TGW, DX, PrivateLink, híbrido, roteamento).
- **SCS-C02 (Security):** Aprofundar 06 (IAM, KMS, detecção, resposta a incidentes).
- **DEA-C01 (Data Engineer):** Aprofundar 10 (ingestão, data lake, ETL, warehouse).

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Trilha e Pesos dos Dominios]]
- [[./02 - Tecnica de Prova em Cenarios]]
- [[./03 - Plano de Estudo Ativo e Simulados]]
