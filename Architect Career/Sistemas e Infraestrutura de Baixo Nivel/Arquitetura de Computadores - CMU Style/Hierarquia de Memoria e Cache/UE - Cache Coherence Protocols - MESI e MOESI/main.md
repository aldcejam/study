---
tema: Cache Coherence Protocols - MESI e MOESI
tipo: unidade-estudo
tags: [arquitetura, multicore, cache, baixo-nivel]
---
# 🧪 UE - Cache Coherence Protocols - MESI e MOESI

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas multicore, cada núcleo tem sua própria cache L1/L2. O que acontece quando dois núcleos tentam ler e escrever no mesmo endereço de memória simultaneamente? Sem um **Protocolo de Coerência de Cache**, um núcleo veria um dado obsoleto enquanto o outro atualiza a memória, levando a bugs de concorrência impossíveis de depurar. Entender o **MESI** é fundamental para entender como a comunicação entre núcleos funciona e por que o compartilhamento de dados pode destruir a performance devido ao "ping-pong" de cache.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Estados do MESI]:** Modified (M), Exclusive (E), Shared (S) e Invalid (I).
2. **[Bus Snooping vs Directory Based]:** Os dois principais métodos de comunicação entre caches para manter a sincronia.
3. **[False Sharing]:** O fenômeno onde dois threads acessam variáveis diferentes que estão na mesma linha de cache, causando invalidações desnecessárias.
4. **[Protocolo MOESI]:** A adição do estado *Owned* (O) para otimizar a transferência de dados entre caches sem passar pela memória principal.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *A New Solution to Coherence Problems in Multicache Systems* (Papamarcos and Patel, 1984).
- **System Blueprint:** Implementação de barramentos de interconexão como Intel QPI ou AMD Infinity Fabric.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Escrever um programa em C++ que demonstre o impacto do **False Sharing** e usar contadores de performance do hardware (PCM) para medir as invalidações de cache.
- [ ] Criar uma struct onde dois campos são acessados por threads diferentes.
- [ ] Medir o tempo de execução com e sem `padding` (alinhamento de linha de cache).
- [ ] Analisar os eventos de `L1-dcache-load-misses` e `L1-dcache-store-misses` via `perf`.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Transicoes de Estado no MESI - Teoria e Fundamentos]]
- [[./Arquitetura de Barramentos e Snooping - Funcionamento Interno e Arquitetura]]
- [[./Impacto do False Sharing na Escalabilidade - Casos de Falha e Análise Amortizada]]
- [[./Alinhamento de Memoria e Cache Padding - Implementacao de Referencia e Benchmarks]]
