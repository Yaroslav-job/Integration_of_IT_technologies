# Makefile — команды для сборки и тестов golang-assignment

# Output директория для сборки и деплоя
BIN_DIR := bin
DEPLOY_DIR := dist

# Entry points
CMDS := cmd/camelcase cmd/network

.PHONY: all clear build run test deploy

all: clear build test

clear:
	@echo "🧹 Очистка старых бинарников..."
	@rm -rf $(BIN_DIR) $(DEPLOY_DIR)

build:
	@echo "🔨 Сборка всех приложений..."
	@mkdir -p $(BIN_DIR)
	@for dir in $(CMDS); do \
		name=$$(basename $$dir); \
		go build -o $(BIN_DIR)/$$name $$dir/main.go || exit 1; \
		echo "✔️ Собрано: $$name"; \
	done

run: build
	@echo "🚀 Запуск camelcase:"
	@./$(BIN_DIR)/camelcase
	@echo "🚀 Запуск network:"
	@./$(BIN_DIR)/network

test:
	@echo "🧪 Запуск unit-тестов..."
	@go test ./... -v

deploy: build
	@echo "📦 Деплой бинарников..."
	@mkdir -p $(DEPLOY_DIR)
	@cp $(BIN_DIR)/* $(DEPLOY_DIR)/
	@echo "✅ Бинарники скопированы в $(DEPLOY_DIR)/"
