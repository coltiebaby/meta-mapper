package main

import (
	"fmt"
	"strings"

	"github.com/coltiebaby/og-parser/parser"
)

func main() {
	keys := strings.Split("og:image:url", ":")[1:]

	x := make(map[string]interface{})
	key, z := parser.Fetch(keys, x)
	parser.Set(key, z, "Colton")
	parser.Set(key, z, "Robertson")

	fmt.Println(x)
}
