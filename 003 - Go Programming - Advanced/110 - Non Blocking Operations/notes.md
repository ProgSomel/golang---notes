# 110 - Non Blocking Operatons
Non-Blocking operations on channels allow a go routine to perform a channel operation like send or receive without getting stuck, if the channel is not ready, they help maintain responsiveness and prevent go routines from getting blocked indefinitely.

We need to use non=blocking operations to avoid deadlocks, prevent go routines from waiting indefinitely on channel operations, and also to improve efficiency, allow goroutines to continue processing or handle tasks, handle other tasks if channels are not immediately ready. Most importantly, to enhance **concurrency**.

Manage multiple concurrent operations more effectively by not blocking on individual channel operations.

**---**

## Non-Blocking receive operations
We do this by using the default case in select. The select statement with a default case allows for Non-Blocking receives by immediately executing the default case if no other channels are ready.

```go
package main

import "fmt"

func main() {
	//? === Non-blocking Receive Operaiton
	ch := make(chan int)

	select {
	case msg := <- ch:
		fmt.Println("Received: ", msg)
	default:
		fmt.Println("No message is available.")
	}

}
```
```bash
No message is available.
```
No message is available because we did not send any message into the channel, if we do not have a select case then we would encounter an error.

**---**
## Non-Blocking Send operation
```go
package main

import "fmt"

func main() {
	//? === Non-blocking Receive Operaiton
	ch := make(chan int)

	select {
	case msg := <- ch:
		fmt.Println("Received: ", msg)
	default:
		fmt.Println("No message is available.")
	}

	//? === Non-blocking Send Operaiton
	select {
	case ch <- 1:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is not ready to receive.")
	}

}
```
```bash
No message is available.
Channel is not ready to receive.
```

**---**

## Non-Blocking operation in real time systems
We use Non-Blocking operations to handle real time data processing, where timely responses are critical.
```go
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
```
```bash
Data received:  0
Waiting for data...
Waiting for data...
Data received:  1
Waiting for data...
Waiting for data...
Data received:  2
Waiting for data...
Waiting for data...
Data received:  3
Waiting for data...
Waiting for data...
Data received:  4
Waiting for data...
Waiting for data...
Stopping...
```

**---**

