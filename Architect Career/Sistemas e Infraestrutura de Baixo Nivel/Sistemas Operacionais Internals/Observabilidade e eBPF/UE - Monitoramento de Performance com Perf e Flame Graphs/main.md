---
tema: Monitoramento de Performance com Perf e Flame Graphs
tipo: unidade-estudo
tags: [performance, linux, observabilidade, baixo-nivel]
---
# 🧪 UE - Monitoramento de Performance com Perf e Flame Graphs

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Por que meu servidor está lento? Onde exatamente a CPU está gastando tempo? O problema é que o monitoramento tradicional (`top`, `htop`) é macro demais. O **perf** resolve isso capturando amostras da stack de execução milhares de vezes por segundo. Os **Flame Graphs** resolvem o problema de visualização, transformando milhares de linhas de log em um gráfico intuitivo que mostra os gargalos de código em "chamas". Sem isso, otimizar performance é apenas tentativa e erro baseada em palpites.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Sampling vs Tracing]:** A diferença entre capturar o estado em intervalos regulares e capturar cada evento.
2. **[Hardware Performance Counters (PMU)]:** Como a CPU conta cache misses, branch mispredictions e ciclos de clock internamente.
3. **[Subcomando perf record e report]:** O fluxo de trabalho para coletar e analisar dados de perfil.
4. **[Flame Graphs]:** A hierarquia de chamadas e a largura proporcional ao tempo de CPU.
5. **[Off-CPU Analysis]:** Por que às vezes o problema não é a CPU, mas a espera por I/O ou locks (e como medir isso).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Systems Performance: Enterprise and the Cloud* (Brendan Gregg).
- **System Blueprint:** Metodologia USE (Utilization, Saturation, and Errors) e RED para serviços.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Gerar um Flame Graph de uma aplicação C++ que realiza cálculos intensivos e identificar qual função é o "gargalo".
- [ ] Compilar uma aplicação com símbolos de debug (`-g`).
- [ ] Executar `perf record -g ./app`.
- [ ] Utilizar os scripts do Brendan Gregg para converter o `perf.data` em um `flamegraph.svg`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Arquitetura de PMUs e Contadores de Hardware - Teoria e Fundamentos]]
- [[./Interpretacao de Flame Graphs e Pilhas de Chamada - Funcionamento Interno e Arquitetura]]
- [[./Overhead de Instrumentacao e Amostragem - Casos de Falha e Analise Amortizada]]
- [[./Analise de Latencia de Cauda com Perf e BPF - Implementacao de Referencia e Benchmarks]]
