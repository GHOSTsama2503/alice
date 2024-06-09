install:
	go mod tidy

run:
	go run cmd/main.go

build:
	go build -ldflags "-s -w" -o build/bot cmd/main.go
