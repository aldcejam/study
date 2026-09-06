# Mentor de Obras - Escopo do Sistema

## Objetivo Principal
Garantir o cumprimento do cronograma estabelecido por meio de acompanhamento e realinhamento dinâmico.

## Objetivo Secundário
Possibilitar economias ao longo da obra (otimização de recursos e compras just-in-time).

---
### FASE 1: Ingestão de Dados e Geração Base
* **Importação Inteligente:** Upload de planilhas de orçamento e insumos (ex: Obra Prima). Classificação automática em materiais, mão de obra e serviços.
* **Mapeamento de Dependências:** Aplicação de regras lógicas da construção civil (ex: não iniciar alvenaria sem fundação; respeitar tempo de cura).
* **Geração do Cronograma:** Criação automática do cronograma preliminar (Gantt) baseado nos quantitativos importados e índices de produtividade padrão.

### FASE 2: Validação e Linha de Base (Baseline)
* **Edição Assistida:** Usuário ajusta durações, equipes e datas. O sistema emite alertas em tempo real se a edição gerar gargalos ou restrições de execução.
* **Aprovação Visual:** Exibição do cronograma consolidado e acompanhamento de progresso planejado.
* **Congelamento da Baseline:** O sistema salva esta versão aprovada como o "cenário ideal" (linha de base) para comparar com a execução real.

### FASE 3: Acompanhamento, Realinhamento e Instruções (Core do Mentor)
* **Input de Progresso:** Usuário informa o que foi concluído ou reporta atrasos (ex: chuva paralisou a obra).
* **Relatório de Diagnóstico:** Sistema calcula o impacto no caminho crítico da obra e mensura o reflexo na data de entrega final.
* **Instruções e Plano de Ação:** O mentor sugere cenários de mitigação:
  * *Opção A (Fast-tracking):* Sobrepor etapas que originalmente seriam sequenciais (se viável).
  * *Opção B (Crashing):* Injetar mais recursos (mão de obra/equipamento) em etapas futuras para recuperar o atraso.
* **Recalibragem:** Ao selecionar uma opção, o sistema reordena e reavalia todas as etapas, projetando a nova data de entrega.

---