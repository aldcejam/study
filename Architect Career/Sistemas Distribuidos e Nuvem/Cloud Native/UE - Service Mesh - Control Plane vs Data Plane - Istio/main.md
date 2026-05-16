---
tema: Service Mesh - Control Plane vs Data Plane - Istio
tipo: unidade-estudo
tags: [cloud-native, infraestrutura, redes, segurança]
---
# 🧪 UE - Service Mesh - Control Plane vs Data Plane - Istio

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Quando você tem centenas de microserviços, como garantir que a comunicação entre eles seja segura (mTLS), resiliente (retries/circuit breakers) e observável sem precisar mudar uma única linha de código da aplicação? O **Service Mesh** resolve isso injetando um proxy (Sidecar) ao lado de cada container. O **Data Plane** (Envoy) lida com o tráfego real, enquanto o **Control Plane** (Istio/Istiod) gerencia as configurações e políticas. Sem o service mesh, cada desenvolvedor precisaria implementar segurança e resiliência na mão, gerando inconsistência e vulnerabilidades.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[O Padrao Sidecar]:** Como interceptar o tráfego de entrada e saída do pod usando `iptables`.
2. **[mTLS Automatico]:** Como o mesh gerencia a rotação de certificados e garante que apenas serviços autorizados se comuniquem.
3. **[Traffic Management]:** Canário deployments, espelhamento de tráfego (Shadowing) e injeção de falhas para testes.
4. **[Circuit Breaking e Outlier Detection]:** Como impedir que um serviço lento derrube o sistema inteiro.
5. **[Overhead de Performance]:** O custo de latência adicional introduzido pelo proxy e como mitigá-lo (eBPF/Ambient Mesh).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Istio: Service Mesh for Modern Microservices* (Google Technical Reports).
- **System Blueprint:** Arquitetura do **Envoy Proxy** e sua integração com o Istio.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Instalar o Istio em um cluster local e configurar uma regra de `VirtualService` para realizar um roteamento baseado em pesos (Canary Release: 90% v1, 10% v2).
- [ ] Deploy de duas versões de um serviço.
- [ ] Aplicar o manifesto de roteamento do Istio.
- [ ] Testar a distribuição de tráfego usando `curl` em loop.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Intercepcao de Trafego e Padrao Sidecar - Teoria e Fundamentos]]
- [[./Arquitetura do Envoy e xDS API - Funcionamento Interno e Arquitetura]]
- [[./Impacto de Latencia e Recursos em Meshes - Casos de Falha e Analise Amortizada]]
- [[./Configuracao de Zero-Trust com Istio - Implementacao de Referencia e Benchmarks]]
