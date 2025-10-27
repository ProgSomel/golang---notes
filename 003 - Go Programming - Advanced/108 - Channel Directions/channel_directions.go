package main

import "fmt"

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

//! Send only channel
func producer(ch chan <- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

//! Receive only channel
func consumer(ch <- chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}