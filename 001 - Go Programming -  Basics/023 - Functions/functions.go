package main

import "fmt"

func add(a, b int) int{
	return  a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int{
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int{
	return func(x int) int {
		return x*factor
	}
}

func main(){
	result := applyOperation(5, 3, add)
	fmt.Println(result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println(multiplyBy2(6))
}