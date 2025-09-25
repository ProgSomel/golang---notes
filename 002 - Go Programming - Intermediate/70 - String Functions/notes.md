# 70 - String Functions
```go
package main

import "fmt"

func main(){

	str := "Hello Go!"
	fmt.Println(len(str))
}
```
```bash
9
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

}
```
```bash
Hello World
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result[0])

}
```
```bash
72 //? ASCII Value
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result[1:])

}
```
```bash
ello World
```

**-------------------------------------------------------------------------------------------------------------------------**

## Standard Library Functions
### String Conversion --> strconv.Itoa(variable)
```go
package main

import (
	"fmt"
	"strconv"
)

func main(){
	
	num := 18
	str := strconv.Itoa(num) //? convert to string(ASCII Value)
	fmt.Println(str)

}
```
```bash
18 //? ASCII Value
```

-------------------------------------------------------------------------------------------------------------------------

### String Splitting --> strings.Split(variableName, "operaton") --> returns an Slice
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)
}
```
```bash
[apple  orange  banana]
```

-------------------------------------------------------------------------------------------------------------------------

### String Join --> strings.Join() --> concatenates elements of a slice into a single string with a sepatator.
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, "-")
	fmt.Println(joined)
}
```
```bash
Germany-France-Italy
```

-------------------------------------------------------------------------------------------------------------------------

### checks stirng contains a subset or characters. It can be one character or multiple characters.
### strings.Contains() --> return true / false
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	str := "Hello Go!"
	fmt.Println(strings.Contains(str, "Go"))
}
```
```bash
true
```

-------------------------------------------------------------------------------------------------------------------------

### strings.Replace() --> replaces the occurrence of a substring within a string with a another substring
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	str := "Hello Go!"
	fmt.Println(str)
	replaced := strings.Replace(str, "Go", "World", 1) //? Here, 1 means how many occurances of Go word
	fmt.Println(replaced)
}
```
```bash
Hello Go!
Hello World!
```

-------------------------------------------------------------------------------------------------------------------------

### Trimming leading and trailing whitespaces from the string --> strings.TrimSpace()
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	str := " Hello Everyone "
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))
}
```
```bash
 Hello Everyone 
Hello Everyone
```

-------------------------------------------------------------------------------------------------------------------------

### strings.ToLower(), strings.ToUpper()
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	str := " Hello Everyone "
	fmt.Println(str)
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
}
```
```bash
 Hello Everyone 
 hello everyone 
 HELLO EVERYONE
 ```


-------------------------------------------------------------------------------------------------------------------------

### strings.Repeat() --> repeats something a fixed number of times.
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println(strings.Repeat("Foo", 3))
}
```
```bash
FooFooFoo
```

-------------------------------------------------------------------------------------------------------------------------

### we can also count the occurances of an alphabet or a substring inside another string.
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println(strings.Count("Hello Hello world", "Hello"))
}
```
```bash
2
```

-------------------------------------------------------------------------------------------------------------------------

### Checking suffix and prefix in a string, strings.HasSuffix(), strings.HasPrefix().
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
}
```
```bash
true
false
```

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "Lo"))
}
```
```bash
true
false
```

**-------------------------------------------------------------------------------------------------------------------------**

## Regular Expression
Go offers us a Regular Expression package which allows pattern matching and manipulation of strings based on complex rules.
```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	
	str1 := "Hello, 123 Go!"
	// d for digits, + for one or more multiple digits
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str1, -1) //? -1 for all the matches //? will return a slice
	fmt.Println(matches)


}
```
```bash
[123]
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	
	str1 := "Hello, 123 Go 11!"
	// d for digits, + for one or more multiple digits
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str1, -1) //? -1 for all the matches //? will return a slice
	fmt.Println(matches)
}
```
```bash
[123 11]
```

**-------------------------------------------------------------------------------------------------------------------------**

## Unicode UTF Eight Package --> Let us work on Unicode characters and strings.
```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	
	str := "Hello, "
	fmt.Println(utf8.RuneCountInString(str))
}
```
```bash
7
```

**-------------------------------------------------------------------------------------------------------------------------**

## string Builder
```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var builder strings.Builder
	//? write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//? convert builder to string
	result := builder.String()
	fmt.Println(result)
}
```
```bash
Hello, World!
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var builder strings.Builder
	//? write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//? convert builder to string
	result := builder.String()
	fmt.Println(result)

	//? using Writerune to add a character 
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result = builder.String()
	fmt.Println(result)
}
```
```bash
Hello, World!
Hello, World! How are you?
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var builder strings.Builder
	//? write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//? convert builder to string
	result := builder.String()
	fmt.Println(result)

	//? using Writerune to add a character 
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result = builder.String()
	fmt.Println(result)

	builder.Reset()
	builder.WriteString("Starting Fresh")
	result = builder.String()
	fmt.Println(result)
}
```
```bash
Hello, World!
Hello, World! How are you?
Starting Fresh
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var builder strings.Builder
	//? write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//? convert builder to string
	result := builder.String()
	fmt.Println(result)

	//? using Writerune to add a character 
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result = builder.String()
	fmt.Println(result)

	builder.Reset()
	builder.WriteString("Starting Fresh")
	fmt.Println(builder.String())
}
```
```bash
Hello, World!
Hello, World! How are you?
Starting Fresh
```