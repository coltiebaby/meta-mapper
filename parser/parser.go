package parser

import (
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

const (
	SPLITTER string = `:`
)

func createOrGet(key string, v map[string]interface{}) map[string]interface{} {
	if x, ok := v[key]; ok {
		return x.(map[string]interface{})
	} else {
		v[key] = make(map[string]interface{})
		return v[key].(map[string]interface{})
	}
}

func Set(key string, v map[string]interface{}, value interface{}) {
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

func Fetch(keys []string, v map[string]interface{}) (string, map[string]interface{}) {
	var key string
	for _, key = range keys {
		v = createOrGet(key, v)
	}

	return key, v
}

func Parse(document io.Reader, storage interface{}) map[string]interface{} {
	nodes, _ := html.Parse(document)
	meta := FetchMetaInformation(nodes)

	for _, m := range meta {
		Rename(strings.Split(m[0], SPLITTER), m[1], storage.(map[string]interface{}))
	}

	return storage.(map[string]interface{})
}

func Rename(keys []string, value string, storage map[string]interface{}) {
	k, s := Fetch(keys, storage)
	Set(k, s, value)
}

func FetchMetaInformation(nodes *html.Node) [][2]string {
	var scraper func(*html.Node)
	storage := [][2]string{}

	scraper = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			if len(n.Attr) == 2 {
				keyType := n.Attr[0].Key
				if keyType == "property" || keyType == "name" {
					// keys := strings.Split(n.Attr[0].Val, SPLITTER)
					keys := n.Attr[0].Val
					value := n.Attr[1].Val

					storage = append(storage, [2]string{keys, value})
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			scraper(c)
		}
	}

	scraper(nodes)
	return storage
}
