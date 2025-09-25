# 020 - Slices
Slices are dynamic, flexible views into arrays.
They provide a more powerful and convenient interface to sequence of data compared to arrays.

**Slices are references to underlying arrays. They do not store any data themselves but provide a window into the array's elements.**

Slice can grow and shrink dynamically.

```go
package main

func main(){
	//! Slice do not have fixed length
	var numbers []int
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{9, 8, 7}
	
}
```

Slice has a cap function, which can check the capacity of the slice. It will check the number of Elements in the underlying array, starting from the slices first element.

**-------------------------------------------------------------------------------------------------------**

## we can also make slice, using make() function
```go
package main

func main(){
	//! Slice do not have fixed length
	var numbers []int
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{9, 8, 7}

	slice := make([]int, 5) //? length: 5, Capacity: 5


}
```

**-------------------------------------------------------------------------------------------------------**

## We can make a slice from an Existing Array
```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:4]

	fmt.Println(slice)


}
```
```bash
[2 3 4]
```

---------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:5]

	fmt.Println(slice)


}
```
```bash
[2 3 4 5]
```

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:len(a)] //? should omit second index in slice, s[a:len(s)] is identical to s[a:]

	fmt.Println(slice)


}
```
```bash
[2 3 4 5]
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	fmt.Println(slice)


}
```
```bash
[2 3 4 5]
```

**---------------------------------------------------------------------------------------------------------**

## we can also add values using to slice using append
```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	fmt.Println(slice)

	slice = append(slice, 100, 200)

	fmt.Println(slice)


}
```
```bash
[2 3 4 5]
[2 3 4 5 100 200]
```


**---------------------------------------------------------------------------------------------------------**

## Coping a slice
```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	fmt.Println("Slice: ", slice)

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	fmt.Println("SliceCopy: ", sliceCopy)



}
```
```bash
Slice:  [2 3 4 5]
SliceCopy:  [2 3 4 5]
```

**---------------------------------------------------------------------------------------------------------**

## Slices also have a concept of nil slices. A nil slice has a zero value and does not reference any underlying array. 
```go
package main

import "fmt"

func main(){

	var nilSlice []int
	fmt.Println(nilSlice)


}
```
```bash
[]
```

**---------------------------------------------------------------------------------------------------------**

## iterating over slice
```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	for i, v := range slice{
		fmt.Println(i, v)
	}
}
```
```bash
0 2
1 3
2 4
3 5
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	for i, v := range slice{
		fmt.Println(i, v)
	}

	fmt.Println("Element at Index 3 of slice1: ", slice[3])

	slice[3] = 50

	fmt.Println("Element at Index 3 of slice1: ", slice[3])
}
```
```bash
0 2
1 3
2 4
3 5
Element at Index 3 of slice1:  5
Element at Index 3 of slice1:  50
```

**---------------------------------------------------------------------------------------------------------**

## Comparing two slices
```go
package main

import (
	"fmt"
	"slices"
)

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	fmt.Println("Slice: ", slice)

	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	fmt.Println("SliceCopy: ", sliceCopy)

	if slices.Equal(slice, sliceCopy){
		fmt.Println("Slice is Equal to sliceCopy")
	}

}
```
```bash
Slice:  [2 3 4 5]
SliceCopy:  [2 3 4 5]
Slice is Equal to sliceCopy
```

**---------------------------------------------------------------------------------------------------------**

## Multidimensional Slice
```go
package main

import (
	"fmt"
)

func main(){

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++{
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

}
```

**---------------------------------------------------------------------------------------------------------**

## slice operator. slice[low:high]

```go
package main

import (
	"fmt"
)

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:]

	slice2 := slice[2:4]
	fmt.Println(slice2)

}
```
```bash
[4 5]
```

--------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){

	a := [5]int{1, 2, 3, 4, 5}

	slice := a[1:] //? 2 3 4 5

	slice2 := slice[2:4] //? 4 5
	fmt.Println(slice2)

	fmt.Println("Length of slice2: ", len(slice2))
	fmt.Println("The Capacity of slice2: ", cap(slice2))

}
```
```bash
[4 5]
Length of slice2:  2
The Capacity of slice2:  2
```
