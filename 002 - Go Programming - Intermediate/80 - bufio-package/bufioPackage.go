package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	str := "This is string.\n"
	n, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes. \n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error")
	}
}