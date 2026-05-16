---
tema: Side-Channel Attacks - Meltdown e Spectre
tipo: unidade-estudo
tags: [segurança, arquitetura, baixo-nivel, hardware]
---
# 🧪 UE - Side-Channel Attacks - Meltdown e Spectre

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> E se a própria CPU, para ser mais rápida, vazar seus segredos? O problema é a **Execução Especulativa**: a CPU tenta prever o futuro e executa instruções antes mesmo de saber se elas deveriam ser executadas. Se a previsão falha, a CPU descarta o resultado, mas o **Cache** mantém um rastro físico (latência) dessa execução. **Meltdown** e **Spectre** resolvem o "problema" de como ler a memória do kernel ou de outros processos medindo o tempo de acesso ao cache. Sem entender ataques de canal lateral, você não entende por que isolamento de software (containers, VMs) pode ser quebrado pelo hardware.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Execucao Especulativa e Out-of-Order]:** O ganho de performance e o rastro deixado no hardware.
2. **[Cache Side-Channels]:** As técnicas `Flush+Reload` e `Prime+Probe` para medir latência e inferir dados.
3. **[Meltdown]:** A quebra do isolamento entre o espaço do usuário e o kernel (específico de Intel).
4. **[Spectre]:** O uso de envenenamento de Branch Predictor para induzir a execução de código que acessa dados sensíveis.
5. **[Mitigacoes]:** KPTI (Kernel Page-Table Isolation), Retpolines e barreiras de memória.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Meltdown: Reading Kernel Memory from User Space* (Lipp et al., 2018) e *Spectre Attacks: Exploiting Speculative Execution* (Kocher et al., 2018).
- **System Blueprint:** O impacto dessas vulnerabilidades em provedores de Cloud (AWS, Azure, Google Cloud).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um ataque simples de latência de cache em C para detectar se um determinado valor foi carregado na memória pela diferença de tempo entre um acesso ao cache (hit) e um acesso à RAM (miss).
- [ ] Usar a instrução `rdtscp` para medir ciclos de clock de forma precisa.
- [ ] Usar `clflush` para limpar a linha de cache.
- [ ] Demonstrar a diferença estatística de tempo entre o hit e o miss.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Arquitetura de Cache e Latencia de Memoria - Teoria e Fundamentos]]
- [[./Tecnicas de Medicao de Tempo e Ataques de Canal Lateral - Funcionamento Interno e Arquitetura]]
- [[./Microcode Updates e Degradacao de Performance - Casos de Falha e Analise Amortizada]]
- [[./Analise de Vulnerabilidades em Cloud Multi-Tenant - Implementacao de Referencia e Benchmarks]]
