# 67 - Generics

Generics in programming languages provide a way to write functions, data structures, and algorithms
that can handle various types without specifying each type explicitly.

This promotes code reuse, type safety and enhance the flexibility of programs.

In Go, As of version 1.18 has introduced support for generics, enabling developers to write more versatile and reusable code.

```go
package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)
}
```
```bash
2 1
```

---------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	x1, y1 := "John", "Jane"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)
}
```
```bash
2 1
Jane John
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return false;
	}
	return true
}

func (s Stack[T]) printAll() {
	if !s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}

	fmt.Println("Printing Stack Elements")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.printAll()
}
```
```bash
Printing Stack Elements
123
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return false;
	}
	return true
}

func (s Stack[T]) printAll() {
	if !s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}

	fmt.Println("Printing Stack Elements")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
}
```
```bash
Printing Stack Elements
123
3 true
Printing Stack Elements
12
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return true;
	}
	return false
}

func (s Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}

	fmt.Println("Printing Stack Elements")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty: ", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty:", intStack.isEmpty())
}
```
```bash
Printing Stack Elements
123
3 true
Printing Stack Elements
12
2 true
Is stack empty:  false
1 true
Is stack empty: true
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return true;
	}
	return false
}

func (s Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("The Stack is Empty")
		return
	}

	fmt.Println("Printing Stack Elements")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {
	
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty: ", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty:", intStack.isEmpty())

	fmt.Println("----------------------------------------------")

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("Somel")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack empty: ", stringStack.isEmpty())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack empty: ", stringStack.isEmpty())
}
```
```bash
Printing Stack Elements
123
3 true
Printing Stack Elements
12
2 true
Is stack empty:  false
1 true
Is stack empty: true
----------------------------------------------
Printing Stack Elements
HelloWorldSomel
Somel true
Is stringStack empty:  false
Printing Stack Elements
HelloWorld
World true
Hello true
Is stringStack empty:  true
```