# Makefile для Worker Pool (Windows-версия)
SHELL = cmd

# Настройки проекта
PROJECT_NAME = worker-pool
BIN_DIR = bin
MODULE_PATH = github.com/Alina0101/worker-pool

# Директории
CMD_DIR = cmd
INTERNAL_DIR = internal
PKG_DIR = pkg

# Команды
GO = go
MKDIR = if not exist "$(BIN_DIR)" mkdir "$(BIN_DIR)"
RMDIR = if exist "$(BIN_DIR)" rmdir /s /q "$(BIN_DIR)"

# Флаги
GO_BUILD_FLAGS = -ldflags="-s -w"
GO_TEST_FLAGS = -v

.PHONY: all fmt vet build test clean run

all: build

fmt:
	@echo [1/4] Formatting code...
	@$(GO) fmt ./$(CMD_DIR)/... ./$(ALGO_DIR)/... ./$(TESTS_DIR)/...

vet:
	@echo [2/4] Code analysis...
	@$(GO) vet ./$(CMD_DIR)/... ./$(ALGO_DIR)/... ./$(TESTS_DIR)/...

build:
	@echo [3/4] Building...
	@$(MKDIR)
	@$(GO) build $(GO_BUILD_FLAGS) -o $(BIN_DIR)/$(PROJECT_NAME).exe ./$(CMD_DIR)
	@echo [3/4] Build complete: $(BIN_DIR)\$(PROJECT_NAME).exe

test:
	@echo [4/4] Running tests...
	@$(GO) test $(GO_TEST_FLAGS) ./$(TESTS_DIR)/...

run: build
	@echo Starting application...
	@$(BIN_DIR)\$(PROJECT_NAME).exe

clean:
	@echo Cleaning...
	@$(RMDIR)
	@echo Cleanup done

tidy:
	@$(GO) mod tidy