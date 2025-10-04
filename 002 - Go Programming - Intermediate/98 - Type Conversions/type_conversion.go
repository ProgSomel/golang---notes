package main

import "fmt"

func main() {
	//! Type(value)
	i := []byte{255, 100, 72}
	j := string(i)
	fmt.Println(j)
}