# gogen

gogen is a simple utility to generate starter template for any language and project type.

## Installation

```bash
go install github.com/zeindevs/gogen@latest

# OR
git clone https://github.com/zeindevs/gogen.git
cd gogen
make install
```

## Usage

```bash
gogen <type> <lang> <name:default main>

# ex:
gogen Makefile go main
gogen CMake cpp main
```

List of types:

- Makefile go
- Makefile cpp
- Makefile c
- CMake cpp
- CMake c

## License

Copyright (c) 2025 zeindevs and Lisenced under MIT License
