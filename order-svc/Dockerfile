ARG GO_VERSION=1.24.3

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order-svc ./cmd/server/main.go

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/order-svc ./order-svc
COPY --from=builder /app/config/production.yml ./config/production.yml

ENV CONFIG_PATH=/app/config/production.yml

EXPOSE 4020

COPY .env /app/.env

ENTRYPOINT ["./order-svc"]
