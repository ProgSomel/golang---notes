package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	loc, _ := time.LoadLocation("America/New_york")
	fmt.Println("Newyork time: ", t.In(loc))
}