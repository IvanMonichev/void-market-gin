# Void Market Gin

```
void-market-gin/
├── README.md
├── go.work                     # Go workspace file для объединения всех сервисов
├── build/                      # docker-compose, Makefile и окружение
│   ├── docker-compose.yml
│   └── env/
│       ├── user.env
│       ├── order.env
│       └── payment.env
├── deployments/                # (опционально) манифесты для Kubernetes
│   └── k8s/
├── scripts/                    # миграции и утилиты
│   └── migrate-user-db.sql
├── shared/                     # общие библиотеки (опционально)
│   ├── logger/
│   ├── config/
│   └── utils/
├── api/                        # OpenAPI / protobuf контракты
│   ├── openapi.yaml
│   └── user.proto
├── gateway/                    # API Gateway на Gin
│   ├── cmd/
│   │   └── gateway/
│   │       └── main.go
│   ├── internal/
│   │   ├── delivery/http/
│   │   ├── config/
│   │   └── middleware/
│   ├── go.mod
│   └── go.sum
├── user-svc/                   # Сервис пользователей
│   ├── cmd/user-svc/main.go
│   ├── internal/
│   │   ├── app/
│   │   ├── delivery/http/
│   │   ├── domain/
│   │   ├── infra/db/
│   │   └── config/
│   ├── go.mod
│   └── go.sum
├── order-svc/                  # Сервис заказов
│   ├── cmd/order-svc/main.go
│   ├── internal/
│   │   ├── app/
│   │   ├── delivery/http/
│   │   ├── domain/
│   │   ├── infra/db/
│   │   └── infra/kafka/
│   ├── go.mod
│   └── go.sum
└── payment-svc/                # Сервис платежей
    ├── cmd/payment-svc/main.go
    ├── internal/
    │   ├── app/
    │   ├── delivery/consumer/
    │   ├── domain/
    │   ├── infra/db/
    │   └── infra/kafka/
    ├── go.mod
    └── go.sum

```