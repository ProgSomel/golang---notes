# 111 - Closing Channels

## Why do we need to close channel?

We close channels to signal completion. It indicates that no more data will be sent on the channel, which helps goroutines that are receveing data know when to stop waiting. Other than that it also prevents resource leaks. Closing channels ensures that resources associated with the channel are properly cleaned up.
For closing the channel, we use the close function and after that a channel is closed, no more values can be sent to it. However we can receive values from a closed channel it it is a buffered channel. A buffered channel may have some values stored in it in the buffer, and those values can be received even if the channel is closed, because closing a channel means that the channel is closed for sending data into the channel, not for receiving values from a channel. We can always receive values from a channel if it has some values, if it is not empty, that only apply to **buffer channel**.

## Simple closing channel

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 5{
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
```

```bash
0
1
2
3
4
```

**---**

## Receiving from a closed channel - a closed unbuffered channel

```go
package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	val, ok := <- ch

	if !ok {
		fmt.Println("Channel is closed")
	}else {
		fmt.Println(val)
	}
}
```

```bash
Channel is closed
```

**---**

## loop over a closed channel

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
```

```bash
0
1
2
3
4
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/111 - Closing Channels/closing_channels.go:14 +0xc5
exit status 2
```

# Go Channel Deadlock - Missing Channel Closure Analysis

## Overview

This document analyzes Go code that causes a deadlock due to a missing channel closure operation. The code demonstrates why proper channel lifecycle management is crucial for preventing deadlocks in producer-consumer patterns.

## Code Under Analysis

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
```

## Output

```
0
1
2
3
4
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /path/to/file.go:14 +0xc5
exit status 2
```

## Root Cause Analysis

### The Missing Operation

The deadlock occurs because the **channel is never closed** after the producer finishes sending data.

```go
go func() {
    for i := range 5 {
        ch <- i
    }
    // ❌ Missing: close(ch)
}()
```

## Step-by-Step Flow Analysis

### Step 1: Channel Creation

```go
ch := make(chan int)
```

**What happens:**

- Creates unbuffered channel for integer communication
- Channel state: Open and empty

### Step 2: Producer Goroutine Launch

```go
go func() {
    for i := range 5 {
        ch <- i
    }
}()
```

**What happens:**

- Launches producer goroutine
- Producer will send values 0, 1, 2, 3, 4
- Producer does NOT close channel after sending

### Step 3: Consumer Range Loop

```go
for val := range ch {
    fmt.Println(val)
}
```

**What happens:**

- Main goroutine enters range loop
- Range loop waits for values from channel
- Will continue until channel is closed

### Step 4: Successful Data Transmission

```go
// Producer sends: 0, 1, 2, 3, 4
// Consumer receives and prints each value
```

**What happens:**

- All 5 values successfully transmitted
- Consumer prints: 0, 1, 2, 3, 4
- Producer goroutine completes its loop

### Step 5: Producer Goroutine Terminates

```go
// Producer goroutine finishes loop and exits
// Channel remains open but no more data will be sent
```

**What happens:**

- Producer goroutine ends execution
- Channel is still open (never closed)
- No more data will ever be sent

### Step 6: Consumer Waits Forever (Deadlock)

```go
for val := range ch {  // Still waiting for more data or channel closure
    fmt.Println(val)
}
```

**What happens:**

- Consumer's range loop continues waiting
- Expects either more data or channel closure
- Neither condition will ever be met
- **Deadlock detected** by Go runtime

## Execution Timeline

```
Time 0ms:
├─ Main: Creates unbuffered channel
├─ Main: Launches producer goroutine
├─ Producer: Starts loop, sends 0
├─ Main: Enters range loop, receives 0
└─ Main: Prints "0"

Time ~1ms:
├─ Producer: Sends 1
├─ Main: Receives 1, prints "1"
└─ Pattern continues for 2, 3, 4

Time ~5ms:
├─ Producer: Completes loop (sent 0,1,2,3,4)
├─ Producer: Goroutine terminates
├─ Main: Continues waiting in range loop
└─ Main: BLOCKS forever waiting for more data

Time ∞:
├─ Main: Still blocked on range loop
├─ Producer: Terminated (no more senders)
├─ Channel: Open but no writers
└─ Go Runtime: Detects deadlock, terminates program
```

## Why Range Loop Doesn't Exit

### Range Loop Behavior

```go
for val := range ch {
    // Continues until channel is closed AND empty
}
```

**Range loop exits when:**

- Channel is closed by sender: `close(ch)`
- AND all buffered data has been consumed

**Range loop continues when:**

- Channel is open (even if no senders)
- Waiting for potential future data

### The Problem

```go
// Producer never signals completion
go func() {
    for i := range 5 {
        ch <- i
    }
    // Missing close(ch) - consumer doesn't know we're done
}()
```

## Fixed Version

```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)  // ✅ Added: Signal completion
	}()

	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("All data received")
}
```

**Correct Output:**

```
0
1
2
3
4
All data received
```

## Alternative Solutions

### Solution 1: Add Channel Closure (Recommended)

```go
go func() {
    defer close(ch)  // Ensures closure even if panic
    for i := range 5 {
        ch <- i
    }
}()
```

### Solution 2: Use Done Channel

```go
func main() {
    ch := make(chan int)
    done := make(chan struct{})

    go func() {
        defer close(done)
        for i := range 5 {
            ch <- i
        }
    }()

    go func() {
        for val := range ch {
            fmt.Println(val)
        }
    }()

    <-done  // Wait for producer completion
}
```

### Solution 3: Use sync.WaitGroup

```go
func main() {
    ch := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := range 5 {
            ch <- i
        }
    }()

    go func() {
        for val := range ch {
            fmt.Println(val)
        }
    }()

    wg.Wait()
}
```

### Solution 4: Known Count Loop

```go
func main() {
    ch := make(chan int)

    go func() {
        for i := range 5 {
            ch <- i
        }
    }()

    // Receive exactly 5 values
    for i := 0; i < 5; i++ {
        val := <-ch
        fmt.Println(val)
    }
}
```

## Channel Lifecycle Best Practices

### 1. Sender Responsibility

```go
// ✅ Good: Sender closes channel
go func() {
    defer close(ch)
    // send data
}()
```

### 2. Use defer for Safety

```go
// ✅ Good: Ensures closure even with panic
go func() {
    defer close(ch)
    // risky operations
}()
```

### 3. Document Channel Ownership

```go
// ✅ Good: Clear responsibility
func producer(output chan<- int) {
    defer close(output)  // Producer owns closure
    // generate data
}
```

## Deadlock Prevention Rules

### 1. Every Range Loop Needs Closure

```go
// Consumer expects closure
for val := range ch {
    // process val
}

// Producer must close
go func() {
    defer close(ch)
    // send data
}()
```

### 2. One Closer Per Channel

```go
// ❌ Wrong: Multiple closers
go func() { defer close(ch) }()
go func() { defer close(ch) }()  // Panic!

// ✅ Right: Single closer
go func() { defer close(ch) }()
```

### 3. Close After All Sends

```go
// ✅ Correct order
func producer(ch chan<- int) {
    for i := range 5 {
        ch <- i
    }
    close(ch)  // Close after sending complete
}
```

## Conclusion

The deadlock occurs because the producer goroutine terminates without closing the channel, leaving the consumer's range loop waiting indefinitely for either more data or a closure signal that will never come. The Go runtime detects this impossible situation and terminates the program. The solution is to ensure the producer always closes the channel when finished sending data, typically using `defer close(ch)` for safety. This signals to the consumer that no more data will be sent, allowing the range loop to exit gracefully.

**---**

- **close channel only from the sender**. Do not close channel from the receiving end. Only the go routine that is sending data should close the channel. Other goroutines that are receiving should only read from the channel.
- Sometimes we close a channel more than once and that results in a runtime panic. So always ensure that channels are closed exactly once.

```go
package main

import "time"

func main() {
	ch := make(chan int)

	go func() {
		close(ch)
		close(ch)
	}()

	time.Sleep(time.Second)
}
```

```bash
panic: close of closed channel

goroutine 17 [running]:
main.main.func1()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/111 - Closing Channels/closing_channels.go:10 +0x26
created by main.main in goroutine 1
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/111 - Closing Channels/closing_channels.go:8 +0x5f
exit status 2
```

- there are some common patterns to close channels. Channels are used to pass data through a series of stages, which are pipelines. Each stage closes the channel when it is done processing, so we are going to have a producer and a filter and these functions, which are producer and filter, are going to close the channels respectively which they should. They are not going to close every channel that they are working with but only the channels that they are concerned.
```go
package main

import "fmt"

func producer(ch chan <- int) {
	for i := range 5{
		ch <- i
	}
	close(ch)
}

func filter(in <- chan int, out chan <- int ) {
	for val := range in {
		if(val % 2 == 0) {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)
	for val := range ch2 {
		fmt.Println(val)
	}
}
```
```bash
0
2
4
```

**---**

Whether it is a buffered channel or an unbuffered channel, every channel needs to be closed manually by us and a channel is closed by using the close function, and once a channel is closed, it sends out a value, a boolean value to the receiver, and we can receive that value and check if the channel is open or closed, and if the channel is closed, we can handle channel gracefully in our program.