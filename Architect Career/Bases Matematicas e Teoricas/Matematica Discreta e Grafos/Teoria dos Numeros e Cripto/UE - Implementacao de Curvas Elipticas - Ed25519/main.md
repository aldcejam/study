---
tema: Implementacao de Curvas Elipticas - Ed25519
tipo: unidade-estudo
tags: [matematica, criptografia, segurança, baixo-nivel]
---
# 🧪 UE - Implementacao de Curvas Elipticas - Ed25519

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> RSA exige chaves gigantescas (3072+ bits) para ser seguro hoje, o que é ineficiente para IoT e conexões rápidas. A **Criptografia de Curvas Elípticas (ECC)** oferece a mesma segurança com chaves muito menores (256 bits). A curva **Ed25519** resolve especificamente o problema de performance e resistência a ataques de canal lateral (side-channel). Ignorar a matemática das curvas elípticas ou usar implementações inseguras (que não são *constant-time*) permite que atacantes extraiam sua chave privada apenas medindo o tempo de execução ou o consumo de energia.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-concepts fundamentais:
1. **[Grupos de Curvas Elipticas]:** A lei de adição de pontos e por que ela forma um grupo abeliano.
2. **[Coordenadas Projetivas e Edwards]:** Como evitar divisões (caras computacionalmente) usando formas alternativas da equação da curva.
3. **[Multiplicacao de Escalar em Tempo Constante]:** Técnicas como o algoritmo de Montgomery para evitar vazamentos de informação via tempo.
4. **[O Esquema de Assinatura EdDSA]:** Como o Ed25519 gera assinaturas determinísticas e por que isso é mais seguro que o ECDSA tradicional.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *High-speed high-security signatures* (Daniel J. Bernstein et al., 2012).
- **System Blueprint:** Uso de Ed25519 no protocolo SSH, no Signal Messenger e no ecossistema de criptomoedas (Solana, Cardano).

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar a multiplicação escalar básica em uma curva de Edwards simplificada e verificar se a propriedade associativa se mantém.
- [ ] Definir a equação da curva e os parâmetros de campo.
- [ ] Implementar a soma de dois pontos e a duplicação de ponto.
- [ ] Criar uma função de multiplicação escalar e validar contra uma biblioteca de referência (como a `libsodium`).

## 🗃️ Notas Heutagógicas Atômicas
- [[./Aritmetica de Curvas e Lei de Adicao - Teoria e Fundamentos]]
- [[./Coordenadas de Edwards e Otimizacao de Performance - Funcionamento Interno e Arquitetura]]
- [[./Ataques de Side-Channel em ECC - Casos de Falha e Analise Amortizada]]
- [[./Geracao de Chaves e Assinaturas Ed25519 - Implementacao de Referencia e Benchmarks]]
