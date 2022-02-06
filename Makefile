include .env.makefile

all: lint test

test:
	go test ./...

lint:
	@gofmt -l $(shell go list -f '{{.Dir}}' ./...)
	@golint -set_exit_status $(shell go list -f '{{.Dir}}' ./...)

server:
	go build -o floor-service ./cmd/

clean:
	@find . -name '*.orig' -exec rm {} \;
	rm -f floor-service

serve:
	@go run ./cmd/

.PHONY:	all lint test server