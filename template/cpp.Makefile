.PHONY: build

all: setup build run

setup:
	cmake -B build -DCMAKE_BUILD_TYPE=Debug

build:
	cmake --build build

install:
	cmake --install build

run:
	./build/$name

clean:
	rm -rf build
