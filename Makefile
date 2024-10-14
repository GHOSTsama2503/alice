.PHONY: install run test build

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
