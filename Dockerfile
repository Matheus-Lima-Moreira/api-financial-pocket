# Stage 1: Builder
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Copia arquivos de dependências primeiro (cache layer)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copia o código fonte
COPY . .

# Build do binário com flags de otimização
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main cmd/api/main.go

# Stage 2: Runtime (imagem leve)
FROM alpine:latest

# Instala apenas o necessário para executar o binário e healthcheck
RUN apk --no-cache add ca-certificates tzdata wget

WORKDIR /app

# Copia o binário do stage builder
COPY --from=builder /app/main .

# Variáveis de ambiente padrão para produção
ENV ENV=production
ENV PORT=8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/health || exit 1

# Executa o binário
CMD ["./main"]