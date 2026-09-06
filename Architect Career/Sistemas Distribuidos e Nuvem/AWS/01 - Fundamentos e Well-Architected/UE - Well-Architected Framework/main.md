---
tema: AWS Well-Architected Framework
tipo: unidade-estudo
tags: [aws, well-architected, arquitetura, governanca]
---
# 🧪 UE - Well-Architected Framework

## ⚖️ O Core Problem (Por que estudamos isso?)
> [!NOTE]
> Sem um framework de avaliação, "boa arquitetura" vira opinião. O Well-Architected dá uma linguagem comum e uma checklist de trade-offs para justificar decisões sob pressão de custo, prazo e risco. As provas Professional são, na prática, testes de aplicação dos 6 pilares em cenários ambíguos onde toda opção "funciona", mas apenas uma respeita o pilar priorizado no enunciado.

## 🧬 Grade Atômica de Tópicos
1. **Os 6 pilares:** Excelência Operacional, Segurança, Confiabilidade, Eficiência de Performance, Otimização de Custos e Sustentabilidade. O que cada um otimiza e como conflitam entre si.
2. **Trade-offs entre pilares:** Ex.: alta confiabilidade (multi-região) vs custo; performance (cache agressivo) vs consistência. Como o contexto de negócio define a prioridade.
3. **Processo de revisão (WAFR) e Trusted Advisor:** Como uma revisão Well-Architected é conduzida, o papel das perguntas e das lentes (lenses) específicas por domínio.
4. **Design Principles transversais:** Automação, elasticidade, testar em escala, evoluir arquitetura, mecanismos de "stop guessing capacity".

> [!TIP] Material aprofundado
> Explicação completa de todos os tópicos (baseada na doc oficial da AWS): [[2-Detalhamento|📚 Detalhamento Técnico]]

## 📜 Fronteira Acadêmica e Referências
- **Documentação oficial:** [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html) — leitura obrigatória dos 6 whitepapers de pilares.
- **System Blueprint:** Uso da ferramenta AWS Well-Architected Tool para revisar uma workload real.

## 🛠️ Sandbox Prática (Do Teórico ao Código)
**Objetivo do Protótipo:** Aplicar o framework a uma arquitetura própria.
- [ ] Desenhar uma arquitetura simples (web app 3 camadas) e submetê-la à Well-Architected Tool.
- [ ] Para cada risco apontado, classificar qual pilar ele afeta e propor mitigação.
- [ ] Documentar 3 trade-offs conscientes que você aceitou e por quê.

## 🗃️ Notas Heutagógicas Atômicas
- [[./01 - Os Seis Pilares em Profundidade]]
- [[./02 - Trade-offs e Priorizacao entre Pilares]]
- [[./03 - Processo WAFR e Ferramentas]]
