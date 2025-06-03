all: build test

build:
	go build -o bin/$(unreliable-network) ./cmd

run:
	go run ./cmd/main.go

test:
	go test ./...

clear:
	rm -rf bin

deploy:
	@echo "Развёртывание симулировано."
