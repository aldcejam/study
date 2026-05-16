---
tema: Semantic Analysis e Symbol Table Design
tipo: unidade-estudo
tags: [compiladores, algoritmos, linguagens, baixo-nivel]
---
# 🧪 UE - Semantic Analysis e Symbol Table Design

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Um código pode estar gramaticalmente correto (sintaxe), mas ser um lixo lógico (semântica). O problema é garantir que variáveis sejam declaradas antes do uso, que os tipos sejam compatíveis e que o escopo seja respeitado. A **Análise Semântica** resolve isso percorrendo a árvore sintática (AST) e a **Symbol Table (Tabela de Símbolos)** é a estrutura de dados central que mantém o controle de tudo o que foi definido. Sem uma tabela de símbolos eficiente, a compilação de projetos grandes se torna impossível.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Verificacao de Tipos (Type Checking)]:** Regras de coerção, casting e polimerismo.
2. **[Escopo e Visibilidade]:** Como gerenciar variáveis locais, globais e membros de classes (Árvores de Escopo).
3. **[Design de Symbol Tables]:** O uso de tabelas hash encadeadas para buscas em $O(1)$ e como representar metadados complexos (assinaturas de funções, tipos genéricos).
4. **[Atributos de AST]:** Como "anotar" a árvore sintática com informações derivadas da análise semântica.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Compilers: Principles, Techniques, and Tools* (Dragon Book) - Capítulo sobre Semantic Analysis.
- **System Blueprint:** Implementação da Tabela de Símbolos no Clang/LLVM.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar uma Tabela de Símbolos simples em C++ que suporta escopos aninhados e detectar erros de "variável não declarada" ou "declaração duplicada".
- [ ] Criar uma classe `SymbolTable` com métodos `enterScope`, `exitScope`, `insert` e `lookup`.
- [ ] Implementar o encadeamento de escopos (parent pointer).
- [ ] Simular a análise de um pequeno trecho de código com blocos `{ }`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Gramaticas Atribuidas e Decoracao de AST - Teoria e Fundamentos]]
- [[./Arquitetura de Tabelas de Simbolos para Linguagens OO - Funcionamento Interno e Arquitetura]]
- [[./Lentidao de Busca em Projetos Massivos - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Verificador de Tipos Estatico - Implementacao de Referencia e Benchmarks]]
