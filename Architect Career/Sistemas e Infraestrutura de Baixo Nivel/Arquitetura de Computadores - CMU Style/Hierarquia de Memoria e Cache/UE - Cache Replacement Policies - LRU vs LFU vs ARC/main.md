---
tema: Cache Replacement Policies - LRU vs LFU vs ARC
tipo: unidade-estudo
tags: [arquitetura, performance, algoritmos, cache]
---
# 🧪 UE - Cache Replacement Policies - LRU vs LFU vs ARC

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> O espaço na cache é extremamente caro e limitado. Quando a cache enche, qual dado devemos jogar fora? A escolha errada aumenta drasticamente o número de acessos à memória lenta (RAM/Disco). **LRU (Least Recently Used)** é o padrão, mas falha miseravelmente em scans de dados. **LFU (Least Frequently Used)** ignora a recência. O **ARC (Adaptive Replacement Cache)** resolve isso equilibrando dinamicamente recência e frequência. Sem dominar essas políticas, você não consegue otimizar sistemas de storage, bancos de dados ou proxies web.

## 🧬 Grade Atômica de Tópicos
Para dominar esta UE, é obrigatório esgotar os seguintes sub-conceitos fundamentais:
1. **[LRU e a Teoria da Recencia]:** Por que dados usados recentemente tendem a ser usados novamente (Localidade Temporal).
2. **[LFU e o Contador de Frequencia]:** A vantagem de manter dados "populares" mesmo que antigos, e o problema da obsolescência de frequência.
3. **[O Algoritmo ARC]:** Como usar duas listas (recência e frequência) e uma "janela de adaptação" para se ajustar à carga de trabalho em tempo real.
4. **[Aproximacoes de Hardware (Pseudo-LRU)]:** Como CPUs implementam LRU com poucos bits por linha de cache.
5. **[Políticas Resistentes a Scans]:** Por que um "SELECT * FROM table" não deve destruir a cache do banco de dados.

## 📜 Fronteira Acadêmica e Referências
Diretrizes de leitura e fundamentação (Padrão MIT/Stanford/CMU/Papers classics):
- **Paper Clássico/Livro:** *ARC: A Self-Tuning, Low Overhead Replacement Cache* (Megiddo and Modha, 2003).
- **System Blueprint:** Implementação de cache no ZFS (que usa ARC) e no Redis.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Implementar um simulador de cache em Python que recebe uma sequência de acessos e compara a taxa de acerto (Hit Rate) entre LRU e ARC.
- [ ] Criar a estrutura de dados para LRU (OrderedDict ou Double Linked List).
- [ ] Implementar a lógica de adaptação do ARC.
- [ ] Rodar um benchmark com um padrão de "loop" (onde LRU falha se a cache for menor que o loop) e observar o comportamento do ARC.

## 🗃️ Notas Heutagógicas Atômicas
- [[./Localidade de Referencia e Hit Rates - Teoria e Fundamentos]]
- [[./Arquitetura do Algoritmo ARC e Auto-Tuning - Funcionamento Interno e Arquitetura]]
- [[./O Problema de Cache Pollution por Scans - Casos de Falha e Analise Amortizada]]
- [[./Benchmarks de Politicas de Cache - Implementacao de Referencia e Benchmarks]]
