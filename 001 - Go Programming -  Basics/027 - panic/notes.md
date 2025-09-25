# 027 - Panic
In Go, panic is a built in function that stops normal execution of a function immediately.
When a function encounters a panic, it stops executing its current activities, unwinds the stack, and then executes any deferred functions.
The syntax of panic function is called with an optional argument of any type, which represents the value associated with the panic. 
```go
package main

import "fmt"

func process(input int){
	if input < 0{
		panic("Input must be a non-negative number")
	}
	fmt.Println("Processing Input: ", input)
}

func main(){

	//! panic(interface{})
	process(10)

}
```
```bash
Processing Input:  10
```

-----------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func process(input int){
	if input < 0{
		panic("Input must be a non-negative number")
	}
	fmt.Println("Processing Input: ", input)
}

func main(){

	//! panic(interface{})
	process(-3)

}
```
```bash
panic: Input must be a non-negative number

goroutine 1 [running]:
main.process(0xc000002380?)
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/027 - panic/panic.go:7 +0x8f
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/027 - panic/panic.go:15 +0x1a
exit status 2
```

**-----------------------------------------------------------------------------------------------------------**

## panic with defer function
It stated that defer will execute when the function returns a value, but it will also execute even when the function is panicking. after defer function's executtion then the panic wll execute
```go
package main

import "fmt"

func process(input int){
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0{
		panic("Input must be a non-negative number") //! panic(interface{})
	}
	fmt.Println("Processing Input: ", input)
}

func main(){

	process(10)

}
```
```bash
Processing Input:  10
Deferred 2
Deferred 1
```

-----------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func process(input int){
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0{
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number") //! panic(interface{})
	}
	fmt.Println("Processing Input: ", input)
}

func main(){

	process(-10)

}
```
```bash
Before Panic
Deferred 2
Deferred 1
panic: Input must be a non-negative number

goroutine 1 [running]:
main.process(0xc000002380?)
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/027 - panic/panic.go:10 +0x165
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/027 - panic/panic.go:17 +0x1a
exit status 2
```
