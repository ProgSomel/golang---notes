# 71 - String Formatting
String Formatting in Go, refers to the techniques used to create formatted output from variables or constants.

Go provides several mechanisms for formatting strings, including the FMT package, String interpolation, or format specifiers.

```go
package main

import "fmt"

func main(){
	num := 42
	fmt.Printf("%05d\n", num)
}
```
```bash
00042
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	num := 42345
	fmt.Printf("%05d\n", num)
}
```
```bash
42345
```

**-------------------------------------------------------------------------------------------------------------------------**

## string Allignment
```go
package main

import "fmt"

func main(){
	message := "Hello"
	//? string right allignment 
	fmt.Printf("|%10s|\n", message)
}
```
```bash
|     Hello|
```

--------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	message := "Hello"
	//? string right allignment 
	fmt.Printf("|%10s|\n", message)
	//? string left allignment
	fmt.Printf("|%-10s|\n", message)
}
```
```bash
|     Hello|
|Hello     |
```

**-------------------------------------------------------------------------------------------------------------------------**

## string Interpolation --> supports using Backticks
when you need to embed special character or multiple lines of text without interpreting escape sequence, Backticks are very useful.

This is particularly handy when dealing with regular expressions.