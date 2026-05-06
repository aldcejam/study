# 📚 Guia de Metadados de Estudo

Para que o sistema consiga extrair as datas de revisão e te notificar corretamente, você deve usar o formato **YAML Frontmatter** no topo das suas notas Markdown.

---

## 🚀 Formato Suportado

O sistema utiliza propriedades dinâmicas nativas do Obsidian. Cada data é uma chave separada.

```yaml
---
tema: Arquitetura Client-Server
subtema: Orquestração de Container
revision_18-04-2026: false
revision_20-04-2026: false
references:
  - "tal_coisa: link"
  - "tal_coisa_la:relative_path"
---
```

*   `true` -> Revisão concluída.
*   `false` -> Revisão pendente.
*   O prefixo `revision_` é obrigatório, seguido da data no formato `DD-MM-YYYY`.

---

## 🛠️ Campos Disponíveis

| Campo | Descrição | Obrigatório |
| :--- | :--- | :--- |
| `tema` | O nome do assunto que aparecerá na notificação. | Não (usa o nome do arquivo se vazio) |
| `revision_DD-MM-YYYY` | Data e status da revisão. | **Sim** (pelo menos uma revisão) |
| `subtema` | Contexto adicional para o estudo. | Não |
| `references` | Links ou descrições de materiais de apoio. | Não |

---

## 💡 Dicas Importantes
*   **Formatos de Data:** Utilize o formato `DD-MM-YYYY` após o prefixo `revision_`.
*   **Notificações Inteligentes:** O sistema só te avisará se a data for **HOJE** ou estiver **ATRASADA**.
*   **Cache:** Se você não fizer a revisão, ele te lembrará novamente a cada **2 dias** para não poluir seu WhatsApp. Se você editar o arquivo, o alerta é enviado imediatamente na próxima execução!
