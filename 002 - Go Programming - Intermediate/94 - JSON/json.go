package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	FullName string `json:"full_name"`
	EmpID string 	`json:"emp_id"`
	Age int `json:"age"`
	Address Address `json:"address"`
}

func main() {

	jsonData := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	//? as this is an unknown JSON structure, we can not make a struct

	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}
	fmt.Println("Decoded/unMarshalled JSON: ", data)
	fmt.Println("Decoded/unMarshalled JSON address: ", data["address"])
	fmt.Println("Decoded/unMarshalled JSON name: ", data["name"])

}

