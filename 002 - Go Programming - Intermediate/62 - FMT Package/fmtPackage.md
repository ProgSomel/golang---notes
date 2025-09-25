# 62 - FMT Package
The FMT Package in Go is a fundamental package that provides formatted input/output functions.
It is widely used for printing and formatting strings, numbers, and other data types.

- Printing Functions
  - Print()
  - Println()
  - Printf()

- Formatting Functions
  - Sprint()
  - Sprintf()
  - Sprintln()

- Scanning Functions
  - Scan()
  - Scanf()
  - Scanln()

---------------------------------------------------------------------------------------------------------------------

## Printing Functions
```go
package main

import "fmt"


func main(){

	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Print(12,456)

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println(12,456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Name: %b, Age: %X\n", age, age)

}
```
```bash
HelloWorld!12 456Hello
World!
12 456
Name: John, Age: 25
Name: 11001, Age: 19
```

**---------------------------------------------------------------------------------------------------------------------**

## Formatting Functions
```go
package main

import "fmt"


func main(){

	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Println(s)

}
```
```bash
Hello World!123 456
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){

	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Println(s)

}
```
```bash
HelloWorld!123 456
Hello World! 123 456
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){

	name := "Somel"
	age := 25

	sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	fmt.Print(sf)

}
```
```bash
Name: Somel, Age 25%
```

**---------------------------------------------------------------------------------------------------------------------**

## Scanning Functions
```go
package main

import "fmt"


func main(){
	
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&name, &age)
	fmt.Printf("Name: %s, Age: %d", name, age)

}
```
```bash
Enter your name and age:somel 25
Name: somel, Age: 25%
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	fmt.Printf("Name: %s, Age: %d", name, age)

}
```
```bash
Enter your name and age:somel

```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	// fmt.Scan(&name, &age)
	fmt.Scanln(&name, &age)
	fmt.Printf("Name: %s, Age: %d", name, age)

}
```
```bash
Enter your name and age:somel
Name: somel, Age: 0%
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	
	var name string
	var age int

	fmt.Print("Enter your name and age:")
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d", name, age)

}
```

**---------------------------------------------------------------------------------------------------------------------**

## Error Formatting Functions
```go
package main

import "fmt"


func main(){
	
	err := checkAge(15)
	if err != nil{
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error{
	if age < 18{
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
```
```bash
Error:  Age 15 is too young to drive.
```