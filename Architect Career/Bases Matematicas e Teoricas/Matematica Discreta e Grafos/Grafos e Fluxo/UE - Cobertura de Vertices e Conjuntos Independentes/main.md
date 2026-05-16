---
tema: Cobertura de Vertices e Conjuntos Independentes
tipo: unidade-estudo
tags: [matematica, algoritmos, grafos, complexidade]
---
# 🧪 UE - Cobertura de Vertices e Conjuntos Independentes

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Qual é o menor número de pontos necessários para monitorar todas as conexões em uma rede? Ou qual o maior número de tarefas que podem ser executadas sem conflitos? Estes são os problemas de **Vertex Cover** e **Independent Set**. Eles são fundamentais para otimização de infraestrutura e detecção de colisões. No entanto, em grafos gerais, eles são **NP-Completos**, exigindo que o engenheiro conheça aproximações ou casos especiais (como grafos bipartidos) para não travar o sistema.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Definicao de Vertex Cover]:** O conjunto mínimo de vértices tal que toda aresta é incidente a pelo menos um deles.
2. **[Conjuntos Independentes]:** O subconjunto máximo de vértices onde nenhum par é adjacente.
3. **[Dualidade de Gallai]:** A relação fundamental: $VertexCover + IndependentSet = TotalVertices$.
4. **[Teorema de König]:** Em grafos bipartidos, o tamanho do matching máximo é igual ao tamanho do vertex cover mínimo.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Combinatorial Optimization: Algorithms and Complexity* (Papadimitriou and Steiglitz).
- **System Blueprint:** Alocação de registros em compiladores e posicionamento de sensores em redes IoT.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Resolver o problema de Vertex Cover Mínimo em um grafo bipartido usando a redução para Matching Máximo (Algoritmo de Hopcroft-Karp).
- [ ] Construir o grafo bipartido.
- [ ] Rodar o algoritmo de matching.
- [ ] Derivar o Vertex Cover a partir do matching usando a construção de König.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Complementaridade e Dualidade de Gallai - Teoria e Fundamentos]]
- [[./Construcao de Konig para Grafos Bipartidos - Funcionamento Interno e Arquitetura]]
- [[./Inviabilidade em Grafos Gerais e Aproximacoes - Casos de Falha e Analise Amortizada]]
- [[./Otimizacao de Localizacao de Sensores - Implementacao de Referencia e Benchmarks]]
