package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch)

		fmt.Println("2 second Go routine finished")
	}()
	ch <- 1;

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second Go routine finished")
	// }()

	// go func() {
	// 	// ch <- 1;
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("2 second Go routine finished")
	// }()

	// receiver := <- ch
	// fmt.Println(receiver)
	fmt.Println("End of the program")
}