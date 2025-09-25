package main

import "fmt"

func process(input int){
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0{
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number") //! panic(interface{})
	}
	fmt.Println("Processing Input: ", input)
}

func main(){

	process(-10)

}