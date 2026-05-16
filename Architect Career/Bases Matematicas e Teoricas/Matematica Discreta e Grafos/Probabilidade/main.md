---
tema: Probabilidade e Analise Estocastica em Sistemas
tipo: unidade-estudo
tags: [matematica, probabilidade, algoritmos, performance]
---
# 🧪 UE - Probabilidade e Analise Estocastica em Sistemas

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Sistemas reais não são determinísticos. A carga de usuários, a latência de rede e a ocorrência de falhas seguem distribuições probabilísticas. Sem dominar a **Probabilidade**, um engenheiro não consegue prever a performance de cauda (P99), projetar sistemas de carga balanceada eficientes ou entender por que um algoritmo randomizado (como o Quicksort ou tabelas Hash) pode degradar para tempo catastrófico em cenários raros. Ignorar a estatística leva a infraestruturas subdimensionadas ou superdimensionadas e caras.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Espacos Amostrais e Variaveis Aleatorias]:** Definições básicas e a distinção entre variáveis discretas e contínuas.
2. **[Distribuicoes de Probabilidade Cruciais]:** Poisson (chegada de pacotes/usuários), Exponencial (tempo entre falhas) e Normal (Central Limit Theorem).
3. **[Esperanca Matematica e Variancia]:** O que esperar na média e quão imprevisível o sistema pode ser.
4. **[Cadeias de Markov e Processos Estocasticos]:** Modelagem de sistemas que evoluem no tempo, essencial para entender filas e estados de protocolos.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Probability and Computing: Randomized Algorithms and Probabilistic Analysis* (Mitzenmacher and Upfal).
- **System Blueprint:** Análise de tráfego de rede em roteadores e sistemas de monitoramento que usam amostragem (Sampling) para reduzir overhead.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Criar um simulador de fila $M/M/1$ (chegadas de Poisson, tempo de serviço exponencial) e medir como a latência explode conforme a utilização da CPU se aproxima de 100%.
- [ ] Gerar eventos de chegada usando uma distribuição de Poisson.
- [ ] Implementar uma fila simples de processamento.
- [ ] Plotar um gráfico de Utilização vs Latência Média para visualizar a curva de saturação.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Variaveis Aleatorias e Distribuições Fundamentais - Teoria e Fundamentos]]
- [[./Teorema do Limite Central e Amostragem - Funcionamento Interno e Arquitetura]]
- [[./Analise de Cauda e Outliers P99 - Casos de Falha e Analise Amortizada]]
- [[./Simulador de Filas e Teoria das Filas - Implementacao de Referencia e Benchmarks]]
