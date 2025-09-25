# 73 - Regular Expressions
Regular Expressions provide a powerful way to search, maipulate, and validate text strings based on patterns.

In Go, Regular Expressions package provides support for working with Regular Expressions.

To work with a regex pattern in Go, you first compile it using the regexp.Compile or regexp.MustCompile

```go
package main

import "fmt"

func main(){
	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)
}
```
```bash
He said, "I am great"
He said, "I am great"
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	//? compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`) //? + means multiple occurance. {2,} means 2 or more

	//? test string 
	email1 := "user@email.com"
	email2 := "invalid_email"
	
	//? Match
	fmt.Println("Email1: ", re.MatchString(email1))
	fmt.Println("Email2: ", re.MatchString(email2))
}
```
```bash
Email1:  true
Email2:  false
```

**-----------------------------------------------------------------------------------------------------------------------**

## Capturing Groups
```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	//! capturing Groups 
	//? Compile a regex pattern to capture date components
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	// Test string
	date := "2025-09-07"
	subMatches := re.FindStringSubmatch(date)
	fmt.Println(subMatches)
	fmt.Println(subMatches[0])
	fmt.Println(subMatches[1])
	fmt.Println(subMatches[2])
	fmt.Println(subMatches[3])
}
```
```bash
[2025-09-07 2025 09 07]
2025-09-07
2025
09
07
```

**-----------------------------------------------------------------------------------------------------------------------**

## replace character in out target string
```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	//? target string
	str := "Hello World"
	re := regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)
}
```
```bash
H*ll* W*rld
```

**-----------------------------------------------------------------------------------------------------------------------**

## flag and options
### flag
- i -> case insensitive
If we use i in our regular expression that means we are accepting alphabets in all cases.
```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	re := regexp.MustCompile(`(?i)go`) //? ?i --> means case insensitive
	text := "Golang is great"

	fmt.Println("Match:", re.MatchString(text))
}
```
```bash
Match: true
```

```go
package main

import (
	"fmt"
	"regexp"
)

func main(){
	re := regexp.MustCompile(`go`)
	text := "Golang is great"

	fmt.Println("Match:", re.MatchString(text))
}
```
```bash
Match: false
```

- multi line model
- s - dot matches all

