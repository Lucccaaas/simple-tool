package converter

import (
	"testing"
	"encoding/json"
	"time"
)

func TestMapper(t *testing.T) {
	var jsonStr string = `
{
  "bool": true,
  "number": 1,
  "float": 1.1,
  "str": "string",
  "array": [
    1,
    2,
    3
  ]
}
`
	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &jsonObj)
	//javaStr := Convert("Demo", jsonObj)
	//fmt.Println(javaStr)
}

func TestTime(t *testing.T) {
	println(time.Now().Format("2006/01/02 03:04:05"))
}
