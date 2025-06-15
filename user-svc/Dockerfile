FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-svc ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-svc ./user-svc

COPY config/production.yml ./config/production.yml

EXPOSE 4010

COPY .env /app/.env

CMD ["./user-svc"]
