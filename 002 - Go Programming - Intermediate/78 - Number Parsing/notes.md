# 78 - Number Parsing
Number parsing is the process of converting textual representations of numbers into their
corresponding numeric values in Go.

Number parsing is typically done using functions from the string conversion package for basic types, ans specialized functions for specific needs.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num)
}
```

--------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num)

	//? converts a string to an integer with specified base and bit size
	num1, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num1) 
}
```
```bash
Parsed Integer:  12345
Parsed Integer:  12345
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num)

	//? converts a string to an integer with specified base and bit size
	num1, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value: ", err)
	}
	fmt.Println("Parsed Integer: ", num1)

	floatstr := "3.14"
	floatVal, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing value: ", err)
	}
	fmt.Printf("Parsed Float: %.2f\n", floatVal)
}
```
```bash
Parsed Integer:  12345
Parsed Integer:  12345
Parsed Float: 3.14
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	binaryStr := "1010" //? 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value: ", err)
		return
	}
	fmt.Println("Parsed binary to decimal value: ", decimal)
}
```
```bash
Parsed binary to decimal value:  10
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	hexStr := "FF"
	decimal, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hex value: ", err)
		return
	}
	fmt.Println("Parsed binary to decimal value: ", decimal)
}
```
```bash
Parsed binary to decimal value:  255
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	invalidNum := "456abc"
	invalidNumParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value: ", err)
		return
	}
	fmt.Println("Parsed invalid number: ", invalidNumParse)
}
```
```bash
Error parsing value:  strconv.Atoi: parsing "456abc": invalid syntax
```