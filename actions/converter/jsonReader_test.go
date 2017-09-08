package converter

import (
	"testing"
)

func TestJsonReader(t *testing.T) {
	jsonObj := ReadJson("./static/1.json")
	if jsonObj == nil {
		t.Fatal("json reader fail")
	}
	t.Log("json reader success")
}
