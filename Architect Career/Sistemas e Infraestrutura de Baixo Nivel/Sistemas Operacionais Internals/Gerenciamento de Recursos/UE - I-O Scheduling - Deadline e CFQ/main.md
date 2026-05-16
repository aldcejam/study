---
tema: I-O Scheduling - Deadline e CFQ
tipo: unidade-estudo
tags: [sistemas-operacionais, storage, performance, baixo-nivel]
---
# 🧪 UE - I-O Scheduling - Deadline e CFQ

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O disco é ordens de magnitude mais lento que a CPU. Se várias aplicações pedem dados ao mesmo tempo, a ordem em que o SO atende esses pedidos decide se o sistema será fluido ou se vai travar. O problema é evitar que pedidos de leitura (que bloqueiam o app) fiquem presos atrás de escritas massivas. O **Deadline Scheduler** garante um tempo máximo para cada pedido, enquanto o **CFQ (Completely Fair Queuing)** tenta dar uma fatia justa de banda para cada processo. Sem entender o I/O scheduler, você não consegue debugar por que um banco de dados fica lento durante um backup.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Elevator Algorithm (SCAN)]:** A base de todos os schedulers: mover a cabeça do disco em uma direção para minimizar o tempo de busca (seek time).
2. **[Deadline Scheduler]:** Como usar filas de expiração para garantir que nenhuma leitura espere mais que X milissegundos.
3. **[CFQ e BFQ]:** A alocação de fatias de tempo de disco por processo e por que isso é importante para desktops vs servidores.
4. **[Schedulers para SSD (Noop/None)]:** Por que em discos sem partes móveis (SSDs), a reordenação para minimizar seek time é inútil e até prejudicial.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Understanding the Linux Kernel* (Daniel P. Bovet and Marco Cesati) - Capítulo sobre I/O Scheduling.
- **System Blueprint:** Configuração de schedulers no Linux (`/sys/block/sdX/queue/scheduler`).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Usar o comando `fio` (Flexible I/O Tester) para medir a latência de leitura em um disco enquanto um processo de escrita pesada está rodando, alternando entre os schedulers `deadline` e `cfq`.
- [ ] Criar um cenário de contenção de I/O.
- [ ] Alterar o scheduler em tempo real.
- [ ] Coletar as métricas de P99 latência para cada caso.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Algoritmos de Escalonamento de Disco e Seek Time - Teoria e Fundamentos]]
- [[./Arquitetura de Filas de Deadline e Prioridades - Funcionamento Interno e Arquitetura]]
- [[./A Transicao para NVMe e a Obsolescencia de Schedulers - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de I-O com FIO e Iostat - Implementacao de Referencia e Benchmarks]]
