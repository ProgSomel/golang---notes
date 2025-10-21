# 106 - Buffered Channels

Buffered Channels allow Channels to hold a limited number of values before blocking the sender.

Buffered Channels are useful for managing data flow and controlling concurrency, so buffer essentially means storage. We are allowing channels to store values.

Earlier in unbuffered channels, we saw that channels can not store value. They need an immediate outflow of the value that they receive. As soon as they have an incoming value, they need a channel to receive. They need a receiver to receive the value because they can not hold a value. However buffered channels can hold value inside them, they do not require an immediate receiver.

We were getting blocked because the unbuffered channel required an immediate receiver.

What buffer channel offers us is an **asynchronous communication**. They allow senders to continue working without blocking until the buffer is full, and they do not require an immediate receiver. They will only block when the buffer is full, and they will only block when we are trying to insert more value, when we are trying to send more value inside the channel when it is already full. So the channel will not return an error or will not block as soon as it gets full. It will only block or return an error when we try to send more value. When we try to feed more value inside the buffer when it is already full.

Other than that, we also use buffer channels when we are handling bursts of data without immediate synchronization. immediate synchronization means that we are receiving values and sending values to a receiver. The channel is receiving values and then sending values to somewhere else. That is immediate synchronization. But we do not need immediate synchronization because we are storing values inside a buffer.

When talking about buffered channels, lets take the same examples like a down flowing pipe or a conveyor belt or a slide thay you see in a children's park. We have a closing mechanism at the end of each of these things. Let's consider a tap, a tap at the end of the pipe, so we can take the values out as soon as we open the tap, and then close it so those values as they are fed in, they will get out when we open the tap. And that pipe, that water pipe which is inclined towards one end and the downward end is the receiving end. So the receiving end has a tap. and we can receive values whenever we want. But this pipe has a capacity and the capacity can be decided by us during the time of declaration. So we can declare a capacity of fixed capacity of this pipe, and then it will be able to hold that much data inside it without passing those values out of the other end. The other end is closed and we can extract, we can take out those values whenever we want.

Another aspect of buffer channels is that buffered channels help us in flow control. This tap, this holding mechanism, this buffer is kind of a flow control. We are controlling the outward flow of data from the channel. So we have created a mechanism to control the flow of data, that is flow control. We are managing the rate of data transfer between producers and consumers. Consumers will be receivers.

---

**Like unbuffered channels, similarly buffered channels also have some blocking priciples involved with.**

Buffered channels block on send only if the buffer is full, and buffer channels also block on receive only if the buffer is empty.

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println("Buffered Channel")
}
```

```bash
Buffered Channel
```

---

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Buffered Channel")
}
```

```bash
Buffered Channel
```

---

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Buffered Channel")
}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/106 - Buffered Channels/buffered_channels.go:10 +0x58
exit status 2
```

We store one value inside this channel and there is space for another. We store two and then there is no more space, so this statement:

```go
ch <- 3
```

it tries to store a third value, a third integer number three, inside this channel:

```go
ch := make(chan int, 2)
```

it goes to the channel, it reaches out to the channel and it tries to store value inside the channel, but the channel is logged, kind of logged, it is full. Consider a jar, a jar that is full, and we try to put some more candies into that jar and it is full of candies and there is no more space left, not even an inch, a millimeter, not even a millimeter of space left in that jar to store anything else now. If we try to put any more candy inside that jar, that is a deadlock.
So, this is where we are getting blocked:

```go
ch <- 3
```

---

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	ch <- 3
	fmt.Println("Buffered Channel")
}
```

```bash
Value:  1
Value:  2
Buffered Channel
```

We see here, that bufferd channels do not block our operations, they do not need an immediate receiver. As soon as they are full in that case, they need a receiver, otherwise they will block.

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	ch <- 3
	fmt.Println("Value: ", <-ch)
	fmt.Println("Buffered Channel")
}
```

```bash
Value:  1
Value:  2
Value:  3
Buffered Channel
```

---

```go
package main

import "fmt"

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("Buffered Channel")
}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/106 - Buffered Channels/buffered_channels.go:10 +0x5e
```

This gives an Error. This is also blocking mechanism as well.

---

**In order for us to see the blocking mechanism to demonstrate the blocking mechanism, we are going to introduce **goroutines\*\*\*\* and we are going to receive values inside the goroutine

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2*time.Second)
		fmt.Println("Received: ", <- ch)
	}()
	ch <- 3
	fmt.Println("Received: ", <- ch)
	fmt.Println("Received: ", <- ch)
	fmt.Println("Buffered Channel")
}
```

```bash
Received:  1
Received:  2
Received:  3
Buffered Channel
```

---

## Blocking on the SEND only, if the Buffer is Full

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//! make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2*time.Second)
		fmt.Println("Received1: ", <- ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3 //? Blocks, because Buffer is Full
	fmt.Println("Blocking ends")
	fmt.Println("Received2: ", <- ch)
	fmt.Println("Received3: ", <- ch)
	fmt.Println("Buffered Channel")
}
```

```bash
Blocking starts
Received1:  1
Blocking ends
Received2:  2
Received3:  3
Buffered Channel
```

**Step 1: Channel Creation**

```go
go := make(chan int, 2)
```

- Creates a **buffered channel** with capacity 2
- Can hold 2 values without blocking

**Step 2: Fill the Buffer**

```go
ch <- 1 // Buffer: [1, _] - No blocking
ch <- 2 // Buffer: [1, 2] - No blocking, buffer full
```

- Both values sent successfully
- No blocking because buffer has space
- Channel is now full

**Step 3: Start Goroutine**

```go
go func() {
time.Sleep(2\*time.Second)
fmt.Println("Received1: ", <- ch) // Will receive 1
}()
```

- Goroutine starts but sleeps for 2 seconds
- Doesn't read from channel yet
- Main goroutine continues

**Step 4: Print Message**

```go
fmt.Println("Blocking starts")
```

**Output**: **Blocking starts**

**Step 5: Blocking Send**

```go
ch <- 3 // BLOCKS HERE!
```

- **Channel is full** (already has [1, 2])
- **Main goroutine blocks** waiting for space
- **Waits 2 seconds** until goroutine reads

**Step 6: Goroutine Wakes Up (After 2 seconds)**

```go
fmt.Println("Received1: ", <- ch) // Reads 1
```

- Goroutine reads 1 from channel
- Output: **Received1: 1**
- Channel now has space: [2, _]
- **Unblocks the main goroutine**

**Step 7: Main Goroutine Continues**

```go
fmt.Println("Blocking ends") // Buffer: [2, 3]
fmt.Println("Received2: ", <- ch) // Reads 2
fmt.Println("Received3: ", <- ch) // Reads 3
fmt.Println("Buffered Channel")

```

```bash
## **Complete Output:**
Blocking starts
(2 second delay)
Received1: 1
Blocking ends
Received2: 2
Received3: 3
Buffered Channel
```

## **Visual Timeline:**

Time 0s:
├─ ch buffer: [1, 2] (full)
├─ Goroutine: sleeping for 2s
├─ Main: prints "Blocking starts"
└─ Main: tries ch <- 3 (BLOCKS - buffer full)

Time 2s:
├─ Goroutine: wakes up, reads 1
├─ ch buffer: [2, _] (has space)
├─ Main: unblocks, ch <- 3 succeeds
├─ ch buffer: [2, 3]
└─ Main: continues execution

Time 2s+:
├─ Main: prints "Blocking ends"
├─ Main: reads 2, ch buffer: [3]
├─ Main: reads 3, ch buffer: []
└─ Main: prints "Buffered Channel"

**---**

## Blocking on Receive only, if the Buffer is Empty

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("Value: ", <- ch)

}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/106 - Buffered Channels/buffered_channels.go:7 +0x38
exit status 2
```

---

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func(){
		time.Sleep(2 * time.Second)
	}()

	fmt.Println("Value: ", <- ch)

}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/106 - Buffered Channels/buffered_channels.go:15 +0x4e
exit status 2
```

---

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	fmt.Println("Value: ", <- ch)
	fmt.Println("End of Program")

}
```

```bash
Value:  1
End of Program
```

---

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func(){
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()

	fmt.Println("Value: ", <- ch)
	fmt.Println("Value: ", <- ch)
	fmt.Println("End of Program")

}
```

```bash
Value:  1
Value:  2
End of Program
```

Here, in this above example, we have two senders and two receivers.
So, we see here that there is a blocking mechanism when we have receivers ready and we do not have senders. the receivers wait for the senders in another goroutine to send values to them.

If all the goroutines have finished and there was no sender which sent a value to them, then they throw an error.

**---**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch)
	}()

	fmt.Println("Blocking starts")
	ch <- 3 //? blocking, because buffer is full
	fmt.Println("Blocking Ends")
	fmt.Println("Received: ", <- ch)
	fmt.Println("Received: ", <- ch)

	fmt.Println("Buffered Channels")

}
```

In this example, we have a channel which has a capacity of two, it is a buffered channel. We are sending one and two, these two values into the channel.

```go
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
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch)
	}()

	fmt.Println("Blocking starts")
	ch <- 3 //? blocking, because buffer is full
	fmt.Println("Blocking Ends")
	// fmt.Println("Received: ", <- ch)
	// fmt.Println("Received: ", <- ch)

	// fmt.Println("Buffered Channels")

}
```

```bash
Receiving from buffer
Blocking starts
Blocking Ends
Received:  1
```

---

```go
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
		fmt.Println("Go routine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch)
	}()

	fmt.Println("Blocking starts")
	ch <- 3 //? blocking, because buffer is full
	fmt.Println("Blocking Ends")
	// fmt.Println("Received: ", <- ch)
	// fmt.Println("Received: ", <- ch)

	// fmt.Println("Buffered Channels")

}
```

```bash
Receiving from buffer
Blocking starts
Go routine 2 second timer started
Blocking Ends
Received:  1
```

When this go function, this goroutine:

```go
go func() {
		fmt.Println("Go routine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch)
}()
```

is encountered by the go runtime scheduler, it extracts it from the main thread. It reaches here:

```go
	ch <- 3 //? blocking, because buffer is full
```

and it blocks our main thread because there is a goroutine and it may transmit data, it may free up some resource for this channel to receive this value, it waits for the goroutine to be over. It is go runtime scheduler that waits for all the goroutines to be over, and then it lets us know if there is an error as per the situation. If the channel is not able to receive this value because the resources were not free, maybe we did not have this statement:

```go
	fmt.Println("Received: ", <- ch)
```

here and the channel buffer was full, in that case, the go runtime scheduler is the one that will throw us an error.

Go runtime scheduler extracts this go routine:

```go
go func() {
		fmt.Println("Go routine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received: ", <- ch)
}()
```

assigns it to a new thread and moves on to the next line, the additional thread, it moves to the next line.

First we take the scenario when received:

```go
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
		fmt.Println("Received: ", <- ch)
	}()

	fmt.Println("Blocking starts")
	ch <- 3 //? blocking, because buffer is full
	fmt.Println("Blocking Ends")
	// fmt.Println("Received: ", <- ch)
	// fmt.Println("Received: ", <- ch)

	// fmt.Println("Buffered Channels")

}
```

```go
	fmt.Println("Received: ", <- ch)
```

was not getting printed, sometimes it was and sometimes it wasn't. **Why was that?**
The thread that is handling the go routine, as soon as it starts reading this function:

```go
go func() {
	fmt.Println("Go routine 2 second timer started")
	time.Sleep(2 * time.Second)
	fmt.Println("Received: ", <- ch)
}()
```

the first line is:
```go
	time.Sleep(2 * time.Second)
```
as soon as it encounters this statement, it pauses reading this function and it puts this goroutine to sleep for two seconds it has not read the next line:
```go
	fmt.Println("Received: ", <- ch)
```
it does not know what is written in the next line. after two seconds, as soon as the sleep is over, it moved on to the next line:
```go
	fmt.Println("Received: ", <- ch)
```
Now. one principle that we need to understand, and that principle has been there since a very long time, since the oldest programming languages to the latest programming languages, it is the same thing, the execution starts from right, towards left when we are facing our code, our right towards left:
```go
	fmt.Println("Received: ", <- ch) //? ends <- starts
```
as soon as this come to this line, it reads this and emits the value out, emits one value from the channel buffer
```go
<- ch
```
as soon as one value is extracted from this channel buffer, our value three:
```go
	ch <- 3 //? blocking, because buffer is full
```
was waiting to enter into the buffer. as soon as one value gets out of the channel buffer, another value enters. This statement:
```go
	ch <- 3 //? blocking, because buffer is full
```
gets executed, and we reach the end of the function, and the remaining part of the println statement:
```go
fmt.Println("Received: ",
```
fails to get executed in time to be printed into the terminal, our main thread ends before the rest of it could be executed. Thats why we sometimes get value and sometimes not. This is how fast things are happening behind the scenes.

**---**
- While sending data to a channel, it will implement a blocking mechanism when the buffer is full.
- When we are receiving data from a buffer channel, we will encounter a blocking mechanism, when the buffer is empty, we are trying to receive values from a buffer which is empty.
- other than that buffer channels have non-blocking operations. They allow non-blocking sends and receives as long as the buffer is not full or empty.
- Buffer channel can improve performance by reducing synchronization overhead, and when it comes to unbuffered channels, they typically use strict synchronization where sender and receiver must be synchronized, they must be ready at the same time to send and receive values.
- when using buffered channels, we should be considered about choosing a buffer size, we should choose a buffer size based on the expected volume of data and concurrency requirements, and that is because large buffers reduce the likelihood of blocking, but increase memory usage, and smaller buffers increase the likelihood of blocking, but use less memory. So always consider this when choosing your buffer size for the channel. And that is why you always should avoid over buffering. Large buffers may mask issues related to design and data flow.
- While buffer channels are easy to implement, always try and use logging or debugging tools to monitor the behavior of buffered channels in a complex application.
- unbuffered channels are dependent on both the ends of channels, and both the ends should be ready at the same time, otherwise we will encounter a blocking mechanism or an error in our application. It is because of this phenomenon that unbuffered channels can not store values that we have to use them in go routines. However, when it comes to buffer channels, they have a possibility of getting used in our main thread without using go routines in a very simpler way.