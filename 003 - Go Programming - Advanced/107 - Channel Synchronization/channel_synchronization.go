package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + strconv.Itoa(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	close(data)

	for value := range data {
		fmt.Println("Received Value: ", value, ":", time.Now())
	}
}