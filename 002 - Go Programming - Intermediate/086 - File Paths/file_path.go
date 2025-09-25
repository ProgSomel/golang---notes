package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	relativePath := "./data/file.txt"
	// absolutePath := "/home/user/docs/file.txt"

	fmt.Println("Relative Path: ", relativePath)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}else{
		fmt.Println("Absolute Path: ", absPath)
	}
}