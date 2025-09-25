# 69 - Custom Errors
```go
package main

import "fmt"

type customError struct {
	code int
	message string
}

// Implementing Error() method of error interface 
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

// Function that return a custom Error 
func doSomething() error {
	return &customError{ //? we are accessing memory address of our struct. And we are not using
						//? a duplicate copy so that we modify the properties, the values inside
						//? the actual struct which is outside of the scope of this function
						//? And thats why we are using ampersand sign to accesss the memory address.
		code: 500,
		message: "Something went wrong!",
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation Completed Successfully")
}
```
```bash
Error 500: Something went wrong!
```

-------------------------------------------------------------------------------------------------------------------------

The key difference is that if your Error() method needed to modify the struct, you'd need a pointer receiver. But since your method only reads the fields, both approaches work fine.
The automatic invocation of Error() happens because of Go's interface system - it's one of the elegant features that makes error handling in Go so clean!

```go
package main

import "fmt"

type customError struct {
	code int
	message string
}

// Implementing Error() method of error interface 
func (e customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func doSomething() error {
	return customError{ 
		code: 500,
		message: "Something went wrong!",
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation Completed Successfully")
}
```

**-------------------------------------------------------------------------------------------------------------------------**

## wrapped Errors --> Introduced after go version 1.13
```go
package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code int
	message string
	err error
}

// Implementing Error() method of error interface 
func (e customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.err)
}

// func doSomething() error {
// 	return customError{ 
// 		code: 500,
// 		message: "Something went wrong!",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code: 500,
			message: "Somthing went wrong",
			err: err,
		}
	}
	return nil

}

func doSomethingElse() error {
	return errors.New("Internal Error")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation Completed Successfully")
}
```
```bash
Error 500: Somthing went wrong, Internal Error
```