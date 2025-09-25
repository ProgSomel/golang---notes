# 084 - Reading Files
Reading files is a common operation in programming for tasks such as configuration loading, data parsing, or processing large data sets.

In Go, the OS and Buff IO package provide powerful functionalities to efficiently read data from files.

In order for us to read a file in Go, we first need to open the file and the OS package is used for file operation, including opening files.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()
	fmt.Println("File opened successfully")
}
```
```bash
File opened successfully
Closing open file
```

--------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()
	fmt.Println("File opened successfully")

	//? Read the contents of the opened file
	data := make([]byte, 1024) // Buffer to read data into
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println("Content of the file: ", string(data))
}
```
```bash
File opened successfully
Content of the file:  Hello World!

Closing open file
```

**--------------------------------------------------------------------------------------------------------------------------**

## Reading a File line by line
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()
	fmt.Println("File opened successfully")

	scanner := bufio.NewScanner(file)

	//? Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
}
```
```bash
File opened successfully
Line:  Hello World!
Line:  Hello Somel
Closing open file
```