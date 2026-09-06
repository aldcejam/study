---
tema: Amazon DynamoDB
tipo: unidade-estudo
tags: [aws, dynamodb, nosql, chave-valor, escala]
---
# 🧪 UE - DynamoDB

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> DynamoDB entrega latência de milissegundos de dígito único em qualquer escala — mas só se a modelagem de dados for feita ao redor dos padrões de acesso, não da normalização relacional. A escolha errada de partition key causa "hot partitions" e throttling; modelar como se fosse SQL destrói a performance. É o exemplo prático de particionamento por consistent hashing e do trade-off do teorema CAP aplicado a um produto gerenciado. Deriva do paper Dynamo. Conecta com [[../../../Arquitetura de Alta Escala/UE - Consistent Hashing e Virtual Nodes/main|Consistent Hashing]].

## 🧬 Grade Atômica de Tópicos
1. **Modelo de dados e chaves:** Partition key (hash) e sort key (range), item collections, single-table design, distribuição uniforme para evitar hot partitions.
2. **Índices:** Global Secondary Index (GSI, novo espaço de chave, eventual) vs Local Secondary Index (LSI, mesma partition, strong); custos e limites.
3. **Capacidade e performance:** On-Demand vs Provisioned + Auto Scaling, RCU/WCU, adaptive capacity, DAX (cache), leitura eventual vs fortemente consistente.
4. **Recursos avançados:** DynamoDB Streams (CDC), Global Tables (multi-região ativo-ativo), TTL, transactions, PITR e backup.

## 📜 Fronteira Acadêmica e Referências
- **Paper fundacional:** [Dynamo: Amazon's Highly Available Key-value Store](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf) (SOSP 2007).
- **Documentação oficial:** [Amazon DynamoDB Developer Guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html) e o clássico "The DynamoDB Book" (Alex DeBrie).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Modelar single-table e observar particionamento.
- [ ] Criar uma tabela com PK+SK e modelar 2-3 access patterns com single-table design.
- [ ] Criar um GSI e consultar por atributo alternativo.
- [ ] Ativar Streams e disparar uma Lambda a cada mudança; testar TTL.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Modelagem, Partition Keys e Single-Table Design]]
- [[./02 - GSI, LSI e Consistencia]]
- [[./03 - Capacidade, Streams e Global Tables]]
