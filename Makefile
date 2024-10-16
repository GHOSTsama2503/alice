.PHONY: install run test build db/create

INPUT=.
OUTPUT=build/bot

MODULE=$(shell cat go.mod | grep module | cut -d " " -f 2)
VERSION=$(shell git describe --tags 2>/dev/null)

CMD=go build -trimpath -ldflags "-s -w -X $(MODULE)/common.version=$(VERSION)"

install:
	go mod download && go mod verify

run:
	go run -ldflags "-X $(MODULE)/common.version=$(VERSION)" ${INPUT}

test:
	go test ./...

build:
	${CMD} -o ${OUTPUT} ${INPUT}

db/create:
	migrate create -ext sql -dir database/migrations -seq $(SEQ)
