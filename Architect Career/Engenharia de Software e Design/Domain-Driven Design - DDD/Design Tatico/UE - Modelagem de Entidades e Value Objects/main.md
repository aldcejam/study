---
tema: Modelagem de Entidades e Value Objects
tipo: unidade-estudo
tags: [ddd, design-tatico, oop, modelagem]
---
# 🧪 UE - Modelagem de Entidades e Value Objects

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como diferenciar um objeto que representa um conceito único no tempo de um objeto que representa apenas um valor? O problema é a confusão entre identidade e estado. **Entidades** resolvem a necessidade de rastrear algo pelo seu ID (ex: um `Pedido` que muda de status). **Value Objects (VO)** resolvem a necessidade de representar medidas ou descrições imutáveis (ex: um `Preco` ou um `Endereco`). Sem essa distinção, seu código terá bugs de mutação acidental e lógica de comparação complexa espalhada por todo o sistema.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Identidade vs Igualdade]:** Por que duas Entidades com o mesmo estado são diferentes, mas dois VOs com o mesmo estado são iguais.
2. **[Imutabilidade de Value Objects]:** Por que nunca devemos alterar um VO, mas sim substituí-lo.
3. **[Auto-Validacao]:** O padrão de não permitir a criação de objetos inválidos (Always-Valid model).
4. **[Comportamento Rico]:** Onde colocar a lógica: no objeto (Model Rico) ou em serviços (Model Anêmico).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Domain-Driven Design Reference* (Eric Evans).
- **System Blueprint:** Implementação de VOs no C# (record types) e Java (Value Types/Records).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um Value Object `EmailAddress` que se valida no construtor e uma Entidade `Usuario` que utiliza esse VO.
- [ ] Criar o VO imutável.
- [ ] Implementar a lógica de igualdade (equals/hashcode ou equivalente).
- [ ] Tentar criar um email inválido e garantir que uma exceção seja lançada.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Identidade de Objetos e o Ciclo de Vida de Entidades - Teoria e Fundamentos]]
- [[./Padroes de Imutabilidade e Value Objects - Funcionamento Interno e Arquitetura]]
- [[./Modelos Anemicos e o Anti-padrao de Getters e Setters - Casos de Falha e Analise Amortizada]]
- [[./Refatoracao de Primitivos para Value Objects (Primitive Obsession) - Implementacao de Referencia e Benchmarks]]
