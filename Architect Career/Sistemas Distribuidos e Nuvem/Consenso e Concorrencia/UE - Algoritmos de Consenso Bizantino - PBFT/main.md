---
tema: Algoritmos de Consenso Bizantino - PBFT
tipo: unidade-estudo
tags: [distribuídos, consenso, segurança, blockchain]
---
# 🧪 UE - Algoritmos de Consenso Bizantino - PBFT

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Paxos e Raft assumem que os nós falham (param de responder), mas que eles são honestos (não mentem). E se um nó for hackeado e começar a enviar informações falsas para confundir o sistema? O problema é o **Problema dos Generais Bizantinos**. O **PBFT (Practical Byzantine Fault Tolerance)** resolve isso permitindo que o sistema chegue a um consenso correto mesmo que até $1/3$ dos nós sejam maliciosos ou falhem de forma bizantina. Sem consenso bizantino, não existiriam criptomoedas ou sistemas de controle críticos imunes a intrusões.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Faltas Bizantinas vs Crash Faults]:** A diferença entre um nó que cai e um nó que mente.
2. **[O Limite de 3f+1]:** Por que precisamos de $3f+1$ nós para tolerar $f$ falhas bizantinas.
3. **[Fases do PBFT]:** Pre-prepare, Prepare e Commit. O uso de assinaturas digitais para autenticar mensagens.
4. **[Troca de Visao (View Change)]:** Como o sistema detecta que o líder é malicioso e o substitui.
5. **[Consenso de Alta Performance (BFT)]:** Evoluções como HotStuff (usado no Libra/Diem) para melhorar a escalabilidade.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Practical Byzantine Fault Tolerance* (Miguel Castro and Barbara Liskov, 1999).
- **System Blueprint:** Uso de variações de BFT em blockchains como Hyperledger Fabric e Cosmos (Tendermint).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular um cenário bizantino onde um nó tenta enviar valores diferentes para diferentes vizinhos e mostrar como o protocolo de 3 fases do PBFT detecta e ignora essa inconsistência.
- [ ] Implementar a verificação de assinaturas digitais (simulada).
- [ ] Simular o envio de mensagens conflitantes por um nó malicioso.
- [ ] Mostrar que os nós honestos só entram na fase de Commit se receberem mensagens consistentes de uma maioria qualificada ($2f+1$).

## 🗃️ Notas Heutagógicas Atômicas
- [[./Problema dos Generais Bizantinos e Provas de Impossibilidade - Teoria e Fundamentos]]
- [[./Arquitetura de Mensagens de Tres Fases do PBFT - Funcionamento Interno e Arquitetura]]
- [[./Latencia e Sobrecarga de Rede em Protocolos BFT - Casos de Falha e Analise Amortizada]]
- [[./Comparativo PBFT vs Proof of Work - Implementacao de Referencia e Benchmarks]]
