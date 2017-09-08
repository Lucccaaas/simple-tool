package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var jsonString string = `
	{
  "bool": true,
  "number": 1,
  "float": 1.1,
  "str": "string",
  "array": [
    1,
    2,
    3
  ],
  "obj": {
    "key1": "value1"
  }
}
	`

	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonObj)

	for key, value := range jsonObj {
		fmt.Printf("key, value, type :  %v %v %v\n", key, value, reflect.TypeOf(value))
	}
}



