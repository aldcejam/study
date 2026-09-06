---
tema: Transferencia de Dados em Massa
tipo: unidade-estudo
tags: [aws, snowball, datasync, transfer-family, transferencia]
---
# 🧪 UE - Transferencia de Dados em Massa

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Mover 100 TB por um link de 1 Gbps leva semanas e pode custar mais em tempo/risco do que enviar discos físicos pelo correio. A matemática de "tempo de transferência vs banda vs volume" define quando usar transferência online (DataSync) vs física (Snow Family). Ignorar essa conta trava cronogramas de migração e explode o custo de data transfer. Conecta com a UE de [[../../04 - Armazenamento/UE - Backup e Transferencia de Dados/main|Backup e Transferencia de Dados]].

## 🧬 Grade Atômica de Tópicos
1. **A conta de transferência:** Cálculo tempo = volume/banda; ponto de cruzamento entre online e físico; custos de data transfer.
2. **Snow Family (física):** Snowcone (pequeno/edge), Snowball Edge (storage/compute optimized), Snowmobile (exabytes); segurança/criptografia e edge computing.
3. **Transferência online (DataSync):** Sync incremental agendado on-premises↔AWS e entre serviços AWS, verificação de integridade, aceleração.
4. **Integração contínua (Transfer Family):** SFTP/FTPS/FTP gerenciado para parceiros externos entregando em S3/EFS; comparação com DataSync.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Snow Family](https://docs.aws.amazon.com/snowball/) e [AWS DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/what-is-datasync.html).
- **System Blueprint:** Migração inicial de 200 TB via Snowball Edge + sincronização contínua do delta via DataSync até o cutover.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Decidir e executar transferência.
- [ ] Calcular o tempo de transferência de 50 TB a 500 Mbps e decidir online vs Snowball.
- [ ] Configurar uma tarefa DataSync entre dois locais (ex.: EFS→S3) com verificação.
- [ ] Subir um servidor SFTP via Transfer Family entregando em um bucket S3.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - A Conta - Online vs Fisico]]
- [[./02 - Snow Family]]
- [[./03 - DataSync e Transfer Family]]
