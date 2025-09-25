package main

import "fmt"

func main(){
	// fmt.Println(factorial(5))
	// fmt.Println(factorial(10))
	fmt.Println(sumOfDigits(5, 0))
}

func factorial(n int) int{
	//? Base case: factorial of 0 is 1
	if n == 0{
		return 1
	}

	//? Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)

}

func sumOfDigits(n int, sum int) int{
	if n == 0{
		return sum
	}
	sum += n;
	return sumOfDigits(n-1, sum) 
}