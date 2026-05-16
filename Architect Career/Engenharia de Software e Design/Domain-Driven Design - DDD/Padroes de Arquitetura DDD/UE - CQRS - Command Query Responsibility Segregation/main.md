---
tema: CQRS - Command Query Responsibility Segregation
tipo: unidade-estudo
tags: [arquitetura, performance, design, ddd]
---
# 🧪 UE - CQRS - Command Query Responsibility Segregation

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas complexos, o modelo usado para escrever dados (Commands) é muitas vezes diferente do modelo necessário para ler dados (Queries). O problema é tentar forçar um único modelo de banco de dados para ambas as tarefas, o que gera queries lentas e um domínio poluído. **CQRS** resolve isso separando a responsabilidade de escrita (focada em regras de negócio e integridade) da responsabilidade de leitura (focada em performance e UI). Sem CQRS, você terá dificuldades gigantescas para escalar a performance de leitura sem comprometer a segurança da escrita.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Comandos vs Consultas]:** Por que comandos não devem retornar dados (exceto talvez um ID) e consultas não devem alterar estado.
2. **[Modelos de Dados Separados]:** O uso de uma base de dados relacional para escrita e um ElasticSearch ou Redis para leitura.
3. **[Sincronizacao de Modelos]:** Como o modelo de leitura é atualizado (consistência eventual) via eventos.
4. **[Task-Based UI]:** Por que o CQRS combina melhor com interfaces que expressam intenção (ex: `MudarEndereco`) do que com CRUDs genéricos.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers classics):
- **Paper Clássico/Livro:** *CQRS Journey* (Microsoft Patterns & Practices).
- **System Blueprint:** Arquitetura de sistemas de alta escala como LMAX Disruptor e sistemas bancários modernos.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um pequeno sistema onde o comando de "CriarProduto" escreve em uma tabela SQL e um evento atualiza uma projeção de leitura em memória (um dicionário otimizado para busca).
- [ ] Criar o Command e o CommandHandler.
- [ ] Criar o Query e o QueryHandler.
- [ ] Demonstrar o descolamento entre os dois modelos.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Segregacao de Responsabilidades e Intent-Driven Design - Teoria e Fundamentos]]
- [[./Projecoes e Sincronizacao de Modelos de Leitura - Funcionamento Interno e Arquitetura]]
- [[./Consistencia Eventual e o Desafio da Experiencia do Usuario - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de CQRS com MediatR e RabbitMQ - Implementacao de Referencia e Benchmarks]]
