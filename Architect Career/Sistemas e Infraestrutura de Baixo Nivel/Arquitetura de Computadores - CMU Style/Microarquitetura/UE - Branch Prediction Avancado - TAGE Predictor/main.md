---
tema: Branch Prediction Avancado - TAGE Predictor
tipo: unidade-estudo
tags: [arquitetura, hardware, performance, baixo-nivel]
---
# 🧪 UE - Branch Prediction Avancado - TAGE Predictor

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em um pipeline profundo (20+ estágios), um erro na predição de um `if` ou `loop` significa descartar dezenas de instruções já em processamento. O problema é que padrões de desvios podem ser complexos e dependentes de um histórico longo. O **TAGE (TAgged GEometric)** é o estado da arte em preditores de desvios, usando tabelas com históricos de comprimentos geométricos para capturar desde padrões curtos até correlações globais distantes. Sem um preditor eficiente, o pipeline fica vazio a maior parte do tempo.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Predicao Bimodal e Local]:** O uso de contadores de 2 bits para prever loops simples.
2. **[Predicao Correlacionada Global]:** Como o histórico de outros desvios influencia o desvio atual.
3. **[O Algoritmo TAGE]:** O uso de múltiplas tabelas (tags) com diferentes tamanhos de histórico e como elas competem para fornecer a predição mais precisa.
4. **[Branch Target Buffer (BTB)]:** Não basta prever *se* vai desviar, é preciso saber *para onde* ir antes mesmo de decodificar a instrução.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *A case for (partially) tagged geometric history length branch prediction* (André Seznec, 2006).
- **System Blueprint:** Implementações modernas de CPUs de alto desempenho (Intel, AMD, Apple M-series) utilizam variações do TAGE.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um preditor TAGE simplificado em Python e testá-lo contra um trace de execução que contém padrões de desvios complexos.
- [ ] Criar tabelas de histórico com diferentes tamanhos.
- [ ] Implementar a lógica de seleção de tabela baseada no maior histórico que deu "match".
- [ ] Comparar a taxa de acerto com um preditor global simples.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Historico Global e Local de Desvios - Teoria e Fundamentos]]
- [[./Arquitetura de Tabelas Tagged Geometric - Funcionamento Interno e Arquitetura]]
- [[./Impacto do Erro de Predicao (Misprediction Penalty) - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Predicao com traces de CPU - Implementacao de Referencia e Benchmarks]]
