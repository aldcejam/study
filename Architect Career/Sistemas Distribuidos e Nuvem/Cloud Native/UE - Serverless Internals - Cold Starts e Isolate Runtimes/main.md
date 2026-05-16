---
tema: Serverless Internals - Cold Starts e Isolate Runtimes
tipo: unidade-estudo
tags: [cloud-native, performance, arquitetura, serverless]
---
# 🧪 UE - Serverless Internals - Cold Starts e Isolate Runtimes

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O modelo **Serverless (FaaS)** promete escala infinita e custo zero quando não está em uso. O problema é a latência de ativação: quando uma função não é chamada há tempo, o provedor desliga o container; a próxima chamada precisa esperar o boot completo. O **Cold Start** resolve o problema de custo, mas cria um problema de UX. Tecnologias como **Isolate Runtimes** (Cloudflare Workers) resolvem isso usando isolates do V8 em vez de containers, reduzindo o tempo de boot para milissegundos. Sem entender o runtime, você não sabe se sua aplicação deve ser um microserviço tradicional ou uma função serverless.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Ciclo de Vida da Funcao]:** Init, Invoke e Shutdown.
2. **[MicroVMs (Firecracker)]:** Como a AWS consegue rodar milhares de funções isoladas com boot em milissegundos.
3. **[V8 Isolates]:** A diferença entre isolamento por processo/container e isolamento por heap de memória (Workers).
4. **[Snapshotting]:** Como salvar o estado da memória pós-inicialização para pular o estágio de Init (ex: AWS Lambda SnapStart).
5. **[Edge Computing]:** Rodar código serverless perto do usuário para eliminar a latência de rede.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Firecracker: Lightweight Virtualization for Serverless Applications* (AWS, 2020).
- **System Blueprint:** Arquitetura do AWS Lambda e Cloudflare Workers.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Medir o impacto do Cold Start em uma função AWS Lambda (ou similar local usando `LocalStack`) variando o tamanho da memória alocada e o tamanho do pacote de código.
- [ ] Criar uma função que demora para inicializar (ex: carregando bibliotecas pesadas).
- [ ] Medir o tempo da primeira chamada (Cold) vs chamadas subsequentes (Warm).
- [ ] Analisar como o aumento de memória alocada reduz o tempo de boot (proporcionalidade de CPU).

## 🗃️ Notas Heutagógicas Atômicas
- [[./Micro-Virtualizacao e Firecracker - Teoria e Fundamentos]]
- [[./Arquitetura de Isolates do V8 vs Containers - Funcionamento Interno e Arquitetura]]
- [[./Estrategias de Pre-warming e Snapshotting - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Provedores Serverless - Implementacao de Referencia e Benchmarks]]
