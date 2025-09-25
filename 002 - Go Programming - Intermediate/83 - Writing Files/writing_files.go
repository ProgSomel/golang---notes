package main

import (
	"fmt"
	"os"
)

func main() {
	//? Creating a file using OS Package
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating a file.", err)
		return
	}
	defer file.Close()

	//? write data to file
	data := []byte("Hello World!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to files: ", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("writingString.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	
}