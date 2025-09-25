# 023 - Functions
Functions are fundamental building blocks in go, encapsulating reusable code blocks that can be invoked multiple times with different inputs. They play a crucial role in structuring go programs by promoting modularity and code reusability.

- A Function name should be valid identifier and should follow ghost naming conventions. When we are making a public function, it should start with an uppercase letter, and if it's a private function, then it needa to start with a lowercase.

**--------------------------------------------------------------------------------------------------------**

## Function's Parameters
Zero or more Parameters can be defined, each with a name and a type. Parameter act as variables inside the function.

**--------------------------------------------------------------------------------------------------------**

## Function's return type
Specifies the type of the values returned by the function. We also have option of returning multiples values, and it is one of the unique aspects of go language.

**--------------------------------------------------------------------------------------------------------**

## Function body
The function body contains the code that needs to be executed when the function is called and in the function body the last thing will be the return statement. a return statement is optional and it returns one or more values to the caller. if return statement is omitted, then functions return default zero values for their types.

**--------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"


func add(a, b int) int{
	return a + b
}

func main(){
	sum := add(4, 5)
	fmt.Println(sum)
}
```
```bash
9
```

---------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func add(a, b int) int{
	return a + b
}

func main(){
	fmt.Println(add(4, 5))
}
```
```bash
9
```

**------------------------------------------------------------------------------------------------------**

## arguments that are passed to a function are copied into the function's parameters. Modification to the parameters inside the function do not affect the original arguments, so never expect the original argument to be updated if the parameters are updated inside the function.

**------------------------------------------------------------------------------------------------------**

## Anonymous Function
The function without a name
```go
package main

import "fmt"

func main(){
	greet := func(){
		fmt.Println("Hello Anonymous Function")
	}
	greet()
}
```
```bash
Hello Anonymous Function
```

**------------------------------------------------------------------------------------------------------**

## Function as a type
```go
package main

import "fmt"

func main(){

	operation := add
	result := operation(3, 5)
	fmt.Println(result)
}

func add(a, b int) int{
	return  a + b
}
```
```bash
8
```

**------------------------------------------------------------------------------------------------------**

## First class citizen or first class objects
Refers to entities that have no restrictions on their use and can be treated uniformly throughout the language. When an entity is a first class citizen, it means that you can perform a wide range of operations on it, just as you would with basic data types like integers or strings. And these operations typically include passing as arguments.
```go
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
}
```
```bash
8
```

------------------------------------------------------------------------------------------------------

```go
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
```
```bash
8
12
```