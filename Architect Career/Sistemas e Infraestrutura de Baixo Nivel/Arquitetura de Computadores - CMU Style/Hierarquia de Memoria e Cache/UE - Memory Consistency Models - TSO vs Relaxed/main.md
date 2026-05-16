---
tema: Memory Consistency Models - TSO vs Relaxed
tipo: unidade-estudo
tags: [arquitetura, concorrência, hardware, baixo-nivel]
---
# 🧪 UE - Memory Consistency Models - TSO vs Relaxed

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas multicore, se a CPU 1 escreve um dado e a CPU 2 lê, o que ela vê? A resposta depende do **Modelo de Consistência de Memória**. O problema é que, para ganhar performance, as CPUs e compiladores reordenam acessos à memória. **TSO (Total Store Ordering)** é o modelo do x86 (mais intuitivo), enquanto o **Relaxed** é o do ARM/RISC-V (mais performante, mas traiçoeiro). Sem entender isso, um desenvolvedor de baixo nível criará bugs de concorrência "impossíveis" que só ocorrem em certas arquiteturas.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Consistencia Sequencial (SC)]:** O modelo ideal (e lento) onde tudo parece ocorrer em ordem global.
2. **[TSO (Total Store Ordering)]:** Por que o x86 permite que leituras passem à frente de escritas anteriores (Store Buffer).
3. **[Modelos Relaxados (Weak Consistency)]:** Como ARM e RISC-V permitem quase qualquer reordenação.
4. **[Memory Barriers (Fences)]:** As instruções que forçam a ordem e sincronizam a visão da memória entre cores.
5. **[Causalidade e Coerência]:** A diferença entre a ordem das operações (Consistência) e a unicidade do valor em um endereço (Coerência).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *A Primer on Memory Consistency and Cache Coherence* (Sorin, Hill, and Wood).
- **System Blueprint:** Implementação de primitivas `std::atomic` no C++ e barreiras de memória no kernel Linux.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Demonstrar o "Store Buffer Reordering" no x86 usando um teste de estresse em C que tenta capturar o estado onde duas CPUs escrevem e ambas leem o valor antigo.
- [ ] Criar duas threads presas em cores diferentes (`sched_setaffinity`).
- [ ] Implementar o algoritmo de Dekker simplificado sem barreiras.
- [ ] Observar a violação da consistência sequencial e depois corrigi-la com `mfence` ou `std::atomic_thread_fence`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Store Buffers e Reordenacao de Instrucoes - Teoria e Fundamentos]]
- [[./Barreiras de Memoria e Primitivas Atomic - Funcionamento Interno e Arquitetura]]
- [[./Litmus Tests e Verificacao de Modelos de Memoria - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Sincronizacao x86 vs ARM - Implementacao de Referencia e Benchmarks]]
