# 91 - Command Line Sub Commands
Subcommands are a way to organize command line interfaces into hierarchical structures, allowing different functionalities or operations to be grouped under main commands.

This approach is useful for organizing and managing complex CLI applications where multiple actions or modes of operations are required.

Subcommands are essentially secondary commands that extend the functionality of a main command.

They are specified after the main command and are typically used to perform specific actions or operations.

When we are using go run --> go is the first command, so it is a command line too Go and run is the sub command to the primary command, which is go.
```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", firstFlag)
		fmt.Println("bytes: ", secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go   
This program requires additional commands
exit status 1
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", firstFlag)
		fmt.Println("bytes: ", secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go hello
no subcommand entered
exit status 1
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", firstFlag)
		fmt.Println("bytes: ", secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go firstSub
subCommand1: 
Processing:  0xc00009600a
bytes:  0xc000096020
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", *firstFlag)
		fmt.Println("bytes: ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", *flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go firstSub
subCommand1: 
Processing:  false
bytes:  1024
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", *firstFlag)
		fmt.Println("bytes: ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", *flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go firstSub -processing=true -bytes=256
subCommand1: 
Processing:  true
bytes:  256
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", *firstFlag)
		fmt.Println("bytes: ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", *flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go secondSub -language="javascript"    
subCommand2: 
language:  javascript
```

**-----------------------------------------------------------------------------------------------------------------------------**

If there is a common flag that we want to use for multiple subcommands, then we can not use this format because this format associates one flag with one subcommand

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(stringFlag)

	subCommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	//! setting flags for these subCommands 
	firstFlag := subCommand1.Bool("processing", false, "Command processing status")
	secondFlag := subCommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subCommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub": 
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1: ")
		fmt.Println("Processing: ", *firstFlag)
		fmt.Println("bytes: ", *secondFlag)
	case "secondSub":
		subCommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2: ")
		fmt.Println("language: ", *flagsc2)
	
	default:
		fmt.Println("no subcommand entered")
		os.Exit(1)
	}
}
```
```bash
go run command_line_sub_commands.go --help          
Usage of /var/folders/cq/p3dykpjj3pv6k60w9dwmblkc0000gn/T/go-build3878731843/b001/exe/command_line_sub_commands:
  -user string
        Name of the user (default "Guest")
```