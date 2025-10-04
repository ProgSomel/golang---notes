# 98 - Type Conversions
In Go, **type Conversion** is a fundamental concept, that allows you to convert a value of one type to another.

This is especially useful when you need to ensure that values are in the correct format for various operations or functions.

## Numeric Type Conversion
```go
package main

func main() {
	//! Type(value)
	var a int = 32;
	b := int32(a)
	c := float64(b)
	e := 3.14
	f := int(e)
}
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main() {
	//! Type(value)
	g := "Hello"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
}
```
```bash
[72 101 108 108 111]
```

-------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main() {
	//! Type(value)
	i := []byte{255, 72}
	j := string(i)
	fmt.Println(j)
}
```
```bash
�H
```

**Byte slice can not be go over 255**
```go
package main

import "fmt"

func main() {
	//! Type(value)
	i := []byte{256, 72} //? cannot use 256 (untyped int constant) as byte value in array or slice literal (overflows)
	j := string(i)
	fmt.Println(j)
}
```