---
tema: Gerenciamento de Interrupcoes e Bottom Halves
tipo: unidade-estudo
tags: [linux, kernel, hardware, baixo-nivel]
---
# 🧪 UE - Gerenciamento de Interrupcoes e Bottom Halves

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Quando um hardware (como uma placa de rede) precisa de atenção, ele "interrompe" a CPU. No entanto, a CPU não pode ficar muito tempo tratando essa interrupção, ou o sistema inteiro trava. O problema é: como processar dados volumosos de hardware sem bloquear o sistema? O Kernel resolve isso dividindo o trabalho em **Top Half** (rápido, desabilita interrupções) e **Bottom Half** (lento, permite preempção). Ignorar essa distinção leva a latência de sistema terrível e deadlocks no kernel.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Contexto de Interrupcao]:** As restrições severas: não pode dormir, não pode acessar espaço de usuário, deve ser atômico.
2. **[IRQ Handlers (Top Half)]:** O registro de interrupções e o processamento mínimo necessário para liberar o hardware.
3. **[Softirqs e Tasklets]:** Os mecanismos legados de bottom half para processamento rápido e paralelo.
4. **[Workqueues]:** O mecanismo moderno que executa em contexto de thread, permitindo que o processamento do bottom half durma ou bloqueie se necessário.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Understanding the Linux Kernel* (Daniel P. Bovet and Marco Cesati).
- **System Blueprint:** O subsistema de rede do Linux, onde pacotes são recebidos via interrupção e processados via Softirqs (NAPI).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Escrever um módulo de kernel simples que registre um handler para uma interrupção fictícia (ou teclado) e delegue o processamento para uma Workqueue.
- [ ] Implementar as funções de init e exit do módulo.
- [ ] Registrar o IRQ handler usando `request_irq`.
- [ ] Agendar uma tarefa na Workqueue e imprimir logs no `dmesg` para validar a ordem de execução.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Arquitetura de IRQs e Vetores de Interrupcao - Teoria e Fundamentos]]
- [[./Mecanismos de Softirqs vs Workqueues - Funcionamento Interno e Arquitetura]]
- [[./Latencia de Interrupcao e Priority Inversion - Casos de Falha e Analise Amortizada]]
- [[./Desenvolvimento de um Driver com Workqueues - Implementacao de Referencia e Benchmarks]]
