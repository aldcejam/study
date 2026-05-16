---
tema: Aggregates e Aggregate Roots - Design de Consistencia
tipo: unidade-estudo
tags: [ddd, design-tatico, concorrência, arquitetura]
---
# 🧪 UE - Aggregates e Aggregate Roots - Design de Consistencia

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em um sistema complexo, como garantir que um grupo de objetos relacionados permaneça consistente após uma alteração? O problema é o acoplamento excessivo e as condições de corrida. **Aggregates** resolvem isso agrupando entidades e VOs sob uma única **Aggregate Root**. Toda alteração deve passar pela raiz, que garante as invariantes do negócio. Sem agregados bem definidos, suas transações de banco de dados serão gigantescas, causando deadlocks e baixa performance, ou seu sistema terá dados inconsistentes.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Invariantes de Negocio]:** As regras que devem ser verdadeiras o tempo todo (ex: "a soma dos itens do pedido não pode exceder o limite do cliente").
2. **[Aggregate Root]:** O único objeto do grupo que pode ser referenciado externamente.
3. **[Design de Agregados Pequenos]:** Por que agregados gigantes são um erro de escalabilidade.
4. **[Referencia por ID]:** Por que agregados não devem conter referências de objetos a outros agregados, mas apenas seus IDs.
5. **[Consistencia Eventual vs Imediata]:** Quando usar transações ACID (dentro do agregado) e quando usar eventos (entre agregados).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Effective Aggregate Design* (Vaughn Vernon).
- **System Blueprint:** Implementação de agregados em frameworks como Axon (Java) ou Orleankka (.NET).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Modelar um Agregado de `Pedido` (Order) com seus `Itens` e garantir que o status do pedido só mude para "Pago" se houver pelo menos um item.
- [ ] Implementar a classe `Order` (Root) e `OrderItem`.
- [ ] Tornar a lista de itens privada e só acessível via métodos da `Order`.
- [ ] Implementar a regra de validação na raiz.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Invariantes e Limites de Transacao - Teoria e Fundamentos]]
- [[./Estrategias de Persistencia de Agregados com ORMs - Funcionamento Interno e Arquitetura]]
- [[./Conflitos de Concorrencia em Agregados Gigantes - Casos de Falha e Analise Amortizada]]
- [[./Guia de Decomposicao de Agregados Complexos - Implementacao de Referencia e Benchmarks]]
