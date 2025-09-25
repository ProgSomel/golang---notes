# 59 - Pointers
- A pointer is a variable that stores the memory address of another variable.
- Use Cases
  - Modify the value of a variable indirectly
  - Pass large data structures efficiently between functions
  - Manage memory dirctly for performance reasons

```go
package main

import "fmt"

func main(){
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(ptr)
	
}
```
```bash
10
0xc0000100b8
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	// var ptr *int
	var a int = 10
	// ptr = &a

	fmt.Println(a)
	fmt.Println(&a)
	
}
```
```bash
10
0xc0000100b8
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(*ptr)
	
}
```
```bash
10
10
```

**-----------------------------------------------------------------------------------------------------------------------**

## Now, a pointer that does not point to any memory address is a null pointer. So the zero value of a pointer is nil

A nil pointer is often used to indicate that a pointer is not currently valid or currently not initialized.

```go
package main

import "fmt"

func main(){
	var ptr *int
	var a int = 10
	// ptr = &a

	fmt.Println(a)
	// fmt.Println(*ptr)
	
	if ptr == nil{
		fmt.Println("Pointer is nil")
		fmt.Println(ptr)
	}
}
```
```bash
10
Pointer is nil
<nil>
```

**-----------------------------------------------------------------------------------------------------------------------**

## We can pass pointer to functions
```go
package main

import "fmt"

func main(){
	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	// fmt.Println(*ptr)
	
	// if ptr == nil{
	// 	fmt.Println("Pointer is nil")
	// 	fmt.Println(ptr)
	// }

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int){
	*ptr++
}
```
```bash
10
11
```

