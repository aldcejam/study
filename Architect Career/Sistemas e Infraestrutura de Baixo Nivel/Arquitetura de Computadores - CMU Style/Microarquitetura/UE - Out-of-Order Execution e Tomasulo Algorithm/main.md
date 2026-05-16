---
tema: Out-of-Order Execution e Tomasulo Algorithm
tipo: unidade-estudo
tags: [arquitetura, hardware, performance, baixo-nivel]
---
# 🧪 UE - Out-of-Order Execution e Tomasulo Algorithm

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Processadores "In-Order" travam se uma instrução demorar (ex: um cache miss). A **Execução Fora de Ordem (OoO)** resolve isso permitindo que instruções subsequentes que não dependem do dado atrasado sejam executadas antes. O **Algoritmo de Tomasulo** é a técnica clássica de hardware (usando *Reservation Stations* e *Register Renaming*) para gerenciar dependências de dados dinamicamente. Sem OoO, a performance dos processadores estaria limitada à velocidade da memória mais lenta.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Dependencias de Dados (RAW, WAR, WAW)]:** Entender os perigos de reordenar instruções.
2. **[Register Renaming]:** Como o hardware usa registros físicos extras para eliminar dependências falsas (WAR/WAW).
3. **[Reservation Stations]:** O local onde as instruções esperam até que seus operandos estejam prontos.
4. **[Common Data Bus (CDB)]:** O mecanismo de broadcast que avisa todas as instruções esperando por um resultado.
5. **[Reorder Buffer (ROB)]:** Como garantir que, embora executadas fora de ordem, as instruções sejam "aposentadas" (commit) na ordem original.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *An Efficient Algorithm for Exploiting Multiple Arithmetic Units* (Robert Tomasulo, 1967).
- **System Blueprint:** Arquiteturas Intel Core e AMD Zen utilizam versões altamente otimizadas deste algoritmo.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular um ciclo de execução OoO para um conjunto de instruções com dependências RAW e WAR, demonstrando o uso de Register Renaming.
- [ ] Definir um banco de registros e uma tabela de renomeação.
- [ ] Simular o despacho de 3 instruções.
- [ ] Mostrar como a segunda instrução pode ser executada antes da primeira se os operandos estiverem prontos.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Dependencias de Dados e Hazards Dinamicos - Teoria e Fundamentos]]
- [[./Arquitetura de Reservation Stations e ROB - Funcionamento Interno e Arquitetura]]
- [[./Explosao de Area e Consumo Energetico em OoO - Casos de Falha e Analise Amortizada]]
- [[./Simulador de Ciclo de Maquina Tomasulo - Implementacao de Referencia e Benchmarks]]
