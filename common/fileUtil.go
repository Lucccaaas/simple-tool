package common

import (
	"path/filepath"
	"strings"
	"path"
	"os"
	"fmt"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func FileName(pathParam string) string {
	fullFileName, error := filepath.Abs(pathParam)
	Check(error)

	filenameWithSuffix := path.Base(fullFileName)
	fileSuffix := path.Ext(filenameWithSuffix)

	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

func ExtensionFiles(pathParam string, ext string) []string {
	fullFileName, error := filepath.Abs(pathParam)
	dir, error := os.OpenFile(fullFileName, os.O_RDONLY, os.ModeDir)
	if error != nil {
		defer dir.Close()
		fmt.Println(error.Error())
		panic(error)
	}
	fileInfo, _ := dir.Stat()
	if !fileInfo.IsDir() {
		fmt.Println(fullFileName, "不是一个有效的目录路径")
	}
	names, _ := dir.Readdir(-1)
	var fileNames []string
	for _, name := range names {
		if !name.IsDir() && strings.HasSuffix(name.Name(), ext) {
			fileNames = append(fileNames, name.Name())
		}
	}
	return fileNames
}
