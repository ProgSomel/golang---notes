# Arithmetic Operations

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi float64 = 22/7;
	fmt.Println(pi)



}
```
```bash
Addition:  15
Substraction:  5
Multiplication:  50
Division:  2
Remainder:  0
3
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi float64 = 22/7.0;
	fmt.Println(pi)



}
```
```bash
Addition:  15
Substraction:  5
Multiplication:  50
Division:  2
Remainder:  0
3.142857142857143
```

------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//? variable declaration
	var a, b int = 10, 5

	var result int

	result = a + b;
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Substraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const pi int = 22/7.0;
	fmt.Println(pi)



}
```
```bash
cannot use 22 / 7.0 (untyped float constant 3.14286) as int value in constant declaration (truncated)
```

**--------------------------------------------------------------------------------------------------------------------------------------------**

# Overflow and underflow

## Overflow
```go
package main

import "fmt"

func main(){

	//! overflow with signed integers
	var maxInt int64 = 9223372036854775807 //? max value that int64 can hold
	fmt.Println(maxInt) //? 9223372036854775807

	maxInt = maxInt + 10 //? overflow eill happen; it will go to negative value; It wraps around to the minimum value.
	fmt.Println(maxInt) //? -9223372036854775799

}
```

----------------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){

	//! overflow with signed integers
	var maxInt int64 = 9223372036854775807 //? max value that int64 can hold
	fmt.Println(maxInt) //? 9223372036854775807

	maxInt = maxInt + 10 //? overflow eill happen; it will go to negative value; It wraps around to the minimum value.
	fmt.Println(maxInt) //? -9223372036854775799

	//! Overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 //? max value that uint64 can hold
	fmt.Println(uMaxInt) //? 18446744073709551615
	uMaxInt = uMaxInt+1 //? 0(lowest value); this is unsigned; so it can not go to negative value
	fmt.Println(uMaxInt)

}
```

**-----------------------------------------------------------------------------------------------------------------------------------------**

## Underflow

```go
package main

import (
	"fmt"
	"math"
)

func main(){

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat) //? 1e-323
	smallFloat = smallFloat/math.MaxFloat64
	fmt.Println(smallFloat) //? 0

}
```