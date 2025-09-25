package main

import "fmt"


func main(){
    number := 15
	switch {
    case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")

	}


}