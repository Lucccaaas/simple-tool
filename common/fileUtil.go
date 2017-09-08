package common

import (
	"path/filepath"
	"strings"
	"path"
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
