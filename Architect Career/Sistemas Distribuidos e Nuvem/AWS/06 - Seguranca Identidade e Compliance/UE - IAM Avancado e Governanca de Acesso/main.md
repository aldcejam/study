---
tema: IAM Avancado e Governanca de Acesso
tipo: unidade-estudo
tags: [aws, iam, sts, scp, seguranca, governanca]
---
# 🧪 UE - IAM Avancado e Governanca de Acesso

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em escala corporativa (centenas de contas), o controle de acesso deixa de ser "que política dar a um usuário" e vira um problema de governança em múltiplas camadas que precisam se intersectar corretamente: SCPs no topo (guardrails), permission boundaries, identity e resource policies, sessões e condições. Um erro de raciocínio na avaliação combinada dessas camadas é a origem de acessos negados misteriosos ou, pior, de privilégios excessivos que passam despercebidos.

## 🧬 Grade Atômica de Tópicos
1. **Camadas de avaliação de permissão:** Ordem completa: SCP → permission boundary → session policy → identity policy → resource policy; interseção vs união; deny sempre vence.
2. **STS e credenciais temporárias:** AssumeRole, cross-account access, external ID (confused deputy), role chaining, federação (SAML/OIDC/Web Identity).
3. **Permission boundaries e delegação:** Permitir que times criem roles sem escalar privilégio; padrões de self-service seguro.
4. **SCPs e ABAC:** Service Control Policies via Organizations (guardrails de conta); Attribute-Based Access Control com tags para escalar permissões sem explosão de políticas.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [IAM — Policy Evaluation Logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html).
- **System Blueprint:** Landing zone multi-conta com SCPs de guardrail + roles federadas via IAM Identity Center (ver categoria 08).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Exercitar cross-account e boundaries.
- [ ] Configurar AssumeRole cross-account com external ID entre duas contas.
- [ ] Aplicar uma permission boundary e provar que ela limita um role mesmo com identity policy ampla.
- [ ] Escrever uma policy ABAC baseada em tags e testar no Policy Simulator.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Camadas de Avaliacao - SCP, Boundary, Identity, Resource]]
- [[./02 - STS, AssumeRole e Federacao]]
- [[./03 - Permission Boundaries e ABAC]]
