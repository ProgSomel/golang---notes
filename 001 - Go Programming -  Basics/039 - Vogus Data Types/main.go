package main

import "fmt"

func main(){
	var a int8 = 5
	var b int8 = 10

	var x uint8 = 255; //?undigned (0 and only positive)

	var j float32 = 10.23343

	var k float64 = 10.4545

	var flag bool = false;

	var s string = "My name is Somel Ahmed";

	fmt.Printf("%d\n", a);//? 5
	fmt.Printf("%d\n", b);//? 10
	fmt.Printf("%d\n", x);//? 255
	fmt.Printf("%f\n", j);//? 10.233430
	fmt.Printf("%.2f\n", k);//? 10.45
	fmt.Printf("%v\n", flag);//? false
	fmt.Printf("%s\n", s);//? My name is Somel Ahmed
	fmt.Printf("%T\n", s);//? string
	
}

/*
	2 phases:
	  1. compilation phase (compile time)
	  2. execution phase (run time)
	
	********* Compile phase(Binary File) ***************
	** Code Segment **
	main = func (){...}


	go run main.go => compile it => main => ./main
	go build main.go => compile it => main 

*/