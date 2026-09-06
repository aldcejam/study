---
tema: Criptografia e Gestao de Segredos - KMS
tipo: unidade-estudo
tags: [aws, kms, criptografia, secrets-manager, seguranca]
---
# 🧪 UE - Criptografia e Gestao de Segredos

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Criptografar dados é obrigatório por compliance, mas o verdadeiro problema é a gestão de chaves: onde a chave vive, quem pode usá-la e como ela é rotacionada. Chaves hardcoded no código e segredos em variáveis de ambiente em texto plano são vetores de ataque clássicos. KMS resolve isso com envelope encryption e controle de acesso via IAM, mas exige entender key policies (que podem sobrepor o IAM) para não perder acesso irreversivelmente às chaves.

## 🧬 Grade Atômica de Tópicos
1. **KMS e envelope encryption:** CMK/KMS keys (AWS-managed vs customer-managed), data keys, envelope encryption, criptografia em repouso integrada aos serviços (S3, EBS, RDS).
2. **Key policies e controle de acesso:** Key policy vs IAM, grants, rotação automática, multi-Region keys, deletion com waiting period.
3. **CloudHSM e casos de compliance:** HSM dedicado (FIPS 140-2 Level 3), quando é exigido vs KMS; custódia de chaves.
4. **Gestão de segredos e certificados:** Secrets Manager (rotação automática integrada a RDS) vs SSM Parameter Store; ACM para certificados TLS e integração com ALB/CloudFront.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html) e [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
- **System Blueprint:** Criptografia fim-a-fim: TLS via ACM na borda + SSE-KMS em repouso + segredos de banco rotacionados pelo Secrets Manager.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Exercitar envelope encryption e rotação.
- [ ] Criar uma KMS key e cifrar/decifrar dados via CLI observando o data key.
- [ ] Guardar uma credencial de RDS no Secrets Manager com rotação automática.
- [ ] Escrever uma key policy que restringe uso e testar acesso negado.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - KMS e Envelope Encryption]]
- [[./02 - Key Policies e CloudHSM]]
- [[./03 - Secrets Manager, Parameter Store e ACM]]
