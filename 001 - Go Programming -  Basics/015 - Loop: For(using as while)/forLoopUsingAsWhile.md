# For Loop Using as While Loop

## For loop as While Loopp
```go
package main

import "fmt"

func main(){
	i := 1;

	for i <= 5{
		fmt.Println("Iteration:", i)
		i++;
	}
}
```
```bash
//output
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
Iteration: 5
```

---------------------------------------------------------------------------------------------------------------------------------

## For loop as Infinite Loop
```go
package main

import "fmt"


func main(){

	for {
		fmt.Println("Hello")
	}
}
```

--------------------------------------------------------------------------------------------------------------------------------

## For Loop as While loop with break
```go
package main

import "fmt"


func main(){

	sum := 0

	for {
		sum += 10
		fmt.Println("Sum: ", sum)
		if(sum >= 50){
			break
		}
	}
}
```
```bash
//output
Sum:  10
Sum:  20
Sum:  30
Sum:  40
Sum:  50
```

--------------------------------------------------------------------------------------------------------------------------------------------

## For Loop as while loop with continue
```go
package main

import "fmt"


func main(){

	num := 1
	for(num <= 10){
		if(num%2==0){
			num++
			continue
		}
		fmt.Println("Odd Number: ", num)
		num++
	}
}
```
```bash
//output
Odd Number:  1
Odd Number:  3
Odd Number:  5
Odd Number:  7
Odd Number:  9
```

---------------------------------------------------------------------------------------------------------------------------------

## Guessing Game using for loop as while loop