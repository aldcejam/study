---
tema: Identidade Base e Estrutura de Conta AWS
tipo: detalhamento-tecnico
contexto_pai: [[./main|Voltar para a UE]]
tags: [aws, iam, seguranca, conta, policies, sts]
fontes:
  - AWS IAM User Guide (docs oficiais)
  - AWS Well-Architected Security Pillar (whitepaper)
  - AWS IAM Best Practices (docs oficiais)
  - AWS Policy Evaluation Logic (docs oficiais)
---
# 📚 Detalhamento — Identidade Base e Estrutura de Conta

> [!ABSTRACT] Como usar esta nota
> Material de referência denso para acompanhar a `main.md`. IAM é o **coração da segurança na AWS** e o assunto mais transversal de todas as provas. Domine a **lógica de avaliação** (seção 3) — é o que separa quem "sabe IAM" de quem "decora IAM".

---

## 1. Identidades: root, users, groups e roles

### 1.1 A conta AWS e o usuário root

- Toda **conta AWS** nasce com um **usuário root**, identificado pelo **e-mail** usado no cadastro. Ele tem **poder irrestrito** — inclusive ações que *nenhuma* outra identidade pode fazer.
- **Ações exclusivas do root** (não delegáveis a IAM users):
  - Alterar/fechar a conta, mudar o plano de suporte.
  - Alterar e-mail/senha da conta raiz.
  - Restaurar permissões de política IAM revogadas por engano.
  - Configurar MFA na própria raiz, mudar configurações de billing.
  - Registrar como vendedor no Marketplace, algumas ações de S3 com MFA Delete.
- **Regra de ouro:** **Ative MFA no root, guarde as credenciais em cofre, e NUNCA use o root no dia a dia.** Crie um usuário/role administrativo para operação.

### 1.2 IAM Users (identidades de longa duração)

- Representam **uma pessoa ou aplicação** com credenciais **permanentes**:
  - **Senha** (para o Console).
  - **Access Key ID + Secret Access Key** (para CLI/SDK/API).
- **Problema:** credenciais de longa duração são o **maior vetor de vazamento** (chaves em repositórios Git, em `.env`, em logs). A recomendação moderna da AWS é **evitar IAM users** e usar **federação + roles** (IAM Identity Center / SSO).

### 1.3 IAM Groups

- **Coleção de users** para atribuir políticas em massa (ex.: grupo `Developers`, `Admins`).
- **Não** é uma identidade: você não "faz login" em um grupo, nem um grupo pode ser Principal de uma policy. Um grupo **não pode conter outro grupo**.
- Serve só para **organizar permissões** de users.

### 1.4 IAM Roles (identidades de curta duração) — o mecanismo mais importante

- Uma **role** é uma identidade **sem credenciais permanentes**. Ela é **assumida** temporariamente, gerando credenciais de curta duração via **AWS STS (Security Token Service)**.
- Componentes de uma role:
  - **Trust policy** (quem *pode assumir* — o Principal): define quem tem permissão de `sts:AssumeRole`.
  - **Permission policy** (o que a role *pode fazer* uma vez assumida).
- **Casos de uso centrais:**
  - **EC2/Lambda/ECS acessando outros serviços** (Instance Profile / execution role) — em vez de gravar chaves na instância.
  - **Cross-account access** — conta A assume role na conta B.
  - **Federação** — usuários do Active Directory / Google / SAML / OIDC assumem roles.
  - **AWS services** agindo em seu nome (service-linked roles).

> [!IMPORTANT] Por que roles > access keys
> Credenciais de role são **temporárias** (minutos a horas) e **rotacionadas automaticamente**. Se vazarem, expiram sozinhas. Chaves de IAM user são permanentes até você revogá-las manualmente. **Preferir roles é o padrão de segurança #1 da AWS.**

### 1.5 Comparativo de identidades

| | Root | IAM User | IAM Role |
|---|---|---|---|
| Credencial | Permanente (dona da conta) | Permanente | Temporária (STS) |
| Uso diário | ❌ Nunca | Evitar | ✅ Preferir |
| Assumível | — | Não | Sim (via trust policy) |
| Ideal para | Setup inicial + ações exclusivas | Legado / casos específicos | Apps, cross-account, federação |

---

## 2. Anatomia de uma Policy JSON

### 2.1 Estrutura de um statement

Uma policy é um documento JSON com um ou mais **statements**:

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "PermitirLeituraBucketX",
      "Effect": "Allow",
      "Action": ["s3:GetObject", "s3:ListBucket"],
      "Resource": [
        "arn:aws:s3:::meu-bucket",
        "arn:aws:s3:::meu-bucket/*"
      ],
      "Condition": {
        "IpAddress": { "aws:SourceIp": "203.0.113.0/24" }
      }
    }
  ]
}
```

Elementos:

- **Version:** sempre `"2012-10-17"` (a versão da linguagem de policy, não a data de hoje).
- **Sid:** identificador opcional do statement.
- **Effect:** `Allow` ou `Deny`.
- **Action:** operações da API (`service:Operation`, ex.: `s3:GetObject`). Aceita curinga (`s3:*`, `s3:Get*`).
- **Resource:** ARN(s) alvo. `*` = qualquer recurso.
- **Principal:** **quem** (só em resource-based e trust policies — ver abaixo). Ausente em identity-based.
- **Condition:** restrições contextuais (IP, MFA, tags, hora, VPC de origem, etc.).

### 2.2 ARN — Amazon Resource Name

Formato do identificador único de qualquer recurso:

```
arn:partition:service:region:account-id:resource
arn:aws:s3:::meu-bucket/arquivo.txt          (S3 não usa region/account)
arn:aws:iam::123456789012:user/Alice
arn:aws:ec2:sa-east-1:123456789012:instance/i-0abc123
```

### 2.3 Identity-based vs Resource-based policies

Distinção **fundamental**:

| | Identity-based policy | Resource-based policy |
|---|---|---|
| Anexada a | User, group, role | Recurso (bucket S3, fila SQS, role trust, KMS key) |
| Tem `Principal`? | ❌ Não (a identidade É o principal) | ✅ Sim (define *quem* de fora pode acessar) |
| Exemplo | "Alice pode ler o bucket X" | "O bucket X pode ser lido por Alice/Conta B" |
| Cross-account | Precisa de policy dos **dois lados** | Habilita acesso cross-account diretamente |

- **Managed policies:** reutilizáveis, versionadas.
  - **AWS managed** (mantidas pela AWS, ex.: `AmazonS3ReadOnlyAccess`).
  - **Customer managed** (você cria e reutiliza).
- **Inline policies:** embutidas em uma única identidade, relação 1:1 (some se a identidade some). Usar com parcimônia.

### 2.4 Outros mecanismos de controle (visão geral)

- **Permissions Boundaries:** teto máximo de permissões que uma identidade pode ter (delega criação de users sem escalonamento de privilégio).
- **Service Control Policies (SCPs):** aplicadas via **AWS Organizations** a contas inteiras; definem o **teto** de permissões da conta (não *concedem*, apenas *limitam*). Aprofundado na categoria de Organizations/governança.
- **Session policies:** passadas no momento do `AssumeRole` para restringir ainda mais a sessão.

---

## 3. Lógica de Avaliação de Permissões (o coração de IAM)

### 3.1 O algoritmo de decisão

Quando uma requisição chega, a AWS avalia **todas** as políticas aplicáveis e decide segundo esta ordem de precedência:

```
1. Deny explícito?  ──► SIM ──► NEGADO (fim, sempre vence)
        │ NÃO
2. Allow explícito? ──► NÃO ──► NEGADO (deny implícito / default deny)
        │ SIM
3. ──► PERMITIDO (se nenhum Deny explícito bloqueou)
```

Regras de ouro:

1. **Tudo é negado por padrão** (implicit deny).
2. Um **Allow explícito** libera.
3. Um **Deny explícito** **sempre** vence qualquer Allow — é irrevogável na avaliação.

### 3.2 Como múltiplas políticas se combinam

As permissões de uma identidade são a **união** de todas as suas policies (identity-based + group + inline). Mas a decisão final passa por **todas** as camadas presentes:

- **SCP** (Organizations) deve permitir → senão nega, mesmo com Allow no IAM.
- **Permissions boundary** deve permitir → senão nega.
- **Identity-based policy** deve permitir.
- **Resource-based policy** pode conceder de forma independente (cross-account).
- **Qualquer Deny explícito** em qualquer camada → **NEGADO**.

> [!IMPORTANT] Mnemônico de avaliação
> **"Deny explícito ganha de tudo. Sem Allow = negado. Todas as camadas (SCP, boundary, identity, resource) precisam alinhar."**

### 3.3 Exemplo raciocinado

Alice tem policy `Allow s3:*` em todos os buckets, mas um SCP na conta tem `Deny s3:DeleteBucket`.
→ Alice pode listar, ler, escrever... mas **não pode deletar bucket**, porque o Deny explícito do SCP vence o Allow do IAM.

### 3.4 Ferramentas de teste

- **IAM Policy Simulator:** testa se uma ação seria permitida/negada, sem executá-la.
- **`aws sts assume-role` + `--dry-run`** e o próprio erro `AccessDenied` (que informa qual policy negou, quando `Encoded authorization failure message` é decodificado com `aws sts decode-authorization-message`).
- **IAM Access Analyzer:** detecta recursos expostos externamente e gera policies de menor privilégio a partir de logs do CloudTrail.

---

## 4. Princípio do Menor Privilégio e Higiene de Conta

### 4.1 Least Privilege (menor privilégio)

- Conceda **apenas** as permissões necessárias para a tarefa, **nada mais**.
- Comece **restritivo** e **abra conforme a necessidade** (não o contrário: nunca `*:*` "para funcionar por enquanto").
- Use **Access Analyzer** para gerar policies a partir do uso real (CloudTrail).
- Refine com **Conditions** (ex.: exigir MFA, restringir por IP/VPC, por tag).

### 4.2 Checklist de higiene de conta (IAM Best Practices)

- [ ] **MFA em tudo:** obrigatório no root e em usuários privilegiados. Preferir MFA de hardware/FIDO2 para o root.
- [ ] **Não usar o root** no dia a dia; **remover access keys do root** se existirem.
- [ ] **Preferir roles a access keys**; para humanos, usar **IAM Identity Center (SSO)** com federação.
- [ ] **Rotacionar/eliminar credenciais** de longa duração; auditar chaves não usadas.
- [ ] **Menor privilégio** com managed policies bem escopadas + conditions.
- [ ] **Permissions boundaries** ao delegar criação de identidades.
- [ ] **Monitorar:** CloudTrail (auditoria de API), IAM Access Analyzer, credential report.
- [ ] **Centralizar** governança multi-conta com **Organizations + SCPs**.

### 4.3 Estruturas avançadas citadas

- **Break-glass account/role:** identidade de emergência, altamente auditada, usada só quando o acesso normal falha.
- **Federação corporativa:** integra Active Directory / IdP via SAML 2.0 ou OIDC → usuários assumem roles temporárias, **sem IAM users** por pessoa. Aprofundado na categoria 06.
- **IAM Identity Center (ex-AWS SSO):** ponto único de acesso a múltiplas contas e apps com permission sets.

---

## 🧭 Tabela de Decisão Rápida

| Situação | Escolha |
|---|---|
| App em EC2/Lambda precisa acessar S3 | **IAM Role** (instance profile / execution role) |
| Dar acesso a um time humano em várias contas | **IAM Identity Center (SSO)** + federação |
| Conta B precisa acessar recurso da conta A | **Resource-based policy** + role cross-account |
| Limitar o teto de permissões de contas inteiras | **SCP** (Organizations) |
| Delegar criação de users sem risco de escalonamento | **Permissions boundary** |
| Restringir acesso a horário/IP/MFA | **Condition** na policy |
| Ação exclusiva da conta (fechar conta, billing raiz) | **Root** (com MFA) |

---

## 📖 Glossário

- **Root user:** dono absoluto da conta, identificado por e-mail; usar só para ações exclusivas.
- **IAM User:** identidade de pessoa/app com credenciais permanentes.
- **IAM Group:** coleção de users para atribuir políticas em massa.
- **IAM Role:** identidade temporária assumível via STS.
- **STS (Security Token Service):** emite credenciais temporárias (AssumeRole).
- **Trust policy:** define *quem pode assumir* uma role (Principal).
- **Identity-based policy:** anexada a user/group/role, sem Principal.
- **Resource-based policy:** anexada ao recurso, com Principal (habilita cross-account).
- **SCP (Service Control Policy):** teto de permissões por conta via Organizations.
- **Permissions boundary:** teto máximo de permissões de uma identidade.
- **ARN:** identificador único de recurso AWS.
- **Least privilege:** conceder o mínimo necessário.
- **Access Analyzer:** detecta exposição externa e gera policies mínimas.
- **Explicit Deny / Implicit Deny:** negação escrita (vence tudo) / negação padrão (ausência de Allow).

---

## ✅ Active Recall (responda sem olhar)

1. Cite três ações que **apenas o root** pode executar.
2. Por que roles são preferíveis a access keys de IAM user? O que acontece se uma credencial de role vazar?
3. Descreva a ordem de avaliação: onde entram deny explícito, allow explícito e deny implícito?
4. Um user tem `Allow s3:*` mas o SCP da conta tem `Deny s3:DeleteBucket`. O que ele consegue fazer? Por quê?
5. Qual a diferença entre identity-based e resource-based policy? Qual delas tem `Principal`?
6. O que é uma permissions boundary e qual problema de segurança ela previne?
7. Quais os dois componentes de uma role (trust vs permission) e o que cada um controla?
8. Como o Access Analyzer ajuda a implementar least privilege na prática?

---

## 🔗 Fontes e leitura oficial

- [IAM User Guide — Introdução](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction.html)
- [IAM Best Practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
- [Policy Evaluation Logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
- [Identity-based vs Resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)
- [AWS STS — Temporary Credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html)
- [Well-Architected — Security Pillar](https://docs.aws.amazon.com/wellarchitected/latest/security-pillar/welcome.html)
- [IAM Access Analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html)
