---
tema: Amazon S3 Profundo
tipo: unidade-estudo
tags: [aws, s3, armazenamento, objeto, durabilidade]
---
# 🧪 UE - S3 Profundo

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> S3 é o serviço mais usado e mais mal-configurado da AWS: buckets públicos por engano são a causa clássica de vazamentos de dados, e a escolha errada de storage class desperdiça milhares de dólares. Além disso, entender o modelo de durabilidade (11 noves), consistência forte (desde 2020) e particionamento de performance é o que separa quem usa S3 como "pasta" de quem o usa como data lake de alta performance.

## 🧬 Grade Atômica de Tópicos
1. **Storage classes e lifecycle:** Standard, Intelligent-Tiering, Standard-IA, One Zone-IA, Glacier Instant/Flexible/Deep Archive; políticas de lifecycle e custos de recuperação/transição.
2. **Durabilidade, consistência e versionamento:** 11 noves, consistência forte read-after-write, versioning, MFA Delete, replicação (CRR/SRR), Object Lock (WORM).
3. **Segurança:** Block Public Access, bucket policies vs ACLs, OAC para CloudFront, SSE-S3/SSE-KMS/SSE-C, VPC Gateway Endpoint, pre-signed URLs.
4. **Performance e escala:** Particionamento por prefixo, requests/s por prefixo, multipart upload, S3 Transfer Acceleration, S3 Select, requester pays.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon S3 User Guide](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html).
- **System Blueprint:** Data lake sobre S3 com Intelligent-Tiering + lifecycle para Glacier + acesso analítico via Athena (ver categoria 10).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Configurar um bucket seguro e econômico.
- [ ] Criar bucket com Block Public Access, versioning e SSE-KMS.
- [ ] Definir uma lifecycle rule que move objetos para Glacier após 30 dias.
- [ ] Gerar uma pre-signed URL e testar acesso temporário sem credenciais.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Storage Classes e Lifecycle]]
- [[./02 - Durabilidade, Versionamento e Replicacao]]
- [[./03 - Seguranca e Performance do S3]]
