# 66 - Struct Embedding
Struct Embedding allows a struct to inherit fields and methods from another struct type.
It is powerful mechanism for code reuse and structuring data.

## Field Inheritance
```go
package main

import "fmt"

type person struct {
	name string
	age int
}

type Employee struct {
	person //? Embedded Struct
	employeeID string
	salary float64
}

func main() {
	emp := Employee {
		person: person{name: "John", age: 30},
		employeeID: "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name) //? Accessing the Embedded struct field emp.person.name
	fmt.Println("Age: ", emp.age) //? Same as above
	fmt.Println("Emp ID: ", emp.employeeID)
	fmt.Println("Salary: ", emp.salary)
}
```
```bash
Name:  John
Age:  30
Emp ID:  E001
Salary:  50000
```

**--------------------------------------------------------------------------------------------------------------------------**

## Method Inheritance
```go
package main

import "fmt"

type person struct {
	name string
	age int
}

type Employee struct {
	person //? Embedded Struct
	employeeID string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func main() {
	emp := Employee {
		person: person{name: "John", age: 30},
		employeeID: "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name) //? Accessing the Embedded struct field emp.person.name
	fmt.Println("Age: ", emp.age) //? Same as above
	fmt.Println("Emp ID: ", emp.employeeID)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
}
```
```bash
Name:  John
Age:  30
Emp ID:  E001
Salary:  50000
Hi, I'm John and I'm 30 years old.
```

**--------------------------------------------------------------------------------------------------------------------------**

## Overidding Methods
Methods can be overidden by rededining them in the outer struct.
```go
package main

import "fmt"

type person struct {
	name string
	age int
}

type Employee struct {
	person //? Embedded Struct
	employeeID string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce(){
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.employeeID, e.salary)
}

func main() {
	emp := Employee {
		person: person{name: "John", age: 30},
		employeeID: "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.name) //? Accessing the Embedded struct field emp.person.name
	fmt.Println("Age: ", emp.age) //? Same as above
	fmt.Println("Emp ID: ", emp.employeeID)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
	emp.introduce()
}
```
```bash
Name:  John
Age:  30
Emp ID:  E001
Salary:  50000
Hi, I'm John, employee ID: E001, and I earn 50000.00.
Hi, I'm John, employee ID: E001, and I earn 50000.00.
```

**--------------------------------------------------------------------------------------------------------------------------**

## Embedded Struct with Named Field
Anynomous fields promote all fields and methods of the Embedded struct, while named fields require
Accessing fields with their explicit names.

```go
package main

import "fmt"

type person struct {
	name string
	age int
}

type Employee struct {
	employeeInfo person //? Embedded Struct Named Field
    // person //? Embedded Struct with Anynomous Field 
	employeeID string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce(){
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.employeeID, e.salary)
}

func main() {
	emp := Employee {
		employeeInfo: person{name: "John", age: 30},
		employeeID: "E001",
		salary: 50000,
	}

	fmt.Println("Name: ", emp.employeeInfo.name) //? Accessing the Embedded struct field emp.person.name
	fmt.Println("Age: ", emp.employeeInfo.age) //? Same as above
	fmt.Println("Emp ID: ", emp.employeeID)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()
	emp.introduce()
}
```
```bash
Name:  John
Age:  30
Emp ID:  E001
Salary:  50000
Hi, I'm John, employee ID: E001, and I earn 50000.00.
Hi, I'm John, employee ID: E001, and I earn 50000.00.
```