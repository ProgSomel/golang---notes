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