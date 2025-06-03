# Makefile для проекта task2

# Цель по умолчанию
all: build

# Сборка проекта
build:
	go build -o task2 main.go

# Запуск проекта
run: build
	./task2

# Запуск тестов
test:
	go test ./...

# Очистка собранных файлов
clear:
	rm -f task2

# Развертывание (можно адаптировать под нужды)
deploy:
	scp task2 user@server:/path/to/deploy

.PHONY: all build run test clear deploy
