package main

import (
	"flag"
	"simple-tool/actions/converter"
	"fmt"
	"simple-tool/common"
)

type config struct {
	json2java bool
	json      string
	dist      string
}

//st -gen -json=./test.json -dist=./dist/
func main() {
	gen := flag.Bool("gen", true, "is generator")
	json := flag.String("json", "", "json file path")
	dist := flag.String("dist", "", "generator file directory")

	flag.Parse()

	if *json == "" {
		panic("no such file: " + *json)
	}

	fileName := common.UpperCaseFirst(common.FileName(*json))

	if *dist == "" {
		*dist = fileName + ".java"
	}

	inputConfig := config{*gen, *json, *dist}

	jsonObj := converter.ReadJson(inputConfig.json)
	javaClass := converter.JsonToJavaClass(fileName, *jsonObj)
	javaFile := converter.ClassToJavaFile(&javaClass)
	javaString := converter.FileDescriptionToString(javaFile)
	converter.WriteToFile(javaString, inputConfig.dist)

	fmt.Printf("generated %s from %s \n", *dist, *json)
}
