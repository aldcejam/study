---
tema: Distributed ID Generation - Snowflake e UUID v7
tipo: unidade-estudo
tags: [distribuídos, arquitetura, performance, bancos-de-dados]
---
# 🧪 UE - Distributed ID Generation - Snowflake e UUID v7

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em um sistema distribuído com múltiplos bancos de dados (sharding), você não pode usar `auto_increment` porque dois servidores podem gerar o ID `100` simultaneamente. O problema é gerar bilhões de IDs únicos, de forma descentralizada, que sejam rápidos de gerar e que mantenham a ordenação temporal (k-sortable). **Snowflake (Twitter)** resolve isso usando um formato de bits que inclui timestamp e ID de máquina. **UUID v7** resolve isso padronizando IDs de 128 bits que são amigáveis a índices de bancos de dados. Sem IDs distribuídos eficientes, seus índices de banco de dados sofrem com fragmentação massiva e sua performance de inserção despenca.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Formatos de ID]:** UUID v4 (aleatório) vs Snowflake vs UUID v7 (time-ordered).
2. **[Anatomia do Snowflake]:** Timestamp (41 bits), Machine ID (10 bits) e Sequence (12 bits).
3. **[O Problema do Clock Skew]:** O que acontece se o relógio de um servidor andar para trás e como evitar IDs duplicados.
4. **[Fragmentacao de B-Tree]:** Por que IDs aleatórios destroem a performance de inserção em bancos de dados relacionais.
5. **[Entropia e Colisao]:** A probabilidade matemática de gerar dois IDs iguais e como garantir a unicidade global.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Announcing Snowflake* (Twitter Engineering Blog, 2010).
- **System Blueprint:** Implementação de UUID v7 (RFC 9562) e uso de Snowflake em sistemas como Discord e Instagram.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um gerador de IDs estilo Snowflake em Go que garanta que os IDs gerados em uma mesma milisegunda sejam diferentes e ordenados.
- [ ] Criar a lógica de bitmasking para compor o ID de 64 bits.
- [ ] Implementar o controle de sequência para colisões na mesma milisegunda.
- [ ] Medir quantos IDs podem ser gerados por segundo em uma única thread.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Ordenacao Temporal (K-Sortable) e Performance de Indices - Teoria e Fundamentos]]
- [[./Arquitetura de Geradores Distribuídos e Sincronizacao de Clock - Funcionamento Interno e Arquitetura]]
- [[./Colisoes de UUID e o Paradoxo do Aniversario - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Insercao UUID v4 vs UUID v7 vs Snowflake - Implementacao de Referencia e Benchmarks]]
