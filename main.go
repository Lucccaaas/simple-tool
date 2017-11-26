package main

import (
	"flag"
	"simple-tool/actions/converter"
	"fmt"
	"simple-tool/common"
	"strings"
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

	if *json != "" {

		*dist = distFileName(*dist, fileName(*json))

		inputConfig := config{*gen, *json, *dist}

		processSingleFile(inputConfig)
	} else {
		fileNames := common.ExtensionFiles("./", ".json")
		fmt.Println("检测到.json文件: ", fileNames)
		for _, fileNameItem := range fileNames {
			dist := distFileName(*dist, fileName(fileNameItem))
			inputConfig := config{*gen, fileNameItem, dist}
			processSingleFile(inputConfig)
		}
	}
}

func fileName(json string) string {
	return common.UpperCaseFirst(common.FileName(json))
}

func processSingleFile(inputConfig config) {
	jsonObj := converter.ReadJson(inputConfig.json)
	javaClass := converter.JsonToJavaClass(fileName(inputConfig.json), *jsonObj)
	javaFile := converter.ClassToJavaFile(&javaClass)
	javaString := converter.FileDescriptionToString(javaFile)
	converter.WriteToFile(javaString, inputConfig.dist)

	fmt.Printf("generated %s from %s \n", inputConfig.dist, inputConfig.json)
}

func distFileName(dist string, fileName string) string {
	if dist == "" {
		fileName = strings.TrimSuffix(fileName, ".json")
		return fileName + ".java"
	} else {
		return dist
	}
}
