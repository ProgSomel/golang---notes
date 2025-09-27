package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkErr(err)

	defer os.Remove(tempDir)
	fmt.Println("Temporary Directory created: ", tempDir)
}