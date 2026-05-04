---
tags:
  - devops
  - k8s
  - service
---
``` YAML
apiVersion: v1
kind: Service
metadata:
  name: goserver
  labels:
    app: goserver
spec:
  type: LoadBalancer
  ports:
    - name: goserver-loadbalancer-service
      port: 8080
      targetPort: 8080
      protocol: TCP
  selector:
    app: goserver
```

>  Necessário quando se deseja expor o serviço para fora da infra. É gerado um external-ip, através dele qualquer um da internet poderá acessar o serviço exposto.