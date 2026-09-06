---
tema: Backup e Transferencia de Dados
tipo: unidade-estudo
tags: [aws, backup, storage-gateway, datasync, transferencia]
---
# 🧪 UE - Backup e Transferencia de Dados

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Backup não é opcional, mas backup sem governança centralizada e sem teste de restauração é uma falsa sensação de segurança. Além disso, mover terabytes/petabytes para a nuvem pela internet pode levar semanas e custar caro — por isso existem serviços de transferência (online e físico). Ignorar isso resulta em RPO/RTO impossíveis de cumprir e migrações que estouram prazo.

## 🧬 Grade Atômica de Tópicos
1. **AWS Backup:** Backup centralizado e policy-based (backup plans) para múltiplos serviços, backup vault, cross-Region/cross-account copy, compliance (Vault Lock).
2. **Storage Gateway:** File Gateway (NFS/SMB↔S3), Volume Gateway (iSCSI cached/stored), Tape Gateway (VTL) para hibridizar storage on-premises.
3. **Transferência online (DataSync / Transfer Family):** DataSync para sync massivo agendado; Transfer Family (SFTP/FTPS/FTP) para integrações externas em S3/EFS.
4. **Transferência física (Snow Family):** Snowcone/Snowball Edge/Snowmobile para volumes onde a rede é inviável (ver categoria 11).

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Backup](https://docs.aws.amazon.com/aws-backup/latest/devguide/whatisbackup.html) e [AWS DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/what-is-datasync.html).
- **System Blueprint:** Estratégia de backup cross-Region com Vault Lock para atender compliance e ransomware resilience.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Automatizar backup e testar restauração.
- [ ] Criar um Backup Plan que faz snapshot diário de um volume EBS com cópia cross-Region.
- [ ] Restaurar o recurso e validar integridade.
- [ ] Desenhar quando usaria DataSync vs Snowball para 50 TB.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - AWS Backup e Vault Lock]]
- [[./02 - Storage Gateway]]
- [[./03 - DataSync e Transfer Family]]
