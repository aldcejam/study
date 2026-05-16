---
tema: Algoritmos de Alocacao de Memoria - Buddy System e SLAB
tipo: unidade-estudo
tags: [sistemas-operacionais, kernel, performance, baixo-nivel]
---
# 🧪 UE - Algoritmos de Alocacao de Memoria - Buddy System e SLAB

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como o kernel gerencia a RAM física para evitar que ela se torne uma "colcha de retalhos" (fragmentação)? O problema é alocar memória de tamanhos variados de forma rápida. O **Buddy System** resolve a alocação de grandes blocos (páginas) dividindo e fundindo blocos de potências de 2. O **SLAB Alocator** resolve a alocação de pequenos objetos (como descritores de arquivos) pré-alocando "caches" de objetos do mesmo tamanho. Sem isso, a alocação de memória se tornaria o maior gargalo de performance do sistema operacional.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Fragmentacao Externa vs Interna]:** Por que sobra espaço que não pode ser usado e como minimizar isso.
2. **[Buddy System]:** A técnica de dividir blocos ao meio (buddies) e fundi-los recursivamente na liberação.
3. **[SLAB, SLUB e SLOB]:** A evolução dos alocadores de objetos no Linux; como reduzir o desperdício de metadados.
4. **[Object Caching]:** Por que é mais rápido reutilizar um objeto inicializado do que criar um novo do zero.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *The SLAB Allocator: An Object-Caching Kernel Memory Allocator* (Jeff Bonwick, 1994).
- **System Blueprint:** Alocadores de memória do Kernel Linux (`mm/slab.c`, `mm/page_alloc.c`).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um simulador de Buddy System que gerencia 1024KB de memória e suporta alocações de potências de 2.
- [ ] Criar a estrutura de listas de blocos livres por ordem (2^k).
- [ ] Implementar a função `split` para alocação e `merge` para liberação.
- [ ] Demonstrar como a fragmentação externa é controlada pelo sistema.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Fragmentacao e Gerenciamento de Espaco Livre - Teoria e Fundamentos]]
- [[./Arquitetura de Alocadores de Objetos SLAB - Funcionamento Interno e Arquitetura]]
- [[./Custo de Travamento (Locking) em Alocadores Multicore - Casos de Falha e Analise Amortizada]]
- [[./Monitoramento de Kernel Meminfo e Slabtop - Implementacao de Referencia e Benchmarks]]
