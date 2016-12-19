.PHONY: all run

all: run

run: build
	./mockr

build: fmt
	go-bindata assets/... && go build

fmt:
	go fmt ./...

test:
	go test ./...
