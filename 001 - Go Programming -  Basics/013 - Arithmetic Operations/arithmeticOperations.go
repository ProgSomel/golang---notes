package main

import (
	"fmt"
	"math"
)

func main(){

	var smallFloat float64 = 1.0e-323 //? very small positive number in scientific notation
	fmt.Println(smallFloat) //? 1e-323
	smallFloat = smallFloat/math.MaxFloat64 //? diving with largest number ; the biggest float number which can hold by float
	fmt.Println(smallFloat) //? 0

}