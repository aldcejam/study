---
tema: Deteccao, Protecao e Resposta a Ameacas
tipo: unidade-estudo
tags: [aws, guardduty, waf, shield, security-hub, seguranca]
---
# 🧪 UE - Deteccao, Protecao e Resposta

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Controle de acesso previne, mas não é suficiente — é preciso detectar comportamento anômalo, proteger o perímetro de aplicações web e responder a incidentes automaticamente. Sem detecção contínua, um comprometimento pode passar meses despercebido. O desafio é orquestrar dezenas de serviços de segurança numa visão unificada (Security Hub) e automatizar a resposta, em vez de afogar o time em alertas isolados.

## 🧬 Grade Atômica de Tópicos
1. **Detecção de ameaças (GuardDuty):** Análise de VPC Flow Logs, DNS logs e CloudTrail via ML para detectar atividade maliciosa; findings e integração com resposta automática.
2. **Postura e visão unificada (Security Hub, Config, Inspector):** Security Hub agrega findings e checa contra padrões (CIS/PCI); Inspector escaneia vulnerabilidades de EC2/ECR/Lambda; Config avalia conformidade contínua.
3. **Proteção de aplicação (WAF, Shield, Firewall Manager):** WAF (regras L7, managed rule groups, rate limiting), Shield Standard/Advanced (DDoS), Network Firewall, Firewall Manager (governança central).
4. **Proteção de dados e resposta (Macie, Detective, automação):** Macie (descoberta de PII no S3), Detective (investigação), EventBridge + Lambda para remediação automática.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon GuardDuty](https://docs.aws.amazon.com/guardduty/latest/ug/what-is-guardduty.html) e [AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/what-is-aws-waf.html).
- **System Blueprint:** Pipeline de segurança: GuardDuty finding → EventBridge → Lambda de remediação + notificação no Security Hub.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Detectar e responder automaticamente.
- [ ] Ativar GuardDuty e gerar findings de teste (sample findings).
- [ ] Criar uma regra WAF com rate limiting num ALB e testar bloqueio.
- [ ] Ligar um finding do GuardDuty a uma Lambda de remediação via EventBridge.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - GuardDuty e Deteccao de Ameacas]]
- [[./02 - Security Hub, Config e Inspector]]
- [[./03 - WAF, Shield e Remediacao Automatica]]
