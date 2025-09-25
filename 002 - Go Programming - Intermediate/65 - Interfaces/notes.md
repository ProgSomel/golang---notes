# 65 - Interfaces
In Go, interfaces provide a way to specify behaviors.
They define a set of method signatures that a type must implement to satisfy the interface.
Interfaces promote code reuse, decoupling, and polymorphism without relying on explicit inheritance.

As was the case with structs and methods, interfaces are also declared outside the main function.

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * (math.Pi * c.radius)
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measur(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measur(r)
	measur(c)

}
```
```bash
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	//? Undefined methods
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * (math.Pi * c.radius)
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r1)

}
```
```bash
cannot use r1 (variable of struct type rect1) as geometry value in argument to measure: rect1 does not implement geometry (missing method perim)
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	//? Undefined methods
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (r1 rect1) perim() float64 {
	return 2 * (r1.height + r1.width)
}

func (c circle) perim() float64 {
	return 2 * (math.Pi * c.radius)
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r1)

}
```
```bash
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
{3 4}
12
14
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func main(){
	myPrinter(1, "John", 45.9, true)
}
```
```bash
1
John
45.9
true
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func myPrinterType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: Unknown")
	}
}

func main(){
	myPrinterType(9)
	myPrinterType("John")
	myPrinterType(false)
}
```
``` bash
Type: Int
Type: string
Type: Unknown
```

**----------------------------------------------------------------------------------------------------------------------**

In Summary, interfaces in Go facilitate polymorphism and enable writing modular, testable and maintainable code by promoting loose coupling between types.