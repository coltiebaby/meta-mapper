package parser

import (
	"strings"
	"testing"
)

func getFixture() [][2]string {

	testStrings := [][2]string{
		[2]string{"og:name:first", "John"},
		[2]string{"og:name:first", "Smith"},
		[2]string{"og:image", "ImageUrl"},
	}

	return testStrings
}

func TestCreateOrGet(t *testing.T) {
	wantedKey := "test"
	storage := make(map[string]interface{})

	if _, ok := storage[wantedKey]; ok {
		t.Error("Expected storage to not contain a key")
	}

	createOrGet(wantedKey, storage)

	if _, ok := storage[wantedKey]; !ok {
		t.Error("Expected storage to contain a key")
	}
}

func TestFetch(t *testing.T) {

	testStrings := getFixture()
	storage := make(map[string]interface{})

	for _, ts := range testStrings {
		keys := strings.Split(ts[0], ":")
		key, v := Fetch(keys, storage)

		if _, ok := v[key]; ok {
			t.Errorf("Expected key (%s) to exist!", key)
		}
	}

}
