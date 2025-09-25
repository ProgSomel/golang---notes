# 058 - Interface and Design pattern
**Paradigms:**
1. Structured Programming Language
2. Object Oriented Programming Language -> **Popular**
3. Functional Programming Language -> **More Popular**

## singleton Design pattern
There will be one object and that object will be shared among all.

**-------------------------------------------------------------------------------------------------------------------------------**

# Interface{}
**Abstraction** -> is a concept
**Interface** -> pure Abstraction, no details(no impelmentation only the signature) only concept
```go
package main

import "fmt"

// "ecommerce/cmd"

type People interface{

}

type user struct {
	Name string
	Age int
	Money float64
}

func printDetails(obj user) {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func main(){
	// cmd.Serve()
	usr1 := user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}

	printDetails(usr1)
	printDetails(usr2)
}
```
```bash
Name:  Somel Ahmed
Name:  25
Name:  10
Name:  Habibur Rahman
Name:  30
Name:  30
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

// "ecommerce/cmd"

type People interface{

}

type user struct {
	Name string
	Age int
	Money float64
}

func (obj user) printDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func main(){
	// cmd.Serve()
	usr1 := user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}
	usr1.printDetails()
	usr2.printDetails()
}
```
```bash
Name:  Somel Ahmed
Name:  25
Name:  10
Name:  Habibur Rahman
Name:  30
Name:  30
```

----------------------------------------------------------------------------------------------------------------------

# struct can implement interface{}
```go
package main

import "fmt"

// "ecommerce/cmd"

type People interface{
	PrintDetails()
	// ReceiveMoney(amount float64) float64
}

type user struct {
	Name string
	Age int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func main(){
	// cmd.Serve()
	var usr1 People
	usr1 = user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}
	usr1.PrintDetails()
	usr2.PrintDetails()
}
```
```bash
Name:  Somel Ahmed
Name:  25
Name:  10
Name:  Habibur Rahman
Name:  30
Name:  30
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

// "ecommerce/cmd"

type People interface{
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type user struct {
	Name string
	Age int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func main(){
	// cmd.Serve()
	var usr1 People
	usr1 = user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}
	usr1.PrintDetails()
	usr2.PrintDetails()
}
```
```bash
cannot use user{…} (value of struct type user) as People value in assignment: user does not implement People (missing method ReceiveMoney)compilerInvalidIfaceAssign
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

// "ecommerce/cmd"

type People interface{
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type user struct {
	Name string
	Age int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func (obj user) ReceiveMoney(amount float64) float64 {
	return obj.Money + amount
}

func main(){
	// cmd.Serve()
	var usr1 People
	usr1 = user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}
	usr1.PrintDetails()
	usr2.PrintDetails()
}
```
```bash
Name:  Somel Ahmed
Name:  25
Name:  10
Name:  Habibur Rahman
Name:  30
Name:  30
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
)

// "ecommerce/cmd"

type People interface{
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

type user struct {
	Name string
	Age int
	Money float64
}

func (obj user) PrintDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Name: ", obj.Age)
	fmt.Println("Name: ", obj.Money)
}

func (obj user) WithdrawMoney(amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

func (obj user) ReceiveMoney(amount float64) float64 {
	obj.Money = obj.Money + amount
	return obj.Money
}

func main(){
	// cmd.Serve()
	var usr1 People
	usr1 = user {
		Name: "Somel Ahmed",
		Age: 25,
		Money: 10.00,
	}

	usr2 := user {
		Name: "Habibur Rahman",
		Age: 30,
		Money: 30.00,
	}

	var usr3 BankUser
	usr3 = user {
		Name: "Rahimullah",
		Age: 100,
		Money: 200,
	}

	usr3.WithdrawMoney(10)

	obj, ok := usr3.(user) 
	if !ok {
		fmt.Println("Sorry usr3 is not type of user struct")
		os.Exit(1)
	}

	obj.PrintDetails()
	fmt.Println(obj.Name)


	usr1.PrintDetails()
	usr2.PrintDetails()
	usr1.ReceiveMoney(100)
}
```
```bash
Name:  Rahimullah
Name:  100
Name:  200
Rahimullah
Name:  Somel Ahmed
Name:  25
Name:  10
Name:  Habibur Rahman
Name:  30
Name:  30
```

----------------------------------------------------------------------------------------------------------------------

## Repository Design Pattern