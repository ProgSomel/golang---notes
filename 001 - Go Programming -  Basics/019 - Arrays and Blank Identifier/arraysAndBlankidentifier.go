package main

import "fmt"
func main(){

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int = &originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", originalArray)
}