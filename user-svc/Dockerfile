FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-svc ./cmd/server

FROM alpine:latest

WORKDIR /app

# Копируем бинарник из builder-этапа
COPY --from=builder /app/user-svc .

EXPOSE 4010

# Запуск
CMD ["./user-svc"]
