package main

import "fmt"

func calculate() (result int){

	fmt.Println("First: ", result) //? First: 0

	show:= func(){ //? show will be stored in heap. it will form a closure with result
		result = result + 10
		fmt.Println("Defer: ",result) //? 15 
	}

	defer show()

	result = 5

	p := func(a int){ //?it will be stored in heap with the argument a which is 5 from result. 
	                //? it will not form any closure because it is not using any outer scope variable likde show
		fmt.Println("Ami: ", a) //? 5
	}
	defer p(result)

	defer fmt.Println(result) //? it will stored in heap with argument result which is 5
	                          //? output: 5

	fmt.Println("Second: ", result) //? Second: 5

	defer fmt.Println(5) //? it will stored in heap with argument which is 5
	                     //? output is: 5


	return
}

func main(){
	
	a := calculate()
	fmt.Println(a)
}