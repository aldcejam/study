---
tema: Alocacao de Registradores via Graph Coloring
tipo: unidade-estudo
tags: [compiladores, algoritmos, grafos, baixo-nivel]
---
# 🧪 UE - Alocacao de Registradores via Graph Coloring

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> A CPU tem um número minúsculo de registros (ex: 16 no x86-64), mas um programa pode ter centenas de variáveis. O problema é decidir quais variáveis ficam nos registros e quais vão para a memória lenta (RAM). A **Alocação de Registradores** transforma isso em um problema de **Coloração de Grafos**: duas variáveis que estão "vivas" ao mesmo tempo não podem usar o mesmo registro (cor). Se o compilador falhar nisso, a performance do programa cai drasticamente devido ao excesso de acessos à memória (*Spilling*).

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Analise de Liveness (Vividade)]:** Como determinar o intervalo de tempo em que uma variável contém um valor útil.
2. **[Grafo de Interferencia]:** A construção do grafo onde nós são variáveis e arestas representam sobreposição de liveness.
3. **[Algoritmo de Chaitin-Briggs]:** O processo de simplificação do grafo e a atribuição de cores (registros).
4. **[Spilling]:** A lógica de custo para decidir qual variável deve ser "sacrificada" para a memória quando não há registros suficientes.
5. **[Coalescing]:** Como eliminar instruções de cópia inúteis (`MOV`) unindo nós no grafo.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Register Allocation & Spilling via Graph Coloring* (Gregory Chaitin, 1982).
- **System Blueprint:** Alocador de registros do LLVM (`RegAllocGreedy`) e do GCC.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Construir manualmente um Grafo de Interferência para um pequeno trecho de código e tentar colori-lo com 3 cores (registros).
- [ ] Listar as variáveis e seus intervalos de vida.
- [ ] Desenhar o grafo.
- [ ] Aplicar a heurística de coloração e identificar se houve necessidade de Spill.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Liveness Analysis e Dataflow Equations - Teoria e Fundamentos]]
- [[./Heuristica de Simplificacao de Chaitin - Funcionamento Interno e Arquitetura]]
- [[./Register Spilling e Hierarquia de Memoria - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Alocadores de Registradores - Implementacao de Referencia e Benchmarks]]
