---
tema: Chaos Engineering - Injecao de Falhas em Producao
tipo: unidade-estudo
tags: [resiliência, distribuídos, srer, qualidade, infraestrutura]
---
# 🧪 UE - Chaos Engineering - Injecao de Falhas em Producao

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Em sistemas distribuídos complexos, falhas são inevitáveis. O problema é que não sabemos como o sistema reagirá a falhas raras ou combinadas (ex: um switch de rede lento + um disco cheio). **Chaos Engineering** resolve isso injetando falhas propositais e controladas no ambiente (idealmente produção) para observar se os mecanismos de resiliência (circuit breakers, retries, fallbacks) funcionam. Sem o caos, sua primeira experiência com uma falha sistêmica será durante um incidente real com impacto financeiro.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Hipótese de Estado Estável]:** Como definir métricas de saúde (throughput, erro, latência) antes de começar o experimento.
2. **[Raio de Explosao (Blast Radius)]:** Como limitar o impacto do experimento para que ele não derrube o sistema inteiro.
3. **[Tipos de Experimentos]:** Latência de rede, queda de instâncias, corrupção de dados e picos de tráfego.
4. **[Game Days]:** A prática de reunir a equipe para simular e resolver falhas em tempo real.
5. **[Automacao do Caos]:** O uso de ferramentas para rodar experimentos de forma contínua.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Chaos Engineering* (Casey Rosenthal and Nora Jones).
- **System Blueprint:** O **Chaos Monkey** da Netflix e o **Gremlin**.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Usar uma ferramenta como `Chaos Mesh` ou `Litmus` em um cluster Kubernetes local para simular a queda de um pod de banco de dados e verificar se a aplicação reconecta automaticamente.
- [ ] Definir a métrica de sucesso (disponibilidade da API).
- [ ] Agendar um experimento de "Pod Kill".
- [ ] Analisar os logs da aplicação para ver a detecção e recuperação da falha.

## 🗃️ Notas Heutagógicas Atômicas
- [[./A Teoria de Sistemas Complexos e Falhas em Cascata - Teoria e Fundamentos]]
- [[./Design de Experimentos e Monitoramento de Blast Radius - Funcionamento Interno e Arquitetura]]
- [[./Quando o Experimento de Caos Causa um Incidente Real - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Chaos Mesh em Pipeline de Teste - Implementacao de Referencia e Benchmarks]]
