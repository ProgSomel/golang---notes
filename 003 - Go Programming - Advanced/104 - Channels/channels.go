package main

import (
	"fmt"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"

		for _, c := range "abcde"{
			greeting <- "Alphabet: " + string(c)
		}
	}()
	
	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)
	
	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	fmt.Println("End of Program")
}