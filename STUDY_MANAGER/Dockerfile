# Estágio de Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instala git para o build (se necessário)
RUN apk add --no-cache git

# Copia e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte
COPY . .

# Compila a API principal
RUN go build -o bin/api ./src/api/main

# Compila todos os estágios do pipeline
RUN go build -o bin/scanner ./src/processes/scanner
RUN go build -o bin/summaryNotification ./src/processes/summaryNotification
RUN go build -o bin/notifier ./src/processes/notifier
RUN go build -o bin/pipeline ./src/pipelines/notifySummary

# Estágio Final (Imutável e Leve)
FROM alpine:latest

WORKDIR /app

# Instala apenas o essencial para o runtime, mais nodejs e npm para a gemini CLI
RUN apk add --no-cache git tzdata ca-certificates nodejs npm ripgrep && \
    npm install -g @google/gemini-cli

# Copia os binários compilados do estágio anterior
COPY --from=builder /app/bin /app/bin

# Variáveis de ambiente padrão
ENV PORT=8443
ENV VAULT_PATH=/app/notes
ENV DB_PATH=/app/output/clover_db

EXPOSE 8443

# Roda o binário da API
CMD ["./bin/api"]
