# Defer ---> use LinkedList Data Structure
**There is a cell in stack frame which is called defer list pointer(it holds a address of closure which is stored in heap) actually defer use linkedlist**
In Go, defer is a keyword that delays the execution of a function until the surrounding function returns. when it will get the the return before that it defer will be executed. Runtime will execute it

```go
package main

import "fmt"

func a(){
	i := 0

	fmt.Println("First: ", i) //? 0

	defer fmt.Println("defer: ", i) //? 0

	i = i+1

	fmt.Println("Third: ", i) //? 0


	return
}

func main(){

	a()
}
```
```bash
First:  0
Third:  1
defer:  0
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func a(){

	i := 0

	fmt.Println("First: ", i) //? 0

	defer fmt.Println("defer: ", i) //? 0

	i = i+1

	fmt.Println("Third: ", i) //? 0

	defer fmt.Println("defer2: ", i) //? 1
	

	return

}

func main(){
	
	a()
}
```
```bash
First:  0
Third:  1
defer2:  1
defer:  0
```

**------------------------------------------------------------------------------------------------------------------------------**

## Named Return Values

```go
package main

import "fmt"

func sum(a int, b int) int{
	result := a + b

	return result
}

func main(){
	
	res := sum(3, 4)
	fmt.Println(res)
}
```
```bash
7
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func sum(a int, b int) (result int){
	result = a + b

	return result
}

func main(){
	
	res := sum(3, 4)
	fmt.Println(res)
}
```
```bash
7
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func sum(a int, b int) (result int){
	result = a + b

	return
}

func main(){
	
	res := sum(3, 4)
	fmt.Println(res)
}
```
```bash
7
```

-----------------------------------------------------------------------------------------------------------------------------

In 99% case, when closure created the outer function become deleted and closure will work with the copy of the variable(not address), But in the case of defer, when closure become created, both function outer function and closure will be presented and it will not use different copy of variable, it will work with same variable(with the reference value with the address of the variable).

```go
package main

import "fmt"

func calculate() (result int){
	fmt.Println("First: ", result) //? First: 0
	defer func(){
		result = result + 10
		fmt.Println(result) //? 15 // it will print at last
	}()

	result = 5
	fmt.Println("Second: ", result) //? Second: 5

	return //? 15
}

func main(){
	
	calculate()
}
```
```bash
First:  0
Second:  5
15
```

-------------------------------------------------------------------------------------------------------------------------------

In 99% case, when closure created the outer function become deleted and closure will work with the copy of the variable(not address), But in the case of defer, when closure become created, both function outer function and closure will be presented and it will not use different copy of variable, it will work with same variable(with the reference value with the address of the variable).

### Rules for Named return values
- all codes execute
- defer function store in magic box
- return -> will be executed all defer function
- will return named variable's values
```go
package main

import "fmt"

func calculate() (result int){
	fmt.Println("First: ", result) //? First: 0
	defer func(){
		result = result + 10
		fmt.Println("Defer: ",result) //? 15 // it will print at last
	}()

	result = 5
	fmt.Println("Second: ", result) //? Second: 5

	return
}

func main(){
	
	a := calculate()
	fmt.Println(a)
}
```
```bash
First:  0
Second:  5
Defer:  15
15
```

------------------------------------------------------------------------------------------------------------------------------

### Rules for just return types
- all codes execute
- defer function store in magic box
- return values are evaluated at this time(store the return value)
- all defer functions will be executed
- stored return value will be returned

```go
package main

import "fmt"

func calculate() int{
	result := 0
	fmt.Println("First: ", result) //? First: 0
	defer func(){
		result = result + 10
		fmt.Println("Defer: ",result) //? 15 // it will print at last
	}()

	result = 5
	fmt.Println("Second: ", result) //? Second: 5

	return result //? 5
}

func main(){
	
	a:= calculate()
	fmt.Println(a)
}
```
```bash
First:  0
Second:  5
Defer: 15
5
```

------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func calculate() (result int){

	fmt.Println("First: ", result) //? First: 0

	show:= func(){ //? show will be stored in heap. it will form a closure with result
		result = result + 10
		fmt.Println("Defer: ",result) //? 15 
	}

	defer show()

	result = 5

	p := func(a int){ //?it will be stored in heap with the argument a which is 5 from result. 
	                //? it will not form any closure because it is not using any outer scope variable likde show
		fmt.Println("Ami: ", a) //? 5
	}
	defer p(result)

	defer fmt.Println(result) //? it will stored in heap with argument result which is 5
	                          //? output: 5

	fmt.Println("Second: ", result) //? Second: 5

	defer fmt.Println(5) //? it will stored in heap with argument which is 5
	                     //? output is: 5


	return
}

func main(){
	
	a := calculate()
	fmt.Println(a)
}
```
```bash
First:  0
Second:  5
5
5
Ami:  5
Defer:  15
15
```