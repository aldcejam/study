---
tags:
  - devops
  - k8s
  - "#volumes"
---


Objetivo, levar dados tal como um arquivo env, um txt, ou quaisquer outros para dentro do container sem precisar criar um nova imagem.

``` ts
apiVersion: v1
kind: CondigMap
metadata:
  name goserver-dynamic-source
  labels: 
	  app: goserver-dynamic-source
data: 
	membems: "Aldcejam, Rebeka"
```

``` ts
apiVersion: apps/v1
kind: Deployment
metadata:
  name: goserver
  labels:
    app: goserver-nivel-deployment
spec:
  replicas: 2
  selector:
    matchLabels:
      app: goserver-nivel-replicaset
  template:
    metadata:
      labels:
        app: goserver-nivel-replicaset
    spec:
      containers:
        - name: goserver-container
          image: aldcejam/goserver:v4
          ports:
            - containerPort: 8080
          volumeMounts:
            - name: goserver-volume-nivel-container
              mountPath: /app/volume-criado
      volumes:
        - name: goserver-volume-nivel-container
          configMap:
            name: goserver-dynamic-source
            items:
              - key: members
                path: members.txt


```