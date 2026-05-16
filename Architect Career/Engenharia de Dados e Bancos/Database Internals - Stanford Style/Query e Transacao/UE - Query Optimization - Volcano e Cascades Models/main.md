---
tema: Query Optimization - Volcano e Cascades Models
tipo: unidade-estudo
tags: [bancos-de-dados, algoritmos, performance, internals]
---
# 🧪 UE - Query Optimization - Volcano e Cascades Models

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Uma consulta SQL pode ser executada de milhares de formas diferentes (Join order, Index scan vs Sequential scan). O problema é encontrar o plano de execução mais barato em milissegundos, antes que a própria busca comece. O **Volcano Model** introduziu a otimização baseada em extensibilidade, e o **Cascades Model** (usado no SQL Server e CockroachDB) utiliza busca no espaço de estados e memoização para encontrar o plano ótimo. Sem entender o otimizador, você não consegue explicar por que uma query rápida de repente fica lenta.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Algebra Relacional e Equivalencias]:** Como transformar uma query sem alterar seu resultado (ex: push down de predicados).
2. **[Otimizacao Baseada em Custo (CBO)]:** O uso de estatísticas (histogramas) para estimar a seletividade e o custo de I/O/CPU.
3. **[Volcano Optimizer Framework]:** O uso de regras de transformação e busca top-down.
4. **[Cascades Optimizer]:** A evolução que integra a busca de regras e o cálculo de custo de forma mais eficiente usando *Groups* e *Expressions*.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *The Cascades Framework for Query Optimization* (Graefe, 1995).
- **System Blueprint:** Otimizadores do PostgreSQL (Genetic Query Optimizer) e do CockroachDB (Cascades-based).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Usar o comando `EXPLAIN ANALYZE` no PostgreSQL para comparar o plano de execução de uma query com e sem índices, e forçar diferentes ordens de Join usando `SET enable_nestloop = off`.
- [ ] Criar duas tabelas com 1 milhão de linhas.
- [ ] Executar um Join complexo e analisar o gráfico do plano.
- [ ] Identificar onde o otimizador errou na estimativa de linhas.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Algebra Relacional e Transformacoes Logicas - Teoria e Fundamentos]]
- [[./Estimativa de Cardinalidade e Histogramas - Funcionamento Interno e Arquitetura]]
- [[./Algoritmos de Join - Hash vs Merge vs NestLoop - Casos de Falha e Analise Amortizada]]
- [[./Analise de Planos de Execucao Reais - Implementacao de Referencia e Benchmarks]]
