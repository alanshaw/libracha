.PHONY: gen build test

gen:
	go generate ./...

build:
	go build ./...

test:
	go test ./...
