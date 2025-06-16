## ⚙️ Используемые технологии

- **Go 1.24.3**
- **Gin** — HTTP-фреймворк
- **GORM** — ORM для PostgreSQL
- **MongoDB Go Driver**
- **RabbitMQ** — асинхронная коммуникация
- **Docker + Compose** — контейнеризация и запуск

## 📁 Структура репозитория

- `gateway/` — REST API-шлюз, маршруты `user` и `order`
- `order-svc/` — PostgreSQL, RabbitMQ consumer, бизнес-логика заказов
- `user-svc/` — MongoDB, регистрация/авторизация пользователей
- `payment-svc/` — принимает REST-запрос → публикует статус оплаты в очередь

## 🚀 Запуск

### 1. Установи зависимости и создать `.env` файл из примеров `.env-example` в каждом микросервисе

### 2. Запустить каждый микросервис через docker-compose.yml

```bash
docker-compose up --build
```

### 3. Примеры API

#### 📤 Создание пользователя

```bash
POST /users/create
Content-Type: application/json
{
  "email": "test@example.com",
  "password": "12345678",
  "name": "Иван"
}
```

#### 📦 Создание заказа
```bash
POST api/orders
Content-Type: application/json
{
  "userId": "684bcd0ce13ad1bc843b41cb",
  "productIds": [1, 2, 3]
}
```


#### Запрос на оплату
```bash
POST /payment/orders/{id}/status
Content-Type: application/json
{
  "status": "paid"
}
```
