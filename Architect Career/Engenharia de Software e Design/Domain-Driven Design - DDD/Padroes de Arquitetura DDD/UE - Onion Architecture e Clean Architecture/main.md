---
tema: Onion Architecture e Clean Architecture
tipo: unidade-estudo
tags: [arquitetura, ddd, design, padrões]
---
# 🧪 UE - Onion Architecture e Clean Architecture

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como organizar um projeto de software para que ele seja fácil de entender e manter por anos? O problema é a tendência de misturar lógica de negócio com frameworks e bancos de dados (ex: anotações do Hibernate no meio da Entidade). A **Onion Architecture** e a **Clean Architecture** resolvem isso organizando o código em camadas circulares concêntricas. A regra de ouro é: as dependências só podem apontar para dentro. Sem essa disciplina, seu projeto se tornará uma massa de código impossível de testar e onde cada pequena mudança quebra dez outras coisas.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Regra de Dependencia]:** O código nas camadas internas não pode saber nada sobre as camadas externas.
2. **[Camada de Entidades (The Core)]:** Objetos de negócio puros e regras que mudam raramente.
3. **[Camada de Use Cases]:** Orquestração de fluxos de dados de e para as entidades.
4. **[Interface Adapters]:** Tradutores entre o formato dos casos de uso e o formato externo (ex: ViewModels, Presenters).
5. **[Frameworks e Drivers]:** A camada mais externa onde residem as ferramentas mutáveis (UI, Web, DB).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Clean Architecture: A Craftsman's Guide to Software Structure and Design* (Robert C. Martin).
- **System Blueprint:** A implementação de referência de Jeffrey Palermo para Onion Architecture.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Estruturar um projeto "Todo-List" simples em 4 camadas (Domain, Application, Interface, Infrastructure) e garantir que a camada Domain não tenha nenhuma importação de frameworks externos.
- [ ] Criar a estrutura de pastas.
- [ ] Implementar um Caso de Uso que salva uma tarefa.
- [ ] Validar a direção das dependências usando uma ferramenta de análise estática ou apenas inspeção visual.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Camadas Concentricas e a Regra de Ouro da Dependencia - Teoria e Fundamentos]]
- [[./Mapeamento de Entidades entre Camadas (Demos/DTOs) - Funcionamento Interno e Arquitetura]]
- [[./Overhead de Boilerplate em Arquiteturas Limpas - Casos de Falha e Analise Amortizada]]
- [[./Guia de Migracao de Legado para Clean Architecture - Implementacao de Referencia e Benchmarks]]
