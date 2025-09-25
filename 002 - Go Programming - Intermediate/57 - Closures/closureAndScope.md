# 57 - Closure and Scope
Closure work with lexical scoping, meaning they capture variables
from their surrounding context where they are defined.
This allows Closures to access variables even after the outer function has finished execution.

```go
package main

import "fmt"


func main(){

	sequence := adder()
	fmt.Println(sequence())

}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
```
```bash
Previous value of i:  0
added 1 to i
1
```

---------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
```
```bash
Previous value of i:  0
added 1 to i
1
added 1 to i
2
added 1 to i
3
added 1 to i
4
```

**------------------------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"


func main(){

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
    fmt.Println(sequence2())

}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
```
```bash
Previous value of i:  0
added 1 to i
1
added 1 to i
2
added 1 to i
3
added 1 to i
4
Previous value of i:  0
added 1 to i
1
```

**---------------------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"


func main(){

	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
    // fmt.Println(sequence2())

	substractor := func() func(int) int{
		countDown := 99
		return func(x int) int{
			countDown -= x
			return countDown
		}
	}()

	//? using the closure substracter
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
	fmt.Println(substractor(1))
}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
```
```bash
98
97
96
95
94
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){

	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
    // fmt.Println(sequence2())

	substractor := func() func(int) int{
		countDown := 99
		return func(x int) int{
			countDown -= x
			return countDown
		}
	}()

	//? using the closure substracter
	fmt.Println(substractor(1))
	fmt.Println(substractor(2))
	fmt.Println(substractor(3))
	fmt.Println(substractor(4))
	fmt.Println(substractor(5))
}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
```
```bash
98
96
93
89
84
```