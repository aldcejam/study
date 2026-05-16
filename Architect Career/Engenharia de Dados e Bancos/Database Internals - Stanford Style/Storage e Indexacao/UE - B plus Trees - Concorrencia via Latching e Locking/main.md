---
tema: B plus Trees - Concorrencia via Latching e Locking
tipo: unidade-estudo
tags: [bancos-de-dados, internals, algoritmos, concorrência]
---
# 🧪 UE - B plus Trees - Concorrencia via Latching e Locking

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> A **B+Tree** é a estrutura de dados dominante para indexação em disco. O problema surge quando centenas de threads tentam ler e escrever na mesma árvore simultaneamente. Se uma thread está dividindo um nó enquanto outra está lendo, o banco de dados corrompe. **Latching** (trancas leves de hardware) e **Crabbing/Coupling** (técnicas de trancamento de nós) resolvem o problema de garantir que a estrutura da árvore permaneça íntegra sob alta concorrência. Sem isso, seu banco de dados não escala em máquinas multicore.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Anatomia da B+Tree]:** Por que manter dados apenas nas folhas e ponteiros nos nós internos.
2. **[Latches vs Locks]:** A distinção crucial entre proteger a estrutura física (Latch) e os dados lógicos (Lock).
3. **[Latch Crabbing/Coupling]:** A técnica de adquirir o latch do filho antes de soltar o do pai para garantir travessias seguras.
4. **[Split e Merge Concorrente]:** Como lidar com mudanças estruturais sem bloquear a árvore inteira.
5. **[Otimizacao B-Link Tree]:** O uso de ponteiros laterais para permitir leituras sem travar os nós pais durante splits.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Efficient Locking for Concurrent Operations on B-Trees* (Lehman and Yao, 1981).
- **System Blueprint:** Implementação de índices no InnoDB (MySQL) e Postgres.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar uma B+Tree simples em memória e adicionar RW-Latches nos nós, simulando múltiplas threads inserindo dados e verificando se a integridade da árvore se mantém.
- [ ] Implementar a busca e inserção básica.
- [ ] Adicionar `std::shared_mutex` em cada nó.
- [ ] Implementar a estratégia de "Crabbing" para inserção.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Estrutura de Nos e Localidade de Cache - Teoria e Fundamentos]]
- [[./Algoritmos de Latch Crabbing e B-Link Trees - Funcionamento Interno e Arquitetura]]
- [[./Contencao no No Raiz e Solucoes de Escalonamento - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Indices em Alta Concorrencia - Implementacao de Referencia e Benchmarks]]
