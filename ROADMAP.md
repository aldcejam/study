# 🗺️ Roadmap de Engenharia e Arquitetura de Sistemas

> [!IMPORTANT]
> **Modelo de Estudo:** Este documento é organizado em **Unidades de Estudo (UE)**. Cada UE representa uma célula atômica de conhecimento que deve ser explorada através de:
> 1. **Teoria:** Leitura de papers clássicos ou capítulos de livros referência (MIT/Stanford/CMU).
> 2. **Prática:** Implementação de um protótipo ou resolução de problemas complexos.
> 3. **Fichamento:** Documentação das decisões de design e lições aprendidas.

``` mermaid
mindmap
  root((ENGENHARIA E ARQUITETURA))
    Bases Matematicas e Teoricas
      Matematica Discreta
        Teoria dos Grafos
          UE: Algoritmos de Fluxo (Ford-Fulkerson e Edmonds-Karp)
          UE: Teorema Max-Flow Min-Cut
          UE: Matching em Grafos Bipartidos (Hopcroft-Karp)
        Teoria dos Numeros
          UE: Aritmetica Modular e Teorema Chinês dos Restos
          UE: Testes de Primaridade (Miller-Rabin)
          UE: Implementacao RSA e Curvas Elipticas (ECC)
        Probabilidade
          UE: Cadeias de Markov e Processos Estocasticos
          UE: Teorema de Bayes e Filtros de Bloom Probabilisticos
      Teoria da Computacao
        Automatos e Gramaticas
          UE: Hierarquia de Chomsky
          UE: Automatos de Pilha e Parsing de Linguagens Context-Free
        Decidibilidade e Complexidade
          UE: Maquinas de Turing e Problema da Parada
          UE: Reducoes Polinomiais e NP-Completude
          UE: Analise Amortizada de Algoritmos
      Logica Matematica
        Metodos Formais
          UE: Logica de Predicados e Unificacao
          UE: Model Checking (LTL/CTL e Ferramenta SPIN)
          UE: Especificacao TLA+ para Sistemas Distribuídos
          UE: Prova de Teoremas com Coq/Lean
        Calculo Lambda
          UE: Reducao Beta e Normalizacao
          UE: Sistemas de Tipos (Hindley-Milner)
    Sistemas e Baixo Nivel
      Arquitetura (CMU Style)
        Processamento e Binarios
          UE: Representacao IEEE 754 e Overflow de Inteiros
          UE: Otimizacao Assembly (SIMD/Vectorization)
          UE: Pipeline de Instrucoes e Out-of-Order Execution
        Hierarquia de Memoria
          UE: Cache Coherence (MESI Protocol)
          UE: Memoria Virtual e TLB Internals
          UE: Performance de I-O (Direct Memory Access - DMA)
      Sistemas Operacionais e Kernel
        Kernel Internals
          UE: Implementacao de System Calls e Context Switching
          UE: Gerenciamento de Memória (Slab Allocator e Paging)
          UE: File Systems Internals (VFS e Inodes)
        Observabilidade e Performance
          UE: eBPF Virtual Machine e BPF Compiler Collection
          UE: XDP (Express Data Path) para Redes
          UE: Tracing de Sistemas com LTTng ou DTrace
      Compiladores e Linguagens
        Frontend e Middle-end
          UE: Analise Lexica/Sintatica e Geracao de AST
          UE: Representacoes Intermediarias (LLVM IR)
        Backend e JIT
          UE: Algoritmos de Alocacao de Registradores
          UE: JIT Compilation e Otimizacoes de Loop (LLVM/GraalVM)
      Computacao de Alta Performance
        Paralelismo
          UE: Programacao CUDA (Kernels e Memory Hierarchy)
          UE: Lock-Free Structures (CAS e ABA Problem)
          UE: Memory Barriers e Memory Models (Release-Acquire)
    Seguranca de Sistemas
      Seguranca Ofensiva
        Memory Corruption
          UE: Stack/Heap Overflows e ROP Chains
          UE: Bypass de Mitigacoes (ASLR, DEP/NX, Canary)
        Engenharia Reversa
          UE: Analise Estatica (Ghidra/IDA Pro)
          UE: Analise Dinamica e Unpacking de Malware
      Seguranca de Hardware e Redes
        UE: Side-Channel Attacks (Spectre/Meltdown)
        UE: Trusted Execution Environments (Intel SGX/ARM TrustZone)
        UE: Protocolos Seguros (TLS 1.3 e Noise Protocol)
    Engenharia de Dados e Algoritmos
      Algoritmos Avancados
        UE: Estruturas Persistentes (Persistent Segment Trees)
        UE: Algoritmos de Aproximacao para Problemas NP-Hard
        UE: Geometria Computacional (Convex Hull e Sweep Line)
      Bancos de Dados (Stanford Style)
        Storage Engines
          UE: Implementacao de B+Trees vs LSM-Trees
          UE: Buffer Pool Management e Page Eviction Policies
        Transacoes e Query
          UE: Query Optimizer (Cost-based e Heuristicos)
          UE: Controle de Concorrencia (MVCC e 2PL)
          UE: Protocolo ARIES (WAL e Recovery)
    Sistemas Distribuidos
      Consenso e Ordem
        UE: Relogios Vetoriais e Causalidade
        UE: Algoritmo Paxos (Single e Multi-Paxos)
        UE: Raft Consensus e Membership Changes
      Alta Escala e Disponibilidade
        UE: Gossip Protocols e Failure Detection (SWIM)
        UE: Particionamento e Sharding (Consistent Hashing)
        UE: Consistencia Eventual e CRDTs
      Arquitetura de Nuvem
        UE: Service Mesh Internals (Envoy e Sidecars)
        UE: Protocolos de Comunicacao (gRPC e Protocol Buffers)
        UE: Teorema CAP e Modelos de Consistencia (PACELC)
    Engenharia de Software e Design
      Qualidade Avançada
        UE: Property-Based Testing (Hypothesis/QuickCheck)
        UE: Chaos Engineering e Fault Injection
      Design e Evolucao
        UE: Strategic DDD (Bounded Contexts e Context Mapping)
        UE: Event Sourcing e Reconstrucao de Estado
        UE: Implementacao de Arquitetura Hexagonal (Ports & Adapters)
```
