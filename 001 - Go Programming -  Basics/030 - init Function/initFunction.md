# 030 - Init Function
The init funciton is a special function that can be declared in any package.
It is used to perform initialization tasks for the package before it is used.

**init function has no parameters and no return values.**
**Go executes init functions automatically when the package is initialized. This happens before the main funciton is executed**

if there are multiple init functions, they execute sequentially following their textual order int the package file.

```go
package main

import "fmt"

func init(){
	fmt.Println("Initializing package...")
}

func main(){
	fmt.Println("Inside the main funciton")
}
```
```bash
Initializing package...
Inside the main funciton
```

**-----------------------------------------------------------------------------------------------------------------**

## Practical Use Cases
- Setup Tasks
- Configuration
- Registration Components
- Database Initialization

## Best Practices
- Avoid Side Effects
- Initialization Order
- Documentation

**-----------------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"

func init(){
	fmt.Println("Initializing package1...")
}

func init(){
	fmt.Println("Initializing package2...")
}

func init(){
	fmt.Println("Initializing package3...")
}

func main(){
	fmt.Println("Inside the main funciton")
}
```
```bash
Initializing package1...
Initializing package2...
Initializing package3...
Inside the main funciton
```