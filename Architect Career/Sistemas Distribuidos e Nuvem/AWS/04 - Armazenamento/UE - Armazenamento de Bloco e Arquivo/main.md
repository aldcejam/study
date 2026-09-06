---
tema: Armazenamento de Bloco e Arquivo - EBS, EFS, FSx
tipo: unidade-estudo
tags: [aws, ebs, efs, fsx, armazenamento]
---
# 🧪 UE - Armazenamento de Bloco e Arquivo

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Nem todo dado cabe em armazenamento de objeto: bancos de dados precisam de bloco de baixa latência (EBS), e workloads que compartilham arquivos entre muitas instâncias precisam de sistema de arquivos (EFS/FSx). Confundir os modelos leva a decisões ruins — como tentar compartilhar um volume EBS (impossível entre AZs no modelo clássico) ou usar EFS onde a latência de bloco era necessária, matando a performance do banco.

## 🧬 Grade Atômica de Tópicos
1. **EBS (bloco):** Tipos (gp3, io2 Block Express, st1, sc1), IOPS/throughput, ligado a uma AZ, snapshots (incrementais no S3), volumes multi-attach (io2).
2. **Instance Store (efêmero):** Storage local NVMe, altíssima performance, perde dados ao parar a instância; casos de uso (cache, scratch).
3. **EFS (NFS gerenciado):** Sistema de arquivos elástico multi-AZ, performance/throughput modes, storage classes (IA), montagem por muitas instâncias simultâneas.
4. **FSx (arquivo especializado):** FSx for Windows (SMB), Lustre (HPC/ML), NetApp ONTAP, OpenZFS; quando cada um é obrigatório.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon EBS](https://docs.aws.amazon.com/ebs/latest/userguide/what-is-ebs.html), [Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/whatisefs.html), [Amazon FSx](https://docs.aws.amazon.com/fsx/).
- **System Blueprint:** Banco de dados em EBS io2 multi-AZ via replicação; cluster de app compartilhando assets via EFS.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Comparar bloco vs arquivo na prática.
- [ ] Anexar um volume gp3 a uma EC2, formatar, medir IOPS com `fio`; tirar e restaurar snapshot.
- [ ] Montar um sistema EFS em 2 instâncias em AZs diferentes e escrever de ambas.
- [ ] Comparar latência de escrita EBS vs EFS.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - EBS - Tipos, IOPS e Snapshots]]
- [[./02 - Instance Store e Efemeridade]]
- [[./03 - EFS e FSx - Arquivo Compartilhado]]
