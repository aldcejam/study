---
tema: Design de APIs - GraphQL e API Gateway
tipo: unidade-estudo
tags: [api, graphql, gateway, arquitetura, bff]
---
# 🧪 UE - Design de APIs - GraphQL e API Gateway

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Se clientes web e mobile consumirem microsserviços do backend diretamente, o sistema entra em colapso devido à explosão do número de requests (Over-fetching e Under-fetching de payloads REST) e complexidade no frontend. Pior: você expõe a rede interna, dificultando refatorações estruturais de microsserviços sem quebrar clients (falta de abstração). API Gateways e padrões de Backend for Frontend (BFF) com GraphQL criam uma fachada unificada (Single Point of Entry) que compõe dados, reduz requests de rede, resolve problemas de segurança na borda (Offloading de SSL, Tokens e Rate Limits) e otimiza payload. O perigo é transformar o gateway num gargalo monolítico.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[O Padrão Gateway, BFF e Borda de Rede]:** A abstração de roteamento L7, agregação de payloads e o isolamento seguro da rede interna em relação à rede externa.
2. **[GraphQL e Resolução de Grafos]:** O modelo de query language, schemas tipados em AST (Abstract Syntax Tree) e como o Type System compõe os resolvers em tempo de execução.
3. **[N+1 Query Problem em Grafos]:** A complexidade assintótica invisível do GraphQL, onde 1 requisição no frontend explode em N queries de banco ou N chamadas RPC no backend se não houver um `DataLoader`.
4. **[Edge Security e Offloading]:** Delegação pesada de CPU (TLS Termination, validação JWT criptográfica) da aplicação para a infraestrutura do Gateway.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper/Livro Clássico:** *Microservices Patterns* (Chris Richardson) - Capítulo 8: "External API patterns".
- **System Blueprint:** A criação do GraphQL no Facebook (2012) para resolver gargalos do app Mobile e a arquitetura federada Apollo Federation usada na Netflix.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Construir um Backend for Frontend robusto usando GraphQL que agrega dados de 2 microsserviços independentes (REST/gRPC) demonstrando a solução do problema N+1.
- [ ] Configurar ambiente de isolamento (Docker com Serviços "Usuário" e "Postagens").
- [ ] Implementar a lógica core de resolução do schema em um Gateway GraphQL usando `DataLoaders` e batching.
- [ ] Injetar carga e coletar métricas de comportamento: Rodar a query com vs sem DataLoader e rastrear com traces do Jaeger a quantidade de chamadas RPC geradas nos sub-serviços.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Abstracoes de Fachada e Padrões BFF - Teoria e Fundamentos]]
- [[./AST parsing de GraphQL e Engine de Execucao - Funcionamento Interno e Arquitetura]]
- [[./N+1 Problem e Explosion de Consultas de Rede - Casos de Falha e Análise Amortizada]]
- [[./Apollo Federation e Agregacao em Alto Desempenho - Implementacao de Referencia e Benchmarks]]
