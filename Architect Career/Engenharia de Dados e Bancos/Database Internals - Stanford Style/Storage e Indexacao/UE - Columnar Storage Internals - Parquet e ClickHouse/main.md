---
tema: Columnar Storage Internals - Parquet e ClickHouse
tipo: unidade-estudo
tags: [bancos-de-dados, data-engineering, performance, analytics]
---
# 🧪 UE - Columnar Storage Internals - Parquet e ClickHouse

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Bancos de dados tradicionais (Row-oriented) são ótimos para transações (`SELECT * FROM user WHERE id=1`), mas terríveis para análises (`SELECT AVG(age) FROM users`). O problema é que, para ler uma coluna, o disco precisa ler a linha inteira. O **Armazenamento Colunar** resolve isso agrupando os dados de cada coluna juntos. Isso permite taxas de compressão absurdas e o uso de instruções SIMD para processar milhões de linhas por segundo. Sem entender o armazenamento colunar, você não consegue construir sistemas de Big Data ou Data Warehouses eficientes.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Row-oriented vs Column-oriented]:** Trade-offs de I/O para OLTP vs OLAP.
2. **[Compressao por Coluna]:** Técnicas como RLE (Run-Length Encoding), Delta Encoding e Dictionary Encoding.
3. **[Formato Parquet]:** A estrutura de Row Groups, Column Chunks e metadados de rodapé (footer).
4. **[ClickHouse Internals]:** O uso de MergeTree, índices esparsos e execução vetorial de queries.
5. **[Late Materialization]:** Por que adiar a reconstrução da linha até o último momento economiza CPU e RAM.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *C-Store: A Column-oriented DBMS* (Stonebraker et al., 2005).
- **System Blueprint:** Arquitetura do Apache Parquet e do motor de storage do ClickHouse.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Utilizar o Python com `pyarrow` para criar um arquivo Parquet e comparar o tamanho em disco com um CSV equivalente, além de medir o tempo de leitura de uma única coluna.
- [ ] Gerar um dataset com 1 milhão de linhas e muitas colunas repetitivas.
- [ ] Salvar em CSV e Parquet (com compressão Snappy/Zstd).
- [ ] Comparar tempos de leitura de uma coluna específica usando `pandas`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Encoding e Compressao de Dados Colunares - Teoria e Fundamentos]]
- [[./Execucao Vetorial e SIMD em Consultas Analiticas - Funcionamento Interno e Arquitetura]]
- [[./Custo de Escrita e Updates em Bancos Colunares - Casos de Falha e Analise Amortizada]]
- [[./Otimizacao de Queries no ClickHouse - Implementacao de Referencia e Benchmarks]]
