package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//? Getting Environment Variables
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var: ", user)
	fmt.Println("Home env var: ", home)

	//? Setting Environment Variables
	os.Setenv("FRUIT", "APPLE")

	//? Getting List of Environment variables's slice in form of key=value pairs
	 for _, e := range os.Environ() {
		kvPair := strings.SplitN(e, "=", -1) //?  based on n it will give value
		//? if n = 0; return 0
		//? if n = -1; return all substrings in a slice
		//? if n = 1/ return the value in a string in a slice
		fmt.Println(kvPair)
	 }

}