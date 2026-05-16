---
tema: MVCC Internals - Version Chain e Vacuuming
tipo: unidade-estudo
tags: [bancos-de-dados, postgres, concorrência, performance]
---
# 🧪 UE - MVCC Internals - Version Chain e Vacuuming

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como um banco de dados permite que mil usuários leiam um dado enquanto outro o atualiza, sem que ninguém precise esperar (lock)? O **MVCC (Multi-Version Concurrency Control)** resolve isso mantendo múltiplas versões da mesma linha. No entanto, isso cria um problema de "lixo" (bloat). Se o engenheiro não entender como o **Vacuuming** funciona e como as *Version Chains* impactam a performance, o banco de dados eventualmente degradará até ficar inutilizável devido ao acúmulo de versões mortas.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Snapshot Isolation]:** Como cada transação vê uma "foto" consistente do banco baseada em seu ID de transação (XID).
2. **[Version Chains e Pointer Chasing]:** Como o banco percorre as versões de uma linha (do mais novo para o mais antigo ou vice-versa).
3. **[Visibilidade de Tuplas (xmin/xmax)]:** Os metadados ocultos que decidem se uma linha é visível para uma transação específica.
4. **[Processo de Vacuum e Freezing]:** A limpeza de versões obsoletas e a prevenção do "Transaction ID Wraparound".

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Transaction Management in Postgres* (PostgreSQL Documentation Internals).
- **System Blueprint:** Implementação de MVCC no PostgreSQL vs MySQL (InnoDB) e como elas diferem no armazenamento das versões (In-place vs Undo Log).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Observar o fenômeno do "Bloat" em uma tabela Postgres e monitorar o trabalho do Autovacuum.
- [ ] Realizar milhões de updates em uma tabela pequena sem Vacuum.
- [ ] Medir o tamanho físico do arquivo no disco usando `pg_relation_size`.
- [ ] Executar o `VACUUM FULL` e analisar a recuperação de espaço e o impacto nos índices.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Isolamento de Snapshot e Transacoes ACID - Teoria e Fundamentos]]
- [[./Estrutura de Tuplas e Cabeçalhos no Postgres - Funcionamento Interno e Arquitetura]]
- [[./Bloat de Tabela e Latencia de Vacuum - Casos de Falha e Analise Amortizada]]
- [[./Analise de Performance de Indices sob MVCC - Implementacao de Referencia e Benchmarks]]
