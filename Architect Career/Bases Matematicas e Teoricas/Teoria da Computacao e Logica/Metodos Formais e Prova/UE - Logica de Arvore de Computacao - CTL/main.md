---
tema: Logica de Arvore de Computacao - CTL
tipo: unidade-estudo
tags: [matematica, logica, metodos-formais, verificação]
---
# 🧪 UE - Logica de Arvore de Computacao - CTL

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Enquanto a LTL olha para um único caminho de execução, a **CTL (Computation Tree Logic)** olha para a **árvore de todas as execuções possíveis** a partir de um estado. O problema que a CTL resolve é expressar propriedades como "existe um caminho onde o sistema eventualmente se recupera" ou "em todos os caminhos, se o botão for pressionado, a luz acende". Esta distinção é crucial para sistemas não-determinísticos onde a possibilidade de um evento futuro importa tanto quanto a sua certeza.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Quantificadores de Caminho]:** $A$ (All - em todos os caminhos) e $E$ (Exists - em pelo menos um caminho).
2. **[Operadores de Estado e Caminho]:** Combinações como $AF$ (Inevitabilidade), $EF$ (Possibilidade), $AG$ (Invariante em todos os caminhos) e $EG$ (Possibilidade de persistência).
3. **[Sistemas de Transicao de Kripke]:** Como a CTL é avaliada sobre grafos de estados.
4. **[CTL vs LTL]:** Entender a expressividade: problemas que a CTL resolve e a LTL não, e vice-versa.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Design and Synthesis of Synchronization Skeletons Using Temporal Logic* (Clarke and Emerson, 1981).
- **System Blueprint:** Verificadores de hardware como o NuSMV e verificação de protocolos de cache em processadores.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Modelar um sistema de semáforo simples no model checker **NuSMV** e provar uma propriedade CTL de que "é possível retornar ao estado Verde a partir de qualquer outro estado".
- [ ] Definir as variáveis de estado e as transições (Verde -> Amarelo -> Vermelho).
- [ ] Escrever a especificação CTL `SPEC AG (EF state = Verde)`.
- [ ] Verificar o modelo e analisar o resultado.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Quantificadores e Caminhos de Execucao - Teoria e Fundamentos]]
- [[./Algoritmos de Verificacao de Modelos CTL - Funcionamento Interno e Arquitetura]]
- [[./CTL e o Problema da Explosao de Estados - Casos de Falha e Analise Amortizada]]
- [[./Verificacao de Hardware com NuSMV - Implementacao de Referencia e Benchmarks]]
