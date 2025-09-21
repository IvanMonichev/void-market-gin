FROM golang:1.24 AS builder


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway ./gateway
COPY --from=builder /app/config/production.yml ./production.yml
COPY --from=builder /app/.env /app/.env

EXPOSE 4010

ENV CONFIG_PATH=/app/production.yml

ENTRYPOINT ["./gateway"]
