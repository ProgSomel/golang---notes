package main

import "fmt"

func process(){
	defer func(){
		r := recover()
		if r!=nil{
			fmt.Println("Recovered: ", r)
		} 
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}

func main(){

	process()
	fmt.Println("Returned from process")

}