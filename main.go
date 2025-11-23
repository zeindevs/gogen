package main

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed banner.txt
var banner embed.FS

func printBanner() {
	banner, err := banner.ReadFile("banner.txt")
	if err != nil {
		fmt.Println("Error reading banner.txt", err)
		return
	}
	fmt.Println(string(banner))
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
	lang := os.Args[2]
	name := "main"
	if len(os.Args) > 3 {
		name = os.Args[3]
	}
	templ, ok := template[gtype][lang]
	if !ok {
		fmt.Printf("Type %s and %s not found\n", gtype, lang)
		return
	}
	switch gtype {
	case "CMake":
		gtype = "CMakeLists.txt"
	}
	templ = strings.ReplaceAll(templ, "$name", name)
	if err := os.WriteFile(gtype, []byte(templ), os.ModePerm); err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fmt.Printf("Created %s file for %s 🎉\n", gtype, lang)
}
