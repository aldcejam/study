---
tema: Algoritmo Multi-Paxos e Otimizacoes
tipo: unidade-estudo
tags: [distribuídos, consenso, algoritmos, arquitetura]
---
# 🧪 UE - Algoritmo Multi-Paxos e Otimizacoes

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como fazer com que um grupo de computadores concorde com um valor único mesmo se alguns deles falharem ou a rede oscilar? O **Paxos** foi o primeiro algoritmo de consenso provado correto. O problema é que o Paxos básico é lento (exige dois rounds de mensagens para cada valor). O **Multi-Paxos** resolve isso elegendo um líder estável que pode propor múltiplos valores com apenas um round de mensagens. Sem entender Paxos, você não entende a base sobre a qual quase todos os sistemas de alta disponibilidade (como Chubby, Zookeeper e Spanner) foram construídos.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Papeis do Paxos]:** Proposer, Acceptor, Learner e Client.
2. **[As Duas Fases (Prepare e Accept)]:** O uso de números de proposta para garantir que propostas antigas sejam ignoradas.
3. **[Quorum]:** Por que precisamos de uma maioria ($N/2 + 1$) para garantir que o consenso seja durável.
4. **[Multi-Paxos]:** A otimização de eleição de líder para evitar a fase de Prepare repetidamente.
5. **[Otimizacoes Modernas]:** Fast Paxos (redução de latência) e EPaxos (consenso sem líder).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Paxos Made Simple* (Leslie Lamport, 2001).
- **System Blueprint:** O serviço **Chubby** do Google e o motor de replicação do Microsoft SQL Server.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar o algoritmo de Paxos Básico (fases Prepare e Accept) em um simulador de rede que pode perder ou atrasar mensagens.
- [ ] Definir a estrutura de um `Acceptor` com seu `promised_id` e `accepted_value`.
- [ ] Simular o conflito de dois `Proposers` tentando propor valores diferentes simultaneamente.
- [ ] Demonstrar como o algoritmo garante que apenas um valor seja finalmente aceito por um quorum.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Fase de Preparacao e Promessas de Maioria - Teoria e Fundamentos]]
- [[./Lideranca Estavel e Otimizacao de Performance em Multi-Paxos - Funcionamento Interno e Arquitetura]]
- [[./O Problema de Livelock (Dueling Proposers) - Casos de Falha e Analise Amortizada]]
- [[./Simulador de Consenso Paxos com Perda de Pacotes - Implementacao de Referencia e Benchmarks]]
