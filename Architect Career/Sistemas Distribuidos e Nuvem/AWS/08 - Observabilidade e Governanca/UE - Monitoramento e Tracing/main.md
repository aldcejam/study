---
tema: Monitoramento e Tracing - CloudWatch e X-Ray
tipo: unidade-estudo
tags: [aws, cloudwatch, x-ray, observabilidade, metricas]
---
# 🧪 UE - Monitoramento e Tracing

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em arquiteturas distribuídas, uma requisição atravessa dezenas de serviços; sem métricas agregadas, logs centralizados e tracing distribuído, diagnosticar uma degradação vira adivinhação. CloudWatch e X-Ray são a implementação AWS dos três pilares da observabilidade. Instrumentação insuficiente cega o time em produção; instrumentação demais explode o custo (métricas de alta cardinalidade, logs verbosos). Conecta com [[../../../Observabilidade Plena e DevOps/UE - Metricas e Monitoramento - OpenTelemetry e Prometheus/main|Métricas e OpenTelemetry]].

## 🧬 Grade Atômica de Tópicos
1. **CloudWatch Metrics e Alarms:** Métricas padrão vs custom, dimensões, resolução (standard/high), namespaces, alarms (estados, ações), composite alarms, anomaly detection.
2. **CloudWatch Logs:** Log groups/streams, metric filters, subscription filters (streaming para Kinesis/Lambda), Logs Insights (query), retenção e custo.
3. **X-Ray (tracing distribuído):** Segments/subsegments, service map, trace sampling, integração com Lambda/ECS/API Gateway, correlação com OpenTelemetry (ADOT).
4. **Dashboards e integração:** CloudWatch Dashboards, Container/Lambda Insights, EventBridge para automação baseada em alarme, Amazon Managed Grafana/Prometheus.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon CloudWatch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/WhatIsCloudWatch.html) e [AWS X-Ray](https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html).
- **System Blueprint:** Pipeline de observabilidade com ADOT (AWS Distro for OpenTelemetry) exportando para CloudWatch + X-Ray + Managed Prometheus.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Instrumentar e alertar.
- [ ] Emitir uma métrica custom e criar um alarm que dispara uma ação (SNS/Lambda).
- [ ] Usar Logs Insights para consultar erros e criar um metric filter.
- [ ] Habilitar X-Ray numa cadeia API Gateway→Lambda e ler o service map.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CloudWatch Metrics e Alarms]]
- [[./02 - CloudWatch Logs e Logs Insights]]
- [[./03 - X-Ray e Tracing Distribuido]]
