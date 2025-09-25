# 60 - Strings and Runes
# Strings
String is a sequence of bytes. Bytes are unsigned int8 values(uint8). Sequence of bytes represent text.

Strings are immutable, meaning once created, their values can not be changed.
```go
package main

import "fmt"

func main(){
	message := "Hello, Go!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(rawMessage)
}
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	message := "Hello, Go!"
	rawMessage := "Hello\nGo"

	fmt.Println(message)
	fmt.Println(rawMessage)
}
```
```bash
Hello, Go!
Hello
Go
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	message := "Hello, Go!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)
}
```
```bash
Hello, Go!
Hello,  Go!
Go!lo, 
Hello\nGo
```

**-----------------------------------------------------------------------------------------------------------------**

## find the length of the string
```go
package main

import "fmt"

func main(){
	message := "Hello, Go!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of the message variable is: ", len(message))
}
```
```bash
Hello, Go!
Hello,  Go!
Go!lo, 
Hello\nGo
Length of the message variable is:  10
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	message := "Hello, Go!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of the message variable is: ", len(message))
	fmt.Println("Length of the message variable is: ", len(message1))
	fmt.Println("Length of the message variable is: ", len(message2))
	fmt.Println("Length of the message variable is: ", len(rawMessage))
}
```
```bash
Hello, Go!
Hello,  Go!
Go!lo, 
Hello\nGo
Length of the message variable is:  10
Length of the message variable is:  11
Length of the message variable is:  11
Length of the message variable is:  9
```

**-----------------------------------------------------------------------------------------------------------------**

## We can also extract any alphabet, any letter, any character from a string using the index number
## -- Accessing a character at a specific index returns its byte value.
```go
Hello, Go!
Hello,  Go!
Go!lo, 
Hello\nGo
Length of the message variable is:  10
The first character in message var is:  72 //? This is ASCII Value
```

**-----------------------------------------------------------------------------------------------------------------**

## string concatenation
```go
package main

import "fmt"

func main(){
	
	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting + name)
}
```
```bash
HelloAlice
```

**-----------------------------------------------------------------------------------------------------------------**

## strings can be compared using the relational operators
```go
package main

import "fmt"

func main(){
	
	str1 := "Apple"
	str2 := "Banana"
	fmt.Println(str1 < str2) //? Lexicographical comparsion
}
```
```bash
true
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Apple"
	// str2 := "Banana"
	str3 := "App"
	fmt.Println(str3 < str1)
}
```
```bash
true
```

--------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Apple"
	// str2 := "Banana"
	str3 := "app"
	fmt.Println(str3 < str1) //? Compare ASCII Values
}
```
```bash
false
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	str1 := "Apple"
	// str2 := "Banana"
	str3 := "apple"
	fmt.Println(str3 < str1)
}
```
```bash
false
```

**-----------------------------------------------------------------------------------------------------------------**

## String iteration --> just like iterating over a slice and an array
```go
package main

import "fmt"

func main(){
	
	message := "Hello, Go!" 
	
	for i, char := range message{
		fmt.Printf("Character at index %d is %c\n", i, char)
	}
}
```
```bash
Character at index 0 is H
Character at index 1 is e
Character at index 2 is l
Character at index 3 is l
Character at index 4 is o
Character at index 5 is ,
Character at index 6 is  
Character at index 7 is G
Character at index 8 is o
Character at index 9 is !
```

-------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	message := "Hello, Go!" 
	
	for _, char := range message{
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x\n", char) //? in Hexadecimal value of character
	}
}
```
```bash
48
65
6c
6c
6f
2c
20
47
6f
21
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	
	message := "Hello, Go!" 
	
	for _, char := range message{
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char) //? in ASCII Value. as rune in the uint8 format
	}
}
```
```bash
72
101
108
108
111
44
32
71
111
33
```

**-----------------------------------------------------------------------------------------------------------------**

## If we have a string and we want to count the number of runes in that atring, we also have another function, a built in function for UTF eight, and that is rune count in string

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	
	message := "Hello" 

	fmt.Println("Rune count: ", utf8.RuneCountInString(message))
}
```
```bash
Rune count:  5
```

**-----------------------------------------------------------------------------------------------------------------**

## String manipulation
### strings are immutable means operations like appending, replacing, or modifying require creating new strings.
```go
package main

import (
	"fmt"
)

func main(){
	
	name := "Somel"
	greeting := "Hello" 

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)
}
```
```bash
HelloSomel
```

**-----------------------------------------------------------------------------------------------------------------**

# Rune
A rune is an alias for int32 and it represents a unicode code point, a Unicode value.
So it is not a character, it is an integer value.
A rune is an integer value, and that value represents a Unicode code point and that will be converted into a character

**Runes are declared using single quotes, double quotes and backticks are for strings.**
**Backtics especially are for string literals**
**Double quotes are for the regular strings**
**Single quotes for runes**

```go
package main

import (
	"fmt"
)

func main(){
	
	var ch rune = 'a'
	fmt.Println(ch)//? ASCII Value
	 
}
```
```bash
97
```

--------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){
	
	var ch rune = 'a' 
	fmt.Println(ch) //? ASCII Value
	fmt.Printf("%c", ch)
}
```
```bash
97
a
```

**-----------------------------------------------------------------------------------------------------------------**

## Converts runes to strings
```go
package main

import (
	"fmt"
)

func main(){
	
	var ch rune = 'a' 
	cstr := string(ch)
	fmt.Println(cstr)
}
```
```bash
a //? type is string
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){
	
	var ch rune = 'a' 
	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T", cstr)
}
```
```bash
a
Type of cstr is string
```

**-----------------------------------------------------------------------------------------------------------------**

## iterate over runes
```go
package main

import (
	"fmt"
)

func main(){
	
	japanese := "こんにちは ありがとう"
	for _, runeValue := range japanese{
		fmt.Printf("%c\n", runeValue)
	}

}
```
```bash
こ
ん
に
ち
は
 
あ
り
が
と
う
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){
	
	japanese := "こんにちは ありがとう"
	for _, runeValue := range japanese{
		fmt.Printf("%v\n", runeValue) //? %v -> actual default value
	}

}
```
```bash
12371
12435
12395
12385
12399
32
12354
12426
12364
12392
12358
```

-----------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main(){
	
	l := '😆'
	fmt.Printf("%v\n", l)
	fmt.Printf("%c\n", l)

}
```
```bash
128518
😆
```

**-----------------------------------------------------------------------------------------------------------------**

## Runes in Go and Characters char in C/C++
- Both runes in Go and characters char in C are used to represent individual characters in strings.
- Both runes and characters char typically occupy a fixed amount of memory.
- Runes represent by int32 and occupy 4 bytes of memory, whereas Characters in C represent by char and occupy 1 byte of memory and capable of storing ASCII Characters directly.
- Runes are designed to handle Unicode Characters efficiently, they can represent any Unicode code point from ASCII to more complex Characters like emojis ans non-Latin as well. While C also supports Characters beyond ASCII through multibyte encodings like UTF eight, handling Unicode Characters directly is not as straightforward as in Go. C libraries ans implementations may vary in their support for Unicode