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