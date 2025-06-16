FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-svc cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/user-svc ./user-svc
COPY --from=builder /app/config/production.yml ./config/production.yml

EXPOSE 4010

COPY .env /app/.env

CMD ["./user-svc"]
