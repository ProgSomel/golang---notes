package main

import "fmt"


func main(){
	if 10%2 == 0 || 5%2 == 0{
		fmt.Println("Either 10 or 5 are Even")
	}
	if 10%2 == 0 && 5%2 == 0{
		fmt.Println("Both 10 or 5 are Even")
	}
}