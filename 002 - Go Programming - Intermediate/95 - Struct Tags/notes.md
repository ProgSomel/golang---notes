# 95 - Struct Tags
Struct tags play a crucial role in controlling how data is encoded and decoded in Go. Especially when working with JSON.

They allow you to specify details about how struct fields should be represented in JSON, providing more flexibility and control.

- First of all,struct tags can be used to map struct field names to specific JSON keys, which might not match the go field names, and this is useful when working with APIs or data source where the JSON keys have the different naming conventions.

- Struct tags can also indicate that certain fields should be omitted from the JSON output, either when they have zero values, or always for omitting.
zero values, we use hypen omitempty and for always omitting we use omitempty.

- Apart from that, you can also rename struct fields in the JSON output using struct tags. This is useful for ensuring the JSON output meets specific schema requirements, and schema is very important when we are storing data into a database or accessing data from a database, and a lot of times we receive JSON data and we store that data into a database. And that's when we can rename struct fields in the JSON output using the struct tags.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "Doe", Age: 33}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","last_name":"Doe","age":33}
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "", Age: 33}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","last_name":"","age":33}
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"age"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "", Age: 33}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","age":33}
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name, omitempty"`
	Age int `json:"age"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "", Age: 33}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","last_name":"","age":33}
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"age,omitempty"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "Doe", Age: 33}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","last_name":"Doe","age":33}
```

-----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"age,omitempty"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "", Age: 0}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane"}
```

**-----------------------------------------------------------------------------------------------------------------------------**

## permanently omit field wether there is a zero value or a non-zero value. We will use hyphen instead of name
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"-"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "Doe", Age: 20}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
```bash
Json Data:  {"first_name":"Jane","last_name":"Doe"}
```

**-----------------------------------------------------------------------------------------------------------------------------**

## Expanding the usage of struct tags from JSON to other functionalities as well.
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName string `json:"last_name,omitempty"`
	Age int `json:"-"`
}

func main() {
	person := Person{FirstName: "Jane", LastName: "Doe", Age: 20}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling struct: ", err)
		return
	}

	fmt.Println("Json Data: ", string(jsonData))
}
```
