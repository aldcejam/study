---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, integracao, mensageria, sqs, eventbridge]
---
# 🗂️ 07 - Integracao de Aplicacoes e Mensageria

## 🎯 Escopo da Área
Sistemas distribuídos resilientes são desacoplados: componentes se comunicam por mensagens e eventos em vez de chamadas síncronas frágeis. A AWS oferece filas (SQS), pub/sub (SNS), barramento de eventos (EventBridge), orquestração de workflows (Step Functions) e fachadas de API (API Gateway/AppSync). O arquiteto usa esses serviços para absorver picos, isolar falhas e permitir evolução independente — aplicando padrões como Saga, Outbox e fan-out. Conecta diretamente com [[../../Comunicacao e Mensageria/main|Comunicacao e Mensageria]].

## 🗺️ Mapa de Exploração
- [[./UE - Mensageria - SQS e SNS/main|UE - Mensageria - SQS e SNS]] -> Filas, pub/sub, fan-out, DLQ, FIFO vs Standard.
- [[./UE - Orquestracao e Eventos - Step Functions e EventBridge/main|UE - Orquestracao e Eventos - Step Functions e EventBridge]] -> Workflows, event bus, schema registry.
- [[./UE - Fachadas de API - API Gateway e AppSync/main|UE - Fachadas de API - API Gateway e AppSync]] -> REST/HTTP/WebSocket APIs, GraphQL, throttling, auth.
