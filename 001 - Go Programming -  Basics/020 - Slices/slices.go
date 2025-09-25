package main

import (
	"fmt"
)

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:] //? 2 3 4 5

	slice2 := slice[2:4] //? 4 5
	fmt.Println(slice2)

	fmt.Println("Length of slice2: ", len(slice2))
	fmt.Println("The Capacity of slice2: ", cap(slice2))

}