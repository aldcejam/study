---
tema: Consistent Hashing e Virtual Nodes
tipo: unidade-estudo
tags: [distribuídos, escalabilidade, algoritmos, arquitetura]
---
# 🧪 UE - Consistent Hashing e Virtual Nodes

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Se você tem 1 milhão de usuários e 3 servidores, pode usar `hash(user_id) % 3` para decidir onde guardar os dados. O problema é: o que acontece se você adicionar o 4º servidor? Com o hash modular tradicional, quase todos os dados mudariam de servidor, causando um caos de rebalanceamento. O **Consistent Hashing** resolve isso garantindo que, ao adicionar ou remover um nó, apenas uma pequena fração dos dados precise ser movida ($1/N$). **Virtual Nodes** resolvem o problema de desbalanceamento de carga (nós heterogêneos). Sem isso, é impossível escalar um cache distribuído ou um banco de dados NoSQL de forma suave.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[O Anel de Hash (Hash Ring)]:** Como mapear servidores e dados no mesmo espaço circular de endereçamento.
2. **[Algoritmo de Busca de Sucessor]:** Encontrar o primeiro servidor no sentido horário a partir da posição do dado.
3. **[Virtual Nodes (VNodes)]:** Por que mapear cada servidor físico para múltiplos pontos no anel melhora a distribuição estatística.
4. **[Replicacao em Anel]:** Como garantir alta disponibilidade armazenando o dado no nó sucessor e nos $K$ nós seguintes.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Consistent Hashing and Random Trees* (Karger et al., 1997).
- **System Blueprint:** Uso de Consistent Hashing no Amazon Dynamo, Cassandra e Memcached (libketama).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um anel de hash consistente em Python e medir a porcentagem de chaves que são movidas ao adicionar um novo nó, comparando com o hash modular simples.
- [ ] Implementar a estrutura de anel com busca binária (`bisect`).
- [ ] Adicionar suporte a Virtual Nodes.
- [ ] Simular a inserção de 10.000 chaves e o impacto da entrada de um novo servidor.

## 🗃️ Notas Heutagógicas Atômicas
- [[./O Conceito de Anel de Hash e Espaco de Endereçamento - Teoria e Fundamentos]]
- [[./Estrategias de Virtual Nodes e Balanceamento de Carga - Funcionamento Interno e Arquitetura]]
- [[./Hotspots e Cascading Failures em Anéis Desbalanceados - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de um Router de Sharding com Consistent Hashing - Implementacao de Referencia e Benchmarks]]
