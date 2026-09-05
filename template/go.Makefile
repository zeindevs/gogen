LDFLAGS = -ldflags="-s -w"

all: build run

build:
	go build $(LDFLAGS) -o ./bin/$name .

install:
	go install $(LDFLAGS) .

run:
	./bin/$name

clean:
	rm -rf ./bin
