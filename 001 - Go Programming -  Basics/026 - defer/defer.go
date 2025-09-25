package main

import "fmt"

func process(i int){
	defer fmt.Println("Deffered i value: ", i)
	defer fmt.Println("First Deffered statement executed")
	defer fmt.Println("Second Deffered statement executed")
	defer fmt.Println("Third Deffered statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i is: ", i)
}

func main(){
	process(10)

}