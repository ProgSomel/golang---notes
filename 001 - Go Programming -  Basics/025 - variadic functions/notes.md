# 025 - Variadic Function
Variadic function in go allow you to create function that can accept a variable number of arguments.
```go
package main

import "fmt"



func sum(nums ...int) int{
	total := 0
	for _, v := range nums{
		total += v
		}
	return total
}

func main(){
	fmt.Println("Sum of 1, 2, 3: ", sum(1, 2, 3))
}
```
```bash
Sum of 1, 2, 3:  6
```

-------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"



func sum(returnString string, nums ...int) (string, int){
	total := 0
	for _, v := range nums{
		total += v
		}
	return returnString, total
}

func main(){
	statement, total := sum("The sum of 1, 2, 3 is ", 1, 2, 3)
	fmt.Println(statement, total)
}
```
```bash
The sum of 1, 2, 3 is  6
```

---------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"



func sum(sequence int, nums ...int) (int, int){
	total := 0
	for _, v := range nums{
		total += v
		}
	return sequence, total
}

func main(){
	statement, total := sum(1, 10, 20, 30, 40)
	fmt.Println("Statement; ", statement, "Total: ", total)
}
```
```bash
Statement;  1 Total:  100
```

**---------------------------------------------------------------------------------------------------------**

## passing slices to a variadic function
```go
package main

import "fmt"

func sum(sequence int, nums ...int) (int, int){
	total := 0
	for _, v := range nums{
		total += v
		}
	return sequence, total
}

func main(){
	numbers := []int{1, 2, 3, 4, 5}
	sequence, total := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence, "Total: ", total)
}
```
```bash
Sequence:  3 Total:  15
```