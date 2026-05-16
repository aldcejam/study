---
tema: Pipeline Stages e Hazards - Data e Control
tipo: unidade-estudo
tags: [arquitetura, hardware, performance, baixo-nivel]
---
# 🧪 UE - Pipeline Stages e Hazards - Data e Control

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como um processador consegue executar bilhões de instruções por segundo se cada instrução leva múltiplos ciclos para terminar? A resposta é o **Pipelining**. No entanto, o pipeline introduz **Hazards** (conflitos): o que acontece quando uma instrução precisa de um dado que a instrução anterior ainda não terminou de calcular? Se o engenheiro não entender os hazards de dados e controle, ele escreverá código assembly (ou C otimizado) que causa *stalls* constantes, destruindo a performance que o hardware deveria entregar.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Estagios do Pipeline (RISC)]:** Fetch (IF), Decode (ID), Execute (EX), Memory (MEM) e Write-back (WB).
2. **[Data Hazards e Forwarding]:** Conflitos de dependência de dados e a técnica de *Forwarding/Bypassing* para evitar stalls.
3. **[Control Hazards e Branch Penalty]:** O problema dos desvios condicionais e por que o processador precisa "chutar" o resultado do desvio.
4. **[Structural Hazards]:** Limitações físicas do hardware (ex: um único porto de memória para dados e instruções).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Computer Architecture: A Quantitative Approach* (Hennessy & Patterson).
- **System Blueprint:** Arquiteturas modernas como x86 (Intel/AMD) e ARM, onde os pipelines podem ter 15+ estágios.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Analisar um trecho de código assembly e identificar onde ocorreriam stalls e onde o forwarding poderia salvar ciclos de CPU.
- [ ] Escrever uma sequência de instruções com dependências `RAW` (Read After Write).
- [ ] Desenhar o diagrama de espaço-tempo do pipeline para essa sequência.
- [ ] Simular o comportamento usando um simulador de arquitetura (como o RARS para RISC-V).

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Fluxo de Instrucoes e Estagios de Execucao - Teoria e Fundamentos]]
- [[./Mecanismos de Forwarding e Interlock - Funcionamento Interno e Arquitetura]]
- [[./Impacto de Branches na Vazao do Pipeline - Casos de Falha e Analise Amortizada]]
- [[./Analise de Stalls em Codigo Assembly Otimizado - Implementacao de Referencia e Benchmarks]]
