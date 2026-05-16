---
tema: Linguagem Ubiqua e Glossario de Dominio
tipo: unidade-estudo
tags: [ddd, engenharia-de-software, design, modelagem]
---
# 🧪 UE - Linguagem Ubiqua e Glossario de Dominio

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O maior risco em projetos de software não é a tecnologia, é a má comunicação. O problema é que desenvolvedores usam termos técnicos (ex: `UserRecord`, `UpdateStatus`) e especialistas de negócio usam termos de domínio (ex: `Cliente`, `Ativacao`). Isso gera um "gap" de tradução que causa bugs e funcionalidades erradas. A **Linguagem Ubíqua** resolve isso forçando o uso de um único idioma falado por todos, refletido diretamente no código. Sem uma linguagem ubíqua, você está construindo uma Torre de Babel tecnológica que vai desabar sob a própria complexidade.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Glossario de Dominio]:** A extração de substantivos e verbos das conversas com especialistas.
2. **[Conceitos vs Termos Técnicos]:** Por que o código deve ler como a descrição do negócio (ex: `account.withdraw(amount)` em vez de `service.updateBalance(id, -100)`).
3. **[Evolucao da Linguagem]:** Como o glossário muda conforme o entendimento do domínio se aprofunda.
4. **[Linguagem Ubiqua por Contexto]:** Por que o termo "Produto" pode significar coisas diferentes em "Vendas" e "Entrega".

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Domain-Driven Design: Tackling Complexity in the Heart of Software* (Eric Evans, 2003) - Capítulo 1 e 2.
- **System Blueprint:** O processo de **Event Storming** para descoberta de linguagem e processos de negócio.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Realizar uma mini-sessão de modelagem para um domínio de "Biblioteca" e transformar o glossário resultante em uma interface de código expressiva.
- [ ] Listar 5 termos de negócio.
- [ ] Escrever um trecho de código em Python ou Java que use esses termos sem vazar detalhes de implementação técnica.
- [ ] Refatorar um código "técnico" genérico para um código "rico em domínio".

## 🗃️ Notas Heutagógicas Atômicas
- [[./Comunicacao e o Custo da Traducao de Conceitos - Teoria e Fundamentos]]
- [[./Metodologia de Event Storming para Descoberta de Linguagem - Funcionamento Interno e Arquitetura]]
- [[./Ambiguidade de Termos em Dominios Complexos - Casos de Falha e Analise Amortizada]]
- [[./Guia de Refatoracao para Codigo Expressivo - Implementacao de Referencia e Benchmarks]]
