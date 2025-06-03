# Makefile для проекта

# Название исполняемого файла
BINARY_NAME_TASK1=second_largest
BINARY_NAME_TASK2=fibonacci

# Команда для сборки
build:
	go build -o $(BINARY_NAME_TASK1) ./TASK1/second_largest.go
	go build -o $(BINARY_NAME_TASK2) ./TASK2/fibonacci.go

# Команда для запуска
run: build
	./$(BINARY_NAME_TASK1)
	./$(BINARY_NAME_TASK2)

# Команда для тестирования
test:
	go test ./TASK1
	go test ./TASK2

# Команда для очистки
clean:
	rm -f $(BINARY_NAME_TASK1) $(BINARY_NAME_TASK2)

# Команда для развертывания (можно настроить по необходимости)
deploy: build
	@echo "Deploying the application..."
