package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"slices"
	"strings"
)

var version = "dev"

var banner = `
   ____ _____  ____ ____  ____
  / __  / __ \/ __  / _ \/ __ \
 / /_/ / /_/ / /_/ /  __/ / / /
 \__, /\____/\__, /\___/_/ /_/
/____/      /____/ %s
`

//go:embed template/*
var templateFS embed.FS

func printBanner() {
	fmt.Println(fmt.Sprintf(banner, version))
}

func printUsage() {
	fmt.Println("Usage: gogen <type> <lang> <name:default main>")
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

func main() {
	if slices.Contains(os.Args, "-h") {
		printBanner()
		printUsage()
		fmt.Print("\nList of generators:\n")
		templ, _ := templateFS.ReadDir("template")
		for _, tmp := range templ {
			r := strings.NewReplacer("File", "", "Lists.txt", "")
			tname := r.Replace(tmp.Name())
			parts := strings.Split(tname, ".")
			fmt.Println("-", parts[0], parts[1])
		}
		return
	}
	if len(os.Args) < 3 {
		printUsage()
		return
	}
	gen := os.Args[1]
	lang := os.Args[2]
	name := "main"
	if len(os.Args) > 3 {
		name = os.Args[3]
	}
	templ, fname, err := readTemplate(lang, gen, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Type %s and %s not found\n", fname, lang)
		os.Exit(1)
	}
	msg := "Created"
	if _, err := os.ReadFile(fname); err == nil {
		fmt.Print("File already exist, will be replaced? (y/n): ")
		var y string
		for {
			n, err := fmt.Scan(&y)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Err: %s\n", err)
				os.Exit(1)
			}
			if n != 1 {
				fmt.Printf("Please enter (y/n)")
				continue
			}
			if y != "y" {
				return
			}
			msg = "Replaced"
			break
		}
	}
	if err := os.WriteFile(fname, templ, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s %s for %s 🎉\n", msg, fname, lang)
}
