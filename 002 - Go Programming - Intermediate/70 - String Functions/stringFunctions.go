package main

import (
	"fmt"
	"strings"
)

func main(){
	
	var builder strings.Builder
	//? write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	//? convert builder to string
	result := builder.String()
	fmt.Println(result)

	//? using Writerune to add a character 
	builder.WriteRune(' ')
	builder.WriteString("How are you?")
	result = builder.String()
	fmt.Println(result)

	builder.Reset()
	builder.WriteString("Starting Fresh")
	fmt.Println(builder.String())
}