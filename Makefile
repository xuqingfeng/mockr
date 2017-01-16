.PHONY: all run

deps:
	go get -v -d ./...

all: run

run: build
	./mockr

build: fmt
	go build

fmt:
	go fmt ./...

test:
# http://stackoverflow.com/a/29085684/4036946
	go test -v $$(go list ./... | grep -v /vendor/)

bindata:
	go-bindata assets/...
