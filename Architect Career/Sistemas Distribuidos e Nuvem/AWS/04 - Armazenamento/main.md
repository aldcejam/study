---
tipo: index-categoria
contexto_pai: [[../main|Voltar]]
tags: [aws, armazenamento, s3, ebs, efs]
---
# 🗂️ 04 - Armazenamento

## 🎯 Escopo da Área
Armazenamento na AWS não é um serviço, é um espectro de trade-offs entre durabilidade, latência, throughput, custo e semântica de acesso (objeto vs bloco vs arquivo). O arquiteto escolhe entre S3 (objeto, "infinito", 11 noves de durabilidade), EBS (bloco, latência baixa, ligado a uma AZ) e EFS/FSx (arquivo compartilhado). Errar a classe ou o tipo gera custo alto, latência inaceitável ou perda de dados em falha de AZ.

## 🗺️ Mapa de Exploração
- [[./UE - S3 Profundo/main|UE - S3 Profundo]] -> Storage classes, lifecycle, consistência, segurança e performance.
- [[./UE - Armazenamento de Bloco e Arquivo/main|UE - Armazenamento de Bloco e Arquivo]] -> EBS, Instance Store, EFS, FSx.
- [[./UE - Backup e Transferencia de Dados/main|UE - Backup e Transferencia de Dados]] -> AWS Backup, Storage Gateway, DataSync.
