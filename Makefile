.PHONY: fmt tidy lint test bench install

default: fmt tidy lint test install

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

test:
	go test ./...

bench:
	go test -bench=Bench ./...

install:
	go install .
	perfsprint -h 2>&1 | head -n1
