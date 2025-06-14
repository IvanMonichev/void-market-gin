FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order-svc ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order-svc .

# Копируем конфиг
COPY config/production.yml ./production.yml

EXPOSE 4020

ENV CONFIG_PATH=/app/production.yml

CMD ["./order-svc"]
