# 64 - Methods
```go
package main

import "fmt"

type Rectangle struct {
	length float64
	width float64
}

func(r Rectangle) Area() float64 {
	return r.length * r.width
}

func main(){
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of Rectangle with width 9 and length 10 is: ", area)
}
```
```bash
Area of Rectangle with width 9 and length 10 is:  90
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Rectangle struct {
	length float64
	width float64
}

//? Method with value receiver
func(r Rectangle) Area() float64 {
	return r.length * r.width
}

//? Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func main(){
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of Rectangle with width 9 and length 10 is: ", area)
	rect.Scale(2);
	area = rect.Area()
	fmt.Println("Area of Rectangle with factor 2 is: ", area)

}
```
```bash
Area of Rectangle with width 9 and length 10 is:  90
Area of Rectangle with factor 2 is:  360
```

**---------------------------------------------------------------------------------------------------------------------**

## Methods can be associated with any types
```go
package main

import "fmt"

type MyInt int
//? Method on a user-defined type
func (m MyInt) IsPositive() bool {
		return m > 0
}

func main(){

	num := MyInt(-5)
	num1 := MyInt(9)

	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())

}
```
```bash
false
true
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type MyInt int
//? Method on a user-defined type
func (m MyInt) IsPositive() bool {
		return m > 0
}

func (MyInt) welcomeMessage() string {
	return "welcome to MyInt Type"
}

func main(){

	num := MyInt(-5)
	num1 := MyInt(9)

	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())

}
```
```bash
false
true
welcome to MyInt Type
```

**---------------------------------------------------------------------------------------------------------------------**

## Methods with Embedding
Struct Embedding allows methods of an embedded struct to be promoted to the outer struct.
```go
package main

import "fmt"

type Shape struct {

	Rectangle
}

type Rectangle struct {
	length float64
	width float64
}

//? Method with value receiver
func(r Rectangle) Area() float64 {
	return r.length * r.width
}

func main(){
	
	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())

}
```
```bash
90
```