package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

var version = "dev"

//go:embed template/*
var templateFS embed.FS

func main() {
	var showVer bool

	flag.BoolVar(&showVer, "v", false, "show version")
	flag.Usage = func() {
		fmt.Print(fmt.Sprintf(`
   ____ _____  ____ ____  ____
  / __ '/ __ \/ __ '/ _ \/ __ \
 / /_/ / /_/ / /_/ /  __/ / / /
 \__, /\____/\__, /\___/_/ /_/
/____/      /____/

Usage:
  gogen [template] [lang] [name]

Example:
  gogen Make go
  gogen CMake cpp app

Generators:
%s
Options:
`, generators()))
		flag.PrintDefaults()
	}
	flag.Parse()

	if showVer {
		fmt.Println("version", version)
		return
	}

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	gen := flag.Arg(0)
	lang := flag.Arg(1)

	name := "main"
	if flag.NArg() > 2 {
		name = flag.Arg(2)
	}

	templ, fname, err := readTemplate(lang, gen, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Type %s and %s not found\n", fname, lang)
		os.Exit(1)
	}

	if err := writeTemplate(lang, fname, templ); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
}

func generators() string {
	var sb strings.Builder

	templ, _ := templateFS.ReadDir("template")

	for _, tmp := range templ {
		r := strings.NewReplacer("File", "", "Lists.txt", "")

		tname := r.Replace(tmp.Name())
		parts := strings.Split(tname, ".")

		sb.WriteString(fmt.Sprintf("  %s %s\n", parts[0], parts[1]))
	}

	return sb.String()
}

func readTemplate(lang, temp, name string) ([]byte, string, error) {
	switch temp {
	case "Make":
		temp += "file"

	case "CMake":
		temp += "Lists.txt"

	case "goreleaser", "gorelease":
		temp += ".yaml"
	}

	raw, err := templateFS.ReadFile(fmt.Sprintf("template/%s.%s", lang, temp))
	if err != nil {
		return nil, temp, err
	}

	if slices.Contains([]string{"clang-format", "clang-tidy", "goreleaser.yaml"}, temp) {
		temp = "." + temp
	}

	raw = bytes.ReplaceAll(raw, []byte("$name"), []byte(name))

	return raw, temp, nil
}

func writeTemplate(lang, fname string, templ []byte) error {
	msg := "Created"

	if _, err := os.ReadFile(fname); err == nil {
		var y string

		fmt.Print("File already exist, will be replaced? (y/n): ")

		for {
			n, err := fmt.Scan(&y)
			if err != nil {
				os.Exit(1)
			}

			if n != 1 {
				fmt.Printf("Please enter (y/n)")
				continue
			}

			if y != "y" {
				return nil
			}

			break
		}

		msg = "Replaced"
	}

	if err := os.WriteFile(fname, templ, 0644); err != nil {
		return nil
	}

	fmt.Printf("%s %s for %s 🎉\n", msg, fname, lang)

	return nil
}
