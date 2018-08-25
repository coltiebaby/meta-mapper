package main

import (
	"fmt"
	"strings"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func buildMap(keys []string, v map[string]interface{}) {
	if len(keys) == 0 {
		return
	}

	key := keys[0]
	if x, ok := v[key]; ok {
		buildMap(keys[1:], x.(map[string]interface{}))
	} else {
		v[key] = make(map[string]interface{})
		buildMap(keys[1:], v[key].(map[string]interface{}))
	}

}

func main() {
	keys := strings.Split("og:image:url", ":")[1:]
	fmt.Println(keys)

	x := make(map[string]interface{})
	buildMap(keys, x)

	fmt.Println(x)
}

