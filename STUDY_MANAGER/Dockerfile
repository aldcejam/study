# Estágio de Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instala git e ssh-client para clonar via SSH
RUN apk add --no-cache git openssh-client

# Adiciona o GitHub aos hosts conhecidos para evitar confirmação manual
RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# Clona o repositório usando o SSH agent do host (seguro, sem expor chaves)
RUN --mount=type=ssh git clone git@github.com:aldcejam/study.git .

# Muda para o diretório do gerenciador para compilar
WORKDIR /app/STUDY_MANAGER

# Download das dependências (usando o go.mod clonado)
RUN go mod download

# Compila todos os binários necessários
RUN mkdir -p bin && \
    go build -o bin/scanner ./src/processes/scanner && \
    go build -o bin/summaryNotification ./src/processes/summaryNotification && \
    go build -o bin/notifier ./src/processes/notifier && \
    go build -o bin/api ./src/api/main

# Estágio Final
FROM alpine:latest
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copia os binários compilados do estágio de build
COPY --from=builder /app/STUDY_MANAGER/bin /app/bin
# Copia o código dos pipelines (necessário para a API invocar)
COPY --from=builder /app/STUDY_MANAGER/src/pipelines /app/src/pipelines

# Variáveis de ambiente padrão
ENV PORT=8443
ENV VAULT_PATH=/app/notes
ENV DB_PATH=/app/output/clover_db
ENV SKIP_BUILD=true

EXPOSE 8443

# Inicia a API
CMD ["./bin/api"]
