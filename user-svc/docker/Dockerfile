FROM golang:1.24.3 AS builder

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-svc ./cmd/server

FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/user-svc .

# Копируем конфигурацию (если нужно встроить её в образ)
COPY ../config/production.yaml ./production.yml

# Порт сервиса
EXPOSE 4010

# Указываем переменную окружения, которую будет читать config.MustLoad()
ENV CONFIG_PATH=/app/production.yml

# Запуск приложения
CMD ["./user-svc"]
