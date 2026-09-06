---
tema: OpenTelemetry API vs SDK e OTLP
_revision_14-06-2026: false
_revision_15-06-2026: false
_revision_20-06-2026: false
_revision_13-07-2026: false
_revision_11-09-2026: false
_revision_10-11-2026: false
references:
  - https://opentelemetry.io/docs/specs/otel/
homework:
  - [[homeworks/OpenTelemetry API vs SDK e OTLP-pratica]]
---
## 🧠 Processo de Aprendizado
---

## 📝 Meu Resumo (Feynman)

Explicar OpenTelemetry API vs SDK e OTLP em linguagem simples, como se ensinasse a alguém:

---
## 🎤 Explicação em Voz

Checklist:
* [ ] Gravar explicação (2–5 minutos)
* [ ] Identificar lacunas
* [ ] Atualizar resumo

---

## ❓ Teste-se (Active Recall)

Responda sem consultar o material:

* Qual a principal diferença arquitetural e de acoplamento de dependências ao usar a API versus o SDK do OpenTelemetry no código da sua aplicação?
* O que acontece com as chamadas de instrumentação da API do OpenTelemetry se nenhuma implementação do SDK estiver registrada em runtime (comportamento no-op)?
* Como o protocolo OTLP otimiza o uso de rede e CPU em comparação com protocolos baseados em texto/JSON (ex: gRPC stream vs HTTP payload)?
* O que são Resources no OpenTelemetry e por que eles devem ser associados ao SDK no momento de inicialização e não na API?

---
## 🔎 Perguntas de Elaboração

* Por que a especificação do OpenTelemetry fez uma separação tão rígida entre a API de instrumentação e o SDK de implementação?
* Sob quais condições de arquitetura de software você escolheria exportar métricas diretamente via HTTP/JSON versus gRPC no protocolo OTLP?
* Como a escolha entre amostragem head-based (no SDK) e tail-based (no Collector) impacta o tráfego de rede gerado pelo SDK de instrumentação da aplicação?

---

## 🔗 Conexões

Relacionar com outros tópicos:
* [[../Observabilidade Plena e DevOps/main]]
* [[../../Resiliencia e Confiabilidade/main]]
* [[../../Arquitetura de Alta Escala/main]]

---
## 📅 Protocolo de Revisão

Durante cada revisão:

1. Cubra o material
2. Responda as perguntas sem olhar
3. Explique em voz alta
4. Revise o resumo
5. Marque a revisão como `done` no Frontmatter
