.PHONY: install run test build db/create

INPUT=.
OUTPUT=build/bot

CMD=go build -trimpath -ldflags "-s -w"

install:
	go mod download && go mod verify

run:
	go run ${INPUT}

test:
	go test ./...

build:
	${CMD} -o ${OUTPUT} ${INPUT}

db/create:
	migrate create -ext sql -dir database/migrations -seq $(SEQ)
