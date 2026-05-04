---
tags:
  - k8s
  - devops
---


| **Escopo**      | **Comando**                              | **Descrição**                               |
| --------------- | ---------------------------------------- | ------------------------------------------- |
| **Cluster**     | `kind create cluster`                    | Cria um cluster local.                      |
|                 | `kind get clusters`                      | Lista os clusters criados.                  |
|                 | `kind delete cluster`                    | Deleta um cluster existente.                |
|                 | `kubectl cluster-info`                   | Mostra info do cluster atual.               |
|                 | `kubectl get all`                        | Lista todos os recursos no namespace atual. |
|                 | `kubectl get all`                        |                                             |
| **Nodes**       | `kubectl get nodes`                      | Lista todos os nodes do cluster.            |
|                 | `kubectl describe node <nome-do-node>`   | Mostra detalhes de um node específico.      |
|                 | `kubectl cordon <nome-do-node>`          | Marca o node como _unschedulable_.          |
|                 | `kubectl drain <nome-do-node>`           | Remove os pods de um node (manutenção).     |
|                 | `kubectl uncordon <nome-do-node>`        | Torna o node schedulable novamente.         |
| **Pods**        | `kubectl get pods`                       | Lista todos os pods.                        |
|                 | `kubectl describe pod <nome-do-pod>`     | Mostra detalhes de um pod específico.       |
|                 | `kubectl logs <nome-do-pod>`             | Exibe os logs de um pod.                    |
|                 | `kubectl exec -it <nome-do-pod> -- bash` | Abre terminal interativo dentro do pod.     |
|                 | `kubectl delete pod <nome-do-pod>`       | Deleta um pod específico.                   |
| **Deployments** | `kubectl create deployment`              | Cria um   novo deployment.                  |
|                 | `kubectl get deployments`                | Lista os deployments.                       |
|                 | `kubectl scale deployment`               | Altera o número de réplicas.                |
|                 | `kubectl rollout status deployment`      | Verifica o status do rollout.               |
|                 | `kubectl apply -f file.yaml`             | Inicia um kind `deployment`                 |
| **Serviços**    | `kubectl expose deployment`              | Exponha um deployment como serviço.         |
|                 | `kubectl get services`                   | Lista os serviços.                          |
|                 | `kubectl describe service <nome>`        | Mostra detalhes de um serviço específico.   |


