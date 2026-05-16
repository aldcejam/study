---
tema: Rowhammer e Ataques em Memoria RAM
tipo: unidade-estudo
tags: [segurança, hardware, memoria, baixo-nivel]
---
# 🧪 UE - Rowhammer e Ataques em Memoria RAM

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Você pode ter um software perfeito e um kernel blindado, mas se o seu hardware for fisicamente frágil, nada disso importa. O problema é a densidade das células de memória RAM moderna. **Rowhammer** resolve o "problema" de como alterar dados na memória sem ter permissão de escrita, simplesmente acessando (lendo) repetidamente as linhas vizinhas de memória. A interferência elétrica faz com que bits "virem" (bit flip) na linha alvo. Sem entender Rowhammer, você não entende como ataques de hardware podem contornar todas as proteções lógicas do sistema operacional.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Fisica da DRAM]:** Capacitores, transistores e o processo de Refresh.
2. **[Interferencia Eletromagnetica]:** Como a ativação constante de uma "linha agressora" vaza carga para as "linhas vitimas".
3. **[Bit Flipping]:** A mudança involuntária de um bit de 0 para 1 (ou vice-versa) e suas consequências lógicas.
4. **[Double-Sided Hammering]:** A técnica de atacar uma linha de memória de ambos os lados para maximizar a chance de erro.
5. **[Mitigacoes]:** ECC (Error Correction Code), TRR (Target Row Refresh) e aumento da taxa de refresh da RAM.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *Flipping Bits in Memory Without Accessing Them: An Experimental Study of DRAM Disturbance Errors* (Kim et al., ISCA 2014).
- **System Blueprint:** Exploração do Rowhammer para ganhar privilégios de root via **Drammer** (Android) ou através de JavaScript no browser.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Estudar o código de ferramentas de teste de Rowhammer (como o `rowhammer-test` do Google) e entender como elas utilizam instruções de `clflush` para garantir que as leituras atinjam a RAM e não fiquem apenas no cache.
- [ ] Analisar o loop de martelamento (hammer loop).
- [ ] Entender o mapeamento de endereços virtuais para endereços físicos de memória (DRAM addressing).
- [ ] Simular teoricamente como um bit flip em uma Page Table pode levar ao escalonamento de privilégio.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Arquitetura Interna da DRAM e Refresh Cycles - Teoria e Fundamentos]]
- [[./Mapeamento de Enderecos Fisicos e Canais de Memoria - Funcionamento Interno e Arquitetura]]
- [[./Confiabilidade de Hardware e o Limite de Miniaturizacao - Casos de Falha e Analise Amortizada]]
- [[./Mitigacoes de Hardware (ECC) e Software contra Rowhammer - Implementacao de Referencia e Benchmarks]]
