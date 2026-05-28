---
tema: Service Mesh e Roteamento - Istio e Envoy
tipo: unidade-estudo
tags: [mesh, infra, proxy, istio, envoy]
---
# 🧪 UE - Service Mesh e Roteamento - Istio e Envoy

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Historicamente, ao desenvolver dezenas de microsserviços, os engenheiros precisavam codificar regras de resiliência de rede (Circuit Breaking, Timeout, Retries) e segurança (Autenticação mTLS) na lógica de negócio de cada aplicação, gerando um pesadelo de bibliotecas repetidas (ex: Netflix OSS / Hystrix) nas linguagens mais variadas. Se a rede flutuar, a aplicação congela de forma invisível. O padrão de Service Mesh delega integralmente essas preocupações transversais de tráfego L7 da rede para um Proxy infraestrutural (Sidecar), separando o que é rede complexa (Data Plane) do controle de observabilidade e segurança centralizado (Control Plane).

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Data Plane vs Control Plane]:** A separação conceitual entre os túneis que movem os bytes (Sidecars como Envoy) e o cérebro matemático/político que injeta a configuração (Istiod).
2. **[Traffic Management e Resiliência Ativa]:** Regras dinâmicas para Canary Releases, Mirroring de Tráfego Sombrio, e interrupções preemptivas por Outlier Detection/Circuit Breaking independente do app.
3. **[Segurança Zero Trust e mTLS inter-serviços]:** Criação e rotação de certificados automatizada pelo SPIFFE/SPIRE em cada Sidecar para criptografia ponta a ponta sem modificar o código.
4. **[O Custo do Mesh e a Evolução Sidecarless]:** O overhead significativo em latência (Context Switches de TCP duplo app->proxy->rede->proxy->app) e o avanço arquitetural via eBPF ou Cilium para resolver gargalos da stack do Kernel no networking L4/L7.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Livro Clássico:** *Istio in Action* (Christian Posta).
- **System Blueprint:** A evolução do Lyft desenvolvendo o Envoy (escrito primorosamente em C++11 com arquitetura de thread model single-threaded event loop) como base de rede L7.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular um Circuit Breaking L7 puramente na infraestrutura, protegendo um microsserviço frágil sem alterar o seu código fonte.
- [ ] Configurar ambiente de isolamento (K3s local ou Docker Compose rodando Envoy ao lado de um backend HTTP mockado em Python/Node).
- [ ] Implementar a configuração do Envoy bloqueando tráfego após exceder 50% de taxa de erro de 5XX e retornando "Circuit Open Fast Fail".
- [ ] Injetar carga/falha: Disparar stress test de RPS alto e forçar o mock de backend a começar a dar Timeout ou Erro 500.
- [ ] Analisar os Logs e Traces provando que o Sidecar corta o fluxo antes mesmo da rede despachar o TCP para o container frágil, salvando seus recursos.

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Padroes Sidecar Data Plane e Control Plane - Teoria e Fundamentos]]
- [[./Envoy Thread Model C++ e Roteamento L7 - Funcionamento Interno e Arquitetura]]
- [[./Overhead de Latencia de Hops e Limites de eBPF - Casos de Falha e Análise Amortizada]]
- [[./Implantacao de Circuit Breaking e Zero-Trust mTLS - Implementacao de Referencia e Benchmarks]]
