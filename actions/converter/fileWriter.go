package converter

import (
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteToFile(content string, path string) {
	path, _ = filepath.Abs(path)

	f, err := os.Create(path)

	check(err)
	defer f.Close()
	f.Write([]byte(content))
}
