---
tema: Relogios Logicos - Lamport e Vector Clocks
tipo: unidade-estudo
tags: [distribuídos, matemática, consenso, arquitetura]
---
# 🧪 UE - Relogios Logicos - Lamport e Vector Clocks

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em um sistema distribuído, não existe um "tempo global" confiável. Os relógios de cada máquina divergem (clock drift). O problema é: como saber se o Evento A aconteceu antes do Evento B se eles ocorreram em máquinas diferentes? **Relógios de Lamport** resolvem a ordenação parcial (causalidade básica) incrementando um contador em cada mensagem. **Vector Clocks** resolvem a detecção de conflitos (eventos concorrentes) mantendo um vetor de contadores de todos os nós. Sem relógios lógicos, é impossível garantir a consistência de dados em bancos distribuídos (como Dynamo ou Cassandra).

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Relacao Happened-Before]:** A definição formal de causalidade de Leslie Lamport.
2. **[Algoritmo de Lamport]:** Regras de atualização do contador local e na recepção de mensagens.
3. **[Vector Clocks]:** Como detectar se dois eventos são causais ou se "divergiram" (concorrência).
4. **[Version Vectors]:** A aplicação prática de Vector Clocks para sincronização de estado em sistemas mestre-mestre.
5. **[Relogios Hibridos (HLC)]:** Como combinar o tempo físico (NTP) com o tempo lógico para o Google Spanner e CockroachDB.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Time, Clocks, and the Ordering of Events in a Distributed System* (Leslie Lamport, 1978) - O paper mais citado da computação.
- **System Blueprint:** Uso de Vector Clocks no Amazon Dynamo e Riak.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar uma classe `VectorClock` em Go que suporta as operações de `Increment`, `Send` (gerar timestamp) e `Receive` (fazer merge), e demonstrar a detecção de um conflito.
- [ ] Criar a estrutura de dados (mapa de ID de nó -> contador).
- [ ] Simular três nós trocando mensagens.
- [ ] Gerar dois eventos concorrentes e mostrar que o algoritmo detecta que nenhum "aconteceu antes" do outro.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Causalidade e Ordenacao de Eventos - Teoria e Fundamentos]]
- [[./Algoritmos de Sincronizacao de Vetores de Versao - Funcionamento Interno e Arquitetura]]
- [[./Clock Drift e Limites de NTP - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Relogio Logico Hibrido (HLC) - Implementacao de Referencia e Benchmarks]]
