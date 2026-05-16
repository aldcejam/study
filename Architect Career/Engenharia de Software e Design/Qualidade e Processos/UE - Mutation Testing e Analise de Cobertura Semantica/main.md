---
tema: Mutation Testing e Analise de Cobertura Semantica
tipo: unidade-estudo
tags: [testes, qualidade, engenharia-de-software, métricas]
---
# 🧪 UE - Mutation Testing e Analise de Cobertura Semantica

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Você tem 100% de cobertura de código (Line Coverage), mas seus testes são realmente bons? O problema é que a cobertura de linha só diz que o código foi executado, não que ele foi validado. **Mutation Testing** resolve isso inserindo pequenos bugs propositais no seu código original (mutantes), como trocar um `>` por `>=`. Se os seus testes continuarem passando mesmo com o bug, o mutante "sobreviveu", o que significa que seus testes são falhos. Sem Mutation Testing, você tem uma falsa sensação de segurança baseada em métricas de cobertura superficiais.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Operadores de Mutação]:** Troca de operadores aritméticos, lógicos e relacionais.
2. **[Mutantes Sobreviventes vs Mortos]:** A métrica de eficácia dos testes (Mutation Score).
3. **[Mutantes Equivalentes]:** O problema de mutações que não alteram o comportamento final do código e são impossíveis de "matar".
4. **[Custo Computacional]:** Por que rodar testes para cada mutante é lento e como otimizar (mutação seletiva).
5. **[Cobertura Semantica]:** A diferença entre executar uma linha e validar a lógica contida nela.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Mutation Testing: A Historical Retrospective* (Jia and Harman, 2011).
- **System Blueprint:** Ferramentas como **PITest** (Java), **Stryker** (JS/TS/C#) e **Mutmut** (Python).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Usar o **Stryker** ou **PITest** em um pequeno projeto e analisar o relatório de mutantes sobreviventes para identificar testes que não têm `asserts` significativos.
- [ ] Rodar a ferramenta de mutação no projeto.
- [ ] Encontrar um mutante que sobreviveu em uma lógica crítica.
- [ ] Escrever um novo teste que "mate" esse mutante.

## 🗃️ Notas Heutagógicas Atômicas
- [[./A Falacia da Cobertura de Linha e Metricas de Qualidade - Teoria e Fundamentos]]
- [[./Tipos de Operadores de Mutacao e Analise de Fluxo - Funcionamento Interno e Arquitetura]]
- [[./Explosao Combinatoria e Performance de Testes de Mutacao - Casos de Falha e Analise Amortizada]]
- [[./Integracao de Mutation Testing no Pipeline de CI-CD - Implementacao de Referencia e Benchmarks]]
