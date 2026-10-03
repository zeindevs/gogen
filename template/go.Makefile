.PHONY: test

VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null | echo 'dev')
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

all: build run

build:
	go build $(LDFLAGS) -o ./bin/$name .

install:
	CGO_ENABLED=0 go install $(LDFLAGS) .

run:
	./bin/$name

test:
	go test -v ./... -count=1

clean:
	rm -rf ./bin
