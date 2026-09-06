---
tema: Pipelines CI/CD na AWS
tipo: unidade-estudo
tags: [aws, cicd, codepipeline, codebuild, codedeploy]
---
# 🧪 UE - Pipelines CI/CD

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Deploy manual é lento, propenso a erro humano e não-auditável — e é onde a maioria dos incidentes de produção nasce. Um pipeline de CI/CD automatiza build, teste, aprovação e deploy, tornando releases frequentes, pequenos e reversíveis. O desafio de arquitetura é desenhar o pipeline com gates de qualidade, segurança (secrets, least privilege do próprio pipeline) e artefatos versionados, sem criar um gargalo ou um vetor de ataque (o pipeline tem acesso amplo).

## 🧬 Grade Atômica de Tópicos
1. **Orquestração (CodePipeline):** Stages, actions, transitions, source/build/deploy, approval gates manuais, integração cross-account e com EventBridge.
2. **Build e teste (CodeBuild):** buildspec, ambientes/imagens, artefatos, cache, variáveis/segredos, testes automatizados e análise estática/segurança no pipeline.
3. **Deploy (CodeDeploy):** Deployments para EC2/ASG/Lambda/ECS, appspec, hooks de lifecycle, rollback automático em alarme.
4. **Artefatos e código (CodeArtifact/CodeCommit):** Repositório de pacotes privado, versionamento de artefatos, integração com terceiros (GitHub/GitLab) via source actions.

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS CodePipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html) e [AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/welcome.html).
- **System Blueprint:** Pipeline GitHub→CodeBuild (test+build imagem→ECR)→CodeDeploy (blue/green no ECS) com aprovação manual antes de produção.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Montar um pipeline end-to-end.
- [ ] Criar um pipeline: source (Git) → CodeBuild (testes) → deploy (Lambda/ECS).
- [ ] Adicionar um approval gate manual antes do stage de produção.
- [ ] Provocar uma falha de teste e confirmar que o deploy não avança.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - CodePipeline - Stages e Gates]]
- [[./02 - CodeBuild - Build, Teste e Seguranca]]
- [[./03 - CodeDeploy e CodeArtifact]]
