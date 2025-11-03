package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <- data:
				fmt.Println("Data received: ", d)
			case <- quit :
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true
}