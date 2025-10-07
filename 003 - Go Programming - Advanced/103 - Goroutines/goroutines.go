package main

import (
	"fmt"
	"time"
)

func sayHello() {
	//? It will wait 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Go Routine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
    //? simulate work
    time.Sleep(1 * time.Second)
    return fmt.Errorf("an Error occured in doWork")
}

func main() {
    var err error
	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello function.")
	go printNumbers()
	go printLetters()

    go func(){
        err = doWork()
    }()
    
	time.Sleep(2 * time.Second)

    if err != nil {
        fmt.Println("Error: ", err)
    }else{
        fmt.Println("Work completed successfully")
    }
}