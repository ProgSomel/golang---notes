# 88 - Temporary Files and Directories
Temporary Files and Directories are essential in many programming scenarios where temporary storage is needed for data processing, caching, or other transient operations.

Go provides mechanisms to create, manage, and clean up temporary files and directories efficiently.

In Go, temporary files are created using the OS dot create temp function. This function creates a temporary file in the default location for temporary files. It is the system's default temporary directory, such as root tmp on Unix like systems, and in-order for us to create temporary directories, we use OS dot mkkdir tmp, which functions similarly to os dot create temp.

## creating Temporary File -> os.CreateTemp()
```go
package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkErr(err)
	fmt.Println("Temporary File created: ", tempFile.Name())

	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
}
```
```bash
Temporary File created:  /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/temporaryFile606571660
```

**--------------------------------------------------------------------------------------------------------------------------**

## creating temporary directory -> os.MkdirTemp(path string, pattern string)
```go
package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkErr(err)

	defer os.Remove(tempDir)
	fmt.Println("Temporary Directory created: ", tempDir)
}
```
```bash
Temporary Directory created:  /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/GoCourseTempDir3323868900
```

**--------------------------------------------------------------------------------------------------------------------------**

The best practices associated with temporary files and directories is that we have to be cautious with temporary files containing sensitive data. Always ensure that they are securely handled and cleaned up promptly.