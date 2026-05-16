---
tema: Identificacao de Bounded Contexts
tipo: unidade-estudo
tags: [ddd, arquitetura, design-estrategico]
---
# 🧪 UE - Identificacao de Bounded Contexts

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O maior gargalo em sistemas corporativos não é a performance do código, mas a **ambiguidade semântica**. Quando um termo como "Produto" significa coisas diferentes para o time de Vendas e para o time de Logística, mas ambos usam a mesma classe no código, o sistema colapsa sob o peso de regras de negócio conflitantes e acoplamento descontrolado. Ignorar os Bounded Contexts leva ao "Big Ball of Mud", onde qualquer mudança quebra partes imprevistas do sistema.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Semântica e Contexto]:** Compreender como o significado de um modelo é válido apenas dentro de uma fronteira delimitada.
2. **[Critérios de Decomposição]:** Identificar limites baseados em processos de negócio, autonomia de times e ciclo de vida dos dados.
3. **[Fronteiras e Interfaces]:** Definir como os modelos "viajam" entre contextos (Dists, Mappings e Translators).
4. **[Trade-offs de Granularidade]:** Balancear entre contextos muito grandes (acoplamento) vs muito pequenos (complexidade de integração/chatty APIs).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Domain-Driven Design* (Eric Evans, 2003) - Parte II: Strategic Design.
- **System Blueprint:** Decomposição de monolitos para microsserviços em empresas como Netflix e Amazon, onde cada serviço é idealmente um Bounded Context.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Realizar uma sessão de **Event Storming** para um domínio complexo (ex: E-commerce com Devoluções e Reembolsos) e mapear fisicamente as fronteiras onde os termos mudam de significado.
- [ ] Listar todos os Domain Events do processo.
- [ ] Agrupar eventos por afinidade e mudança de ator/responsabilidade.
- [ ] Desenhar o Context Map preliminar identificando os relacionamentos (Upstream/Downstream).

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Semantica e Limites de Modelagem - Teoria e Fundamentos]]
- [[./Estrategias de Decomposicao de Dominio - Funcionamento Interno e Arquitetura]]
- [[./Ambiguidade e Conflitos de Regras - Casos de Falha e Analise Amortizada]]
- [[./Mapeamento de Contextos Reais - Implementacao de Referencia e Benchmarks]]
