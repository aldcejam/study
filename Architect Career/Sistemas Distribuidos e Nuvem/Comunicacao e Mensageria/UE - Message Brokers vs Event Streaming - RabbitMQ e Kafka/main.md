---
tema: Message Brokers vs Event Streaming - RabbitMQ e Kafka
tipo: unidade-estudo
tags: [mensageria, eventos, kafka, rabbitmq, arquitetura]
---
# 🧪 UE - Message Brokers vs Event Streaming - RabbitMQ e Kafka

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Quando sistemas interagem de forma síncrona, a queda de um serviço afeta todos os dependentes (Cascading Failures) e cria acoplamento temporal. Mensageria resolve isso. Porém, confundir uma Fila Tradicional (Message Broker) com um Log de Eventos (Event Stream) é um erro arquitetural fatal. Usar RabbitMQ para reprocessamento de longo prazo sobrecarrega a memória (pois a mensagem é deletada após o ack). Usar Kafka para tarefas transacionais simples de fila de trabalho (Work Queues) gera extrema complexidade no gerenciamento de partições e balanceamento. Entender a diferença do modelo "Smart Broker / Dumb Consumer" para o "Dumb Broker / Smart Consumer" define a resiliência e a semântica de entrega da sua arquitetura.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[AMQP vs Append-Only Log]:** O mecanismo de filas efêmeras, roteamento e acknowledgments vs o conceito de log imutável persistente no disco com offsets gerenciados por consumidores.
2. **[Particionamento e Consumer Groups]:** Como a escalabilidade é alcançada através de shards lógicos (Kafka Partitions) e a garantia de concorrência massiva sem race conditions.
3. **[Garantias de Entrega e Ordenação]:** As semânticas matemáticas: At-Most-Once, At-Least-Once e a utopia de Exactly-Once (Idempotência vs Transações). O problema de Head-of-Line em ordens rígidas.
4. **[Rebalanceamento e Split-Brain]:** O comportamento caótico quando nós ou brokers perdem heartbeat, forçando eleições de líderes (ZooKeeper/KRaft) e reatribuição de partições.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Livro Clássico:** *Designing Data-Intensive Applications* (Martin Kleppmann) - Capítulos sobre Event Streams.
- **System Blueprint:** O log original criado pelo LinkedIn ("The Log: What every software engineer should know about real-time data's unifying abstraction" - Jay Kreps).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular um cenário de perda de conexão e rebalanceamento em um cluster Kafka validando se ocorre duplicação de mensagens (teste de idempotência).
- [ ] Configurar ambiente de isolamento (Docker Compose com Kafka e múltiplos producers/consumers).
- [ ] Implementar a lógica core com processamento de transações financeiras simples (At-Least-Once).
- [ ] Injetar carga/falha e coletar métricas de comportamento: Matar aleatoriamente 1 de 3 consumers no meio do pico de processamento.
- [ ] Instrumentar e medir quanto tempo dura a pausa do "Stop The World" do rebalanceamento do Kafka.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Modelos Formais de Log vs AMQP - Teoria e Fundamentos]]
- [[./Mecanica de Storage e Indexacao de Offsets no Disco - Funcionamento Interno e Arquitetura]]
- [[./Consumer Lag e Efeito Avalanche em Rebalanceamentos - Casos de Falha e Análise Amortizada]]
- [[./Mecanismos de Exactly-Once e Transacionalidade - Implementacao de Referencia e Benchmarks]]
