---
tema: Raft Consensus - Leader Election e Log Replication
tipo: unidade-estudo
tags: [sistemas-distribuidos, consenso, raft, arquitetura]
---
# 🧪 UE - Raft Consensus - Leader Election e Log Replication

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O consenso é o problema mais difícil e importante em sistemas distribuídos. Sem um algoritmo como o **Raft** (ou Paxos), é impossível manter um estado consistente entre múltiplas máquinas quando falhas ocorrem (e elas sempre ocorrem). Se o engenheiro ignorar as nuances da eleição de líder e da replicação de log, o sistema sofrerá de *split-brain*, perda de dados ou indisponibilidade total durante partições de rede.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Termos e Eleição de Lider]:** Como o Raft divide o tempo em *terms* e garante que apenas um líder seja eleito por termo através de *heartbeats* e *timeouts*.
2. **[Log Replication e Commits]:** O mecanismo de replicação de entradas do log e a definição de "comitado" (maioria dos nós).
3. **[Seguranca e Invariantes]:** Entender por que um nó só vota em um candidato se o log do candidato for pelo menos tão atualizado quanto o dele.
4. **[Gerenciamento de Configuracao e Snapshots]:** Como o cluster lida com a entrada/saída de nós e como o log é compactado para não crescer infinitamente.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *In Search of an Understandable Consensus Algorithm* (Diego Ongaro and John Ousterhout, 2014).
- **System Blueprint:** Implementações reais como etcd (Kubernetes), CockroachDB e HashiCorp Consul.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar uma versão simplificada da máquina de estados do Raft em Go ou Rust que consiga eleger um líder e replicar uma string simples entre 3 processos locais.
- [ ] Implementar o loop de timer para timeouts de eleição.
- [ ] Simular falhas de rede (drop de pacotes) e observar a reeleição.
- [ ] Verificar a consistência dos logs em todos os nós após a recuperação de uma falha.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Estados do Raft e Timers de Eleicao - Teoria e Fundamentos]]
- [[./Mecanismo de AppendEntries e Quorum - Funcionamento Interno e Arquitetura]]
- [[./Particoes de Rede e Split-brain no Raft - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de um KV Store Consistente - Implementacao de Referencia e Benchmarks]]
