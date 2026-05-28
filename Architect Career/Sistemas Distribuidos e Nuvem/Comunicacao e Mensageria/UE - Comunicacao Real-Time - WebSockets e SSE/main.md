---
tema: Comunicação Real-Time - WebSockets e SSE
tipo: unidade-estudo
tags: [real-time, websockets, sse, tcp, concurrencia]
---
# 🧪 UE - Comunicação Real-Time - WebSockets e SSE

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> HTTP é um protocolo intrinsecamente bidirecional síncrono e baseado em requisição/resposta do cliente. Em cenários de colaboração ao vivo, dashboards financeiros ou notificações instantâneas, o servidor precisa ativamente "empurrar" dados para o cliente. Usar truques antigos como "Long Polling" consome ferozmente threads, memória e buffers TCP para conexões ociosas. Sustentar 100 mil a 1 milhão de clientes passivos conectados a um único servidor exige o domínio de WebSockets, Server-Sent Events (SSE) e do famigerado problema do Sistema Operacional chamado C10K/C10M (esgotamento de File Descriptors e Context Switches).

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Handshake de Upgrade e Framings do WebSocket]:** Como a negociação de headers muda um socket HTTP tradicional para um túnel TCP contínuo usando máscara XOR e troca de frames binários e de controle (Ping/Pong).
2. **[Esgotamento de File Descriptors (C10K problem)]:** Como I/O multiplexado não bloqueante (`epoll`, `kqueue`, `io_uring`) e a eliminação do "Thread per Connection" são essenciais para segurar sockets abertos a baixo custo.
3. **[WebSocket (Bidirecional) vs SSE (Unidirecional)]:** O trade-off arquitetural de usar HTTP/2 nativo e SSE apenas para downlink leve contra a customização do framming full-duplex de WebSockets.
4. **[Escalabilidade Horizontal de Sockets Abertos]:** O problema crônico do Statefulness de WebSockets. Se um node falha, centenas de milhares de sessões caem e fazem "Thundering Herd" (reconexão em massa) em outro node.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **RFC:** RFC 6455 (The WebSocket Protocol) e RFC 8895 (Server-Sent Events).
- **System Blueprint:** A arquitetura do Erlang/Elixir Phoenix Channels (conseguindo 2M conexões numa só máquina) e o artigo original The C10K Problem (Dan Kegel).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Demonstrar o problema de esgotamento de recursos em concorrência baseada em threads versus um loop de eventos com chamadas bare-metal.
- [ ] Configurar ambiente de isolamento (Local).
- [ ] Implementar a lógica core de servidor WebSocket em baixo nível utilizando `epoll`/`select` em C, Go ou Rust gerindo 10k requisições falsas ativas.
- [ ] Injetar carga/falha: Fazer um Thundering Herd simulando os 10k clientes reconectando agressivamente ao mesmo tempo.
- [ ] Instrumentar as métricas de uso de memória de thread stack e context switching overhead no Kernel.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Diferenciacao entre SSE, Long Polling e WS - Teoria e Fundamentos]]
- [[./Multiplexacao C10K via epoll e Estrutura dos Frames Binarios - Funcionamento Interno e Arquitetura]]
- [[./Reconexao em Massa Thundering Herd e Stateful Drops - Casos de Falha e Análise Amortizada]]
- [[./Implementacao de Milhoes de Sockets Ociosos em ErlangGo - Implementacao de Referencia e Benchmarks]]
