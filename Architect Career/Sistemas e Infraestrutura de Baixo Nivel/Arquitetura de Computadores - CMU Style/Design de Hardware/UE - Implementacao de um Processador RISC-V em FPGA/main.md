---
tema: Implementacao de um Processador RISC-V em FPGA
tipo: unidade-estudo
tags: [hardware, verilog, risc-v, arquitetura]
---
# 🧪 UE - Implementacao de um Processador RISC-V em FPGA

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> A maioria das arquiteturas de processadores (x86, ARM) é fechada e proprietária. O **RISC-V** é uma ISA (Instruction Set Architecture) aberta que permite que qualquer um projete seu próprio processador. O desafio é implementar o pipeline, os registros e a unidade de controle em um hardware real (FPGA). Sem essa experiência, a arquitetura de computadores permanece apenas uma teoria abstrata; ao construir um RISC-V, você entende exatamente como cada bit de instrução se torna uma operação física.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[A ISA RISC-V (RV32I)]:** O conjunto básico de 32 instruções e o formato de instrução (R-type, I-type, S-type, etc).
2. **[Pipeline de 5 Estagios]:** IF (Fetch), ID (Decode), EX (Execute), MEM (Memory), WB (Write Back).
3. **[Tratamento de Hazards de Hardware]:** Encaminhamento de dados (Forwarding) e inserção de bolhas (Stalls).
4. **[Interface de Memoria e I-O]:** Como o processador se comunica com a memória externa e periféricos (ex: LEDs, UART) através de barramentos.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *The RISC-V Reader: An Open Architecture Atlas* (Patterson and Waterman).
- **System Blueprint:** Projetos como o **PicoRV32** ou **NEORV32** para integração em FPGAs pequenas.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar a Unidade Lógica e Aritmética (ULA) de um RISC-V em Verilog e testá-la com operações básicas.
- [ ] Definir o módulo da ULA.
- [ ] Implementar operações como `ADD`, `SUB`, `AND`, `OR` e `SLT`.
- [ ] Criar um testbench que simule a entrada de instruções e verifique o resultado nos registros.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Formatos de Instrucao e Decodificacao RISC-V - Teoria e Fundamentos]]
- [[./Arquitetura do Pipeline e Encaminhamento de Dados - Funcionamento Interno e Arquitetura]]
- [[./Temporizacao e Frequencia Maxima em FPGA - Casos de Falha e Analise Amortizada]]
- [[./Bootstrapping de Código C no RISC-V customizado - Implementacao de Referencia e Benchmarks]]
