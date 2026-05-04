---
tags:
  - devops
  - k8s
  - "#health"
---
### StartupProbe
Antes de iniciar qualquer outra probe e ao iniciar o container ele verifica através, por exemplo, de um endpoint `/healthz` se aplicação está saudável, se estiver, as outras probes começarão a ser executadas.
*Before starting any other probe and when the container starts, it checks, for example, through a `/healthz` endpoint whether the application is healthy. If it is, the other probes will then be executed.*
### ReadinessProbe
Checks when application is available to receive traffic, if is not available it blocks a requests to the container.
*Ele verifica quando a aplicação está disponível para receber tráfego, se não estiver disponível ele bloqueia as requisições para o container*
### LivenessProbe
Checks if the container is healthy, if not, it will restart the container