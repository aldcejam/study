---
tema: Arquitetura Hexagonal - Ports e Adapters
tipo: unidade-estudo
tags: [arquitetura, ddd, design, padrões]
---
# 🧪 UE - Arquitetura Hexagonal - Ports e Adapters

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como impedir que o seu código de negócio fique "refém" de tecnologias externas (Banco de Dados, UI, APIs de terceiros)? O problema é o acoplamento tecnológico que torna o código difícil de testar e caro de mudar. A **Arquitetura Hexagonal (Ports and Adapters)** resolve isso colocando o domínio no centro e definindo **Ports** (interfaces) para comunicação externa. Os **Adapters** implementam essas interfaces para tecnologias específicas. Sem isso, você não consegue trocar seu banco de dados ou rodar seus testes unitários sem uma conexão de rede ativa.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[O Centro do Hexagono]:** O modelo de domínio e os serviços de aplicação que não dependem de nada externo.
2. **[Driving Ports (Primary)]:** Interfaces chamadas por atores externos (ex: APIs REST, CLI).
3. **[Driven Ports (Secondary)]:** Interfaces que o domínio usa para falar com o mundo (ex: Repositórios, Gateways de SMS).
4. **[Adapters]:** A implementação concreta (ex: `PostgresUserRepository`, `RestCustomerController`).
5. **[Inversao de Dependencia]:** Como garantir que a dependência sempre aponte para o centro do hexágono.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Hexagonal Architecture* (Alistair Cockburn, 2005).
- **System Blueprint:** Estrutura de diretórios recomendada para projetos hexagonais em Go e Java.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Criar uma aplicação de "Notificacoes" onde o domínio define um `NotificationPort` e existem dois adapters: um que imprime no console (mock) e outro que envia um email real.
- [ ] Definir a interface `NotificationPort`.
- [ ] Implementar a lógica de domínio que usa o port.
- [ ] Trocar o adapter em tempo de execução para mostrar a desacoplagem.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Desacoplamento Tecnológico e Testabilidade - Teoria e Fundamentos]]
- [[./Design de Ports e Implementacao de Adapters - Funcionamento Interno e Arquitetura]]
- [[./Complexidade de Mapeamento de Objetos entre Camadas - Casos de Falha e Analise Amortizada]]
- [[./Comparativo Hexagonal vs Layered Architecture - Implementacao de Referencia e Benchmarks]]
