all: build run

build:
	go build -o ./bin/gogen .

install:
	go install -ldflags="-s -w" .

run:
	./bin/gogen

clean:
	rm -rf ./bin
