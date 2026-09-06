---
tipo: diagnostico-temporario
status: temporario
modulo: 01 - Fundamentos e Well-Architected
instrucao: >
  Responda o que souber SEM consultar materiais.
  Se não souber, escreva "NÃO SEI" ou deixe em branco.
  Depois me devolva este arquivo (ou cole as respostas) para eu mapear lacunas por UE.
---
# 🧪 Diagnóstico de Nível — Módulo 01 (TEMP)

> **Como usar**
> 1. Responda de memória, em português ou inglês.
> 2. Não precisa acertar formato perfeito — o que importa é o raciocínio.
> 3. Marque no fim de cada UE: `Confiança: baixa | média | alta`.
> 4. Quando terminar, peça: *"mapeia o que me falta estudar com base nas respostas"*.

**Autoavaliação rápida antes de começar** (opcional):
- Já estudei: UE1 ☐  UE2 ☐  UE3 ☐  UE4 ☐
- Tempo estimado: ~45–90 min se for com calma

---

## UE 1 — Infraestrutura Global e Modelos de Nuvem

### Conceitos básicos
1. Explique, em 2–3 frases, a hierarquia Region → AZ → Data Center e o que cada nível isola.
> **Respostas:** A escrita da pergunta já estabelece corretamente a sequência hierarquica. A região é o nível limite de trafego de dados padrão, ou seja, todo dado/infra ficará restrito a este nível, a menos que seja configurado e intencionalmente queira-se transpor este nível. A AZ, é um componente de uma região, de maneira resumida, elas são unidades de isolamentos a nível geográfico, o intuíto é impedir que um desastre em uma AZ não afete outras. Por conseguinte, os datacenters de uma AZ são as unidades, por assim dizer, de processamento.
2. Por que multi-AZ é o padrão mínimo de alta disponibilidade na AWS, e quando multi-Region se justifica?
> **Respostas:** Devido a sua capacida de prevenir erros locais de um AZ. Multi-região justifica-se perante necessidades de latência, baixa disponibilidade de uma AZ.
3. Qual a latência típica entre AZs da mesma Region e por que isso importa para replicação síncrona vs assíncrona?
> **Respostas:** Latencia de até no máximo 2ms, nível suficiente para ser considerada síncrona.
4. O que é *blast radius*? Dê um exemplo de design que o reduz.
não sei

### Escopo de serviços
5. Classifique como **zonal**, **regional** ou **global**: EC2 instance, EBS, S3, DynamoDB, IAM, Route 53, CloudFront, subnet.
> **Respostas:** EC2: zonal, s3 e DynamoDB: regional, Route 53 = regiona, restante não sei (não sei também o que é Route 53, CloudFront, subnet)
6. `us-east-1a` na minha conta é o mesmo data center físico que `us-east-1a` na sua? Explique AZ name vs AZ ID.
> **Respostas:** Não necessariamente, AZ name não define uma AZ especifica, apenas a AZ ID é que possibilita este reconhecimento de qual AZ fisíca aquele recurso está.

### Edge e extensões
7. Diferencie Edge Location, Regional Edge Cache e Region. Para que serve cada um?
8. CloudFront vs Global Accelerator: quando usar cada um?
9. Local Zones vs Wavelength vs Outposts: problema que cada um resolve e trade-off principal.
10. Em que cenário você escolheria Outposts em vez de uma Region padrão?
^^ não sei nenhuma das acima ^^

### Responsabilidade compartilhada
11. Explique "segurança **DA** nuvem" vs "segurança **NA** nuvem".
 > **Respostas:** segurança "DA" nuvem trata-se da segurança ofertada pela AWS para seus servidores a nível fisíco, já a "NA" nuvem é aquilo cujo responsabilidade aplica-se ao "usuário" que está a gerenciar os recursos AWS.
12. Como essa linha muda entre EC2 (IaaS), RDS (PaaS) e S3 (serviço gerenciado)? Quem patcha o SO em cada caso? 
> **Respostas:** Não sei
13. IaaS vs PaaS vs SaaS: dê um exemplo AWS de cada e o que o cliente ainda controla.
> **Respostas:** Não sei

### Cenários (nível prova)
14. App exige latência < 5 ms para usuários em uma cidade onde não há Region. Quais opções AWS avaliar?
> **Respostas:** Não sei
15. Workload precisa de RPO ≈ 0 e failover automático dentro da Region. Multi-AZ ou multi-Region? Por quê? 
> **Respostas:** Multi-AZ, latencia de no máximo 2ms
16. Dados de saúde não podem sair do país. Qual decisão de infraestrutura isso força?
> **Respostas:** Escolha de uma região no país, sem configurações que possibilitem a saida, de padrão escolher uma zona já a delimita por padrão a ficar na região selecionada

**Confiança nesta UE:** baixa / média / alta  
**Respostas / anotações:**

```
média
```

---

## UE 2 — Identidade Base e Estrutura de Conta

### Identidades
17. Por que a conta root não deve ser usada no dia a dia? Cite 2 ações exclusivas do root. 
> **Respostas:** Pois ela tem todos os acessos a recursos e configurações. 1 fechar a própria, não sei outro
18. Diferença prática entre IAM User e IAM Role (credenciais, duração, casos de uso).
> **Respostas:** IAM User tratase de um usuário com acesso aquela AWS específica. IAM Role são para o que um usuário o máquina pode acessar de recursos e o que pode fazer com estes recursos.
19. O que é um IAM Group? Ele pode ser Principal? Pode conter outro group?
> **Respostas:** IAM Group é uma forma de agrupar permissões de role para atribuir a quem desejar. Sim, pode conter(confiança baixa)
20. O que é STS e o que acontece quando você faz `AssumeRole`?
> **Respostas:** Não sei

### Policies
21. Liste e explique os elementos de um statement IAM: Effect, Action, Resource, Condition, Principal.
> **Respostas:** Não sei
22. Em quais tipos de policy o campo `Principal` aparece? Por que identity-based normalmente não tem?
> **Respostas:** Não sei
23. Diferencie identity-based policy vs resource-based policy. Dê um exemplo de cada.
> **Respostas:** Não sei
24. O que é trust policy de uma role vs permission policy?
> **Respostas:** Não sei

### Avaliação de permissões
25. Explique a ordem de avaliação: Deny explícito, Allow explícito, Deny implícito.
> **Respostas:** Não sei
26. Um user tem Allow `s3:*` numa policy e Deny `s3:DeleteObject` em outra. Pode deletar? Por quê?
> **Respostas:** Não sei
27. Nenhuma policy menciona `ec2:TerminateInstances`. O user pode terminar instâncias? Por quê?
> **Respostas:** Não sei

### Higiene e menor privilégio
28. Cite 5 práticas de higiene mínima de conta AWS (root, MFA, chaves, etc.).
> **Respostas:** Não sei
29. Por que a AWS recomenda roles + federação em vez de access keys de longa duração?
> **Respostas:** Não sei
30. O que o IAM Access Analyzer ajuda a detectar?
> **Respostas:** Não sei
31. Como você escreveria (em palavras) uma policy que permite só `s3:GetObject` em `arn:aws:s3:::logs-prod/*` vindo de um IP corporativo?
> **Respostas:** Não sei

### Cenários
32. EC2 precisa ler um bucket S3. Access key na instância ou instance profile/role? Justifique.
> **Respostas:** instance profile/role, pois atrela permissão específica de um bucket para o EC2 ao invés de dar uma chave de acesso (que pode vazar), além de escapar a permissão para aquele recurso específico.
33. Conta A precisa acessar recurso na Conta B. Qual padrão IAM usar?
> **Respostas:** Não sei
34. Enunciado: "least privilege" e "temporary credentials". Que solução tende a vencer?
> **Respostas:** Não sei

**Confiança nesta UE:** baixa / média / alta  
**Respostas / anotações:**

```

```

---

## UE 3 — Well-Architected Framework

### Os 6 pilares
35. Liste os 6 pilares e, em uma frase cada, o que cada um otimiza.
> **Respostas:** Não sei o que são estes pilares
36. Qual pilar foi adicionado mais recentemente e o que ele endereça?
> **Respostas:** Não sei
37. Para cada pilar abaixo, cite 2 serviços/práticas típicos:
    - Operational Excellence
    - Security
    - Reliability
    - Performance Efficiency
    - Cost Optimization
    - Sustainability
> **Respostas:** Não sei

### Design principles
38. Explique "stop guessing capacity" e como elasticidade realiza isso.
> **Respostas:** Não sei
39. O que são *game days* e em qual pilar se encaixam principalmente?
> **Respostas:** Não sei
40. Cite 3 design principles transversais da nuvem (não só de um pilar).
> **Respostas:** Não sei

### Trade-offs (o que mais cai em prova) > **Respostas:** Não sei nenhum
41. Dê um trade-off concreto Reliability × Cost.
> **Respostas:** Não sei
42. Dê um trade-off Performance × Consistency (ou Performance × Cost).
> **Respostas:** Consistência necessita de verificações de estados, o que gere por sua vez mais processos (confirmação) para um fluxo, e em decorrência disto, acaba por não ter 100% do potencial de perfomance. Performance por sua vez não necessariamente implica em buscar fazer tudo o mais rápido, mas a medida que preciso performar, por muitas vezes preciso abrir um pouco a mão da consistência.  
43. Dê um trade-off Security × Operational Agility.
> **Respostas:** Burocracia, a medida que sistemas buscam ser seguros, muitas medidas e etapas para ler, modificar, criar e deletar são criadas a fim de mitigar brechas de segurança.
44. Como você raciocina quando **todas as opções "funcionam"** numa questão Professional?
> **Respostas:** Considero o contexto, por muitas vezes o objetivo "funcionar" torna-se o mais considerado, porém, a medida que o contexto abordado cita: "performance", "resiliência", "Consistência" como necessidades, então a resposta que melhor se adequa é aquela que compreende objetivos que vão além do "funcionar". É preciso dizer que mesmo quando, por exemplo, o objetivo for performance, ainda deve-se buscar equilibrio entre os outros pilares, como o custo, para que a solução seja viável. 

### Palavras-gatilho > **Respostas:** Não sei nenhum
Para cada frase, diga qual **pilar priorizar** e o tipo de solução que tende a vencer:
45. "most cost-effective"
46. "least operational overhead" / "fully managed"
47. "must not lose data" / "regardless of cost"
48. "compliance / encryption / audit trail"
49. "lowest latency for global users"

### Processo e ferramentas  > **Respostas:** Não sei nenhum
50. O que é WAFR? Quando fazer? O que são HRI e MRI?
> **Respostas:** não sei
51. Diferencie AWS Well-Architected Tool vs Trusted Advisor.
> **Respostas:** não sei
52. O que é uma *lens*? Quando usar Serverless Lens (ex.)?
> **Respostas:** não sei
53. Quais categorias o Trusted Advisor cobre? A cobertura depende de quê?
> **Respostas:** não sei

### Cenários > **Respostas:** Não sei nenhum
54. Arquitetura: web 3 camadas em 1 AZ, deploys manuais, sem backup testado, Spot no banco crítico. Liste riscos por pilar.
55. Negócio prioriza "menor fatura este trimestre", mas compliance exige criptografia e audit. Quais pilares entram em tensão e como decidir?

**Confiança nesta UE:** baixa / média / alta  
**Respostas / anotações:**

```
baixa
```

---

## UE 4 — Billing, FinOps e Precificação

### Fundamentos de preço
56. Quais são os 3 eixos típicos de cobrança AWS (compute, storage, data transfer)?
> **Respostas:** armazenamento: uso, frequência de acesso. compute: tempo de execução, potencia da máquina (RAM e CPU). transferencia de dados pela rede, mais detalhado na próxima pergunta
57. Regra geral: data transfer IN vs OUT vs cross-AZ vs cross-Region — o que é grátis e o que é pago?
> **Respostas:** grátis tudo que entra da internet e comunicação dentro da mesma AZ. Custo a partir da saida de trafego para internet, entre AZ e mais ainda entre regiões (tanto in quanto out)

### Modelos de compra
58. Compare On-Demand, Reserved Instances, Savings Plans e Spot (desconto, compromisso, interrupção, caso de uso).
> **Respostas:**
- on demand: custo sem descontos, recurso criado a medida que necessita, sem interrupções. Para recursos ou quantidade de recursos que não possuem previsibilidade de uso e não podem ser interrompidos.
- Reserved instances: compromisso de 1-3 anos de uso de recurso computacional, custo com desconto de até 70%. sem interrupção. Quando se quer economizar em ec2 e há embasamento para optar por este compromisso, como histórico ou perspectivas futuras bem fundamentadas.
- Saving plans: compromisso de 1-3 anos de uma valor a ser gasto mensalmente, desconto também de até 70%. Quando se economizar e tem sob padrão de custo um valor específico, ou seja, se é sabido que é gasto 40 dol a 1 ano, e exploradicamente 50dol em mes ou outro, pode-se optar por este plano para suprir os 40dol.
- spot instances: possui descontos de até 90%, pode ser interrompido com aviso de até no mínimo 2min antes. Quando deseja-se economizar e o serviço a estar nestas instancias não são criticos

59. Compute Savings Plans vs EC2 Instance Savings Plans: diferença e quando escolher cada.
> **Respostas:** 
60. Standard RI vs Convertible RI: trade-off.
61. Spot: aviso de interrupção, cargas adequadas e cargas **inadequadas**.
62. O que é a estratégia "layered" (baseline + pico + Spot)?

### Custos ocultos
63. Por que NAT Gateway costuma surpreender na fatura?
64. Por que servir mídia via CloudFront costuma ser mais barato que direto do S3/EC2 para internet?
65. Cross-AZ em arquitetura multi-AZ "mal pensada": o que pode inflar custo?

### Storage
66. Quando usar S3 Standard vs Intelligent-Tiering vs Standard-IA vs Glacier Deep Archive?
67. O que lifecycle policies fazem? Dê um exemplo de política.
68. EBS cobra por GB provisionado ou usado? Snapshots vão para onde?

### Governança FinOps
69. Para que servem Cost Explorer, AWS Budgets e CUR?
70. O que são cost allocation tags? Showback vs chargeback (ideia geral).
71. Consolidated billing (Organizations): benefício principal além da fatura única.
72. Cite 4 estratégias práticas de otimização de custo (além de "comprar RI").

### Cenários
73. Baseline estável 24/7 + picos imprevisíveis + jobs batch canceláveis. Como misturar modelos de compra?
74. Padrão de acesso S3 desconhecido. Qual storage class tende a ser a escolha segura e por quê?
75. Enunciado: "minimize cost" para processamento batch tolerante a falha. Spot, On-Demand ou RI? Por quê?

**Confiança nesta UE:** baixa / média / alta  
**Respostas / anotações:**

```
(escreva aqui)
```

---

## Bloco final — Integração entre UEs

Responda só se sobrar energia (muito úteis para mapear nível de arquiteto, não só memorização):

76. Desenhe mentalmente um e-commerce: escolha Region/AZ, identidade (como apps acessam S3/DB), 2 trade-offs Well-Architected conscientes, e modelo de compra do compute. Justifique em ~10 linhas.
77. Uma AZ cai. O que deve continuar funcionando se você aplicou bem UE1 + UE3 (Reliability)?
78. Credencial vazou num GitHub público. Quais práticas da UE2 teriam limitado o dano?
79. A fatura dobrou sem aumento de usuários. Quais 3 lugares olhar primeiro (UE4) e qual pilar WAF isso aciona?

**Confiança geral no módulo 01:** baixa / média / alta

---

## Checklist de entrega (para você)

- [ ] Respondi UE1
- [ ] Respondi UE2
- [ ] Respondi UE3
- [ ] Respondi UE4
- [ ] Marquei confiança em cada UE
- [ ] Pronto para pedir o mapeamento de lacunas

> Arquivo **temporário**. Pode apagar depois do diagnóstico.
>
> Próximo passo sugerido: *"Aqui estão minhas respostas — mapeia lacunas e me diga o que estudar em cada UE."*
