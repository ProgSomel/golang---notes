# 68 - Errors
Errors are fundamental part of any programming language.
Allowing programs to handle exceptional conditions gracefully.

In Go, errors are represented by the error interface, which is a built in type used to indicate 
the presence of an error condition.

Errors are typically created using the errors package or by implementing the error interface.

```go
package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: square root of negative number")
	}
	return 1, nil
}

func main() {
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	result1, err1 := sqrt(-16)

	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(result1)

}
```
```bash
1
Math: square root of negative number
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: square root of negative number")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty Data")
	}
	return nil
}

func main() {
	
	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data process Successfully")

}
```
```bash
Error:  Error: Empty Data
```

**------------------------------------------------------------------------------------------------------------------------**

## custom Errors
```go
package main

import (
	"fmt"
)

type myError struct {
	message string
}

func (m * myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func process() error {
	return &myError{message: "Custom Error Message"}
}

func main() {
	err := process()
	fmt.Println(err)
}
```
```bash
Error: Custom Error Message
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

type myError struct {
	message string
}

func (m * myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func process() error {
	return &myError{"Custom Error Message"}
}

func main() {
	err := process()
	fmt.Println(err)
}
```
```bash
Error: Custom Error Message
```

**------------------------------------------------------------------------------------------------------------------------**

## fmt.errorf
```go
package main

import (
	"errors"
	"fmt"
)

type myError struct {
	message string
}

func readConfig() error {
	return errors.New("Config Error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func main() {
	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read Successfully")
}
```
```bash
readData: Config Error
```