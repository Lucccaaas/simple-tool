package main

import (
	"testing"
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"reflect"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestFetchPath(t *testing.T) {
	t.Log("test fetch current path")
	{
		path, err := os.Getwd()
		if err != nil {
			t.Fatal("path=", err, ballotX)
		} else {
			t.Log("path=\t", path)
			{
				if path == "/Users/yunge/projects/gogland-projects/simple-tool/test" {
					t.Logf("get correct path", checkMark)
				}
			}
		}
	}
	t.Log("\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestFetchJsonFile(t *testing.T) {
	t.Log("test read json file")
	fileName := "./test.json"
	file, error := ioutil.ReadFile(fileName)
	check(error)
	t.Log(string(file))

	str := string(file)

	var res map[string]interface{}

	json.Unmarshal([]byte(str), &res)

	for key, value := range res {
		fmt.Printf("key,vlaue,type = %s, %v, %s\n", key, value, reflect.TypeOf(value))
	}
}
