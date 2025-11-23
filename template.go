package main

var template = map[string]map[string]string{
	"Makefile": {
		"go": `all: build run

build:
	go build -o ./bin/$name .

install:
	go install -ldflags="-s -w" .

run:
	./bin/$name

clean:
	rm -rf ./bin`,
		"cpp": `.PHONY: build
all: setup build run

setup:
		cmake -B build

build:
		cmake --build build

run:
	./build/$name

clean:
	rm -rf build`,
		"c": `.PHONY: build
all: setup build run

setup:
		cmake -B build

build:
		cmake --build build

run:
	./build/$name

clean:
	rm -rf build`,
	},
	"CMake": {
		"cpp": `cmake_minimum_required(VERSION 3.25)

project($name VERSION 0.0.1 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -pedantic")
set(CMAKE_CXX_FLAGS_DEBUG "${CMAKE_CXX_FLAGS_DEBUG} -g")
set(CMAKE_CXX_FLAGS_RELEASE "${CMAKE_CXX_FLAGS_RELEASE} -O3")

add_executable(${PROJCT_NAME} src/main.cpp)

target_include_directories(${PROJCT_NAME} PUBLIC src)`,
		"c": `cmake_minimum_required(VERSION 3.25)

project($name VERSION 0.0.1 LANGUAGES C)

set(CMAKE_C_STANDARD 23)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra -pedantic")
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g")
set(CMAKE_C_FLAGS_RELEASE "${CMAKE_C_FLAGS_RELEASE} -O3")

add_executable(${PROJCT_NAME} src/main.c)

target_include_directories(${PROJCT_NAME} PUBLIC src)`,
	},
}
