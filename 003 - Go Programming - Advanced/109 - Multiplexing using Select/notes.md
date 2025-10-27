# Multiplexing using Select

Multiplexing is the process of handling multiple channel operations simultaneously, allowing a goroutine to wait on multiple channel operations and react to whichever operations is ready first.

**The select statement** in Go facilities multiplexing by allowing a **go routine** to wait on multiple channel.

**Multiplexing is like switch** So select statement is like a switch statement that we have used before.

## Why should we use Multiplexing?

Multiplexing manages multiple concurrent operations within a single goroutine. And it manages those concurrent operations seamlessly with improved readability, and you will enjoy writing the code as well because it is like **switch case** and switch case kind of statement.

Apart from that, multiplexing efficiently handles operations that might block without locking up resources, so it handles the blocking operations pretty efficiently. Most importantly, we get to implement timeouts and cancellation mechanisms. We can timeouts channel and we can cancel channel using multiplexing.

Select syntax is so simple. Just Select, case and default. Just like we had switch case we now have Select case, but Select case is only for **handling channels** to be precise for h**andling multiple channels** operations simultaneously.

```go
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	msg1 := <- ch1
	fmt.Println("Received from ch1:", msg1)

	msg2 := <- ch2
	fmt.Println("Received from ch2:", msg2)
}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/109 - Multiplexing using Select/multiplexing_using_select.go:9 +0x55
exit status 2
```

---

```go
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)
	default:
		fmt.Println("No channels ready...")
	}
}
```

```bash
No channels ready...
```

**Earlier one is statement but this one is condition. Statement executes and condition checks.**

---

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second)
	select {

	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	default:
		fmt.Println("No Channels Ready...")
	}

	fmt.Println("End of the Program")
}
```

```bash
Received from ch1: 1
End of the Program
```

---

```go
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	}
}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/109 - Multiplexing using Select/multiplexing_using_select.go:9 +0x9d
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
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)
	default:
		fmt.Println("No channels ready...")
	}
}
```

```bash
No channels ready...
```

What is happening there is that we have a default case. As soon as the program starts, we are not waiting for these goroutines to finish. Because if we use these conditions:

```go
msg := <- ch1
```

as Statements, then we know that channels will wait for goroutines to finish. but here this is a condition and as soon as this condition is not satisfied execution moves on to next and then next, so by default the default case is being executed. **If we remove the default condition then it will work:**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	}

	fmt.Println("End of the Program")
}
```

```bash
Received from ch2: 2
End of the Program
```

# Go Select Statement with Multiple Channels - Flow Analysis

## Overview

This document analyzes Go code that demonstrates the `select` statement with multiple channels, showing how Go handles concurrent channel operations and non-deterministic channel selection.

## Code Under Analysis

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	}

	fmt.Println("End of the Program")
}
```

## Output

```bash
Received from ch2: 2
End of the Program
```

Note: The output may vary between runs due to the non-deterministic nature of the select statement.

## Step-by-Step Flow Analysis

### Step 1: Channel Creation

```go
ch1 := make(chan int)
ch2 := make(chan int)
```

**What happens:**

- Creates two unbuffered channels for integer communication
- Both channels start empty and ready for communication
- Channel capacity: 0 (synchronous communication required)

### Step 2: First Goroutine Launch

```go
go func() {
	time.Sleep(time.Second)
	ch1 <- 1
}()
```

**What happens:**

- Launches first goroutine that will send to ch1
- Goroutine immediately starts sleeping for 1 second
- After sleep, will send value 1 to ch1

### Step 3: Second Goroutine Launch

```go
go func() {
	time.Sleep(time.Second)
	ch2 <- 2
}()
```

**What happens:**

- Launches second goroutine that will send to ch2
- Goroutine immediately starts sleeping for 1 second
- After sleep, will send value 2 to ch2

### Step 4: Select Statement Execution

```go
select {
case msg := <- ch1:
	fmt.Println("Received from ch1:", msg)

case msg := <- ch2:
	fmt.Println("Received from ch2:", msg)
}
```

**What happens:**

- Main goroutine enters select statement
- Select blocks waiting for any of the case conditions to become ready
- Both channels are empty, so main goroutine waits

### Step 5: Concurrent Sleep Period

```go
// Time 0ms - 1000ms
// Both goroutines sleeping simultaneously
// Main goroutine blocked on select statement
```

**What happens:**

- All goroutines are in waiting state
- Goroutine 1: sleeping, will send to ch1 after 1 second
- Goroutine 2: sleeping, will send to ch2 after 1 second
- Main: blocked on select, waiting for any channel to receive data

### Step 6: Simultaneous Channel Operations

```go
// At approximately 1000ms:
// Goroutine 1: ch1 <- 1
// Goroutine 2: ch2 <- 2
```

**What happens:**

- Both goroutines wake up at nearly the same time
- Both attempt to send their values simultaneously
- Select statement has two ready cases
- Go runtime makes non-deterministic choice

### Step 7: Select Case Execution

```go
case msg := <- ch2:
	fmt.Println("Received from ch2:", msg)
```

**What happens:**

- In this execution, ch2 case was selected
- Value 2 received from ch2
- Prints "Received from ch2: 2"
- Select statement completes - only one case executes

### Step 8: Program Termination

```go
fmt.Println("End of the Program")
```

**What happens:**

- Main goroutine continues after select
- Prints final message
- Program terminates
- Other goroutine that wasn't selected may still be blocked

## Execution Timeline

```
Time 0ms:
├─ Main: Creates ch1 and ch2 channels
├─ Main: Launches goroutine 1 (sleeps 1s, sends to ch1)
├─ Main: Launches goroutine 2 (sleeps 1s, sends to ch2)
├─ Goroutine 1: Begins sleep(1s)
├─ Goroutine 2: Begins sleep(1s)
└─ Main: Enters select statement (blocks)

Time 0ms - 1000ms:
├─ Main: Blocked on select
├─ Goroutine 1: Sleeping...
└─ Goroutine 2: Sleeping...

Time ~1000ms:
├─ Goroutine 1: Wakes up, attempts ch1 <- 1
├─ Goroutine 2: Wakes up, attempts ch2 <- 2
├─ Main: Select has 2 ready cases
└─ Go Runtime: Makes non-deterministic choice

Time ~1000ms (execution):
├─ Selected: ch2 case executes
├─ Main: Receives 2 from ch2
├─ Main: Prints "Received from ch2: 2"
├─ Main: Exits select statement
├─ Main: Prints "End of the Program"
└─ Program: Terminates
```

## Select Statement Mechanics

### How Select Works

1. **Evaluates all cases** simultaneously
2. **Blocks until at least one case is ready**
3. **If multiple cases ready**, chooses one randomly
4. **Executes only the selected case**
5. **Continues with code after select**

### Non-Deterministic Selection

```go
// When both channels ready simultaneously:
// Could execute either case randomly

// Possible outputs:
// "Received from ch1: 1" OR "Received from ch2: 2"
```

### Case Readiness

| Time    | ch1 Status        | ch2 Status         | Select Behavior  |
| ------- | ----------------- | ------------------ | ---------------- |
| 0-999ms | Not ready         | Not ready          | Blocks           |
| 1000ms  | Ready             | Ready              | Random selection |
| 1001ms+ | One case executed | Program terminated |

## Why Output Can Vary

### Race Condition by Design

- Both goroutines wake up after approximately 1 second
- Timing differences in nanoseconds affect which case is ready first
- Go scheduler and OS factors influence execution order
- This creates intentional non-determinism

### Possible Outputs

```bash
# Run 1:
Received from ch1: 1
End of the Program

# Run 2:
Received from ch2: 2
End of the Program

# Run 3:
Received from ch1: 1
End of the Program
```

## What Happens to the Non-Selected Channel

### Important Behavior

When select executes one case:

- **Other goroutine may still be blocked** trying to send
- **Program terminates before other send completes**
- **No goroutine leak** - Go runtime cleans up on program exit

### Example: If ch1 case selected

```go
// If ch1 case executed:
case msg := <- ch1:  // This executes
    fmt.Println("Received from ch1:", msg)

case msg := <- ch2:  // This case skipped
    // Goroutine 2 still trying: ch2 <- 2
    // But main exits, so send never completes
```

## Select Statement Patterns

### Pattern 1: Multiple Channel Communication (Current)

```go
select {
case msg := <-ch1:
    // Handle ch1 data
case msg := <-ch2:
    // Handle ch2 data
}
```

### Pattern 2: Non-blocking Operations

```go
select {
case msg := <-ch:
    // Handle data
default:
    // No data available, continue
}
```

### Pattern 3: Timeout Handling

```go
select {
case msg := <-ch:
    // Handle data
case <-time.After(2 * time.Second):
    // Handle timeout
}
```

### Pattern 4: Channel Priority (Not Guaranteed)

```go
// Note: Go's select is random, not priority-based
select {
case msg := <-highPriority:
    // Process high priority
case msg := <-lowPriority:
    // Process low priority
}
```

## Key Characteristics

### 1. Non-Blocking After Ready

- Select blocks until at least one case is ready
- Once ready, executes immediately
- Does not wait for all cases to be ready

### 2. Random Selection

- No priority or order preference
- Truly random when multiple cases ready
- Prevents starvation of any channel

### 3. Single Case Execution

- Only one case block executes
- Other cases ignored for this select iteration
- Would need loop for multiple case handling

## Common Use Cases

### 1. Fan-in Pattern

```go
// Merge multiple channels into processing
select {
case data := <-input1:
    process(data)
case data := <-input2:
    process(data)
case data := <-input3:
    process(data)
}
```

### 2. Circuit Breaker

```go
select {
case result := <-operationCh:
    return result
case <-time.After(timeout):
    return errors.New("operation timeout")
}
```

### 3. Cancellation

```go
select {
case data := <-dataCh:
    return data
case <-ctx.Done():
    return ctx.Err()
}
```

## Conclusion

This code demonstrates Go's `select` statement for handling multiple channel operations concurrently. The non-deterministic output occurs because both channels become ready simultaneously after the 1-second delay, causing Go's runtime to randomly choose between them. The select statement provides a powerful mechanism for coordinating multiple channel operations while maintaining fairness through random selection when multiple cases are ready.

---

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	}

	fmt.Println("End of the Program")
}
```

```bash
Received from ch1: 1
End of the Program
```

**---**

## if you want to receive both the messages?

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second)

	for range 2{
		select {

	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	default:
		fmt.Println("No Channels Ready...")
	}
	}

	fmt.Println("End of the Program")
}
```

```bash
Received from ch2: 2
Received from ch1: 1
End of the Program
```

---

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second)

	for range 2{
		select {

	case msg := <- ch1:
		fmt.Println("Received from ch1:", msg)

	case msg := <- ch2:
		fmt.Println("Received from ch2:", msg)

	default:
		fmt.Println("No Channels Ready...")
	}
	}

	fmt.Println("End of the Program")
}
```

```bash
Received from ch1: 1
Received from ch2: 2
End of the Program
```

**---**

## timeout -> a great functionality that select statement offers us.

It is not the select statement, but we can use select with timeouts. It makes our job easier to implement cancellation of channels.
What we do is we use **time.after()** to implement timeouts, providing us a way to handle operations that take too long.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
        close(ch)
	}()

	select {
	case msg := <- ch:
		fmt.Println("Received: ", msg)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout.")
	}
}
```

```bash
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
	ch := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch <- 1
	}()

	select {
	case msg := <- ch:
		fmt.Println("Received: ", msg)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout.")
	}
}
```

```bash
Timeout.
```

**---**

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <- ch:
			if !ok {
				fmt.Println("Channel Closed")
				//? cleanup activites if any
				return
			}
			fmt.Println("Received: ", msg)
		}
	}
}
```

```bash
Received:  1
Channel Closed
```

---

**Better Approach (Using for range)**

```go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        ch <- 1
        close(ch)
    }()

    for msg := range ch {
        fmt.Println("Received: ", msg)
    }

    fmt.Println("Channel Closed")
    // cleanup activities if any
}
```

---

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg := <- ch:
			fmt.Println("Received: ", msg)
		}
	}
}
```
```bash
Received: 0
infinite time
```

**---**

## How closing of channel Works
Closing a channel signals that no more values will be sent on that channel. This is important for receiving end to know when to stop waiting for new data.