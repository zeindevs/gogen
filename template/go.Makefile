all: build run

build:
	go build -o ./bin/$name .

install:
	go install -ldflags="-s -w" .

run:
	./bin/$name

clean:
	rm -rf ./bin
