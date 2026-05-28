---
tema: Padrões de Mensageria - Saga e Outbox Pattern
tipo: unidade-estudo
tags: [distribuido, padroes, transacoes, resiliencia, saga]
---
# 🧪 UE - Padrões de Mensageria - Saga e Outbox Pattern

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O calcanhar de Aquiles de qualquer sistema de microsserviços é a Transação Distribuída (Dual-Write Problem). Como salvar os dados no seu banco de dados e simultaneamente colocar um evento num broker de fila garantindo que nenhum dos dois falhe deixando o sistema inconsistente? Bloquear a rede inteira aguardando o commit (via 2-Phase Commit / 2PC) mata o throughput e a disponibilidade do sistema. Ignorar o problema resulta em perda silenciosa de dados ou "Orphan States". O Transactional Outbox Pattern e o Saga Pattern resolvem o problema trocando a consistência transacional ACID imediata pela Consistência Eventual, usando o banco como log transacional primário.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Transactional Outbox Pattern e Dual-Write]:** Em vez de salvar e enviar ao RabbitMQ/Kafka, a transação SQL atômica salva num DB na tabela de 'negócio' e na tabela de 'eventos_outbox'. CDC (Change Data Capture) ou Polling assíncrono enviam os eventos posteriormente.
2. **[Saga Pattern (Coreografia vs Orquestração)]:** Fluxo de eventos assíncronos (eventos locais) ou comandos direcionados através de um Orchestrator stateful encadeando processos sem travar o banco do vizinho.
3. **[Transações de Compensação (Undo Logs)]:** Em caso de falha em etapa `N` de uma Saga distribuída, é fundamental enviar comandos de regressão nas etapas `N-1, N-2, ...` para reverter transações locais efetuadas.
4. **[Idempotência Absoluta]:** Padrões distribuídos assíncronos garantem *At-Least-Once*, ou seja, duplicações de mensagens são certas na lei das probabilidades do particionamento e timeouts. O sistema deve tratar redundâncias de forma invisível.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico:** *SAGAS* (Hector Garcia-Molina and Kenneth Salem, 1987).
- **System Blueprint:** Debezium (ferramenta padrão ouro para CDC conectada ao Log do Banco) e arquitetura de pagamentos globais do Uber e Netflix.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Criar um simulador de Transação Distribuída Falha usando CDC (Change Data Capture) para despachar comandos em SAGA orquestrada.
- [ ] Configurar ambiente de isolamento (PostgreSQL + Debezium + Kafka + Serviços A e B em Docker).
- [ ] Implementar a lógica core de um pedido gerando evento Outbox via Debezium para iniciar a saga e aplicar uma transação no Serviço B.
- [ ] Injetar carga/falha: Gerar um pânico deliberado na transação do Serviço B.
- [ ] Disparar e medir a chamada do Evento de Compensação que deve anular os dados no Banco do Serviço A e colocar a ordem como `CANCELED`.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Transacoes Distribuidas e o Dual-Write Problem - Teoria e Fundamentos]]
- [[./Saga Orquestrada vs Coreografada e Change Data Capture - Funcionamento Interno e Arquitetura]]
- [[./Tratamento de Transacoes Compensatorias e Race Conditions - Casos de Falha e Análise Amortizada]]
- [[./Debezium Kafka Connect como Outbox Relay Garantido - Implementacao de Referencia e Benchmarks]]
