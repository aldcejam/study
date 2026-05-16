---
tema: Speculative Execution e Vulnerabilidades Relacionadas
tipo: unidade-estudo
tags: [arquitetura, hardware, segurança, baixo-nivel]
---
# 🧪 UE - Speculative Execution e Vulnerabilidades Relacionadas

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Processadores modernos tentam adivinhar o futuro para não ficar parados esperando dados da memória. A **Execução Especulativa** executa instruções antes mesmo de saber se elas são necessárias. O problema é que, se o "palpite" estiver errado, o processador descarta o resultado, mas **as alterações na cache permanecem**. Isso cria um canal lateral que permite ler qualquer parte da memória do sistema, quebrando o isolamento entre processos e kernel. Sem entender Spectre e Meltdown, um arquiteto não consegue projetar sistemas seguros em hardware compartilhado.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Execucao Fora de Ordem e Especulacao]:** Como o processador preenche os "bolsões" de tempo ocioso.
2. **[Canais Laterais de Cache]:** A técnica de medir o tempo de acesso para deduzir se um dado foi carregado na cache (Flush+Reload).
3. **[Spectre (Bounds Check Bypass)]:** Como enganar o preditor de desvios para acessar memória fora dos limites de um array.
4. **[Meltdown (Rogue Data Cache Load)]:** Como a execução especulativa ignora permissões de página de kernel em certas CPUs.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Spectre Attacks: Exploiting Speculative Execution* (Kocher et al., 2018).
- **System Blueprint:** Mitigações no kernel Linux como KPTI (Kernel Page Table Isolation) e Retpolines.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um ataque de canal lateral simples (Flush+Reload) para medir a diferença de tempo de acesso entre dados na cache vs RAM.
- [ ] Escrever um código C que limpa uma linha de cache (`clflush`).
- [ ] Acessar o dado e medir os ciclos de CPU usando `rdtsc`.
- [ ] Demonstrar como essa diferença de tempo pode ser usada para vazar 1 bit de informação.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Execucao Especulativa e Predicao de Desvios - Teoria e Fundamentos]]
- [[./Arquitetura de Cache e Timings de Acesso - Funcionamento Interno e Arquitetura]]
- [[./Mitigacoes de Software e Microcode - Casos de Falha e Analise Amortizada]]
- [[./Exploit Proof-of-Concept Spectre - Implementacao de Referencia e Benchmarks]]
