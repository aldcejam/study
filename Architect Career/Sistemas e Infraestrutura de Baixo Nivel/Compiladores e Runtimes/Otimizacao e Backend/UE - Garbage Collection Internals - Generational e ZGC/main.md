---
tema: Garbage Collection Internals - Generational e ZGC
tipo: unidade-estudo
tags: [runtimes, performance, memoria, java, baixo-nivel]
---
# 🧪 UE - Garbage Collection Internals - Generational e ZGC

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Gerenciar memória manualmente (C/C++) é perigoso e lento para produtividade. Runtimes modernos (Java, Go, .NET) usam **Garbage Collection (GC)**. O problema é que o GC tradicional causa pausas "Stop-the-World" que congelam a aplicação. O **GC Geracional** resolve isso focando em objetos jovens (que morrem rápido). O **ZGC (Z Garbage Collector)** resolve o problema das pausas longas mantendo-as abaixo de 1ms, mesmo em heaps de terabytes. Sem entender o GC, você não consegue tunar sistemas de baixa latência ou alta taxa de transferência.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Hipótese Geracional]:** A observação de que a maioria dos objetos morre jovem.
2. **[Mark-and-Sweep vs Copying Collectors]:** As técnicas básicas de identificação e limpeza de memória.
3. **[Barreiras de Escrita (Write Barriers)]:** Como o GC rastreia referências entre gerações (Old to Young).
4. **[ZGC e Colored Pointers]:** O uso de bits extras no ponteiro para realizar o GC de forma concorrente sem parar as threads da aplicação.
5. **[Safe Points]:** Como e quando as threads da aplicação param para permitir o trabalho do GC.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *The Garbage Collection Handbook* (Richard Jones et al.).
- **System Blueprint:** Tuning do G1GC e ZGC na JVM (HotSpot).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Executar uma aplicação Java com carga pesada de alocação e utilizar a flag `-Xlog:gc` para analisar os tempos de pausa e a frequência das coletas entre o G1GC e o ZGC.
- [ ] Criar um programa que gera muitos objetos temporários.
- [ ] Rodar com `-XX:+UseG1GC` e capturar os logs.
- [ ] Rodar com `-XX:+UseZGC` e comparar as pausas máximas.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Hipótese Geracional e Espacos Eden, Survivor e Tenured - Teoria e Fundamentos]]
- [[./Algoritmos de Marcacao Concorrente e Barreiras - Funcionamento Interno e Arquitetura]]
- [[./Fragmentacao de Heap e Compactacao - Casos de Falha e Analise Amortizada]]
- [[./Tuning de GC para Baixa Latencia - Implementacao de Referencia e Benchmarks]]
