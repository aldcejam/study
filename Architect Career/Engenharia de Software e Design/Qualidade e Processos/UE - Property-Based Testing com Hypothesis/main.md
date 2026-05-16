---
tema: Property-Based Testing com Hypothesis
tipo: unidade-estudo
tags: [testes, qualidade, algoritmos, automação]
---
# 🧪 UE - Property-Based Testing com Hypothesis

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Testes unitários tradicionais baseiam-se em exemplos manuais (ex: `test_soma(1, 1)`). O problema é que o desenvolvedor tende a testar apenas os casos que ele imaginou, deixando passar bugs de borda (edge cases) bizarros. **Property-Based Testing (PBT)** resolve isso definindo propriedades que devem ser sempre verdadeiras (ex: `soma(a, b) == soma(b, a)`) e deixando a ferramenta gerar centenas de entradas aleatórias para tentar quebrar o código. **Hypothesis** é a ferramenta líder que faz isso e ainda realiza o **Shrinking** (reduzir uma falha complexa ao menor exemplo possível). Sem PBT, você nunca terá 100% de confiança em algoritmos críticos.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Exemplo vs Propriedade]:** A mudança de mentalidade de testar "valor X" para testar "invariante Y".
2. **[Geradores de Dados (Strategies)]:** Como definir espaços de busca para inteiros, strings, listas e objetos customizados.
3. **[Shrinking]:** O processo automático de simplificar o caso de falha para facilitar o debug.
4. **[Invariantes Comuns]:** Round-trip (encode/decode), Idempotência, Comutatividade e Oráculo (comparar com implementação simples).
5. **[Falsificação]:** Como o PBT atua como um "adversário" que tenta ativamente quebrar seu código.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *QuickCheck: A Lightweight Tool for Random Testing of Haskell Programs* (Claessen and Hughes, 2000).
- **System Blueprint:** Documentação oficial do Hypothesis (Python) e jqwik (Java).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Escrever uma função de ordenação de lista (ou usar uma pronta) e criar um teste de propriedade que valide que a lista resultante está ordenada e contém os mesmos elementos da lista original.
- [ ] Usar `@given(st.lists(st.integers()))`.
- [ ] Definir a propriedade de ordenação.
- [ ] Inserir um bug proposital no código e ver o Hypothesis encontrar o menor contra-exemplo.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Invariantes Matematicas e o Pensamento Baseado em Propriedades - Teoria e Fundamentos]]
- [[./Algoritmos de Shrinking e Geracao de Dados Estocasticos - Funcionamento Interno e Arquitetura]]
- [[./Explosao de Espaco de Busca e Performance de Testes - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de PBT em Sistemas de Alta Criticidade - Implementacao de Referencia e Benchmarks]]
