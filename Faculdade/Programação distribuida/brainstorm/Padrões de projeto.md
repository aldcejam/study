# Padrões de Sistemas Distribuídos — Explicação Detalhada

Baseado no livro **Patterns of Distributed Systems — Unmesh Joshi**

Este documento resume padrões importantes de sistemas distribuídos, focando em:

- **qual problema eles resolvem**
- **como funcionam**
- **onde são usados**
- **nível de dificuldade**

---

# 1. Circuit Breaker (fácil)

## Problema

Em sistemas distribuídos, serviços frequentemente dependem de outros serviços.

Exemplo:

```
API → Serviço de Pagamento → Banco de Dados
```

Se o serviço de pagamento ficar lento ou indisponível, o sistema inteiro pode congestionionar.

Isso pode causar:

- threads bloqueadas
- timeouts em cascata
- queda do sistema

## Solução

O **Circuit Breaker** funciona como um disjuntor elétrico.

Ele monitora falhas e alterna entre três estados.

| Estado | Comportamento |
|------|------|
| Closed | requisições normais |
| Open | requisições falham imediatamente |
| Half-Open | algumas requisições de teste |

Fluxo típico:

```
muitas falhas
     ↓
circuito abre
     ↓
requisições falham rápido
     ↓
serviço se recupera
     ↓
estado half-open testa
     ↓
circuito volta para closed
```

## Uso comum

- arquiteturas de microserviços
- comunicação resiliente entre serviços
- bibliotecas de tolerância a falhas

---

# 2. Retry (fácil)

## Problema

Falhas temporárias são comuns em redes:

- perda de pacotes
- timeout
- congestionamento momentâneo

Uma requisição pode falhar **mesmo quando o serviço está saudável**.

## Solução

Repetir automaticamente a operação.

Tipos de retry:

| Tipo | Descrição |
|-----|-----|
| Imediato | tenta novamente imediatamente |
| Delay fixo | espera um tempo fixo |
| Backoff exponencial | o tempo de espera aumenta exponencialmente |

Exemplo:

```
tentativa 1 → falha
espera 100 ms

tentativa 2 → falha
espera 200 ms

tentativa 3 → sucesso
```

## Problema importante

Retries podem causar **tempestade de retries** se muitos clientes tentarem novamente ao mesmo tempo.

Por isso normalmente se combina:

```
Retry + Circuit Breaker
```

---

# 3. Saga (médio)

## Problema

Sistemas distribuídos não suportam facilmente **transações ACID entre múltiplos serviços**.

Exemplo de fluxo em e-commerce:

```
1. Criar pedido
2. Reservar estoque
3. Cobrar pagamento
```

Se o passo 3 falhar, precisamos **desfazer os passos anteriores**.

## Solução

Uma **Saga** divide a transação em **transações locais com ações compensatórias**.

Exemplo:

```
createOrder
reserveStock
chargePayment
```

Se o pagamento falhar:

```
cancelStockReservation
cancelOrder
```

Cada etapa possui uma **ação de compensação**.

## Dois tipos de Saga

### Choreography

Serviços se coordenam através de eventos.

```
OrderCreated → InventoryService
StockReserved → PaymentService
```

### Orchestration

Um serviço central controla o fluxo.

## Uso comum

- microserviços
- plataformas de e-commerce
- sistemas financeiros

---

# 4. CQRS (médio)

**Command Query Responsibility Segregation**

## Problema

Muitos sistemas possuem:

- muitas leituras
- poucas escritas

Mas o mesmo modelo de dados precisa atender ambos.

## Solução

Separar os modelos de **leitura** e **escrita**.

```
Write Model → comandos
Read Model  → consultas
```

Arquitetura:

```
Command → Write Database
          ↓
         Evento
          ↓
Atualiza Read Model
```

## Benefícios

- consultas muito rápidas
- escalabilidade de leitura
- modelos otimizados para cada uso

## Desvantagens

- consistência eventual
- maior complexidade arquitetural

---

# 5. Event Sourcing (médio)

## Problema

Sistemas tradicionais armazenam apenas **o estado atual**.

Exemplo:

```
saldo = 100
```

Mas isso perde **o histórico completo das mudanças**.

## Solução

Armazenar **eventos que modificaram o estado**.

Exemplo:

```
ContaCriada
Deposito(100)
Saque(20)
Deposito(50)
```

O estado atual é reconstruído aplicando os eventos.

```
estado = fold(eventos)
```

## Benefícios

- histórico completo
- auditoria
- replay de eventos
- depuração facilitada

## Muito usado com CQRS

```
Event Store → Read Models
```

---

# 6. Leader Election (difícil)

## Problema

Algumas operações precisam de **um único líder**:

- coordenação
- agendamento
- escrita
- locks distribuídos

Mas nós podem falhar.

## Solução

Um algoritmo escolhe dinamicamente qual nó será o líder.

Propriedades importantes:

| Propriedade | Significado |
|------|------|
| Safety | apenas um líder |
| Liveness | sempre haverá um líder |
| Fault tolerance | falha do líder é tratada |

Fluxo típico:

```
nós enviam heartbeat
líder falha
timeout detectado
nova eleição acontece
```

Algoritmos comuns:

- Bully Algorithm
- eleição do Raft
- eleição do ZooKeeper

---

# 7. Replicated Log / Consensus (muito difícil)

Este é **o núcleo dos sistemas distribuídos modernos**.

## Problema

Vários nós precisam concordar **na mesma sequência de operações**.

Exemplo:

```
write A
write B
write C
```

Todos os nós precisam aplicar **na mesma ordem**.

Caso contrário, o estado diverge.

## Solução

Manter um **log replicado entre os nós**.

Arquitetura:

```
Leader
  ↓
replica entradas do log
  ↓
Followers
```

Fluxo:

```
cliente → líder
líder → adiciona entrada no log
líder → replica para followers
maioria confirma
entrada é commitada
```

## Regras importantes

- ordem total das operações
- quorum (maioria)
- um líder por termo

## Algoritmos famosos

- Paxos
- Raft
- Zab

## Sistemas que utilizam esse conceito

- Kafka
- etcd
- ZooKeeper

---

# Comparação de Dificuldade

| Padrão | Dificuldade | Motivo |
|------|------|------|
| Circuit Breaker | Baixa | lógica simples |
| Retry | Baixa | repetição |
| Saga | Média | compensações |
| CQRS | Média | separação de modelos |
| Event Sourcing | Média | eventos representam estado |
| Leader Election | Alta | coordenação distribuída |
| Replicated Log | Muito Alta | consenso distribuído |

---

# Insight Principal

O padrão **Replicated Log / Consensus** é a base de muitos sistemas distribuídos modernos porque garante:

- consistência
- tolerância a falhas
- estado coordenado entre múltiplos nós

Entender esse padrão ajuda a compreender a arquitetura de muitos sistemas distribuídos atuais.