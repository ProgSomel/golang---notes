# 019 - Arrays and Blank Identifier
An array is a fixed size collection of elements of the same size.
Arrays are fundamental data structure that allow you to store multiple values under a single variable name.
Understanding arrays is crucial as they provide a way to manage and manipulate ordered data efficiently.

```go
package main
func main(){
	//! var arrayName [size]elementType
}
```

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	//! var arrayName [size]elementType
	var numbers [5]int //? it is a blank array. all values will be initialize with zero
	fmt.Println(numbers)

}
```
```bash
[0 0 0 0 0]
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	//! var arrayName [size]elementType
	var numbers [5]int //? it is a blank array. all values will be initialize with zero
	fmt.Println(numbers)

	//? first index will be zero
	//? and last index will be: sizeOfTheArray - 1
	numbers[4] = 20
	fmt.Println(numbers) 

}
```
```bash
[0 0 0 0 0]
[0 0 0 0 20]
```

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	//! var arrayName [size]elementType
	var numbers [5]int //? it is a blank array. all values will be initialize with zero
	fmt.Println(numbers)

	//? first index will be zero
	//? and last index will be: sizeOfTheArray - 1
	numbers[4] = 20
	fmt.Println(numbers) 

	//! initialzing array with values
	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits Array: ", fruits)

}
```
```bash
[0 0 0 0 0]
[0 0 0 0 20]
Fruits Array:  [Apple Banana Orange Grapes]
```

---------------------------------------------------------------------------------------------------------

## In Go arrays are value types. When you assign an array to a new variable or pass an array as 
## an argument to a function, a copy of the original array is created and modification to the copy do not affect the original array.
```go
package main

import "fmt"
func main(){
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", copiedArray)

}
```
```bash
Original Array:  [1 2 3]
Copied Array:  [100 2 3]
```

--------------------------------------------------------------------------------------------------------

## We can iterate over the elements of an array using a for loop
```go
package main

import "fmt"
func main(){
	numbers := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(numbers); i++{
		fmt.Println("Element at index, ", i, ":", numbers[i])
	}

}
```
```bash
Element at index,  0 : 1
Element at index,  1 : 2
Element at index,  2 : 3
Element at index,  3 : 4
Element at index,  4 : 5
Element at index,  5 : 6
Element at index,  6 : 7
Element at index,  7 : 8
Element at index,  8 : 9
Element at index,  9 : 10
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	numbers := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers{
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
```
```bash
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
Index: 3, Value: 4
Index: 4, Value: 5
Index: 5, Value: 6
Index: 6, Value: 7
Index: 7, Value: 8
Index: 8, Value: 9
Index: 9, Value: 10
```

---------------------------------------------------------------------------------------------------------

## Unerscore is a blank Identifier, used to store unused values
```go
package main

import "fmt"
func main(){
	numbers := [10]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numbers{
		fmt.Printf("Value: %d\n", value)
	}

}
```
```bash
Value: 1
Value: 2
Value: 3
Value: 4
Value: 5
Value: 6
Value: 7
Value: 8
Value: 9
Value: 10
```

---------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	
	a, b := someFunction()
	fmt.Println(a)
	fmt.Println(b)

}

func someFunction() (int, int){
	return 1, 2
}
```
```bash
1
2
```

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	
	a, _ := someFunction()
	fmt.Println(a)

}

func someFunction() (int, int){
	return 1, 2
}
```
```bash
1
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	
	a, _ := someFunction()
	fmt.Println(a)

	b := 10 //? declared and not used: b compiler
	

}

func someFunction() (int, int){
	return 1, 2
}
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"
func main(){
	
	a, _ := someFunction()
	fmt.Println(a)

	b := 10

	_ = b


}

func someFunction() (int, int){
	return 1, 2
}
```

-------------------------------------------------------------------------------------------------------

## Comparing Arrays
```go
package main

import "fmt"
func main(){

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array1 is Equal to Array2: ", array1 == array2)


}
```
```bash
Array1 is Equal to Array2:  true
```

**-------------------------------------------------------------------------------------------------------**

## Go supports Multi-Dimensional Arrays - which are arrays of arrays. They are useful for representing matrices and other structured data.
```go
package main

import "fmt"
func main(){

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
}
```
```bash
[[1 2 3] [4 5 6] [7 8 9]]
```

**-------------------------------------------------------------------------------------------------------**

## to use original array in copied array, we have to use pointers ans addressed
```go
package main

import "fmt"
func main(){

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int = &originalArray
	copiedArray[0] = 100
	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", originalArray)
}
```
```bash
Original Array:  [100 2 3]
Copied Array:  [100 2 3]
```