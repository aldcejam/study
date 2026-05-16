---
tema: Kubernetes Internals - API Server e Controller Manager
tipo: unidade-estudo
tags: [cloud-native, kubernetes, arquitetura, infraestrutura]
---
# 🧪 UE - Kubernetes Internals - API Server e Controller Manager

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O Kubernetes não é um sistema que "manda" as coisas acontecerem, ele é um sistema que tenta fazer com que a realidade corresponda ao desejo do usuário. O problema é como gerenciar o estado de milhares de containers de forma resiliente e extensível. O **API Server** resolve isso servindo como a única fonte da verdade e interface de comunicação. O **Controller Manager** resolve isso através de "loops de controle" que monitoram o estado e tomam ações corretivas. Sem entender o funcionamento interno do Control Plane, você não consegue criar operadores customizados ou debugar por que seus pods não estão escalando.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Declarative API vs Imperative API]:** Por que o K8s usa o modelo de "Estado Desejado".
2. **[Etcd e o API Server]:** Como o estado é armazenado e por que o API Server é o único que fala com o banco de dados.
3. **[Reconciliation Loop]:** A lógica fundamental: `Observe -> Diff -> Act`.
4. **[Admission Controllers e Mutators]:** Como interceptar e modificar requisições antes que elas sejam gravadas no Etcd.
5. **[Custom Resource Definitions (CRDs)]:** Como estender o Kubernetes para gerenciar novos tipos de recursos (ex: um banco de dados).

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Borg, Omega, and Kubernetes* (Brendan Burns et al., ACM Queue 2016).
- **System Blueprint:** O código fonte do Kubernetes (`pkg/controlplane/` e `pkg/controller/`).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Criar um controlador simples em Go ou Python (usando o framework `kopf` ou `client-go`) que monitora a criação de um ConfigMap e automaticamente cria um log de auditoria.
- [ ] Configurar o acesso ao cluster (Minikube ou Kind).
- [ ] Escrever o loop de observação (Informer).
- [ ] Implementar a lógica de reconciliação.

## 🗃️ Notas Heutagógicas Atômicas
- [[./A Filosofia do Estado Desejado e Reconciliacao - Teoria e Fundamentos]]
- [[./Arquitetura do API Server e Watch Events - Funcionamento Interno e Arquitetura]]
- [[./Lentidao no Etcd e Impacto no Control Plane - Casos de Falha e Analise Amortizada]]
- [[./Desenvolvimento de um Operator Customizado - Implementacao de Referencia e Benchmarks]]
