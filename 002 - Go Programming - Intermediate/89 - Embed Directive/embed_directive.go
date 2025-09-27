package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basic
var basicsFolder embed.FS

func main() {
	fmt.Println("Embeded content: ", content)
	content, err := basicsFolder.ReadFile("basic/hello.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("Embeded file content: ", string(content))

	err = fs.WalkDir(basicsFolder, "basic", func(path string, d fs.DirEntry, err error) error{
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}