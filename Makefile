.PHONY: fmt tidy lint test install

default: fmt tidy lint test install

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

test:
	go test ./...

install:
	go install .
	gostrconv -h 2>&1 | head -n1
