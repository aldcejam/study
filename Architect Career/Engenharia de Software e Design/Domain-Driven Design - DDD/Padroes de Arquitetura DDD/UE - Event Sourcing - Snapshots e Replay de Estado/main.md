---
tema: Event Sourcing - Snapshots e Replay de Estado
tipo: unidade-estudo
tags: [arquitetura, ddd, distribuídos, logs]
---
# 🧪 UE - Event Sourcing - Snapshots e Replay de Estado

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Bancos de dados tradicionais guardam apenas o "estado atual". Se você tem R$ 100 na conta, o banco só sabe isso. Mas como você chegou lá? O problema é a perda do histórico e da intenção do usuário. **Event Sourcing** resolve isso guardando todos os eventos que mudaram o estado (ex: `DepositoRealizado`, `SaqueRealizado`). O estado atual é apenas o resultado de um **Replay** desses eventos. **Snapshots** resolvem o problema de performance ao carregar agregados com milhões de eventos. Sem Event Sourcing, você não tem uma trilha de auditoria perfeita nem a capacidade de "viajar no tempo" para debugar problemas do passado.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Append-Only Store]:** Por que o banco de eventos nunca deve ter `UPDATE` ou `DELETE`.
2. **[Replay de Estado]:** O processo de reconstruir o agregado a partir do log de eventos.
3. **[Snapshots]:** Como salvar o estado periódico (ex: a cada 100 eventos) para acelerar o carregamento.
4. **[Versioning e Upcasting]:** Como lidar com mudanças no formato dos eventos ao longo do tempo.
5. **[Consistencia Eventual e Projecoes]:** Como criar modelos de leitura a partir do stream de eventos.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Exploring CQRS and Event Sourcing* (Microsoft Patterns & Practices).
- **System Blueprint:** Uso de Event Sourcing em sistemas bancários, blockchains e motores de jogos.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um agregado de `ContaBancaria` que não tem um campo `saldo`, mas sim uma lista de eventos. Implementar o método `get_balance()` que faz o replay.
- [ ] Criar a classe `Event` e os tipos de eventos.
- [ ] Implementar a lógica de `apply(event)` no agregado.
- [ ] Implementar uma função de `snapshot` que salva o saldo atual e o último ID de evento processado.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Log de Eventos como Fonte da Verdade - Teoria e Fundamentos]]
- [[./Arquitetura de Event Stores e Estratégias de Snapshots - Funcionamento Interno e Arquitetura]]
- [[./Evolucao de Esquemas de Eventos (Upcasting) - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Event Sourcing com Kafka e Postgres - Implementacao de Referencia e Benchmarks]]
