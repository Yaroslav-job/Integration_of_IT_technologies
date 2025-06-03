# Переменные
BINARY_NAME=goproject
BUILD_DIR=build

# Задача по умолчанию
all: test build

# Очистка временных файлов и директорий
clean:
	rm -rf $(BUILD_DIR)
	go clean
	rm -f $(BINARY_NAME)

# Сборка проекта
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Запуск проекта
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Запуск тестов
test:
	go test goproject/rotation
	go test goproject/calculator

# Развертывание (в данном случае просто копирование бинарного файла)
deploy: build
	cp $(BUILD_DIR)/$(BINARY_NAME) ./$(BINARY_NAME)

.PHONY: all clean build run test deploy
