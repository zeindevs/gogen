package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed banner.txt
var banner embed.FS

var (
	Version = "v0.0.1"
)

func printBanner() {
	text, err := banner.ReadFile("banner.txt")
	if err != nil {
		fmt.Println("Error reading banner.txt", err)
		return
	}
	textStr := strings.ReplaceAll(string(text), "$version", Version)
	fmt.Println(textStr)
}

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printBanner()
		fmt.Println("Usage: gogen <type> <lang> <name:default main>")
		fmt.Print("\nList of types:\n\n")
		for k := range template {
			for kl := range template[k] {
				fmt.Println("-", k, kl)
			}
		}
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Usage: gogen <type> <lang> <name:default main>")
		return
	}
	gtype := os.Args[1]
	glang := os.Args[2]
	gname := "main"
	if len(os.Args) > 3 {
		gname = os.Args[3]
	}
	templ, ok := template[gtype][glang]
	if !ok {
		fmt.Printf("Type %s and %s not found\n", gtype, glang)
		return
	}
	switch gtype {
	case "CMake":
		gtype = "CMakeLists.txt"
	}
	templ = strings.ReplaceAll(templ, "$name", gname)
	if err := os.WriteFile(gtype, []byte(templ), os.ModePerm); err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fmt.Printf("Created %s file for %s 🎉\n", gtype, glang)
}
