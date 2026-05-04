# Introdução aos Sistemas Distribuídos
**Prof. Nélio Cacho** **Universidade Federal do Rio Grande do Norte (UFRN)** **Departamento de Informática e Matemática Aplicada (DIMAP)** **Email:** neliocacho@dimap.ufrn.br

## Evolução dos Sistemas
* **Aplicação Standalone:** Uma aplicação única rodando em uma máquina isolada.
* **Sistema Distribuído (Cliente-Servidor):** A aplicação é dividida entre um Cliente (ex: IP 10.0.0.25) e um Servidor (ex: IP 10.0.0.20:8080), comunicando-se via Protocolo através de uma Rede de Computadores.

---

## Anatomia de um Sistema Distribuído
Em sistemas distribuídos, **System Models** são modelos usados para descrever um sistema. Eles guiam o projeto, a escolha de algoritmos, o desempenho e a tolerância a falhas.

### Principais System Models:
1.  **Modelos de Arquitetura (Architectural Models)**
2.  **Modelos de Interação (Interaction Models)**
3.  **Modelos de Falha (Failure Models)**

---

## 1. Modelos de Arquitetura (Architectural Models)
Definem como os componentes estão organizados, como interagem e como os recursos são distribuídos. A escolha da arquitetura define:
* Complexidade de coordenação.
* Estratégias de balanceamento de carga.
* Possibilidades de escalabilidade e resiliência.

### Principais variações:
* **Cliente-Servidor:**
    * Clientes realizam requisições; Servidores respondem com serviços/resultados.
    * Um servidor pode ser cliente de outros servidores.
    * Baseado no paradigma requisição-resposta (request-reply).
    * **Transparência de localização:** Clientes não precisam saber a localização física do servidor.
    * **Stateless (Sem estado):** Cada requisição é independente. Fácil de escalar e recuperar de falhas (Ex: API de consulta de preço).
    * **Stateful (Com estado):** Mantém informações de sessão (Ex: Carrinho de compras). Recuperação de falhas é complexa.

* **Peer-to-Peer (P2P):**
    * Todas as entidades podem ser clientes e servidoras simultaneamente.
    * Envolve compartilhamento de objetos e simetria de serviços.

* **Modelo de Filtros (Pipes-and-Filters):**
    * Processos transformam entradas em saídas através de estágios.
    * **Filtro:** Componente de processamento.
    * **Tubo (Pipe):** Conector que transporta os dados.
    * Permite reuso e separação de responsabilidades.

* **Publicação-Subscrição (Publish-Subscribe):**
    * Comunicação assíncrona com Publicadores (Publishers) e Subscritores (Subscribers).
    * Uso de um **Broker de mensagens** (barramento de eventos).
    * **Desacoplamento Espacial:** Anonimato entre as partes.
    * **Desacoplamento Temporal:** Ciclos de vida independentes.

* **Microserviços / SOA:**
    * Aplicação dividida em serviços pequenos, independentes e autocontidos.
    * Foco em escalabilidade horizontal e gerenciamento distribuído (Kubernetes/Docker).

---

## 2. Modelos de Interação (Interaction Models)
Descrevem as restrições temporais e garantias de entrega:
* **Modelo Sincrônico:** Tempos máximos de processamento e comunicação conhecidos. Clocks perfeitamente sincronizados.
* **Modelo Assíncrono:** Sem garantias de tempo. Mensagens podem atrasar indefinidamente. Clocks podem divergir.

---

## 3. Modelos de Falha (Failure Models)
Sistemas distribuídos são assíncronos por natureza, tornando difícil distinguir um nó que travou (crash) de um nó apenas lento.

### Categorias de Falhas:
* **Crash Failure:** O nó para abruptamente. Requer replicação de dados e logs (WAL).
* **Omission Failure:** Mensagens perdidas no envio ou recebimento (Ex: pacotes TCP perdidos).
* **Timing Failure:** Resposta chega fora do prazo esperado.
* **Response Failure:** Resposta incorreta ou valor errado (Value Failure) devido a erro interno.
* **Byzantine Failure:** Comportamento arbitrário ou malicioso (respostas falsas/contraditórias). Comum em Blockchains.

### Uso de Timeouts:
Impõem um limite artificial de espera. 
* **Timeout pequeno:** Detecção rápida, mas pode gerar falsos positivos (marcar nó lento como morto).
* **Timeout grande:** Evita erros com nós lentos, mas demora a detectar falhas reais.

---

## Semântica "Exactly-Once"
Garantia de que uma operação será processada exatamente uma vez, mesmo com reenvios de mensagens.
* **Idempotent Receiver:** Operações que podem ser repetidas sem mudar o resultado final.
* **De-duplication:** Uso de identificadores únicos para ignorar mensagens repetidas.

---

## Replicação e Teorema CAP
Dados são replicados para garantir disponibilidade, mas isso gera inconsistência temporal.

### Teorema CAP:
Um sistema distribuído só pode garantir simultaneamente **dois** dos três atributos:
1.  **Consistência (Consistency):** Todos os nós veem o mesmo dado ao mesmo tempo.
2.  **Disponibilidade (Availability):** Toda requisição recebe uma resposta (sucesso ou falha).
3.  **Tolerância a Partições (Partition Tolerance):** O sistema continua operando apesar de falhas na rede.

**Combinações:**
* **CP (Consistência e Partição):** Ex: MongoDB (sacrifica disponibilidade se não houver quórum).
* **AP (Disponibilidade e Partição):** Ex: Cassandra, CouchDB (sacrifica consistência imediata pela eventual).

---

## Partição de Dados (Sharding)
Quando dados não cabem em um único nó, são divididos. Desafios incluem localizar o dado rapidamente e garantir atomicidade em transações multi-partição (Two-Phase Commit).

---

## Tempo em Sistemas Distribuídos
Não há relógio global perfeito. O "drift" de hardware e latência de rede causam ambiguidade na ordem dos eventos.
* **Soluções:** Lamport Clocks, Vector Clocks e Hybrid Logical Clocks.

---

## Atividade
* Estudar o padrão de projeto sorteado.