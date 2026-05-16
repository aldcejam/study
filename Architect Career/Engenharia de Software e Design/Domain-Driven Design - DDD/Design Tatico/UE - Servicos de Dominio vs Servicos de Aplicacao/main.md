---
tema: Servicos de Dominio vs Servicos de Aplicacao
tipo: unidade-estudo
tags: [ddd, design-tatico, arquitetura, padrões]
---
# 🧪 UE - Servicos de Dominio vs Servicos de Aplicacao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Onde colocar a lógica que não pertence naturalmente a uma única Entidade ou Agregado? O problema é o surgimento de "Objetos de Deus" ou lógica espalhada. **Serviços de Domínio** resolvem operações que envolvem múltiplos agregados ou dependem de ferramentas externas de domínio (ex: um `CalculadorDeFrete`). **Serviços de Aplicação** resolvem a coordenação da tarefa (orquestração), como carregar um objeto do banco, chamar o domínio e salvar o resultado. Sem essa distinção, sua camada de domínio ficará poluída com detalhes técnicos (infra) ou sua camada de aplicação ficará cheia de regras de negócio (leaking logic).

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Statelessness]:** Por que serviços devem ser sem estado e agir apenas sobre os dados passados.
2. **[Servicos de Dominio]:** Quando a lógica envolve múltiplos agregados e não "cabe" em nenhum deles.
3. **[Servicos de Aplicacao]:** O papel de "Maestro" da aplicação: segurança, transações e orquestração.
4. **[Injeção de Dependência]:** Como serviços de aplicação injetam interfaces de infraestrutura (repositórios) nos serviços de domínio.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Domain-Driven Design* (Eric Evans) - Capítulo sobre Services.
- **System Blueprint:** O padrão **Use Case** da Clean Architecture como equivalente ao Serviço de Aplicação.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um `TransferService` (Domínio) que realiza a lógica de transferência entre duas contas e um `ProcessTransferUseCase` (Aplicação) que coordena a persistência e notificações.
- [ ] Criar as Entidades `Account`.
- [ ] Implementar o serviço de domínio `TransferService.execute(from, to, amount)`.
- [ ] Implementar o serviço de aplicação que usa um `AccountRepository`.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Orquestracao vs Logica de Negocio - Teoria e Fundamentos]]
- [[./Injecao de Interfaces e Inversao de Controle - Funcionamento Interno e Arquitetura]]
- [[./Vazamento de Dominio para a Camada de Aplicacao - Casos de Falha e Analise Amortizada]]
- [[./Padrao Command Handler e a Camada de Aplicacao - Implementacao de Referencia e Benchmarks]]
