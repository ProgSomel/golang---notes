# 58 - Recursion
Recursion is a fundamental concept in programming, where a function calls itself directly or indirectly in order to solve a problem.

Recursion is the process of a function calling itself. It breaks down a problem into subproblems of the same type, untill they become simple enough to solve directly.

```go
package main

import "fmt"

func main(){
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}

func factorial(n int) int{
	//? Base case: factorial of 0 is 1
	if n == 0{
		return 1
	}

	//? Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)

}
```
```bash
120
3628800
```

-----------------------------------------------------------------------------------------------------------------------

```go
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
```
```bash
15
```