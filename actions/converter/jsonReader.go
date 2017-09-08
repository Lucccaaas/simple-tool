package converter

import (
	"path/filepath"
	"encoding/json"
	"io/ioutil"
)

func ReadJson(name string) *map[string]interface{} {
	path, _ := filepath.Abs(name)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	jsonObj := map[string]interface{}{}

	json.Unmarshal(data, &jsonObj)

	return &jsonObj
}
