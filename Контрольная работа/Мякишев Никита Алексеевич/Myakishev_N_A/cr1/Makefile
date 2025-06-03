# Название исполняемого файла
PROJECT_NAME := cr1

# Каталог с исходными файлами
SRC_DIR := src

# Каталог для исполняемого файла
BIN_DIR := bin

# Каталог с тестами (относительно SRC_DIR)
TEST_DIR := test

# Компилятор Go
GO := go

# Инструменты для форматирования и линтинга
GOFMT := $(GO) fmt
GOLINT := $(GO) vet
GOVET  := $(GO) vet
GOBUILD := $(GO) build
GOTEST  := $(GO) test

# Архитектура и ОС
GOARCH ?= $(shell $(GO) env GOARCH)
GOOS   ?= $(shell $(GO) env GOOS)
GO_BUILD_FLAGS := -ldflags="-s -w"

# Получить все Go файлы исходников
GO_FILES := $(wildcard $(SRC_DIR)/*.go $(SRC_DIR)/**/*.go)

# Получить все Go файлы тестов
TEST_GO_FILES := $(wildcard $(SRC_DIR)/$(TEST_DIR)/*_test.go $(SRC_DIR)/$(TEST_DIR)/**/*.go)

# Получить список объектных файлов
OBJ_FILES := $(patsubst $(SRC_DIR)/%.go,$(OBJ_DIR)/%.o,$(GO_FILES))

# Цели
.PHONY: all fmt lint build test clean run

# Правило для форматирования кода
fmt:
	$(GOFMT) $(GO_FILES)

# Правило для запуска линтера
lint:
	$(GOVET) $(GO_FILES)

# Правило для сборки приложения (без запуска)
build: $(BIN_DIR)/$(PROJECT_NAME)

# Правило для сборки исполняемого файла
$(BIN_DIR)/$(PROJECT_NAME): $(GO_FILES)
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(PROJECT_NAME) $(GO_FILES)

# Правило для запуска тестов (с учетом каталога тестов)
test:
	$(GOTEST) -v $(SRC_DIR)/$(TEST_DIR)

# Правило для запуска приложения
run: build
	$(BIN_DIR)/$(PROJECT_NAME)

# Правило для очистки (удаления исполняемого файла)
clean:
	rm -rf $(BIN_DIR)

# Цель по умолчанию
all: clean run