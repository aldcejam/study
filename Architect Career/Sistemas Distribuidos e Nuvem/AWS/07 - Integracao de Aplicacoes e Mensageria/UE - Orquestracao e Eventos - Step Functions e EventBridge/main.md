---
tema: Orquestracao e Eventos - Step Functions e EventBridge
tipo: unidade-estudo
tags: [aws, step-functions, eventbridge, orquestracao, eventos]
---
# 🧪 UE - Orquestracao e Eventos - Step Functions e EventBridge

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Coordenar múltiplos passos distribuídos (com retries, timeouts, compensação e ramificações) dentro de código de aplicação gera "spaghetti" frágil e sem visibilidade. Orquestração (Step Functions) externaliza a máquina de estados; coreografia por eventos (EventBridge) desacopla produtores e consumidores. Escolher entre orquestração e coreografia — e errar — define se o sistema é observável e evolutivo ou um emaranhado impossível de debugar.

## 🧬 Grade Atômica de Tópicos
1. **Step Functions (orquestração):** State machines (Standard vs Express), estados (Task, Choice, Parallel, Map, Wait), retry/catch, integração direta com serviços (SDK integrations), padrão Saga com compensação.
2. **EventBridge (coreografia):** Event bus (default/custom/partner), regras e padrões de evento, targets, schema registry, archive & replay.
3. **Orquestração vs Coreografia:** Trade-offs de acoplamento, observabilidade, complexidade; quando cada uma é a escolha certa.
4. **EventBridge avançado:** Pipes (source→enrichment→target), Scheduler, integração SaaS, cross-account event routing.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) e [Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html).
- **System Blueprint:** Saga distribuída com Step Functions coordenando serviços via compensação; arquitetura event-driven com EventBridge como espinha dorsal.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Orquestrar um workflow com compensação.
- [ ] Criar uma state machine (Standard) com Choice, Parallel e Retry/Catch chamando Lambdas.
- [ ] Implementar um passo de compensação (padrão Saga) para uma falha simulada.
- [ ] Publicar um evento custom no EventBridge e roteá-lo por padrão para 2 targets.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Step Functions - State Machines e Saga]]
- [[./02 - EventBridge - Bus, Regras e Replay]]
- [[./03 - Orquestracao vs Coreografia]]
