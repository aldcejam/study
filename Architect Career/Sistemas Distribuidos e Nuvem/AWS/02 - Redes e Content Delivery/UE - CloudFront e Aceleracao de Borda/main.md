---
tema: CloudFront e Aceleracao de Borda
tipo: unidade-estudo
tags: [aws, cloudfront, cdn, edge, global-accelerator]
---
# 🧪 UE - CloudFront e Aceleracao de Borda

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Servir conteúdo a partir de uma única Region para usuários globais significa latência alta, custo de transferência elevado e origem sobrecarregada. A camada de borda (CDN e rede backbone da AWS) resolve isso cacheando perto do usuário e encurtando o caminho de rede. Errar TTLs, cache keys ou a escolha entre CloudFront (HTTP, cacheável) e Global Accelerator (TCP/UDP, não-cacheável) degrada performance e segurança.

## 🧬 Grade Atômica de Tópicos
1. **CloudFront (CDN):** Distribuições, origens (S3/ALB/custom), cache behaviors, TTLs, cache key/policies, invalidações, OAC (Origin Access Control) para proteger buckets.
2. **Computação na borda:** CloudFront Functions (leve, viewer) vs Lambda@Edge (mais pesado, 4 gatilhos) para personalização, auth e reescrita de requisições.
3. **Global Accelerator:** IPs anycast estáticos, roteamento pelo backbone AWS para tráfego TCP/UDP não-cacheável, failover rápido entre Regions.
4. **Segurança na borda:** Integração com AWS WAF, Shield (DDoS), Signed URLs/Cookies e geo-restriction.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [Amazon CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html) e [Global Accelerator](https://docs.aws.amazon.com/global-accelerator/latest/dg/what-is-global-accelerator.html).
- **System Blueprint:** Site estático S3 + CloudFront + OAC + WAF; API global com Global Accelerator + ALB multi-região.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Publicar e proteger conteúdo na borda.
- [ ] Servir um site estático de um bucket S3 privado via CloudFront + OAC.
- [ ] Adicionar uma CloudFront Function para redirecionar/headers e medir a redução de latência por Edge.
- [ ] Comparar quando usaria Global Accelerator vs CloudFront para uma API.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CloudFront - Cache e Origens]]
- [[./02 - Computacao na Borda - Functions e Lambda@Edge]]
- [[./03 - Global Accelerator e Seguranca de Borda]]
