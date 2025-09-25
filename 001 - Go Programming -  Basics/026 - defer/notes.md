# 026 - Defer
In Go, defer is a mechanism that allows you to postpone the execution of a function untill the surrounding function returns.
The Defer function always part of another function. The surrounding function means the function that encloses the defer function. The defer function is always encloseed in another function.
```go
package main

import "fmt"

func process(){
	defer fmt.Println("Deffered statement executed")
	fmt.Println("Normal execution statement")
}

func main(){
	process()

}
```
```bash
Normal execution statement
Deffered statement executed
```

------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func process(){
	defer fmt.Println("First Deffered statement executed")
	defer fmt.Println("Second Deffered statement executed")
	defer fmt.Println("Third Deffered statement executed")
	fmt.Println("Normal execution statement")
}

func main(){
	process()

}
```
```bash
Normal execution statement
Third Deffered statement executed
Second Deffered statement executed
First Deffered statement executed
```

**------------------------------------------------------------------------------------------------------------**

## arguments to defered functions are evaluated immediately when the defer statement is encountered.
## So just because defer function is getting executed at the end, does not mean that it is getting evaluated at the end. Defer function will be evaluated immediately as soon as it is encountered.
## So if it is encountered in the beginning of the function of the enclosing function, then it will be evaluated in the beginning of the enclosing function.
```go
package main

import "fmt"

func process(i int){
	defer fmt.Println("Deffered i value: ", i)
	defer fmt.Println("First Deffered statement executed")
	defer fmt.Println("Second Deffered statement executed")
	defer fmt.Println("Third Deffered statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i is: ", i)
}

func main(){
	process(10)

}
```
```bash
Normal execution statement
Value of i is:  11
Third Deffered statement executed
Second Deffered statement executed
First Deffered statement executed
Deffered i value:  10
```

