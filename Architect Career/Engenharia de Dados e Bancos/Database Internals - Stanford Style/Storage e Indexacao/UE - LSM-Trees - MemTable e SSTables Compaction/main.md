---
tema: LSM-Trees - MemTable e SSTables Compaction
tipo: unidade-estudo
tags: [bancos-de-dados, storage, performance, lsm-trees]
---
# 🧪 UE - LSM-Trees - MemTable e SSTables Compaction

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O gargalo fundamental de sistemas de escrita intensiva é a latência de I/O de disco aleatório. As **LSM-Trees (Log-Structured Merge-Trees)** resolvem isso transformando escritas aleatórias em escritas sequenciais em lote. Se o engenheiro não entender como o processo de *Compaction* funciona, o sistema sofrerá de *Write Amplification* (escrita excessiva), *Space Amplification* (disco cheio com dados obsoletos) ou picos de latência imprevisíveis durante o merge de arquivos.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[MemTable e Write-Ahead Log (WAL)]:** Como as escritas são bufferizadas na memória e protegidas contra falhas.
2. **[SSTables (Sorted String Tables)]:** A estrutura imutável de arquivos em disco e por que a ordenação é crucial para buscas rápidas.
3. **[Niveis de Compactacao (Leveled vs Tiered)]:** Diferentes estratégias para mesclar arquivos e remover dados deletados (tombstones).
4. **[Filtros de Bloom e Index Sparse]:** Como evitar leituras desnecessárias em disco quando uma chave não existe.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *The Log-Structured Merge-Tree (LSM-Tree)* (Patrick O'Neil et al., 1996).
- **System Blueprint:** Bancos de alta performance como RocksDB, Apache Cassandra, ScyllaDB e LevelDB.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um "Mini-LSM" em Rust ou Go que suporte `Put`, `Get` e um processo de compactação manual disparado quando a MemTable atinge um limite.
- [ ] Criar uma MemTable (ex: usando uma SkipList ou B-Tree em memória).
- [ ] Implementar o flush da MemTable para um arquivo SSTable ordenado.
- [ ] Criar um mecanismo simples de Merge Sort para compactar duas SSTables em uma nova.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Arquitetura de MemTable e WAL - Teoria e Fundamentos]]
- [[./Estrutura de SSTables e Busca Binaria - Funcionamento Interno e Arquitetura]]
- [[./Estrategias de Compactacao e Amplificacao de Escrita - Casos de Falha e Analise Amortizada]]
- [[./Benchmark de RocksDB vs B-Trees - Implementacao de Referencia e Benchmarks]]
