---
tema: Identidade de Aplicacao e Federacao
tipo: unidade-estudo
tags: [aws, cognito, identity-center, federacao, sso]
---
# 🧪 UE - Identidade de Aplicacao e Federacao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> IAM controla acesso de identidades AWS (funcionários/serviços), mas não é feito para milhões de usuários finais de uma aplicação nem para o SSO corporativo. Confundir os planos — usar IAM users para usuários de app, ou não federar identidades corporativas — gera arquiteturas inseguras e insustentáveis. É preciso separar identidade de workforce (Identity Center) de identidade de clientes (Cognito).

## 🧬 Grade Atômica de Tópicos
1. **Cognito (customer identity):** User Pools (autenticação, JWT, MFA, social login) vs Identity Pools (troca de token por credenciais AWS temporárias); fluxos OAuth2/OIDC.
2. **IAM Identity Center (workforce SSO):** SSO centralizado multi-conta, permission sets, integração com IdP corporativo (Okta/Entra ID) via SAML/SCIM.
3. **Directory Service:** Managed Microsoft AD, AD Connector, Simple AD; integração com workloads Windows e RDS/FSx.
4. **Federação e padrões de token:** SAML 2.0, OIDC, Web Identity Federation; mapeamento de claims para roles; escopo de sessão.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon Cognito](https://docs.aws.amazon.com/cognito/latest/developerguide/what-is-amazon-cognito.html) e [IAM Identity Center](https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html).
- **System Blueprint:** App SPA + API Gateway autenticada por Cognito User Pool (JWT authorizer); acesso administrativo via Identity Center federado.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Autenticar usuários e trocar por credenciais AWS.
- [ ] Criar um Cognito User Pool e proteger uma API do API Gateway com JWT authorizer.
- [ ] Usar Identity Pool para dar a um usuário logado acesso temporário a um bucket específico.
- [ ] Desenhar quando usaria Identity Center vs Cognito.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Cognito - User Pools e Identity Pools]]
- [[./02 - IAM Identity Center e SSO Corporativo]]
- [[./03 - Directory Service e Federacao]]
