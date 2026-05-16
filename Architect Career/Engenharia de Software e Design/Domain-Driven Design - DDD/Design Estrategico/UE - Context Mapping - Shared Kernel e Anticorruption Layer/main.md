---
tema: Context Mapping - Shared Kernel e Anticorruption Layer
tipo: unidade-estudo
tags: [ddd, arquitetura, design-estrategico, integraçao]
---
# 🧪 UE - Context Mapping - Shared Kernel e Anticorruption Layer

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas grandes, o modelo de domínio é inevitavelmente dividido em vários **Bounded Contexts**. O problema é: como esses contextos se integram sem que um corrompa o outro? O **Shared Kernel** resolve a necessidade de compartilhar código comum, mas com o custo de acoplamento forte. A **Anticorruption Layer (ACL)** resolve a integração com sistemas legados ou externos, garantindo que a "sujeira" do modelo externo não vaze para o seu domínio limpo. Sem o mapeamento de contextos, sua arquitetura se torna um "Big Ball of Mud".

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Relacoes entre Contextos]:** Upstream (provedor) vs Downstream (consumidor).
2. **[Shared Kernel]:** Quando é aceitável compartilhar uma biblioteca de domínio entre duas equipes.
3. **[Anticorruption Layer (ACL)]:** O uso de tradutores (Mappers) e adaptadores para isolar o domínio.
4. **[Separate Ways]:** A decisão estratégica de não integrar e duplicar lógica para ganhar independência.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Implementing Domain-Driven Design* (Vaughn Vernon) - Capítulo sobre Context Mapping.
- **System Blueprint:** Estratégias de migração de legados usando o padrão **Strangler Fig**.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Projetar uma ACL que consome uma API legada de "LegadoUsuarios" e a traduz para o objeto de domínio "MembroDaComunidade".
- [ ] Definir o objeto de domínio limpo.
- [ ] Criar a interface da ACL.
- [ ] Implementar o tradutor que lida com campos estranhos ou nulos do sistema legado.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Relacoes Upstream-Downstream e a Politica de Integracao - Teoria e Fundamentos]]
- [[./Implementacao de Tradutores e Adaptadores em ACL - Funcionamento Interno e Arquitetura]]
- [[./Explosao de Shared Kernels e Acoplamento Rigido - Casos de Falha e Analise Amortizada]]
- [[./Mapa de Contexto de um Sistema de E-commerce Real - Implementacao de Referencia e Benchmarks]]
