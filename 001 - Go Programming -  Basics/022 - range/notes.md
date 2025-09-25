# 022 - Range Keyword
The range Keyword in go, provides a convenient way to iterate over various data structures like arrays, slices, map and channels. It simplifies the process of iterating and accessing elements without needing to deal directly with indices or iterators. 

## range over strings
```go
package main

import "fmt"

func main(){
	message := "Hello World"
	for i, v := range message{
		fmt.Println(i, v)
	}
}
```
```bash
0 72
1 101
2 108
3 108
4 111
5 32
6 87
7 111
8 114
9 108
10 100
```
**Here, we get index value and unicode value of characters**

------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	message := "Hello World"
	for i, v := range message{
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}
```
```bash
0 72
Index: 0, Rune: H
1 101
Index: 1, Rune: e
2 108
Index: 2, Rune: l
3 108
Index: 3, Rune: l
4 111
Index: 4, Rune: o
5 32
Index: 5, Rune:  
6 87
Index: 6, Rune: W
7 111
Index: 7, Rune: o
8 114
Index: 8, Rune: r
9 108
Index: 9, Rune: l
10 100
Index: 10, Rune: d
```

**------------------------------------------------------------------------------------------------------**

## Now there is somthing that we need to consider while using range
Range Keyword operates on a copy of the data structure it iterates over, therefore modifying index or value inside the loop does not affect the original data structure.

**------------------------------------------------------------------------------------------------------**

## Now discuss how range behaves with different types
For arrays, slices and strings range iterates in order from the first element to the last. For maps, range iterates over key value pairs, but in an unspecified order, and for channels range iterates untill the channel is closed. So if the channel is not closed, range will keep iterating over that channel.

So in conclusion, the range Keyword in go provides a clean and efficient way to iterate over data structures. It enhance readability and simplifies code while maintaining efficieny.