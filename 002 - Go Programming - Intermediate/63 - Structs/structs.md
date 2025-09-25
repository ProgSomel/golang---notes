# 63 - Structs
Structs in Go are composite data types that allow you to group together different types of variables
under a single name.

They are similar to classes in object oriented languages, but they are more lightweight and do not
support inheritance.

```go
package main

import "fmt"

func main(){
	type Person struct {
		firstName string
		lastName string
		age int
	}

	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
	}

	p1 := Person{
		firstName: "Jane",
		age: 25,
	}

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
}
```
```bash
John
Jane
```

**--------------------------------------------------------------------------------------------------------------------**

## Anynomous Struct
```go
package main

import "fmt"

func main(){
	
	user := struct{
		userName string
		email string
	}{
		userName: "Somel",
		email: "somelahmed",
	}

	fmt.Println(user)
}
```
```bash
{Somel somelahmed}
```

**--------------------------------------------------------------------------------------------------------------------**

## Struct Method
```go
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
	}
	
	fmt.Println(p.fullName())
}

func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}
```
```bash
John Doe
```

**--------------------------------------------------------------------------------------------------------------------**

## Pointer Receiver
To modify struct fields within a method we use a Pointer Receiver instead of a value Receiver.
```go
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
	}
	
	fmt.Println(p.fullName())
	p.incrementAgeByOne()
	fmt.Println(p.age)
}

func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
```
```bash
John Doe
31
```

**--------------------------------------------------------------------------------------------------------------------**

## Structs Embedding
```c++
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
		address Address
}

type Address struct{
	city string
	country string
}



func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "Dhaka",
			country: "Bangladesh",
		},
	}

	p2 := Person{
		firstName: "Jan",
		lastName: "Mia",
		age: 45,
	}
	p2.address.city = "New York"
	p2.address.country = "USA"
	
	fmt.Println(p)
	fmt.Println(p2)
}
```
```bash
{John Doe 30 {Dhaka Bangladesh}}
{Jan Mia 45 {New York USA}}
```

**--------------------------------------------------------------------------------------------------------------------**

## We can define structs with Anynomous fields as well
```go
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
		address Address
		PhoneHomeCell
}

type Address struct{
	city string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}



func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "Dhaka",
			country: "Bangladesh",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}

	p2 := Person{
		firstName: "Jan",
		lastName: "Mia",
		age: 45,
	}
	p2.address.city = "New York"
	p2.address.country = "USA"
	
	fmt.Println(p)
	fmt.Println(p2)
}
```
```bash
{John Doe 30 {Dhaka Bangladesh} {46576 4556778}}
{Jan Mia 45 {New York USA} { }}
```

**--------------------------------------------------------------------------------------------------------------------**

## struct comparing --> structs are comparable if all their fields are comparable
```go
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
		address Address
		PhoneHomeCell
}

type Address struct{
	city string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}



func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "Dhaka",
			country: "Bangladesh",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}

	p2 := Person{
		firstName: "Jan",
		lastName: "Mia",
		age: 45,
	}
	p2.address.city = "New York"
	p2.address.country = "USA"
	
	fmt.Println(p==p2)
}
```
```bash
false
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Person struct {
		firstName string
		lastName string
		age int
		address Address
		PhoneHomeCell
}

type Address struct{
	city string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}



func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func main(){
	
	p := Person {
		firstName: "John",
		lastName: "Doe",
		age: 30,
		address: Address{
			city: "Dhaka",
			country: "Bangladesh",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}

	p2 := Person{
		firstName: "John",
		lastName: "Doe",
		age: 30,
		PhoneHomeCell: PhoneHomeCell{
			home: "46576",
			cell: "4556778",
		},
	}
	p2.address.city = "Dhaka"
	p2.address.country = "Bangladesh"
	
	fmt.Println(p==p2)
}
```
```bash
true
```