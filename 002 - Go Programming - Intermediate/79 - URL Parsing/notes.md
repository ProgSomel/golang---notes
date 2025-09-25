# 79 - URL Parsing
URL Parsing in Go, involves extracting various components like scheme, host path, query parametes, etc from a URL string

This is crusial for tasks like building web applications, API endpoints, or processing URLs in general.

URL Structure --> [scheme://][userinfo@]host[:port][/path][?query][#fragment]
**scheme** -> is also called protocol(can be HTTP or Https or FTP)
**userinfo@** -> contains username and password, which is optional
**host** -> domain name or ip address
**port** -> optional
**path** -> resource on the server
**query parameters** -> Query parameters are in the key value pairs.
**fragment identifier** -> This are optional and used for specifying a location within the resource.

Go's Net URL provides a comprehensive package for parsing URLs and manipulating their components.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path?query=param$fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Port: ", parsedURL.Port())
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("Raw Query: ", parsedURL.RawQuery)
	fmt.Println("Fragment: ", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=john&age=30"
	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))
}
```
```bash
Scheme:  https
Host:  example.com:8080
Port:  8080
Path:  /path
Raw Query:  query=param$fragment
Fragment:  
map[age:[30] name:[john]]
Name:  john
Age:  30
```

**------------------------------------------------------------------------------------------------------------------------**

## building URL
```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "25")
	baseURL.RawQuery = query.Encode() //? Enocde in url format

	fmt.Println("Built URL: ", baseURL.String())
}
```
```bash
Built URL:  https://example.com/path?age=25&name=John
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	values := url.Values{}
	//? Add key value pairs to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	//? Encode
	encodedQuery := values.Encode()
	fmt.Println(values)
	fmt.Println(encodedQuery)
}
```
```bash
map[age:[30] city:[London] country:[UK] name:[Jane]]
age=30&city=London&country=UK&name=Jane
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	values := url.Values{}
	//? Add key value pairs to the values object
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "London")
	values.Add("country", "UK")

	//? Encode
	encodedQuery := values.Encode()
	fmt.Println(values)
	fmt.Println(encodedQuery)

	//? Build a URL
	baseURL := "https://example.com/search"
	fullURL := baseURL + "?" + encodedQuery

	fmt.Println(fullURL)
}
```
```bash
map[age:[30] city:[London] country:[UK] name:[Jane]]
age=30&city=London&country=UK&name=Jane
https://example.com/search?age=30&city=London&country=UK&name=Jane
```

**-----------------------------------------------------------------------------------------------------------------------**

Net URL Package is essential for extracting and manipulating various components of URLs, and understanding how to parse, build, and handle URLs ensures robust handling of web related tasks in Go applications.