---
tema: Fluxo Maximo - Algoritmo Edmonds-Karp
tipo: unidade-estudo
tags: [matematica, algoritmos, grafos]
---
# 🧪 UE - Fluxo Maximo - Algoritmo Edmonds-Karp

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O problema de fluxo máximo resolve gargalos de **capacidade de transporte** em redes. Seja para otimizar a largura de banda em uma rede de computadores, o fluxo de logística em uma cadeia de suprimentos ou até o balanceamento de carga em sistemas distribuídos, ignorar a teoria de fluxo leva a subutilização de recursos ou saturação catastrófica de links críticos. Sem Edmonds-Karp (ou Ford-Fulkerson), você não consegue garantir o limite teórico de dados que podem atravessar um sistema complexo.

## 📖 Material Teórico Aprofundado
- [[./aula|👉 Acessar Apostila / Aula Completa desta UE]]

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Rede de Fluxo e Capacidade]:** Entender a definição formal de uma rede de fluxo, as leis de conservação e as restrições de capacidade.
2. **[Rede Residual e Caminhos Aumentantes]:** O mecanismo de "voltar atrás" através de arestas residuais para encontrar caminhos adicionais.
3. **[Busca em Largura (BFS) como Diferencial]:** Por que Edmonds-Karp usa BFS em vez de DFS para garantir convergência em tempo polinomial $O(VE^2)$.
4. **[Teorema Max-Flow Min-Cut]:** A dualidade fundamental: o fluxo máximo é limitado pelo corte mínimo da rede.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Introduction to Algorithms* (CLRS) - Capítulo sobre Maximum Flow.
- **System Blueprint:** Otimização de roteamento BGP em backbones de internet e algoritmos de segmentação de imagem em visão computacional (Graph Cuts).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar o algoritmo de Edmonds-Karp em Python ou C++ para resolver um problema de alocação de largura de banda em uma topologia de rede arbitrária.
- [ ] Construir a matriz de adjacência ou lista de capacidades.
- [ ] Implementar a BFS para encontrar o caminho aumentante mais curto.
- [ ] Atualizar as capacidades residuais e calcular o fluxo total acumulado.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Redes de Fluxo e Leis de Conservacao - Teoria e Fundamentos]]
- [[./Algoritmo de BFS e Complexidade de Edmonds-Karp - Funcionamento Interno e Arquitetura]]
- [[./Ciclos de Capacidade Zero e Estagnacao - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Roteamento Otimizado - Implementacao de Referencia e Benchmarks]]
