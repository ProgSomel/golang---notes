# 80 - Bufio Package
The bufio package in Go, provides buffered input/output operations, which can significantly improve
performance when reading or writing data, especially for large volumes of data.

## reader
```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))
	//? reading byte byte
	data := make([]byte, 20)
	n, err := reader.Read(data) //? returns number of bytes and error
	if err!=nil {
		fmt.Println("Error reading: ", err)
	}
	fmt.Println(data) //? ASCII Value
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n') //? will store after 20
	if err!=nil {
		fmt.Println("Error reading string: ", err)
	}
	fmt.Println("Read string: ", line)

}
```
```bash
[72 101 108 108 111 44 32 98 117 102 105 111 32 112 97 99 107 97 103 101]
Read 20 bytes: Hello, bufio package
Read string:  !
```

**-----------------------------------------------------------------------------------------------------------------------**

## writer
### writing byte slice
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	writer := bufio.NewWriter(os.Stdout)

	//? write byte slice
	data := []byte("Hello, bufio package!\n")
	nn, err := writer.Write(data)
	if err!=nil {
		fmt.Println("Error writing: ", err)
	}
	fmt.Printf("Wrote %d bytes\n", nn)

	//? Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error writer: ", err)
		return
	}
}
```
```bash
Wrote 22 bytes
Hello, bufio package!
```

-----------------------------------------------------------------------------------------------------------------------

### writing string
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	str := "This is string.\n"
	n, err := writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes. \n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error")
	}
}
```
```bash
Wrote 16 bytes. 
This is string.
```