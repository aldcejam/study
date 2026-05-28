---
tema: Evolução Web e RPC - HTTP3 e gRPC
tipo: unidade-estudo
tags: [redes, rpc, grpc, http, performance]
---
# 🧪 UE - Evolução Web e RPC - HTTP3 e gRPC

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O design inicial da Web baseado em HTTP/1.1 sobre TCP introduziu o problema de Head-of-Line (HoL) Blocking, onde uma requisição enfileirada no socket trava todas as subsequentes. Além disso, a serialização de dados de texto (como JSON) tem enorme overhead de parsing e tamanho de payload. Em arquiteturas de microsserviços modernos, de altíssimo throughput, ignorar essas restrições leva a latências inaceitáveis, timeouts em cascata e exaustão silenciosa de CPU devido ao parsing contínuo de texto. HTTP/3 (sobre QUIC) e gRPC resolvem o bloqueio de rede e o peso da serialização.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Evolução Protocolar e QUIC]:** Transição de conexões TCP estado-pesadas (HTTP/1 e 2) para conexões UDP com controle de fluxo no nível da aplicação (QUIC), resolvendo o TCP HoL Blocking.
2. **[Remote Procedure Call e Protobuf]:** Funcionamento da geração de stubs via gRPC e o mecanismo interno do Protocol Buffers (serialização binária ultrarrápida vs parsers DOM/JSON).
3. **[Concorrência e Multiplexação de Streams]:** Como o HTTP/2 e HTTP/3 lidam com centenas de requisições paralelas sem abrir múltiplos file descriptors (sockets) no Sistema Operacional.
4. **[REST vs gRPC e Trade-offs Arquiteturais]:** Acoplamento rígido de schemas (gRPC) versus acoplamento fraco e cacheabilidade via verbos HTTP nativos (REST).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/RFC:** RFC 9000 (QUIC: A UDP-Based Multiplexed and Secure Transport).
- **System Blueprint:** Arquitetura do Google "Stubby" (antecessor do gRPC) e o motor de balanceamento interno de microsserviços do Envoy Proxy configurado para gRPC.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Demonstrar o impacto do Head-of-Line Blocking e do payload binário em condições de rede degradadas.
- [ ] Configurar ambiente de isolamento (Docker / Local Sandbox).
- [ ] Implementar um servidor gRPC com streaming bidirecional e um servidor REST em paralelo, processando payloads volumosos.
- [ ] Injetar carga/falha: Usar o `tc` (Linux Traffic Control) para simular 20% de perda de pacotes e 100ms de latência na rede.
- [ ] Coletar métricas de comportamento: Comparar os tempos de resposta finais no p99 e o uso de CPU entre REST/JSON e gRPC/Protobuf.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Evolucao TCP para QUIC e Solucao de HoL Blocking - Teoria e Fundamentos]]
- [[./Protocol Buffers e Serializacao Binaria - Funcionamento Interno e Arquitetura]]
- [[./Sobrecarga de Sockets e Problema de Concorrencia em REST - Casos de Falha e Análise Amortizada]]
- [[./Comparativo de Latencia gRPC vs HTTP em Redes Degradadas - Implementacao de Referencia e Benchmarks]]
