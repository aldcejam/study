---
tema: Domain Events - Publicacao e Consumo
tipo: unidade-estudo
tags: [ddd, design-tatico, arquitetura, eda]
---
# 🧪 UE - Domain Events - Publicacao e Consumo

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> "Algo aconteceu que outros especialistas de negócio se importam." Como notificar diferentes partes do sistema sobre uma mudança de estado sem acoplá-las rigidamente? O problema é a dependência direta (ex: a `Order` chamando o `EmailService`). **Domain Events** resolvem isso permitindo que o Agregado publique um evento (ex: `OrderPlaced`) e qualquer interessado reaja a ele. Isso possibilita a consistência eventual e a extensibilidade do sistema. Sem eventos de domínio, seu código será um emaranhado de chamadas síncronas que tornam o sistema frágil e lento.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Definicao de Domain Event]:** Um objeto imutável que representa um fato passado no domínio.
2. **[Publicacao de Eventos]:** Onde disparar o evento (dentro do Agregado ou no Serviço de Aplicação?).
3. **[Consumo e Side Effects]:** Como os handlers processam o evento (ex: enviar email, atualizar uma projeção de leitura).
4. **[Padrao Outbox]:** Como garantir que o evento seja enviado mesmo se a transação do banco de dados falhar (ou vice-versa).
5. **[Idempotencia]:** Garantir que processar o mesmo evento duas vezes não cause problemas.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Domain Events: Salvation* (Udi Dahan).
- **System Blueprint:** Implementação de eventos no Spring Boot (ApplicationEventPublisher) ou no MediatR (.NET).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um sistema simples de eventos em memória onde a criação de um `User` dispara um evento `UserRegistered` que é ouvido por um `WelcomeEmailHandler`.
- [ ] Criar a classe do evento com timestamp e dados do usuário.
- [ ] Implementar um `EventBus` simples.
- [ ] Demonstrar a execução assíncrona do handler.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Fatos Passados e a Semantica de Eventos - Teoria e Fundamentos]]
- [[./Padrao Transactional Outbox e Mensageria Confiavel - Funcionamento Interno e Arquitetura]]
- [[./Explosao de Eventos e Loop de Recursao - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Event Sourcing Minimalista - Implementacao de Referencia e Benchmarks]]
