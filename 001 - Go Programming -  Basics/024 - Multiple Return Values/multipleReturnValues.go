package main

import (
	"fmt"
)

//! func functionName(parameter1 type1, paramater2 type2, ...) (returnType1, returnType2){}
func divide(a, b int) (quotient int, remainder int){
	quotient = a / b
	remainder = a % b
	return
}

func main(){
	quotient, remainder := divide(4, 4)
	fmt.Println(quotient, remainder)

}
