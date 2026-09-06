O **Config Server** em conjunto com o `/refresh` é extremamente útil no dia a dia, principalmente em ambientes de produção onde reiniciar a aplicação pode gerar indisponibilidade.

Aqui está uma lista dos contextos onde o **Config Server + `/refresh` dinâmico** é ideal, versus onde você ainda **precisará reiniciar**:

### 1. Quando usar o `/refresh` (Funciona 100% dinâmico em tempo de execução)

Você pode alterar e aplicar dinamicamente sem reiniciar em casos como:

- **Chaves de API e Tokens de Integração**: Alterar chaves de terceiros (ex: atualizar o token `${DEEPSEEK_API_KEY}` se o atual expirar).
- **Níveis de Log (Logging Levels)**: Mudar o log de `INFO` para `DEBUG` para depurar um problema em produção temporariamente, e depois voltar para `INFO`.
- **Feature Toggles / Flags**: Ativar ou desativar funcionalidades específicas no sistema (ex: `feature.new-analysis-flow.enabled: true/false`).
- **Regras de Negócio e Parâmetros**: Configurar limites, tempos de timeout, taxas de cálculo, ou valores padrão que mudam frequentemente.
- **Prompts de IA**: Ajustar as instruções enviadas para o modelo LLM do Spring AI no arquivo de configuração sem precisar recompilar ou reiniciar o microsserviço.

---

### 2. Quando você DEVE reiniciar o microsserviço

Existem configurações de infraestrutura básica que o Spring Boot inicializa apenas **uma vez** durante o boot da aplicação e que não suportam atualização em tempo de execução nativamente:

- **Dados de Conexão com Banco de Dados**: `url`, `username`, `password` (pois o pool de conexões já foi aberto no início).
- **Porta do Servidor**: Alterar o `server.port` (o Tomcat já está ouvindo na porta antiga; ele não vai fechar e abrir uma porta nova dinamicamente).
- **Configurações do JPA / Hibernate**: Alterar o `ddl-auto` (ex: de `update` para `validate`).
- **Perfis Ativos (Profiles)**: Mudar o profile ativo de `local` para `dev` ou `prod` (os Beans criados dependem do profile inicial).
- **Configurações fundamentais do Eureka/RabbitMQ**: Endereços dos servidores de mensageria e discovery.