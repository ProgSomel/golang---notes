package main

import (
	"fmt"
)

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func myPrinterType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: Unknown")
	}
}

func main(){
	myPrinterType(9)
	myPrinterType("John")
	myPrinterType(false)
}