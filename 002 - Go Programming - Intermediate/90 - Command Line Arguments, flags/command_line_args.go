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