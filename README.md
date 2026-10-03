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
```

```bash
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
- github gorelease
- go Makefile
- go goreleaser
- qt clang-format

## License

Licensed under the Apache License, Version 2.0 (the "License"); See [LICENSE](LICENSE) for more details.
