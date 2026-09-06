---
tipo: sessao-atual
ciclo: [[main]]
progresso: [[progresso]]
status: ativo
atualizado: 2026-08-17
bloco: A1
instrucao_ia: >
  Este e o unico ponto de partida de uma sessao neste ciclo.
  Sempre: (1) ler Onde estou, (2) confirmar Tema atual, (3) ensinar pela
  secao Explicacao profunda abaixo — nao pular para o proximo bloco.
  ESTILO OBRIGATORIO em toda explicacao (chat e este arquivo):
  exemplos reais (empresa/workload, nao so definicao); variacoes do mesmo
  mecanismo (o que muda e o que e igual); Terraform com arvore de arquivos
  e HCL da trust vs permission vs attachment.
  Depois da explicacao, fazer recall curto so deste tema.
  Quando o criterio de fechamento do bloco for atingido, atualizar
  Onde estou / Tema atual / Explicacao (trocar pelo proximo item de [[progresso]])
  e marcar o checkbox em [[progresso]].
---
# Sessão atual — ciclo pós-diagnóstico

Arquivo para **sempre** abrir primeiro. Fluxo fixo:

1. **Onde estou** — posição no ciclo
2. **Tema atual** — a lacuna desta sessão
3. **Explicação profunda** — estudar isto agora (não a UE inteira)

Estilo fixo da explicação: **exemplo real + variações + Terraform (pasta + HCL)**.

Checklist sequencial: [[progresso]] · Direção do ciclo: [[main]]

---

## 1. Onde estou

| Campo | Valor |
|---|---|
| Módulo | 01 — Fundamentos e Well-Architected |
| Ciclo | 2026-08-13 pós-diagnóstico |
| Nível do módulo | inicial / lacunas grandes (~15% no diagnóstico) |
| Sessão | **A — UE2 Identidade** (prioridade 1, maior buraco) |
| Bloco aberto | **A1** — User vs Role vs Group vs STS |
| Próximos (não abrir agora) | A2 policy → A3 avaliação → A4 higiene/cross-account → A5 fechar UE2 |
| Já sólido neste tema | EC2 acessa S3 com instance profile/role, não com access key na instância |
| Erro a corrigir cedo | IAM User ≠ Role; Group agrupa users, não é Principal, não aninha |

**Critério para fechar A1:** explicar `AssumeRole` em 3 frases; dizer por que instance profile vence access key na EC2.

---

## 2. Tema atual

**User vs Role vs Group vs STS**

Por que este bloco existe: no diagnóstico a única âncora sólida de IAM foi “EC2 não leva chave no disco”. O resto do vocabulário (o que é User, o que é Role, o que o Group não é, o que o STS faz) estava errado ou em branco. Sem isso, policy e avaliação de permissão (A2/A3) não grudam.

Material-fonte (não substituir esta explicação; usar se quiser o texto original da UE): [[../UE - Identidade Base e Estrutura de Conta/main|UE2 main]] · [[../UE - Identidade Base e Estrutura de Conta/3-Detalhamento|detalhamento §1]]

---

## 3. Explicação profunda — A1

### 3.1 O problema que o IAM resolve

Uma conta AWS é um recipiente de recursos (buckets, instâncias, bancos). Alguém precisa **agir** sobre esses recursos: pessoa no console, pipeline de CI, função Lambda, instância EC2.

O IAM responde só isto: **quem** é, **o que pode fazer**, e **por quanto tempo a prova de identidade vale**.

Três erros clássicos nascem de confundir as peças:

- Tratar **Role** como “um tipo de User” (são identidades de natureza diferente).
- Tratar **Group** como identidade que faz login (não é).
- Gravar **access key** em disco/código porque “precisa autenticar na API” (a Role + STS existem exatamente para não fazer isso).

---

### 3.2 Quatro peças, quatro funções

Pense em um prédio corporativo:

| Peça IAM | Analogia | O que é de verdade |
|---|---|---|
| **Root** | Dono do prédio com a chave-mestra | Identidade nascida com a conta; poder irrestrito; **não é este bloco** (volta em A4) |
| **IAM User** | Crachá nominal permanente de um funcionário | Identidade **com credencial de longa duração** (senha e/ou access key) |
| **IAM Group** | Lista do departamento no RH | **Não é identidade.** Só agrupa Users para anexar a mesma policy |
| **IAM Role** | Crachá temporário de visitante / de um cargo | Identidade **sem senha permanente**; você a **assume** e o **STS** emite credencial que **expira** |
| **STS** | Máquina que imprime o crachá temporário | Serviço que gera tokens de curta duração quando alguém assume uma Role |

A confusão User vs Role não é “os dois têm permissões”. Os dois **recebem policies**. A diferença é **como provam quem são** e **quanto tempo essa prova vive**.

---

### 3.3 IAM User — identidade de longa duração

Um User representa **uma pessoa ou (legado) uma aplicação** que precisa de credenciais **permanentes** até alguém revogá-las:

- **Senha** → console.
- **Access Key ID + Secret Access Key** → CLI, SDK, API.

Enquanto a chave existir e não for desativada, quem a tiver **é** aquele User para a AWS. Não há expiração automática. Por isso vazamento em Git, `.env` ou log é o vetor nº 1: a chave continua válida amanhã e daqui a um ano.

A AWS moderna **evita Users no dia a dia**. Pessoas entram por federação (Identity Center / SSO) e recebem Roles. Users ainda existem, mas o default de arquitetura é: **pessoa e workload assumem Role**.

O que um User **não** é:

- Não é um “grupo com login”.
- Não é assumível. Ninguém faz `AssumeRole` *em* um User. User **é** a identidade; Role **é assumida**.

---

### 3.4 IAM Group — organizador, não ator

Group existe para um único trabalho: **anexar a mesma identity-based policy a vários Users**.

Exemplo: Users `ana`, `bruno`, `carla` no Group `Developers`. A policy “pode abrir tickets no console e ler S3 de staging” vai no Group uma vez.

Limites que caem em prova e que o diagnóstico errou:

1. **Não é identidade.** Não existe login no Group. Não existe access key do Group.
2. **Não pode ser Principal.** Principal = “quem”. Policies que dizem *quem pode assumir / quem pode acessar este recurso* apontam para User, Role, conta, serviço — **não** para Group.
3. **Não aninha.** Group não contém Group. Não há hierarquia tipo AD.
4. **Só organiza Users**, não Roles. Role não “entra em Group”.

Se você precisa de permissão compartilhada entre workloads, isso é **a mesma Role** (ou policies gerenciadas anexadas a várias Roles) — não Group.

---

### 3.5 IAM Role — identidade sem chave permanente

Uma Role é uma identidade que **não tem senha nem access key própria gravada**. Ninguém “é” a Role o tempo todo. Alguém **assume** a Role por um período.

Isso desdobra a Role em **dois documentos**:

1. **Trust policy** (quem *pode assumir*)  
   Resource-based na própria Role. Diz o **Principal**: User X, conta A, serviço `ec2.amazonaws.com`, IdP SAML, etc. A ação implícita é `sts:AssumeRole` (ou equivalentes de federação).
2. **Permission policy** (o que *pode fazer depois de assumida*)  
   Identity-based na Role: `s3:GetObject` neste bucket, `dynamodb:Query` nesta tabela, etc.

Separar os dois é o clique mental. “A Role pode ler S3” é permission. “A instância EC2 / a conta A / o usuário federado pode virar essa Role” é trust.

Casos de uso que justificam a Role existir (todos o mesmo mecanismo, Principals diferentes):

- **Compute → outros serviços:** EC2, Lambda, ECS assumem uma Role em vez de carregar chave.
- **Cross-account:** identidade na conta A assume Role na conta B (detalhe no A4).
- **Federação:** humano autenticado no IdP corporativo assume Role na AWS.
- **Serviço AWS no seu nome:** service-linked role.

---

### 3.6 STS e o que `AssumeRole` realmente faz

**STS (Security Token Service)** não é “mais um tipo de identidade”. É o serviço que **emite credenciais temporárias**.

Fluxo, em ordem:

1. Uma identidade já autenticada (User, outra Role, ou o serviço EC2 via instance profile) chama `sts:AssumeRole` apontando o ARN da Role alvo.
2. O STS avalia a **trust policy**: este Principal tem permissão de assumir?
3. Se sim, o STS devolve um pacote de curta duração:
   - `AccessKeyId`
   - `SecretAccessKey`
   - `SessionToken` (obrigatório em credencial temporária — sem ele a chamada falha)
4. A partir daí, as APIs são assinadas **como a Role**, não como o ator original. A permission policy da Role é o que vale para essas chamadas.
5. Quando a sessão expira (minutos a horas, configurável até o máximo da Role), o token **morre sozinho**. Assumir de novo gera outro token.

Três frases para gravar (critério de fechamento):

> AssumeRole é o pedido ao STS para **trocar** uma identidade atual por uma Role.  
> O STS só aceita se a **trust policy** da Role listar aquele Principal.  
> O retorno é credencial **temporária** (inclui session token) com a qual você age **como a Role** até expirar.

Detalhe de prova: credencial temporária **sempre** leva `SessionToken`. Copiar só Access Key + Secret de uma sessão STS, sem o token, não funciona.

---

### 3.7 Por que Role vence access key longa (e o caso EC2 → S3)

Access key de User:

- Vive até revogação manual.
- Costuma ser copiada para servidor, CI, notebook, Git.
- Vazou = atacante é aquele User até você achar e desativar.

Credencial de Role:

- Nascida no STS, TTL curto, **rotação automática**.
- Na EC2, o hypervisor injeta a Role via **instance profile**: o agente/SDK consulta o **Instance Metadata Service (IMDS)** e recebe token fresco. Nada é gravado no AMI, no user-data, nem no Git.
- Vazou um token de 1 hora: a janela de abuso é aquela hora. Não há chave eterna no disco.

Por isso o padrão da Q32 (que você já acertou) não é preferência estética: **a instância assume a Role; o SDK pega credencial no IMDS; o S3 vê a Role, não um User.**

O que está errado em “criar User, gerar access key, colar na EC2”:

- Chave permanente no disco (snapshot, AMI, logs, Git).
- Rotação é trabalho humano.
- Não há vínculo automático instância ↔ identidade: a chave funciona de **qualquer lugar** da internet se vazar, não só daquela EC2.

Instance profile = jeito da AWS **anexar uma Role a uma instância**. A Role continua sendo Role; o profile é só o gancho na EC2.

---

### 3.8 Mapa mental único (para não misturar de novo)

```
Conta
 ├── Root          → dono; não usar no dia a dia
 ├── User          → identidade + credencial LONGA (evitar)
 ├── Group         → pasta de Users; não loga; não é Principal
 └── Role          → identidade SEM credencial longa
        ├── trust policy      → quem pode AssumeRole
        └── permission policy → o que pode fazer depois
              ↑
         STS emite token curto quando AssumeRole passa na trust
```

Workload (EC2/Lambda) **nunca** deveria ser um User com chave no código. Workload **é** uma Role.

---

### 3.9 Erros do diagnóstico, reescritos certo

| O que costuma aparecer | Correção |
|---|---|
| “Role é um User especial” | User tem credencial permanente. Role é assumida; STS emite credencial temporária. |
| “Group é uma identidade / Principal” | Group só agrupa Users. Sem login, sem Principal, sem aninhamento. |
| “Policy é o que diferencia User de Role” | Policy descreve **ações**. A diferença é o **ciclo de vida da credencial** e o **AssumeRole**. |
| “EC2 precisa de access key para chamar S3” | EC2 usa instance profile → Role → STS/IMDS. |

---

### 3.10 Empresa de referência (um cenário, várias peças)

Empresa **LojaX**, conta `123456789012`, Region `sa-east-1`.

| Ator | Precisa | Peça IAM certa | Peça errada |
|---|---|---|---|
| Ana e Bruno, time de staging | Console + CLI no dia a dia | Group `developers-staging` com Users **ou** (melhor) Identity Center → Role | Access key de cada um no Slack |
| App PHP na EC2 `web-01` | Ler imagens em `s3://lojax-prod-imagens` | Role `role-ec2-web` + **instance profile** | User `app-user` com key no `.env` da instância |
| Lambda que gera thumbnail | Escrever em `s3://lojax-prod-thumbs` | Role `role-lambda-thumbs` (trust `lambda.amazonaws.com`) | A mesma key da EC2, reutilizada |
| GitHub Actions (deploy) | `ecs:UpdateService` 10 min por pipeline | Role `role-gha-deploy` com trust **OIDC** do GitHub | Access key de User `ci` no GitHub Secrets |
| Conta Analytics `999...` | Ler o bucket de prod (depois, A4) | Role na conta LojaX com trust na conta Analytics | Copiar key de prod para a outra conta |

O mecanismo é sempre o mesmo: **trust (quem assume) + permission (o que faz) + STS (token curto)**. O que muda é o **Principal** da trust.

---

### 3.11 Variações do mesmo `AssumeRole`

| Variação | Quem é o Principal da trust | Como o token chega | TTL típico |
|---|---|---|---|
| **EC2 instance profile** | `ec2.amazonaws.com` | IMDS (`169.254.169.254`), SDK busca sozinho | ~6 h, renovado pela instância |
| **Lambda execution role** | `lambda.amazonaws.com` | Runtime injeta no ambiente da invocação | Vida da invocação / sessão interna |
| **ECS task role** | `ecs-tasks.amazonaws.com` | Agent da task (não confundir com *task execution role*, que só puxa imagem/logs) | Vida da task |
| **Humano no CLI** | ARN do User ou da Role federada | `aws sts assume-role` / profile `role_arn` no `~/.aws/config` | 1 h default (max 12 h se a Role permitir) |
| **CI GitHub OIDC** | IdP `token.actions.githubusercontent.com` + `sub` do repo | Action troca JWT por credencial STS. **Zero access key no GitHub.** | Minutos |
| **Cross-account** | `arn:aws:iam::CONTA_A:root` (depois restringe) | Conta A chama AssumeRole na Role da conta B | Igual a humano/CI |

O que **não** muda entre essas linhas: não existe senha da Role; STS emite Access Key + Secret + **SessionToken**; permission policy da Role é o que o S3/ECS enxerga.

Variação perigosa (anti-padrão ainda comum):

```hcl
# NÃO FAÇA: User de aplicação + access key no disco
resource "aws_iam_access_key" "app" {
  user = aws_iam_user.app.name
}
# e depois: export AWS_ACCESS_KEY_ID=... no user-data da EC2
```

Essa key funciona de um laptop em outro país. Instance profile só entrega token **daquela** instância.

Detalhe Terraform/IAM que aparece no mundo real: quem **anexa** a Role na EC2 precisa `iam:PassRole` naquela Role. Não é AssumeRole. PassRole = “posso entregar este crachá a este serviço”. AssumeRole = “eu visto o crachá”. Volta com mais rigor em A2/A4.

---

### 3.12 Terraform — estrutura de pastas

Não misture User, Role e compute no mesmo arquivo. Separe **identidade** de **quem a usa**:

```
infra/
├── providers.tf
├── variables.tf
└── iam/
    ├── users_groups.tf          # pessoas (legado / exceção)
    ├── policy_s3_imagens.tf     # permission reutilizável
    ├── role_ec2_web.tf          # trust EC2 + attachment + instance profile
    ├── role_lambda_thumbs.tf    # mesma permission pattern, trust diferente
    └── role_gha_deploy.tf       # trust OIDC (variação CI)
# compute (outro módulo) só referencia o instance profile / role.arn
```

Regra de leitura do HCL:

| Recurso Terraform | É o quê no IAM |
|---|---|
| `aws_iam_role.assume_role_policy` | **Trust** — quem pode AssumeRole |
| `aws_iam_policy` + `aws_iam_role_policy_attachment` | **Permission** — o que pode fazer |
| `aws_iam_role_policy` (inline) | Permission colada na Role (ok em demo; em prod prefira policy gerenciada sua) |
| `aws_iam_instance_profile` | Gancho EC2 → Role. A Role continua sendo Role |
| `aws_iam_group` + `aws_iam_group_policy_attachment` | Permission em massa para **Users** |
| `aws_iam_user` / `aws_iam_access_key` | Credencial longa — evitar em workload |

Use `data.aws_iam_policy_document` em vez de JSON cru: o Terraform valida o documento e o diff fica legível.

---

### 3.13 Terraform — configs do cenário LojaX

`iam/policy_s3_imagens.tf` — **permission** (não diz quem assume):

```hcl
data "aws_iam_policy_document" "ler_imagens" {
  statement {
    sid       = "LerObjetosDoBucketImagens"
    effect    = "Allow"
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::lojax-prod-imagens/*"]
  }
}

resource "aws_iam_policy" "ler_imagens" {
  name   = "lojax-prod-s3-ler-imagens"
  policy = data.aws_iam_policy_document.ler_imagens.json
}
```

`iam/role_ec2_web.tf` — **trust EC2** + attachment + instance profile:

```hcl
data "aws_iam_policy_document" "ec2_assume" {
  statement {
    sid     = "EC2AssumeThisRole"
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ec2_web" {
  name               = "role-ec2-web"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume.json
  # assume_role_policy = TRUST, não a policy de S3
}

resource "aws_iam_role_policy_attachment" "ec2_web_s3" {
  role       = aws_iam_role.ec2_web.name
  policy_arn = aws_iam_policy.ler_imagens.arn
}

resource "aws_iam_instance_profile" "ec2_web" {
  name = "profile-ec2-web"
  role = aws_iam_role.ec2_web.name
}

# no módulo compute:
# resource "aws_instance" "web" {
#   iam_instance_profile = aws_iam_instance_profile.ec2_web.name
#   ...
# }
```

Variação **Lambda** — mesma permission, **trust muda**:

```hcl
data "aws_iam_policy_document" "lambda_assume" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_thumbs" {
  name               = "role-lambda-thumbs"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume.json
}

resource "aws_lambda_function" "thumbs" {
  function_name = "lojax-thumbs"
  role          = aws_iam_role.lambda_thumbs.arn
  # ...
}
```

Se você colar `ec2.amazonaws.com` na trust da Lambda, a função **não sobe** (ou não assume). Trust errada = STS recusa. Permission errada = assume, mas S3 nega.

Variação **humano + Group** (legado; ainda cai em conta antiga):

```hcl
resource "aws_iam_user" "ana" {
  name = "ana"
}

resource "aws_iam_group" "developers_staging" {
  name = "developers-staging"
}

resource "aws_iam_group_membership" "developers_staging" {
  name  = "developers-staging-members"
  group = aws_iam_group.developers_staging.name
  users = [aws_iam_user.ana.name]
}

# Policy no Group — Ana herda. O Group NÃO loga e NÃO é Principal.
resource "aws_iam_group_policy_attachment" "dev_s3" {
  group      = aws_iam_group.developers_staging.name
  policy_arn = aws_iam_policy.ler_imagens.arn
}
```

Variação **Ana assume Role pelo CLI** (User existe, mas o poder está na Role):

```hcl
data "aws_iam_policy_document" "ana_assume_web" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      type        = "AWS"
      identifiers = [aws_iam_user.ana.arn]
    }
  }
}

resource "aws_iam_role" "dev_breakglass_s3" {
  name               = "role-dev-breakglass-s3"
  assume_role_policy = data.aws_iam_policy_document.ana_assume_web.json
  max_session_duration = 3600
}
```

```ini
# ~/.aws/config  (não é Terraform; é o cliente)
[profile lojax-s3]
role_arn         = arn:aws:iam::123456789012:role/role-dev-breakglass-s3
source_profile   = ana
```

`aws s3 ls` com `--profile lojax-s3` dispara AssumeRole. Sem esse profile, Ana só tem o que o Group deu.

---

### 3.14 Depois desta leitura (ainda neste bloco)

Feche o arquivo e responda, de memória, no chat:

1. User vs Role: credencial e caso de uso.
2. Group: login? Principal? aninha?
3. AssumeRole em 3 frases.
4. Por que Role vence access key se vazar no GitHub.
5. Por que instance profile na EC2, não chave no disco.
6. No Terraform, o que vai em `assume_role_policy` vs `aws_iam_role_policy_attachment`? O que o `aws_iam_instance_profile` faz?
7. Se a trust da Role da Lambda tiver `ec2.amazonaws.com`, o que quebra — assume ou o S3?

Quando estiver estável, este arquivo avança para **A2 — Anatomia de policy**.
