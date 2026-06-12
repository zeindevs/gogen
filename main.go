package main

import (
	"embed"
	"fmt"
	"os"
	"slices"
	"strings"
)

//go:embed banner.txt
var banner embed.FS

var Version = "v0.0.1"

func printUsage() {
	fmt.Println("Usage: gogen <type> <lang> <name:default main>")
}

func printBanner() {
	text, _ := banner.ReadFile("banner.txt")
	fmt.Println(strings.ReplaceAll(string(text), "$version", Version))
}

func main() {
	if slices.Contains(os.Args, "-h") {
		printBanner()
		printUsage()
		fmt.Print("\nList of generators:\n\n")
		for k := range template {
			for kl := range template[k] {
				fmt.Println("-", k, kl)
			}
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

	templ, ok := template[gen][lang]
	if !ok {
		fmt.Printf("Type %s and %s not found\n", gen, lang)
		return
	}

	switch gen {
	case "CMake":
		gen = "CMakeLists.txt"
	}

	templ = strings.ReplaceAll(templ, "$name", name)
	if err := os.WriteFile(gen, []byte(templ), 0644); err != nil {
		fmt.Println("Error creating file", err)
		return
	}

	fmt.Printf("Created %s file for %s 🎉\n", gen, lang)
}
