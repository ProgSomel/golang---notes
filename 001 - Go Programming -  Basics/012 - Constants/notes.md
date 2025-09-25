# Constant

constant must be initialized with values that can be determined at compile time
```go
package main

const PI = 3.14
const GRAVITY = 9.81 //? untyped constant


func main(){

	const days int = 7 //? typed constant

	//! const block 
	const(
		monday = 1
		tuesday = 2
		wednessday = 3
		thursday int = 4
	)
}
```