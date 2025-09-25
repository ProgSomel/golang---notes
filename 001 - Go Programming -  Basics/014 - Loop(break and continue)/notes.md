# Loop - Break and Condition

## Simple iteration over a range
```go
package main

import "fmt"


func main(){
	for i:=1; i <=5; i++{
		fmt.Println(i);
	}
}
```

**--------------------------------------------------------------------------------------------------------------------------------------------**

## Iteration over collection
```go
package main

import "fmt"


func main(){
	//? Iterate over collection 
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers{
		fmt.Printf("Index: %d, Value: %d\n", index, value);
	}
}
```

**--------------------------------------------------------------------------------------------------------------------------------------------**

## Break(terminate the loop), Continue(skip the current iteration and we move to next iteration of the loop)

```go
package main

import "fmt"


func main(){
	for i:=1; i <=10; i++{
		if i%2 == 0{
			continue
		}
		fmt.Println("Odd Number ", i)
		if i == 5{
			break
		}
	}
}
/*
Odd Number  1
Odd Number  3
Odd Number  5
*/
```