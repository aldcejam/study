---
tema: Criptografia Pos-Quantica - Kyber e Dilithium
tipo: unidade-estudo
tags: [criptografia, matematica, segurança, futuro]
---
# 🧪 UE - Criptografia Pos-Quantica - Kyber e Dilithium

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Computadores quânticos em larga escala conseguirão quebrar quase toda a criptografia atual (RSA, ECC, Diffie-Hellman) em segundos usando o Algoritmo de Shor. O problema "Harvest now, Decrypt later" significa que dados interceptados hoje podem ser lidos no futuro. A **Criptografia Pós-Quântica (PQC)** resolve isso usando problemas matemáticos que são difíceis até para computadores quânticos (como o aprendizado com erros em redes latice). Ignorar a migração para PQC coloca em risco a longevidade da privacidade de dados críticos.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Problemas em Redes (Lattices)]:** O problema do vetor mais curto (SVP) e o aprendizado com erros (LWE).
2. **[CRYSTALS-Kyber]:** O mecanismo de encapsulamento de chaves (KEM) escolhido pelo NIST como padrão para troca de chaves.
3. **[CRYSTALS-Dilithium]:** O esquema de assinatura digital baseado em redes latice.
4. **[Trade-offs de Performance e Tamanho]:** Por que chaves PQC são significativamente maiores que as de ECC e como isso afeta protocolos de rede.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *NIST Post-Quantum Cryptography Standardization Process* (NIST IR 8413).
- **System Blueprint:** Implementação de suporte PQC no Google Chrome e Cloudflare (handshakes híbridos).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Utilizar a biblioteca `liboqs` (Open Quantum Safe) para realizar uma troca de chaves Kyber entre dois processos e comparar o tamanho das chaves com o ECDSA.
- [ ] Configurar o ambiente com a biblioteca OQS.
- [ ] Gerar um par de chaves Kyber-512.
- [ ] Encapsular e desencapsular um segredo e medir o tempo de CPU gasto.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Matematica de Lattices e Problema LWE - Teoria e Fundamentos]]
- [[./Arquitetura do Mecanismo Kyber KEM - Funcionamento Interno e Arquitetura]]
- [[./Impacto do Tamanho de Chave em Redes e MTU - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Handshake Hibrido - Implementacao de Referencia e Benchmarks]]
