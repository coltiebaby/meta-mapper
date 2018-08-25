package main

import (
	"fmt"
	"reflect"
	"strings"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Foo(key string, v map[string]interface{}) map[string]interface{} {
	if x, ok := v[key]; ok {
		return x.(map[string]interface{})
	} else {
		v[key] = make(map[string]interface{})
		return v[key].(map[string]interface{})
	}
}

func buildMap(keys []string, v map[string]interface{}) (string, map[string]interface{}) {
	var key string
	for _, key = range keys {
		v = Foo(key, v)
	}

	return key, v
}

func setValue(key string, v map[string]interface{}, value interface{}) {
	if v[key] == nil {
		v[key] = value
		return
	}

	rt := reflect.TypeOf(v[key])
	switch rt.Kind() {
	case reflect.Int, reflect.String:
		oldVal := v[key]
		v[key] = []interface{}{oldVal, value}
	case reflect.Slice:
		v[key] = append(v[key].([]interface{}), value)
	}
}

func main() {
	// var name string = `Colton`

	keys := strings.Split("og:image:url", ":")[1:]

	x := make(map[string]interface{})
	key, z := buildMap(keys, x)
	setValue(key, z, "Colton")
	setValue(key, z, "Robertson")

	fmt.Println(x)
}
