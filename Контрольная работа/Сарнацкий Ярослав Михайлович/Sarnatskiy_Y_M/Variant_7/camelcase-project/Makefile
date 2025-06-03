# Makefile для проекта
# Описание: Упрощает сборку, тестирование и запуск программы

all: build

build:
	go build -o main cmd/main.go

run: build
	./main

test:
	go test ./...

clear:
	rm -f main

deploy: build
	# Пример деплоя на сервер через SSH
	@echo "Deploying to server..."
	scp ./main <username>@<hostname>:<path/to/deploy/directory>
	@echo "Deployment complete."
