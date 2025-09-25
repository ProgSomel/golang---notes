package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()
	fmt.Println("File opened successfully")

	scanner := bufio.NewScanner(file)

	//? Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
}