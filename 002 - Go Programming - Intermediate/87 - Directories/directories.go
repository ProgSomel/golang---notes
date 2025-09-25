package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	pathFile := "subdir/parent"
	err := filepath.WalkDir(pathFile, func(path string, d os.DirEntry, err error) error {
		checkError(err)
		fmt.Println(path)
		return nil
	})
	checkError(err)

	checkError(os.RemoveAll("./subdir"))
}