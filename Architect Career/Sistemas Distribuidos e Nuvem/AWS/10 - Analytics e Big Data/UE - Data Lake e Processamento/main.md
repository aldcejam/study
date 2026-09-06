---
tema: Data Lake e Processamento - Glue, Athena, EMR
tipo: unidade-estudo
tags: [aws, data-lake, glue, athena, emr, lake-formation]
---
# 🧪 UE - Data Lake e Processamento

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Um data lake sem catálogo e sem governança vira um "data swamp": petabytes ilegíveis, sem esquema conhecido e sem controle de quem acessa o quê. O desafio é catalogar dados heterogêneos (Glue), consultá-los sem ETL prévio (Athena), processá-los em escala (EMR/Glue jobs) e governar acesso granular a colunas/linhas (Lake Formation). Errar formato de arquivo (JSON vs Parquet) ou particionamento multiplica o custo de query por ordens de magnitude.

## 🧬 Grade Atômica de Tópicos
1. **Fundação do data lake (S3 + formatos):** Layout de dados, particionamento, formatos colunar (Parquet/ORC), compressão, table formats (Iceberg/Hudi/Delta).
2. **Catálogo e ETL (Glue):** Glue Data Catalog, crawlers, Glue ETL jobs (Spark serverless), Glue Studio, schema evolution.
3. **Consulta serverless (Athena):** SQL sobre S3 via Presto, particionamento e projeção, custo por dado escaneado, federated queries, Athena for Spark.
4. **Processamento em escala e governança (EMR + Lake Formation):** EMR (Spark/Hive/Presto em cluster), EMR Serverless; Lake Formation para permissões finas (coluna/linha/tag) e data sharing.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/what-is-glue.html), [Amazon Athena](https://docs.aws.amazon.com/athena/latest/ug/what-is.html) e [AWS Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/what-is-lake-formation.html).
- **System Blueprint:** Data lake governado: S3 (Parquet particionado) + Glue Catalog + Lake Formation (permissões por coluna) + Athena para analistas.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Construir um mini data lake consultável.
- [ ] Carregar dados no S3, rodar um Glue crawler e consultar via Athena.
- [ ] Comparar custo/tempo de query em JSON vs Parquet particionado.
- [ ] Aplicar uma permissão de coluna via Lake Formation e validar o filtro.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Data Lake no S3 - Formatos e Particionamento]]
- [[./02 - Glue - Catalogo e ETL]]
- [[./03 - Athena, EMR e Lake Formation]]
