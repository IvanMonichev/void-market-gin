FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gateway ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway .

COPY config/production.yml ./config.yml
COPY .env .env

EXPOSE 4000

ENV CONFIG_PATH=/app/config.yml

ENTRYPOINT ["./gateway"]
