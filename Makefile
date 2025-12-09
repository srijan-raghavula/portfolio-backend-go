.PHONY: run build test fmt tidy clean

dev:
	air

run:
	go run ./cmd/web

build:
	go build -o bin/server ./cmd/web

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy


clean:
	rm -rf bin/
