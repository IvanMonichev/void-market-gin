# Архитектура

```
myapp/
├── cmd/
│   └── server/             # Точка входа (main.go)
├── internal/
│   ├── app/                # Инициализация зависимостей (роутеры, middleware, DI)
│   ├── handler/            # HTTP-хендлеры (Gin)
│   ├── service/            # Бизнес-логика
│   ├── repository/         # Доступ к данным (Postgres, Redis и т.д.)
│   ├── model/              # DTO/Entity/Request/Response
│   └── transport/          # HTTP-обёртки, сериализация/десериализация
├── config/                 # Конфигурации (YAML, ENV, dotenv)
├── pkg/                    # Переиспользуемые библиотеки (логгер, валидация, JWT и т.д.)
├── migrations/             # SQL миграции (если есть)
├── go.mod
└── README.md

```