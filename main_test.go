package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTemplate(t *testing.T) {
	templs, err := templateFS.ReadDir("template")
	if err != nil {
		t.Fatal(err)
	}

	for _, tmp := range templs {
		r := strings.NewReplacer("File", "", "Lists.txt", "")
		tname := r.Replace(tmp.Name())
		parts := strings.Split(tname, ".")
		t.Log(parts)
	}

	langs := []string{"c", "cpp", "go", "qt"}
	tmps := []string{"Make", "CMake", "clang-format", "clang-tidy"}
	name := "test"

	tpl := func(lang, temp, name string) ([]byte, string, error) {
		switch temp {
		case "Make":
			temp += "file"
		case "CMake":
			temp += "Lists.txt"
		}
		raw, err := templateFS.ReadFile(fmt.Sprintf("template/%s.%s", lang, temp))
		if err != nil {
			return nil, temp, fmt.Errorf("File does not exist")
		}
		raw = bytes.ReplaceAll(raw, []byte("$name"), []byte(name))
		return raw, temp, nil
	}

	for _, lang := range langs {
		for _, tmp := range tmps {
			raw, temp, err := tpl(lang, tmp, name)
			if err != nil {
				t.Log(temp, "not available", "for", lang)
			} else {
				_ = raw
				t.Log(temp, "available for", lang)
			}
		}
	}
}
