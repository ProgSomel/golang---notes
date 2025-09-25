# 024 - Multiple Return Values
```go
package main

import "fmt"

//! func functionName(parameter1 type1, paramater2 type2, ...) (returnType1, returnType2){}
func divide(a, b int) (int, int){
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main(){

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d", q, r)

}
```
```bash
Quotient: 3, Remainder: 1
```

**---------------------------------------------------------------------------------------------------------**

## return errors along with return value
```go
package main

import (
	"errors"
	"fmt"
)

//! func functionName(parameter1 type1, paramater2 type2, ...) (returnType1, returnType2){}
func divide(a, b int) (int, int){
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error){
	if(a > b){ 
		return "a is greater than b", nil
	}else if(b > a){
		return "b is greater than a", nil
	}else{
		return "", errors.New("Unable to compare which is greater")
	}
}

func main(){

	result, err :=	compare(3, 2)
	if err != nil{
		fmt.Println("Error: ", err)
	}else{
		fmt.Println("Result: ", result)
	}

}
```
```bash
Result:  a is greater than b
```

---------------------------------------------------------------------------------------------------------

```go
package main

import (
	"errors"
	"fmt"
)

//! func functionName(parameter1 type1, paramater2 type2, ...) (returnType1, returnType2){}
func divide(a, b int) (int, int){
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error){
	if(a > b){ 
		return "a is greater than b", nil
	}else if(b > a){
		return "b is greater than a", nil
	}else{
		return "", errors.New("Unable to compare which is greater")
	}
}

func main(){

	result, err :=	compare(2, 2)
	if err != nil{
		fmt.Println("Error: ", err)
	}else{
		fmt.Println("Result: ", result)
	}

}
```
```bash
Error:  Unable to compare which is greater
```

**---------------------------------------------------------------------------------------------------------**

## named return Values
```go
package main

import (
	"fmt"
)

//! func functionName(parameter1 type1, paramater2 type2, ...) (returnType1, returnType2){}
func divide(a, b int) (quotient int, remainder int){
	quotient = a / b
	remainder = a % b
	return
}

func main(){
	quotient, remainder := divide(4, 4)
	fmt.Println(quotient, remainder)

}
```
```bash
1 0
```
