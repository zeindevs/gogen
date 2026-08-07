# gogen - Build system generator

## Installation

```bash
go install github.com/zeindevs/gogen@latest
```

Download:

- <https://github.com/zeindevs/gogen/releases>

## Usage

```bash
gogen <type> <lang> <name:default main>
```

Example:

```bash
gogen Makefile go main
gogen CMake cpp main
```

List of generators:

- c CMake
- c Makefile
- c clang-format
- c clang-tidy
- cpp CMake
- cpp Makefile
- cpp clang-format
- cpp clang-tidy
- go Makefile
- qt clang-format

## License

Copyright 2025-2026 Husein Assidiq. All Rights Reserved.

[Apache-2.0](./LICENSE)
