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