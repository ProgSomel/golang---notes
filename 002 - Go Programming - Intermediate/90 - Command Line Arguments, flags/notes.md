# 90 - Command Line Arguments/flags
Command Line Arguments are a common way to pass parameters to a program when it is executed from a terminal or command prompt.

In Go, command line arguments is straightforward, leveraging the OS package for accessing arguments and flags.

In Go, command line arguments are accessible via the OS dot arg's slice, where os dot args zero is the name of the command or the name of the program itself
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
}
```
```bash
Command:  /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/go-build2460727064/b001/exe/command_line_args
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
	fmt.Println("Argument1: ", os.Args[1])
}
```
```bash
 90 - Command Line Arguments, flags go run command_line_args.go hello
Command:  /Users/progsomel/Library/Caches/go-build/06/068b85ad8bf4ea00dd6ea5bc9be649b485484d1b6efe2719957af3b20afd07af-d/command_line_args
Argument1:  hello
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Command: ", os.Args[0])
	
	for i, arg := range os.Args {
	fmt.Println("Argument", i, ":", arg)
	}

}
```
```bash
run command_line_args.go hello world
Command:  /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/go-build4261400217/b001/exe/command_line_args
Argument 0 : /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/go-build4261400217/b001/exe/command_line_args
Argument 1 : hello
Argument 2 : world
```

------------------------------------------------------------------------------------------------------------------------

## flags -> flag package allows defining flags with various types like int, string, bool, etc. 
and it automically parses command line arguments into these flags.
```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	//? Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "name of the user")
	flag.IntVar(&age, "age", 30, "age of the user")
	flag.BoolVar(&male, "male", true, "gender of the user")

	flag.Parse()

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Male: ", male)
}
```
```bash
go run command_line_args.go -name "Jane Doe" -age 50 -male false 
Name:  Jane Doe
Age:  50
Male:  true
```

------------------------------------------------------------------------------------------------------------------------

**When we have command line flags with their usage description available for the main command, that is, the main program, we can use **--help** to get them all printed on the terminal**
```bash
go run command_line_args.go --help              
Usage of /Users/progsomel/Library/Caches/go-build/37/373a84b058f9aed01cc963ff121e51db81a58a6ba7784ef547206c5d01605543-d/command_line_args:
  -age int
        age of the user (default 30)
  -male
        gender of the user (default true)
  -name string
        name of the user (default "John")
```