all: build run

build:
	go build -o ./bin/gogen .

install:
	CGO_ENABLED=0 go install -ldflags="-s -w" .

test:
	 go test -v ./ -count=1

run:
	./bin/gogen

clean:
	rm -rf ./bin
