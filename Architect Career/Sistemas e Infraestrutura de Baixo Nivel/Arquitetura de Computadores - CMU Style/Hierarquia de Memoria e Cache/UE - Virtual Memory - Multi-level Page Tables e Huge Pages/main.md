---
tema: Virtual Memory - Multi-level Page Tables e Huge Pages
tipo: unidade-estudo
tags: [arquitetura, sistemas-operacionais, performance, baixo-nivel]
---
# 🧪 UE - Virtual Memory - Multi-level Page Tables e Huge Pages

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como dar a cada programa a ilusão de que ele possui toda a memória do computador sem que eles sobrescrevam uns aos outros? A **Memória Virtual** resolve isso através de um mapeamento (Page Tables) entre endereços lógicos e físicos. O problema é que esse mapeamento consome memória e tempo de CPU (Page Walks). **Multi-level Page Tables** economizam espaço, enquanto **Huge Pages** reduzem o overhead do TLB para aplicações massivas. Sem entender isso, um arquiteto não consegue otimizar bancos de dados ou máquinas virtuais que sofrem com latência de memória.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Paginacao e Segmentacao]:** A divisão da memória em blocos fixos vs blocos variáveis.
2. **[Multi-level Page Tables]:** Como a estrutura de árvore (ex: 4 níveis no x86-64) economiza RAM para processos com memória esparsa.
3. **[TLB (Translation Lookaside Buffer)]:** A cache de hardware para traduções de endereços e o impacto de um *TLB Miss*.
4. **[Huge Pages (2MB/1GB)]:** Quando e por que usar páginas maiores para reduzir a pressão no TLB.
5. **[Page Faults e Swapping]:** O mecanismo de carregar dados do disco sob demanda e o custo catastrófico do *thrashing*.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Operating Systems: Three Easy Pieces* (Remzi and Andrea Arpaci-Dusseau) - Capítulos sobre Virtual Memory.
- **System Blueprint:** Gerenciamento de memória no Kernel Linux e suporte a `Transparent Huge Pages (THP)`.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Escrever um programa C que aloca uma grande quantidade de memória e comparar o tempo de acesso aleatório com e sem o uso de Huge Pages.
- [ ] Usar `mmap` com a flag `MAP_HUGETLB`.
- [ ] Medir a latência de acesso aleatório em um array de 1GB.
- [ ] Analisar as estatísticas do sistema (`/proc/meminfo`) para verificar o uso de Huge Pages.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Estrutura de Tabelas de Paginas em 4 Niveis - Teoria e Fundamentos]]
- [[./Funcionamento do TLB e Context Switch - Funcionamento Interno e Arquitetura]]
- [[./Thrashing e Algoritmos de Substituicao de Paginas - Casos de Falha e Analise Amortizada]]
- [[./Configuracao e Benchmarks de Huge Pages no Linux - Implementacao de Referencia e Benchmarks]]
