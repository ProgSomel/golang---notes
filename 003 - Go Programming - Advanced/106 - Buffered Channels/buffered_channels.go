package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		// fmt.Println("Go routine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch) //? ends <- starts
	}()

	fmt.Println("Blocking starts")
	ch <- 3 //? blocking, because buffer is full
	fmt.Println("Blocking Ends")
	// fmt.Println("Received: ", <- ch)
	// fmt.Println("Received: ", <- ch)

	// fmt.Println("Buffered Channels")

}
