# 89 - Embed Directive
The Embed Directive is a feature introduced in go version 1.16 to Embed static files or directories into go binaries at build time.

This Directive provides a convenient and efficient way to include assets directly within your go programs.

The reason, we use Embed directive is for simplicity. Embeding files simplifies deployment as it reduces the number of separate files to manage and also for efficiency.

Embeding files into binaries makes distribution and execution straightforward without worrying about file paths or external dependencies, because all our dependencies, all our assets, all the extra files that are needed are combined into our executablek.

And another important factor we use embeded dirctive is for security.

Embeded files are bundled within the binary, reducing exposure to external manipulation or unauthorized access.

**Now what types does the embed directive support?**
- The embed dirctive supports files, the individual files that we have and directories, entire directories and their contents can be embeded recursively and that means that we can embed almost anything into our executable.

The Embed dirctive in go is not a command, it is a dirctive.

## Embed a file in final executable file
```go
package main

import (
	_"embed"
	"fmt"
)

//go:embed example.txt
var content string

func main() {
	fmt.Println("Embeded content: ", content)
}
```
```bash
Embeded content:  Hello World!
```

**-------------------------------------------------------------------------------------------------------------------------**

## Embed a folder to final executable file
```go
package main

import (
	"embed"
	"fmt"
)

//go:embed example.txt
var content string

//go:embed basic
var basicsFolder embed.FS

func main() {
	fmt.Println("Embeded content: ", content)
	content, err := basicsFolder.ReadFile("basic/hello.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("Embeded file content: ", string(content))
}
```
```bash
Embeded content:  Hello World!
Embeded file content:  Hello Basic Folder
```

**-------------------------------------------------------------------------------------------------------------------------**

## We can walk through the folder and list all the contents of the folder
```go
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basic
var basicsFolder embed.FS

func main() {
	fmt.Println("Embeded content: ", content)
	content, err := basicsFolder.ReadFile("basic/hello.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	fmt.Println("Embeded file content: ", string(content))

	err = fs.WalkDir(basicsFolder, "basic", func(path string, d fs.DirEntry, err error) error{
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
```
```bash
Embeded content:  Hello World!
Embeded file content:  Hello Basic Folder
basic
basic/hello.txt
```

**-------------------------------------------------------------------------------------------------------------------------**

The Embed directive finds its use in web servers for embeding static HTML, CSS, and JavaScript files for serving web content and also in configuration files.

Consider that embedded files can not be modified at runtime, and you may need to rebuild the binary for any updates.

In conclusion, the embed directive in go provides a powerful mechanism to include static files and directories into go binaries, simplifying deployment and distribution of applications.

By leveraging the embed package, developers can enhance application portability and security while maintaining ease of use and management of embedded resources.