---
tema: Sintese Verilog para Maquinas de Estado
tipo: unidade-estudo
tags: [hardware, verilog, fpga, eletrônica]
---
# 🧪 UE - Sintese Verilog para Maquinas de Estado

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Quase todo hardware digital é, no fundo, uma máquina de estados finitos (FSM). O problema é converter a lógica sequencial (onde o tempo e a ordem importam) em circuitos combinacionais e flip-flops. Usar **Verilog** para descrever FSMs exige um rigor extremo para evitar "latches" indesejados e problemas de temporização. Sem dominar a síntese de FSMs, um engenheiro não consegue projetar controladores de periféricos, processadores ou protocolos de comunicação em silício.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Maquinas de Moore vs Mealy]:** A diferença fundamental entre saídas dependentes apenas do estado ou também das entradas.
2. **[Descricao em Dois Processos]:** O padrão de design Verilog: um bloco `always` para lógica combinacional (próximo estado) e outro para lógica sequencial (atualização de estado).
3. **[Codificacao de Estados]:** Binary vs One-Hot e como isso afeta a velocidade vs área ocupada no chip.
4. **[Metastabilidade e Debouncing]:** Como lidar com sinais assíncronos (como botões) em uma máquina de estados síncrona.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Digital Design and Computer Architecture* (Harris and Harris).
- **System Blueprint:** Implementação de controladores SPI, I2C e UART em FPGAs.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Projetar e simular um controlador de semáforo com detecção de pedestres em Verilog.
- [ ] Desenhar o diagrama de estados.
- [ ] Escrever o código Verilog usando o padrão de dois processos.
- [ ] Testar no simulador (Icarus Verilog ou Vivado) usando um testbench para verificar as transições.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Moore vs Mealy e Logica Combinacional - Teoria e Fundamentos]]
- [[./Flip-Flops e Arvore de Clock - Funcionamento Interno e Arquitetura]]
- [[./Latches Indesejados e Race Conditions - Casos de Falha e Analise Amortizada]]
- [[./Controlador UART em Verilog - Implementacao de Referencia e Benchmarks]]
