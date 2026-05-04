---
tags:
  - devops
  - k8s
  - configmap
---
## Definição
``` YAML
apiVersion: v1
kind: ConfigMap
metadado:
	name: goserver-env
data:
	NAME: "aldcejam",
	AGE: "21"
```
> Necessário aplicar este yaml: `kubectl apply -f pasta/configmap.yaml`

## Implementação
#### 1°
``` YAML
api: v1
kind: deployment
...
spec:
	containers:
		...
		envFrom:
			- configMapRef:
				name: goserver-env
```

#### 2º
``` YAML
api: v1
kind: deployment
...
spec:
	containers:
		...
		env
			- name: NAME
			  valueFrom:
				  configMapKeyRef:
					  name: goserver-ref
					  key: NAME
```

