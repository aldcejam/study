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
    app: goserver-service
spec:
  type: ClusterIP
  ports:
    - port: 8080
      targetPort: 8080
  selector:
    app: goserver
```

- spec > selector > metadado do deployment
>serve para referenciar para quais pods ele mandará as requisições (`app: goserver` é um metadado de um deployment). Logo, deve haver um deployment para quem referenciar neste service. Por conseguinte, todas as requisições deste service serão mandadas de forma balanceada entres os pods deste kind `deployment`

