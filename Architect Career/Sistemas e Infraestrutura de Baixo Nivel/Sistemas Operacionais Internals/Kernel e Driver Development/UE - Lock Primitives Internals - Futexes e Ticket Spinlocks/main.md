---
tema: Lock Primitives Internals - Futexes e Ticket Spinlocks
tipo: unidade-estudo
tags: [sistemas-operacionais, concorrência, kernel, baixo-nivel]
---
# 🧪 UE - Lock Primitives Internals - Futexes e Ticket Spinlocks

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Quando duas threads tentam acessar o mesmo dado, elas precisam de um "cadeado" (lock). O problema é que locks são caros. Se a thread ficar girando em loop esperando (Spinlock), ela gasta CPU; se ela for dormir (Mutex), o SO gasta tempo trocando o contexto. O **Futex (Fast Userspace Mutex)** resolve isso tentando fazer tudo no espaço do usuário e só chamando o kernel se houver contenção real. O **Ticket Spinlock** resolve o problema de justiça (fairness) no kernel. Sem entender como locks funcionam "por baixo do capô", você criará gargalos massivos de escalabilidade em sistemas multicore.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Instrucoes Atomicas de Hardware]:** `Compare-and-Swap` (CAS) e `Test-and-Set`.
2. **[Spinlocks vs Mutexes]:** Quando é melhor gastar ciclos de CPU esperando vs quando é melhor dormir.
3. **[Ticket Spinlocks]:** O mecanismo de "fila de padaria" para garantir que a thread que chegou primeiro seja a próxima a entrar na seção crítica.
4. **[Futexes]:** O funcionamento do `futex()` no Linux: como o espaço do usuário e o kernel colaboram para minimizar as System Calls.
5. **[Contencao e Falsa Compartilhamento (False Sharing)]:** Como o layout de memória pode destruir a performance de locks mesmo sem contenção lógica.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Fuss, Futexes and Furwocks: Fast Userlevel Locking in Linux* (Hubertus Franke et al., 2002).
- **System Blueprint:** Implementação de `pthread_mutex` na Glibc e `mutex_lock` no kernel Linux.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um Spinlock simples em C usando a instrução `__sync_lock_test_and_set` e comparar sua performance com um Mutex padrão sob alta contenção.
- [ ] Implementar o Spinlock.
- [ ] Criar 10 threads incrementando um contador global.
- [ ] Medir o tempo total e a utilização de CPU, comparando com `pthread_mutex_t`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Aritmetica Atomica e Cache Coherence - Teoria e Fundamentos]]
- [[./O Mecanismo de Espera e Acorda do Futex - Funcionamento Interno e Arquitetura]]
- [[./Inversao de Prioridade e Deadlocks - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Locks e Escalabilidade Multicore - Implementacao de Referencia e Benchmarks]]
