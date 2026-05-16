---
tema: Logaritmo Discreto e Troca de Chaves Diffie-Hellman
tipo: unidade-estudo
tags: [matematica, criptografia, redes, segurança]
---
# 🧪 UE - Logaritmo Discreto e Troca de Chaves Diffie-Hellman

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Como duas pessoas podem concordar em um segredo compartilhado através de um canal público e inseguro (como a Internet) sem que ninguém mais o descubra? O problema do **Logaritmo Discreto** resolve isso. Sem o protocolo **Diffie-Hellman**, não haveria HTTPS/TLS seguro, e cada conexão exigiria uma troca física de chaves. Ignorar a segurança dos parâmetros de DH (como o uso de primos fracos) permite ataques de interceptação que quebram a privacidade de milhões de usuários.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[Grupos Ciclicos e Geradores]:** Entender como um gerador $g$ percorre todos os elementos de um campo finito.
2. **[A Exponenciacao Modular]:** Por que calcular $g^x \pmod{p}$ é fácil, mas encontrar $x$ dado o resultado (Logaritmo Discreto) é computacionalmente inviável.
3. **[O Protocolo Diffie-Hellman]:** O fluxo de troca: Alice envia $g^a$, Bob envia $g^b$, ambos chegam a $(g^b)^a = (g^a)^b = g^{ab}$.
4. **[Forward Secrecy]:** Por que o uso de chaves efêmeras (DHE) é superior à troca de chaves estática para proteção contra vazamentos futuros.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers clássicos):
- **Paper Clássico/Livro:** *New Directions in Cryptography* (Whitfield Diffie and Martin Hellman, 1976).
- **System Blueprint:** Implementação do TLS 1.3, que obriga o uso de trocas de chaves efêmeras baseadas em logaritmo discreto.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Simular um handshake Diffie-Hellman entre dois scripts independentes usando sockets e validar a geração do segredo compartilhado.
- [ ] Implementar a exponenciação modular rápida (Binary Exponentiation).
- [ ] Gerar parâmetros seguros (primos de Sophie Germain).
- [ ] Realizar a troca de mensagens e comparar as chaves finais derivadas (HKDF).

## 🗃️ Notas Heutagógicas Atômicas
*(Links para os arquivos de estudo que serão populados individualmente)*
- [[./Grupos de Corpos Finitos e Geradores - Teoria e Fundamentos]]
- [[./Exponenciacao Rapida e Logaritmo Discreto - Funcionamento Interno e Arquitetura]]
- [[./Ataque Man-in-the-Middle e Primos Fracos - Casos de Falha e Analise Amortizada]]
- [[./Implementacao de Handshake TLS Simplificado - Implementacao de Referencia e Benchmarks]]
