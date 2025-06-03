.PHONY: all build run test clean deploy

all: build run

build:
	go build -o string-processor cmd/main.go

run:
	./string-processor.exe

test:
	go test ./...

clean:
	rm -f string-processor

deploy:
	echo "Deploying application..."
