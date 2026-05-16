---
tema: Scheduler Design - Multi-Level Feedback Queue
tipo: unidade-estudo
tags: [sistemas-operacionais, performance, algoritmos, baixo-nivel]
---
# 🧪 UE - Scheduler Design - Multi-Level Feedback Queue

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como o sistema operacional decide quem usa a CPU agora? O problema é equilibrar processos interativos (que precisam de resposta rápida) e processos de cálculo (que precisam de throughput). O **MLFQ (Multi-Level Feedback Queue)** resolve isso sem precisar saber o comportamento do processo de antemão. Ele "aprende" o comportamento: se um processo usa muita CPU, ele é rebaixado; se ele faz muito I/O, ele é priorizado. Sem entender o scheduler, você não entende por que sua aplicação "engasga" sob carga.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[As Regras do MLFQ]:** Como as filas de prioridade funcionam e como os processos se movem entre elas.
2. **[Time Slicing e Quanta]:** Por que filas de alta prioridade têm fatias de tempo menores e filas de baixa prioridade têm fatias maiores.
3. **[Starvation e Priority Boost]:** O perigo de processos de baixa prioridade nunca rodarem e a solução de subir todos de prioridade periodicamente.
4. **[Contabilidade de Tempo de CPU]:** Como o sistema impede que processos "enganem" o scheduler fazendo I/O falso para manter a prioridade alta.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *A Multi-Level Feedback Queuing Scheduling Algorithm* (Fernando J. Corbató, 1962).
- **System Blueprint:** O scheduler do Windows e as versões antigas do scheduler do FreeBSD utilizam variações do MLFQ.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um simulador de MLFQ em Python que gerencia 3 processos (um interativo e dois CPU-bound) e medir o tempo de resposta e turnaround.
- [ ] Criar 3 filas de prioridade.
- [ ] Implementar a lógica de rebaixamento por uso de fatia de tempo.
- [ ] Implementar o "Priority Boost" e observar a recuperação dos processos CPU-bound.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Filas de Prioridade e Time Slices - Teoria e Fundamentos]]
- [[./Mecanismos de Aprendizado de Comportamento de Processos - Funcionamento Interno e Arquitetura]]
- [[./Starvation e o Problema da Injustica - Casos de Falha e Analise Amortizada]]
- [[./Simulador de Scheduler MLFQ vs Round Robin - Implementacao de Referencia e Benchmarks]]
