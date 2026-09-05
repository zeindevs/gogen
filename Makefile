VERSION = $(shell git describe --tags --abbrev=0)
LDFALGS = -ldflags="-s -w -X main.version=$(VERSION)"

all: build run

build:
	go build $(LDFALGS) -o ./bin/gogen .

install:
	CGO_ENABLED=0 go install $(LDFALGS) .

run:
	./bin/gogen

clean:
	rm -rf ./bin
