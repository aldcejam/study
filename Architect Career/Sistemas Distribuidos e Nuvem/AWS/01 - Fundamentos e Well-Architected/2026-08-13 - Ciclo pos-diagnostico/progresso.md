---
tipo: progresso-ciclo
ciclo: [[main]]
data: 2026-08-13
status: ativo
proximo: A1 — User vs Role vs Group vs STS (explicacao em [[sessao-atual]])
instrucao_ia: >
  Ponto de partida: [[sessao-atual]] (onde estou → tema → explicacao profunda).
  Siga a sequencia de cima para baixo. Nao pule bloco.
  Marque [x] so depois de explicacao + active recall + correcao (nao so leitura).
  Atualize `proximo` neste frontmatter, [[sessao-atual]] e a tabela Estado em [[main]].
---
# Progresso — ciclo 2026-08-13

Ordem obrigatoria. Um item aberto por vez. Sessao viva: [[sessao-atual]]. Direcao: [[main]].

**Como marcar:** explicacao profunda no `sessao-atual` → recall no chat → correcao. Leitura sem recall nao conta.

---

## Sessao A — UE2 Identidade

Material: [[../UE - Identidade Base e Estrutura de Conta/main|UE2 main]] · [[../UE - Identidade Base e Estrutura de Conta/3-Detalhamento|detalhamento]]

### A1. User vs Role vs Group vs STS
- [ ] Recall: 5 perguntas (User vs Role; Group/Principal/aninhamento; STS/AssumeRole; roles vs access key longa; EC2→S3)
- [ ] Correcao das Q18–Q20 (User ≠ Role; Group nao e Principal e nao aninha)
- [ ] Criterio: explicar AssumeRole em 3 frases; dizer por que instance profile vence access key na EC2

### A2. Anatomia de policy
- [ ] Recall: Effect, Action, Resource, Condition, Principal; onde Principal aparece
- [ ] Correcao: identity-based vs resource-based vs trust policy (Q21–Q24)
- [ ] Criterio: escrever em palavras policy `s3:GetObject` em `arn:aws:s3:::logs-prod/*` vinda de IP corporativo (Q31)

### A3. Avaliacao de permissoes
- [ ] Recall: Deny explicito, Allow explicito, Deny implicito
- [ ] Correcao Q26 (Allow `s3:*` + Deny `s3:DeleteObject` → nao deleta) e Q27 (silencio em `TerminateInstances` → nao termina)
- [ ] Criterio: responder Q26 e Q27 sem hesitar

### A4. Higiene e cross-account
- [ ] Recall: 5 praticas de higiene; Access Analyzer; padrao conta A → recurso na conta B
- [ ] Correcao Q17 (2 acoes exclusivas do root), Q28–Q30, Q33
- [ ] Criterio: descrever role na conta B + trust na conta A

### A5. Fechar UE2 neste ciclo
- [ ] Recitar os 4 criterios de fechamento da UE2 no [[main]]
- [ ] Atualizar tabela Estado em [[main]]: UE2 → em andamento ou fechada

---

## Sessao B — UE3 + UE1 restante

### B1. Seis pilares (UE3)
Material: [[../UE - Well-Architected Framework/main|UE3 main]] · [[../UE - Well-Architected Framework/2-Detalhamento|detalhamento]]

- [ ] Recall: listar os 6 pilares em uma frase cada; qual foi adicionado mais recentemente
- [ ] Correcao Q35–Q37 (2 servicos/praticas por pilar)
- [ ] Criterio: recitar os 6 sem consultar

### B2. Principios e gatilhos de prova (UE3)
- [ ] Recall: stop guessing capacity; game days; 3 principios transversais
- [ ] Mapear Q45–Q49 (cost-effective, fully managed, must not lose data, encryption/audit, lowest latency global)
- [ ] Criterio: cada frase-gatilho → um pilar, de memoria

### B3. WAFR e ferramentas (UE3)
- [ ] Recall: WAFR, HRI vs MRI; Tool vs Trusted Advisor; o que e lens
- [ ] Correcao Q50–Q53
- [ ] Criterio: distinguir Well-Architected Tool de Trusted Advisor
- [ ] Atualizar [[main]]: UE3 fechada se B1–B3 ok

### B4. Escopo de servicos (UE1 restante)
Material: [[../UE - Infraestrutura Global e Modelos de Nuvem/main|UE1 main]] · [[../UE - Infraestrutura Global e Modelos de Nuvem/1-Detalhamento|detalhamento]]

- [ ] Recall: classificar EC2, EBS, subnet, S3, DynamoDB, IAM, Route 53, CloudFront (zonal / regional / global)
- [ ] Correcao Q5 (Route 53, IAM, CloudFront = globais)
- [ ] Criterio: a lista acima de memoria

### B5. Borda, hibrido, responsabilidade (UE1 restante)
- [ ] Recall: blast radius; Edge vs REC vs Region; CloudFront vs Global Accelerator
- [ ] Recall: Local Zones vs Wavelength vs Outposts; cidade sem Region (Q14)
- [ ] Recall: IaaS/PaaS/SaaS e quem pacha o SO (EC2 / RDS / S3)
- [ ] Correcao Q4, Q7–Q13
- [ ] Criterio: os tres criterios de fechamento da UE1 no [[main]]
- [ ] Atualizar [[main]]: UE1 restante fechada

---

## Sessao C — UE4 restante + integracao

### C1. Variantes de compra (UE4)
Material: [[../UE - Billing FinOps e Precificacao/main|UE4 main]] · [[../UE - Billing FinOps e Precificacao/4-Detalhamento|detalhamento]]

- [ ] Recall: Compute SP vs EC2 Instance SP; Standard vs Convertible RI; layered baseline + pico + Spot
- [ ] Correcao Q59–Q62, Q73, Q75 (batch falivel → Spot)
- [ ] Criterio: montar mixed purchasing em um cenario 24/7 + pico + batch

### C2. Custos ocultos e storage (UE4)
- [ ] Recall: por que NAT surpreende; CloudFront vs S3/EC2 OUT; cross-AZ chatty
- [ ] Recall: Standard vs Intelligent-Tiering vs IA vs Glacier; lifecycle; EBS GB provisionado vs usado
- [ ] Correcao Q63–Q68, Q74
- [ ] Criterio: Intelligent-Tiering quando padrao S3 e desconhecido; NAT e CloudFront na fatura

### C3. Governanca FinOps (UE4)
- [ ] Recall: Cost Explorer vs Budgets vs CUR; tags; showback vs chargeback; consolidated billing
- [ ] Correcao Q69–Q72
- [ ] Criterio: 4 estrategias de otimizacao alem de “comprar RI”
- [ ] Atualizar [[main]]: UE4 restante fechada

### C4. Integracao Q76–Q79
- [ ] Recall: e-commerce (Region/AZ, identidade app→S3/DB, 2 trade-offs WAF, modelo de compra)
- [ ] Recall: uma AZ cai — o que continua se UE1 + Reliability
- [ ] Recall: credencial no GitHub — o que da UE2 limita o dano
- [ ] Recall: fatura dobrou sem usuarios — 3 lugares (UE4) e qual pilar WAF
- [ ] Atualizar [[main]]: Integracao fechada

---

## Fim do ciclo

- [ ] Todos os blocos A–C marcados
- [ ] Tabela Estado em [[main]] sem `pendente`
- [ ] Frontmatter deste arquivo: `status: encerrado`
