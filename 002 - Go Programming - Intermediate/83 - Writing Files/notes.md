# 83 - Writing Files
Writing to Files involves creating or opening a file, Writing data to it, and handling any errors that may occur during these operations.

We have OS packages that provide convenient methods of handling file operations.

## OS Package
The OS Package in go provides functions for operating system functionality, including file operations.

**-------------------------------------------------------------------------------------------------------------------------**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	//? Creating a file using OS Package
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating a file.", err)
		return
	}
	defer file.Close()

	//? write data to file
	data := []byte("Hello World!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to files: ", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")


}
```
```bash
Data has been written to file successfully.
```
```output.txt
Hello World!
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	//? Creating a file using OS Package
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating a file.", err)
		return
	}
	defer file.Close()

	//? write data to file
	data := []byte("Hello World!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to files: ", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("writingString.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go\n")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	
}
```

**-------------------------------------------------------------------------------------------------------------------------**

So, whenever we are performing all these operations on files,
always check for errors returned by file operations and handle them appropriately. Log them or send them to your event management mechanism.

And always always defer closing. Do not forget to close files, when you are working on files. Use defer to ensure files are properly closed automatically when you are done using them to prevent resource leaks.

Consider using buffered writers for better performance when writing large amounts of data.
So we should use buff.io new writer when writing large chunks of data to our files.

So overall, writing files in go involves working with the operating system, the OS package, and understanding how to create, open, write to, and handle errors when working with files is crusial for building robust applications that interact with the file system effectively.