package main

import (
	"fmt"
	"os"

	"github.com/coltiebaby/og-parser/parser"
)

func main() {
	document, err := os.Open("./meta.html")
	if err != nil {
		fmt.Println(err)
	}

	defer document.Close()

	storage := make(map[string]interface{})
	meta := parser.Parse(document, storage)

	fmt.Println(meta)
}
