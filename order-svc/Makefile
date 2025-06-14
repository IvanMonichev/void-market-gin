# Название бинарника
BINARY_NAME=order-svc

# Путь до main.go
MAIN_PATH=./cmd/server

# Конфигурация
CONFIG_PATH=config/production.yml

# Targets

.PHONY: run build clean docker-up docker-down

## Запуск с учетом .env и config
run:
	@echo "Running $(BINARY_NAME)..."
	@source .env && CONFIG_PATH=$(CONFIG_PATH) go run $(MAIN_PATH)

## Сборка бинарника
build:
	@echo "Building binary $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)

## Очистка бинарника
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

## Запуск docker-compose
docker-up:
	docker-compose up --build

## Остановка docker-compose
docker-down:
	docker-compose down
