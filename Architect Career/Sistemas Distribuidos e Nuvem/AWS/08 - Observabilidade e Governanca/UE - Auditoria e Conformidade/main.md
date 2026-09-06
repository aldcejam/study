---
tema: Auditoria e Conformidade - CloudTrail e Config
tipo: unidade-estudo
tags: [aws, cloudtrail, config, auditoria, compliance]
---
# 🧪 UE - Auditoria e Conformidade

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> "Quem fez o quê, quando e o recurso está em conformidade?" — sem responder a isso, não há auditoria, investigação de incidente nem prova de compliance. CloudTrail registra as chamadas de API (o "log de segurança"); Config registra o estado e a conformidade dos recursos ao longo do tempo. A ausência ou má configuração desses serviços (ex.: trail não organizacional, sem proteção contra tampering) deixa a organização cega em incidentes e reprovada em auditorias.

## 🧬 Grade Atômica de Tópicos
1. **CloudTrail:** Management vs data events, trails organizacionais, entrega para S3 (com Object Lock/KMS), log file integrity validation, CloudTrail Lake.
2. **AWS Config:** Configuration items e histórico, config rules (managed/custom), conformance packs, remediação automática, agregadores multi-conta.
3. **Detecção e resposta baseada em auditoria:** CloudTrail→EventBridge para alertas em ações sensíveis; correlação com GuardDuty/Security Hub.
4. **AWS Health e Trusted Advisor:** Eventos de saúde da conta/serviço, checks de custo/segurança/limites do Trusted Advisor.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS CloudTrail](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html) e [AWS Config](https://docs.aws.amazon.com/config/latest/developerguide/WhatIsConfig.html).
- **System Blueprint:** Trail organizacional centralizado num bucket de log account com Object Lock + Config rules aplicadas a todas as contas via conformance pack.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Auditar e reforçar conformidade.
- [ ] Habilitar CloudTrail e consultar quem criou um recurso específico.
- [ ] Criar uma Config rule (ex.: "S3 não pode ser público") com remediação automática.
- [ ] Disparar um alerta via EventBridge quando um Security Group abrir 0.0.0.0/0 na porta 22.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CloudTrail e Auditoria de API]]
- [[./02 - AWS Config e Conformidade Continua]]
- [[./03 - AWS Health e Trusted Advisor]]
