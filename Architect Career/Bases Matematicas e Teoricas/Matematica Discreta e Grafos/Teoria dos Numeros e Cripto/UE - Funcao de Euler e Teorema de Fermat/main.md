---
tema: Funcao de Euler e Teorema de Fermat
tipo: unidade-estudo
tags: [matematica, criptografia, teoria-dos-numeros]
---
# 🧪 UE - Funcao de Euler e Teorema de Fermat

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> A segurança da internet moderna repousa sobre a dificuldade de reverter certas operações matemáticas. O **Pequeno Teorema de Fermat** e a **Função Totiente de Euler** são os pilares que permitem a existência do RSA. Sem entender como essas propriedades permitem que um número seja elevado a uma potência e depois "revertido" via aritmética modular, um engenheiro nunca compreenderá como chaves públicas e privadas se relacionam matematicamente, tratando a criptografia como uma "caixa mágica" perigosa.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Aritmetica Modular]:** O conceito de congruência e operações em campos finitos $\mathbb{Z}_n$.
2. **[Pequeno Teorema de Fermat]:** A propriedade $a^{p-1} \equiv 1 \pmod{p}$ para $p$ primo e sua aplicação em testes de primaridade.
3. **[Funcao Totiente de Euler ($\phi$)]:** Contagem de números coprimos e a generalização do teorema de Fermat para $a^{\phi(n)} \equiv 1 \pmod{n}$.
4. **[Inverso Multiplicativo Modular]:** Uso do Algoritmo de Euclides Estendido para encontrar inversos, essencial para a geração de chaves.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *A Course in Number Theory and Cryptography* (Neal Koblitz).
- **System Blueprint:** Geração de parâmetros para RSA em bibliotecas como OpenSSL e Mbed TLS.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um pequeno sistema de cifragem RSA manual (usando números pequenos) para validar as propriedades de Euler.
- [ ] Implementar a função para calcular $\phi(n)$.
- [ ] Implementar o Algoritmo de Euclides Estendido para achar o expoente privado $d$.
- [ ] Cifrar e decifrar uma mensagem numérica curta para verificar a corretude.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Congruencia Modular e Propriedades - Teoria e Fundamentos]]
- [[./Algoritmo de Euclides Estendido e Inversos - Funcionamento Interno e Arquitetura]]
- [[./Ataques de Fatoracao de Primos - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de RSA Simplificado - Implementacao de Referencia e Benchmarks]]
