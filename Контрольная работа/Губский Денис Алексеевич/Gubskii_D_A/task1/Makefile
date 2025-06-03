# Makefile для проекта task1

# Цель по умолчанию
all: build

# Сборка проекта
build:
	go build -o task1 main.go

# Запуск проекта
run: build
	./task1

# Запуск тестов
test:
	go test ./...

# Очистка собранных файлов
clear:
	rm -f task1

# Развертывание (можно адаптировать под нужды)
deploy:
	scp task1 user@server:/path/to/deploy

.PHONY: all build run test clear deploy
