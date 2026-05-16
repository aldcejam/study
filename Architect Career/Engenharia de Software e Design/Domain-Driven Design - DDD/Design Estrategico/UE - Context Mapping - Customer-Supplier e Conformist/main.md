---
tema: Context Mapping - Customer-Supplier e Conformist
tipo: unidade-estudo
tags: [ddd, arquitetura, design-estrategico, gestao]
---
# 🧪 UE - Context Mapping - Customer-Supplier e Conformist

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Integrar sistemas não é apenas um problema técnico, é um problema político e organizacional. O problema é como lidar com a dependência entre equipes. No padrão **Customer-Supplier**, a equipe Upstream (Supplier) aceita requisitos da equipe Downstream (Customer). No padrão **Conformist**, a equipe Downstream não tem poder de negociação e precisa se adaptar totalmente ao modelo do Upstream. Sem entender essas relações, você tentará implementar soluções técnicas (como ACLs) onde o problema é na verdade a falta de alinhamento entre as pessoas.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Dinamica de Equipes]:** Como a estrutura organizacional (Lei de Conway) afeta o design do software.
2. **[Customer-Supplier]:** O uso de testes de aceitação conjuntos (Contract Testing) para garantir a integridade da integração.
3. **[Conformist]:** Quando vale a pena simplesmente copiar o modelo de uma API externa em vez de criar uma camada de isolamento cara.
4. **[Open Host Service (OHS)]:** Como um provedor simplifica a vida dos consumidores definindo um protocolo padrão (ex: REST API bem documentada).
5. **[Published Language]:** O uso de padrões de troca de dados (JSON Schema, Protobuf) para reduzir o atrito entre contextos.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Strategic Monoliths and Microservices* (Vaughn Vernon and Tomasz Jaskula).
- **System Blueprint:** A relação entre o sistema de "Pagamentos" (Supplier) e o "Checkout" (Customer) em grandes e-commerces.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Definir um contrato de API (Consumer-Driven Contract) usando uma ferramenta como **Pact** para simular uma relação Customer-Supplier.
- [ ] Escrever a expectativa do consumidor.
- [ ] Verificar se o provedor atende a essa expectativa.
- [ ] Simular uma quebra de contrato e ver o teste falhar.

## 🗃️ Notas Heutagógicas Atômicas
- [[./A Lei de Conway e a Estrutura de Bounded Contexts - Teoria e Fundamentos]]
- [[./Implementacao de Consumer-Driven Contracts (Pact) - Funcionamento Interno e Arquitetura]]
- [[./Custos de Manutencao de ACL vs Conformidade Direta - Casos de Falha e Analise Amortizada]]
- [[./Guia de Negociacao Técnica entre Equipes de Plataforma - Implementacao de Referencia e Benchmarks]]
