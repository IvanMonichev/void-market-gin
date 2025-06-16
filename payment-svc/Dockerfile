ARG GO_VERSION=1.24.3

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o payment-svc ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/payment-svc ./payment-svc
COPY --from=builder /app/.env ./.env
EXPOSE 4030

ENTRYPOINT ["./payment-svc"]
