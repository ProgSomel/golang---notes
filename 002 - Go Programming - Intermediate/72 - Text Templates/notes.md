# 72 - Text Templates
Text Templates in Go are a powerful feature that allow you to define and execute templates for generating text output

They are particularly usefull when you need to generate structured text such as HTML, JSON, SQL Queries.

A Template is a string or a file that contains one or more action sequences.

These actions control the template execution, such as inserting values, iterating over data, or executing conditionals.

[Text template](assets/image.png)

Templates are part of two Packages:
- Text Templates Package
- HTML Template Package --> Has some advanced features that text template package does not have.
  text template package has basic features of templating 

## Text Template
```go
package main

import (
	"os"
	"text/template"
)

func main() {
	// tmpl := template.New("example")
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?")
	if err != nil {
		panic(err)
	}

	//? Define data for the welcome message template
	data := map[string]interface{} {
		"name" : "John",
	}

	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}


}
```
```bash
Welcome, John! How are you doing?%                                                                                            
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	// tmpl := template.New("example")
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?")
	if err != nil {
		panic(err)
	}

	//? Define data for the welcome message template
	data := map[string]interface{} {
		"name" : "John",
	}

	err = tmpl.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}


}
```
```bash
Welcome, <no value>! How are you doing?%
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	// tmpl := template.New("example")
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?")
	if err != nil {
		panic(err)
	}

	//? Define data for the welcome message template
	data := []map[string]interface{} {
			{
			"name": "John",
			"age":  25,
			"city": "New York",
		},
		{
			"name": "Alice",
			"age":  30,
			"city": "London",
		},
		{
			"name": "Bob",
			"age":  35,
			"city": "Tokyo",
		},
		{
			"name": "Sarah",
			"age":  28,
			"city": "Paris",
		},
		{
			"name": "Mike",
			"age":  32,
			"city": "Sydney",
		},
		
	}
		for _, person := range data {
			err = tmpl.Execute(os.Stdout, person)
			if err != nil {
				panic(err)
			}
			fmt.Println()
		}
}
```
```bash
Welcome, John! How are you doing?
Welcome, Alice! How are you doing?
Welcome, Bob! How are you doing?
Welcome, Sarah! How are you doing?
Welcome, Mike! How are you doing?
```

--------------------------------------------------------------------------------------------------------------------

## if we use template.Must --> then we do not need to use panic, it will do panic automically
```go
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	
	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	//? Define data for the welcome message template
	data := []map[string]interface{} {
			{
			"name": "John",
			"age":  25,
			"city": "New York",
		},
		{
			"name": "Alice",
			"age":  30,
			"city": "London",
		},
		{
			"name": "Bob",
			"age":  35,
			"city": "Tokyo",
		},
		{
			"name": "Sarah",
			"age":  28,
			"city": "Paris",
		},
		{
			"name": "Mike",
			"age":  32,
			"city": "Sydney",
		},
		
	}
	for _, person := range data {
		err := tmpl.Execute(os.Stdout, person)
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}
}
```
```bash
Welcome, John! How are you doing?

Welcome, Alice! How are you doing?

Welcome, Bob! How are you doing?

Welcome, Sarah! How are you doing?

Welcome, Mike! How are you doing?
```