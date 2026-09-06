---
tema: Identidade Base e Estrutura de Conta AWS
tipo: unidade-estudo
tags: [aws, iam, seguranca, conta, fundamentos]
---
# 🧪 UE - Identidade Base e Estrutura de Conta

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> A causa raiz da maioria dos incidentes de segurança na nuvem não é uma falha da AWS, mas uma permissão excessiva concedida pelo cliente. Usar a conta root para o dia a dia, chaves de acesso de longa duração espalhadas e políticas com `*:*` transformam qualquer vazamento de credencial em comprometimento total. Entender IAM desde a base (identidades, políticas, avaliação de permissões) é pré-requisito para não desenhar sistemas inseguros por padrão.

## 🧬 Grade Atômica de Tópicos
1. **Identidades (root, users, groups, roles):** Por que a conta root é única e deve ser protegida com MFA e travada; diferença entre credenciais de longa duração (users) e temporárias (roles/STS).
2. **Anatomia de uma policy JSON:** Effect, Action, Resource, Condition, Principal. Diferença entre identity-based e resource-based policies.
3. **Lógica de avaliação de permissões:** Deny explícito > Allow explícito > Deny implícito (default). Como múltiplas políticas se combinam.
4. **Princípio do menor privilégio e higiene de conta:** MFA obrigatório, rotação de credenciais, uso de roles em vez de chaves, IAM Access Analyzer.

> [!TIP] Material aprofundado
> Explicação completa de todos os tópicos (baseada na doc oficial da AWS): [[3-Detalhamento|📚 Detalhamento Técnico]]

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [IAM User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction.html) e [Security Pillar — Well-Architected](https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/welcome.html).
- **System Blueprint:** Estratégia de "break-glass account" e federação corporativa via IAM Identity Center (aprofundado na categoria 06).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Configurar uma conta com higiene mínima de segurança e testar avaliação de políticas.
- [ ] Ativar MFA no root, criar um usuário admin e nunca mais usar o root.
- [ ] Escrever uma policy que permite `s3:GetObject` só num bucket específico e testar com `aws sts` / IAM Policy Simulator.
- [ ] Criar uma role assumível e trocar credenciais via `aws sts assume-role`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Identidades - Root, Users, Groups e Roles]]
- [[./02 - Anatomia e Avaliacao de Policies IAM]]
- [[./03 - Menor Privilegio e Higiene de Conta]]
