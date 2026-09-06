---
tema: Fachadas de API - API Gateway e AppSync
tipo: unidade-estudo
tags: [aws, api-gateway, appsync, graphql, api]
---
# 🧪 UE - Fachadas de API - API Gateway e AppSync

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Expor serviços diretamente à internet sem uma fachada gerenciada significa reimplementar autenticação, throttling, cache, versionamento e proteção em cada serviço. API Gateway e AppSync centralizam essas preocupações transversais. Escolher o tipo errado (REST vs HTTP API vs WebSocket, ou REST vs GraphQL) gera custo desnecessário, latência extra ou limitações funcionais difíceis de contornar depois de publicado. Conecta com [[../../../Comunicacao e Mensageria/UE - Design de APIs - GraphQL e API Gateway/main|Design de APIs]].

## 🧬 Grade Atômica de Tópicos
1. **Tipos de API Gateway:** REST API (features completas, cache, WAF), HTTP API (mais barato/rápido, menos features), WebSocket API (bidirecional); quando usar cada.
2. **Preocupações transversais:** Authorizers (IAM, Cognito, Lambda authorizer), throttling/usage plans/API keys, request/response mapping, stages e canary deployments.
3. **AppSync (GraphQL gerenciado):** Resolvers, data sources (DynamoDB/Lambda/RDS), subscriptions em tempo real, resolução de over/under-fetching.
4. **Integração e performance:** Integração com Lambda/serviços, VPC Link para backends privados, caching, edge-optimized vs regional vs private endpoints.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html) e [AWS AppSync](https://docs.aws.amazon.com/appsync/latest/devguide/what-is-appsync.html).
- **System Blueprint:** API serverless: API Gateway (HTTP API) + Cognito authorizer + Lambda + DynamoDB; frontend mobile com AppSync + subscriptions.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Publicar uma API segura e com throttling.
- [ ] Criar uma HTTP API integrada a uma Lambda, protegida por Cognito JWT authorizer.
- [ ] Configurar throttling/usage plan e testar o rate limit.
- [ ] (Opcional) Modelar a mesma API em AppSync (GraphQL) com resolver para DynamoDB.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Tipos de API Gateway - REST, HTTP e WebSocket]]
- [[./02 - Authorizers, Throttling e Stages]]
- [[./03 - AppSync e GraphQL Gerenciado]]
