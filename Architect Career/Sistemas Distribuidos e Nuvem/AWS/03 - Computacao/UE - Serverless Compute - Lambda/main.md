---
tema: Serverless Compute - AWS Lambda
tipo: unidade-estudo
tags: [aws, lambda, serverless, computacao]
---
# 🧪 UE - Serverless Compute - Lambda

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Serverless promete escala a zero e custo por uso, mas transfere problemas para o design: cold starts adicionam latência imprevisível, limites de concorrência causam throttling sob pico, e o modelo stateless força repensar conexões de banco (pooling) e estado. Sem entender o modelo de execução (isolamento por microVM Firecracker), arquiteturas serverless falham exatamente quando mais são exigidas. Conecta com [[../../../Cloud Native/UE - Serverless Internals - Cold Starts e Isolate Runtimes/main|Serverless Internals]].

## 🧬 Grade Atômica de Tópicos
1. **Modelo de execução:** Invocação síncrona vs assíncrona vs event source mapping (poll); ambiente de execução, Firecracker microVMs, reutilização de contexto.
2. **Cold starts e otimização:** Causas (init de runtime/deps), Provisioned Concurrency, SnapStart, impacto de VPC/ENI, tamanho do pacote.
3. **Concorrência e limites:** Concorrência reservada vs provisionada, throttling, limites de payload/tempo/memória, relação memória↔CPU.
4. **Integração e padrões:** Triggers (S3, SQS, API Gateway, EventBridge), destinations, DLQ, idempotência, gestão de conexões (RDS Proxy).

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
- **Paper/Blueprint:** [Firecracker: Lightweight Virtualization for Serverless Applications](https://www.usenix.org/conference/nsdi20/presentation/agache) (NSDI 2020) — o motor de isolamento por trás do Lambda/Fargate.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Medir e mitigar cold starts.
- [ ] Criar uma Lambda simples e medir latência de cold vs warm start.
- [ ] Ativar Provisioned Concurrency e comparar; testar throttling com concorrência reservada baixa.
- [ ] Integrar com SQS (event source mapping) e implementar idempotência.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Modelo de Execucao e Invocacao]]
- [[./02 - Cold Starts, Provisioned Concurrency e SnapStart]]
- [[./03 - Concorrencia, Limites e Padroes de Integracao]]
