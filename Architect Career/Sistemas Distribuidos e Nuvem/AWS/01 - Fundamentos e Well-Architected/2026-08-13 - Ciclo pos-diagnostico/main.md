---
tipo: ciclo-pos-diagnostico
status: ativo
data: 2026-08-13
modulo: 01 - Fundamentos e Well-Architected
fonte: [[../UE - Well-Architected Framework/_TEMP-Diagnostico-Nivel-UEs|_TEMP-Diagnostico-Nivel-UEs]]
contexto_pai: [[../main|01 - Fundamentos e Well-Architected]]
instrucao_ia: >
  Ponto de partida de QUALQUER sessao: [[sessao-atual]].
  Fluxo obrigatorio: Onde estou → Tema atual → Explicacao profunda (no sessao-atual).
  Nao recomece a UE do zero se o nivel ja lista o que esta solido.
  Depois da explicacao do tema aberto, faca recall curto so daquele bloco.
  Estilo obrigatorio: exemplos reais, variacoes do mesmo mecanismo, Terraform
  (estrutura de pastas + HCL: trust vs permission vs attachment).
  Atualize [[sessao-atual]] e a secao "Estado do ciclo" quando o aluno fechar um bloco.
  Checklist sequencial: [[progresso]].
---
z# Ciclo 2026-08-13 — pos-diagnostico (Modulo 01)

Documento de direcao para estudo **depois** do diagnostico. Uma IA (e voce) deve abrir **[[sessao-atual]] primeiro**. Este `main` e o mapa de lacunas. Sequencia de checkboxes: [[progresso]].

## Como a IA deve dirigir

1. Abrir [[sessao-atual]]: confirmar **Onde estou** e o **Tema atual**.
2. Ensinar pela **Explicacao profunda** daquele arquivo — um bloco, nao a UE inteira. Sempre com **exemplos reais**, **variacoes** (mesmo mecanismo, Principals/servicos diferentes) e **Terraform** (arvore de arquivos + HCL).
3. Depois da explicacao, 3–6 perguntas de memoria **so desse tema**. Se errar, voltar a secao da explicacao (e ao `*-Detalhamento.md` da UE se precisar de fonte).
4. O que esta em **Solido** nao precisa de aula; so use para conectar.
5. Ao fechar um bloco: atualizar [[sessao-atual]] (onde/tema/explicacao do proximo item) e **Estado do ciclo**.
6. Material: `main` da UE = mapa; `Detalhamento` = fonte. Diagnostico original: [[../UE - Well-Architected Framework/_TEMP-Diagnostico-Nivel-UEs|_TEMP-Diagnostico-Nivel-UEs]].

**Nivel do modulo neste ciclo:** inicial / lacunas grandes. Aproveitamento do diagnostico ~15% (20/79 itens tentados). Nao e zero: ha modelo mental em isolamento, soberania, role na EC2 e trade-off de prova.

---

## Estado do ciclo

| Bloco | Status | Notas |
|---|---|---|
| UE2 Identidade | em andamento | A1 aberto em [[sessao-atual]] — User vs Role vs Group vs STS |
| UE3 Well-Architected | pendente | Vocabulario; julgamento de prova ja existe |
| UE1 restante (borda, escopo, responsabilidade) | pendente | Region/AZ/soberania ja ok |
| UE4 restante (FinOps operacional) | pendente | Modelos de compra ja ok |
| Integracao Q76–79 | pendente | So depois dos quatro blocos |

Sessoes sugeridas: **A = UE2** · **B = UE3 + UE1 restante** · **C = UE4 restante + Q76–79**.

---

## Ordem de estudo (alavancagem, nao numeracao das UEs)

1. [[../UE - Identidade Base e Estrutura de Conta/main|UE2 — Identidade Base]]
2. [[../UE - Well-Architected Framework/main|UE3 — Well-Architected Framework]]
3. [[../UE - Infraestrutura Global e Modelos de Nuvem/main|UE1 — Infraestrutura Global]] (so lacunas)
4. [[../UE - Billing FinOps e Precificacao/main|UE4 — Billing, FinOps e Precificacao]] (so lacunas)

---

## UE 1 — Infraestrutura Global e Modelos de Nuvem

- **Nivel:** intermediario baixo (nota ~31%). Confianca autoavaliada: media — coerente.
- **Material:** [[../UE - Infraestrutura Global e Modelos de Nuvem/main|main]] · [[../UE - Infraestrutura Global e Modelos de Nuvem/1-Detalhamento|detalhamento]]
- **Solido:** hierarquia Region → AZ → DC; Region como limite padrao de dados; AZ name embaralhado vs AZ ID; residencia no pais sem replicacao para fora; escolha Multi-AZ para RPO≈0 na Region (motivo ainda raso).
- **Parcial / corrigir:** Multi-Region e para queda da Region, DR, usuarios globais, compliance — nao “AZ fraca”. Latencia ~1–2 ms **habilita** replicacao sincrona; cross-Region costuma ser assincrona. “Seguranca DA nuvem” inclui hypervisor/instalacoes, nao so predio. Route 53 **nao** e regional.
- **Em branco (abordar):** blast radius (Q4); Edge Location / Regional Edge Cache / CloudFront vs Global Accelerator (Q7–8); Local Zones vs Wavelength vs Outposts (Q9–10, Q14); IaaS/PaaS/SaaS e quem pacha o SO (Q12–13); classificacao zonal/regional/global completa (Q5).

**Ate quando esta UE esta “fechada” neste ciclo:** classificar de memoria EC2, EBS, subnet, S3, DynamoDB, IAM, Route 53, CloudFront; explicar shared responsibility nos tres modelos; escolher LZ / Wavelength / Outposts num cenario de cidade sem Region.

---

## UE 2 — Identidade Base e Estrutura de Conta

- **Nivel:** inicial (nota ~10%). Confianca: nao marcada no diagnostico; tratar como baixa.
- **Material:** [[../UE - Identidade Base e Estrutura de Conta/main|main]] · [[../UE - Identidade Base e Estrutura de Conta/3-Detalhamento|detalhamento]]
- **Solido:** EC2 acessa S3 com instance profile/role, nao access key na instancia (Q32). Intuicao de nao usar root no dia a dia (Q17 incompleta).
- **Errado (corrigir cedo):** IAM User ≠ Role. User = identidade com credencial longa. Role = assumida, credencial temporaria (STS). Policy e o que pode fazer. Group agrupa **users**, **nao** e Principal, **nao** aninha.
- **Em branco (eixo da sessao A):** STS / `AssumeRole` (Q20, Q34); elementos do statement (Q21–24); identity-based vs resource-based vs trust policy; avaliacao Deny explicito > Allow > Deny implicito (Q25–27: nao deleta se houver Deny; nao termina se ninguem Allow); higiene (MFA, sem access key no root, Access Analyzer, Condition de IP) (Q28–31); cross-account via role+trust (Q33).

**Ate quando esta UE esta “fechada” neste ciclo:** explicar AssumeRole; escrever em palavras uma policy `s3:GetObject` + IP corporativo; responder Q26 e Q27 sem hesitar; descrever o padrao conta A → recurso na conta B.

---

## UE 3 — Well-Architected Framework

- **Nivel:** inicial em vocabulario, intermediario em julgamento (nota ~12%). Confianca autoavaliada: baixa — coerente para pilares; Q42–Q44 ja estao no nivel de prova.
- **Material:** [[../UE - Well-Architected Framework/main|main]] · [[../UE - Well-Architected Framework/2-Detalhamento|detalhamento]]
- **Solido:** trade-off performance × consistencia; security × agilidade; quando todas as opcoes “funcionam”, o enunciado escolhe o pilar e ainda pede equilibrio com custo.
- **Em branco:** listar os 6 pilares e o que cada um otimiza (Q35–37); Sustainability como pilar mais recente; 2 servicos/praticas por pilar; stop guessing capacity, game days, principios transversais (Q38–40); gatilhos Q45–Q49; WAFR, HRI/MRI, Tool vs Trusted Advisor, lens (Q50–53); cenario 3 camadas 1 AZ (Q54–55).

**Ate quando esta UE esta “fechada” neste ciclo:** recitar os 6 pilares; mapear cada frase-gatilho a um pilar; distinguir WAFR Tool de Trusted Advisor.

---

## UE 4 — Billing, FinOps e Precificacao

- **Nivel:** inicial-intermediario no eixo de compra, inicial no resto (nota ~13%). Confianca: nao marcada.
- **Material:** [[../UE - Billing FinOps e Precificacao/main|main]] · [[../UE - Billing FinOps e Precificacao/4-Detalhamento|detalhamento]]
- **Solido:** tres eixos (compute, storage, data transfer); On-Demand / RI / Savings Plans / Spot com caso de uso e interrupcao (~2 min no Spot).
- **Parcial:** data transfer — IN da internet em geral gratis; same-AZ privado em geral gratis; cross-AZ e cross-Region pagos; IP publico mesmo na AZ pode cobrar. Savings Plans = compromisso de **USD/hora**, nao “valor mensal” literal.
- **Em branco:** Compute SP vs EC2 Instance SP; Standard vs Convertible RI; layered baseline+pico+Spot (Q59–62, Q73, Q75); NAT Gateway, CloudFront vs S3 OUT, cross-AZ chatty (Q63–65); classes S3, lifecycle, EBS provisionado vs usado (Q66–68, Q74); Cost Explorer / Budgets / CUR, tags, showback vs chargeback, consolidated billing (Q69–72).

**Ate quando esta UE esta “fechada” neste ciclo:** montar mixed purchasing (baseline + pico + batch); escolher Intelligent-Tiering quando o padrao S3 e desconhecido; citar por que NAT e CloudFront aparecem na fatura.

---

## Integracao entre UEs (Q76–79)

Nao comecar por aqui. Quando UE2–UE4 tiverem o minimo acima: e-commerce (Region/AZ, identidade app→S3/DB, 2 trade-offs WAF, modelo de compra); queda de uma AZ; credencial no GitHub; fatura dobrou sem usuarios.

---

## Proxima acao (se a IA nao receber outro pedido)

Abrir [[sessao-atual]]. Bloco **A1** ja tem explicacao profunda. Seguir o recall no final daquele arquivo.
