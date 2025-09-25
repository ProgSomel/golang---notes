package main

import (
	"fmt"
	"strconv"
)

func main() {
	invalidNum := "456abc"
	invalidNumParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value: ", err)
		return
	}
	fmt.Println("Parsed invalid number: ", invalidNumParse)
}