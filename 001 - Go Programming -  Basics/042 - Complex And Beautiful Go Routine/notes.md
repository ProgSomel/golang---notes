# Complex And Beautiful Go Routine

**Go Routine is virtual thread. Logically it works like thread but it is not actually a thread also thread is virtual process(because it works like process but actually not a process) also process is called virtual computer(because it has no physical existing like physical computer)**

## Computer -> process -> thread -> Routine

```go
package main

import "fmt"

func printHello(num int){
	fmt.Println("Hello, Habib", num)
}

func main(){
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)
}
```

```bash
//output
Hello, Habib 1
Hello, Habib 2
Hello, Habib 3
Hello, Habib 4
Hello, Habib 5
```

---

```go
package main

import (
	"fmt"
	"time"
)

func printHello(num int){
	fmt.Println("Hello, Habib", num)
}

func main(){
	//? Here go will create go routine
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)
	time.Sleep(5*time.Second) //? 5 * 1 second = 5 seconds
}
```

```bash
// will give different different output flow
Hello, Habib 5
Hello, Habib 4
Hello, Habib 2
Hello, Habib 1
Hello, Habib 3
```

---

```go
package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int){
	fmt.Println("Hello, Habib", num)
}

func main(){
	//? Here go will create go routine
	fmt.Println("Hello")
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

	fmt.Println(a, " ", p)
	time.Sleep(5*time.Second) //? 5 * 1 second = 5 seconds
}
```

## Simulation of the above code:

**1. Compilation Phase(compile time):**

```bash
   go build filename.go
```

it will create a main file --> which is a binary executable file, it will create under a folder and will store in hard-disk
this main, binary file will have a code segment, which will have:

p = 11 //? const variable
printHello = printHello(num int){...}
main = main(){...}

after that when, we run ./main, the binary executable file will load into ram(memory) and program will run then.

**2. Execution Phase**
after that when, we run ./main, the binary executable file will load into ram(memory) and program will run then.
A process will be created. Process will have:

- **Code segment:** Contains all the functions and executable instructions.
- **Data segment:** Contains global and static variables.
- **Stack:** Used for function calls and local variables.
- **Heap:** Used for dynamic memory allocation.

Process will have a initial thread(main thread), when a process will create, a thread will also be created.
when **OS Kernal** starts to execute the thread, then stack will start to be executed.
When stack executes line by line thats means, thread executes. So we can say that thread executes stack and OS Kernal executes the thread. When thread starts execution, it fells that process starts executing.

+-------------------------------------------------------------+
| MEMORY LAYOUT |
+-------------------+-------------------+---------------------+
| Code Segment | Data Segment | Heap |
|-------------------|-------------------|---------------------|
| - main() | - var a = 10 | (dynamic objects) |
| - printHello() | | |
| - time.Sleep() | | |
| - const p = 11 | | |
| (constant value | | |
| embedded as | | |
| instructions) | | |
+-------------------+-------------------+---------------------+
| | |
| | |
V V V

**Go Runtime(Go's Operating System)** -- is a virtual operating system(a os under a os, which is Logically existed). When go program runs(./main, under code segment when binary file stored and data segment become loaded), it runs a mini operating system of it's own which is called **Go Runtime**. When go program runs **go runtime** will load into ram(memory). **Main thread(Every process has a main thread)** will execute the **go runtime**. Go runtime will execute the full OS Thread and will initialize:

1. at first intializes **Go routine scheduler**
2. initialize Heap allocator
3. initialize Garbage collector
4. intialize logical processors(thread)

# **CPU** executes the code

**CPU - Processor** will run **OS**. Under **OS** there is **Kernal(core component of OS)**. **OS Kernal** will execute **thread**(OS Kernal will tell Processor to execute the thread). Thread will then execute stack. Kernal will tell what to do when. **CPU > OS > Kernel > Process > Thread > Stack**

## Logical Processor

Let CPU has 2 cores. Every core has 2 virtual processors, then there is four virtual processors.
if 4 virtual processors run OS, OS Kernal will run Thread, Thread will run stack(means go runtime). Then Go runtime will create four Logical processors. **If CPU has n processors, then go runtime will create n Logical processors(these will store under go runtime)**. These Logical processors not exist in real world, exists Logicallly.
**Process runs under operating system**. Operating system will create 4 threads(Every threads will create stack frame, so four threads will create extra fout stacks) for these four logical processors, these 4 thread are called **OS thread**

**Now process have total five threads. One is main thread and others four are supporting threads.**. These 5 threads will be tracked by **OS Kernal**. CPU runs OS and OS runs Kernal.

## Go routine scheduler

Go runtime will have scheduler
Go runtime is a mini os. So it will also have virtal processor. In this case there will be 4 virtual process as there is four os threads, these four virtual os threads will map with four os threads.
Programmer creates go routine.
If programmers create 100 go routines,
then go runtime's scheduler will decide which logical processor will run which go routine. If there is huge go routines and these logical processors can not handle it, then go runtime's scheduler will create another virtal processor , if virtal processor creates then another os thread will also be created. If os thread creates then a stack will also be created into the ram. If ram is full and there is no 8mb space left in ram, then we can not create new thread, so other left go routine can not be run. For this there will be huge load on virtual processor and application will be hanged.

**The first go routine which runs first called main go routine**. Means, under which go routine, the main function executes that go routine is called **main go routine**. Go routine is virtual thread, mini thread.

**Go routine only knows about Go's Operating System(Go runtime). It does not know about other things like: about main OS.**

At first **Go runtime(mini OS)** - will create a **Go routine**(mini thread, virtual thread, logical thread, not real thread, it is logical thread). This go routine will store on **Heap**.

When main function runs, main go routine will run. To become a go routine, it must have a stack frame(2kb) on heap. Main will be started to execute. **At first it will run fmt.Println("Hello")**

then it will go to second line: **go printHello(1)**.Here it will see **go** keyword. So it will create another go routine on heap, this routine also need a stack, So it will create a stack on heap also, under this stack a function **printHello()** will be executed and a stack frame(2KB) called printHello() stack frame will also be created under this stack(Go routine's stack).

+-------------------------------------------------------------+
| Stack (per goroutine) |
+-------------------------------------------------------------+
| (main goroutine) | (go printHello(1)) |
|----------------------|--------------------------------------|
| - local vars in main | - param num = 1 |
| | - call frame for printHello |
|----------------------|--------------------------------------|
| (go printHello(2)) | ... |
|----------------------|--------------------------------------|
| - param num = 2 | ... |
| - call frame | |
+-------------------------------------------------------------+

Legend:

- **Code Segment:** All function code and **const p** (as a literal/instruction)
- **Data Segment:** Global/static variables (e.g., `var a`)
- **Heap:** Dynamically allocated memory
- **Stack:** Each goroutine's function call frames and local variables

**---------------------------------------------------------------------------------------------------------------------------------------------------**

What Actually Happens (Step by Step):

1. You turn on your computer:

- OS kernel (e.g., Linux, Windows) loads into RAM.

2. You double-click (or run) your Go program:

- OS creates a process for your Go program binary.
- The OS allocates a main thread for the new process.
- The OS loads the program’s code, data, and global variables into the process’s virtual memory.

3. The main thread starts running:

- The first thing executed is the entry point of your program’s binary.
- For a Go program, this entry point is actually Go’s runtime startup code (not your main() yet).

4. Go runtime starts:

- The Go runtime initializes memory, garbage collector, scheduler, etc.
- It creates the main goroutine (which will call your main() function).
- The main goroutine gets its own small stack (on the heap).
- Any other goroutines you start are created and scheduled by the Go runtime, not by the OS.

5. Go runtime manages goroutines:

- The Go runtime (not the OS) handles goroutine scheduling, stack management, garbage collection, etc.
- OS only “sees” the process and its threads (not goroutines).

```go
package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int){
	fmt.Println("Hello, Habib", num)
}

func main(){
	//? Here go will create go routine
	var x = 10
	fmt.Println("hello", x)
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

	fmt.Println(a, " ", p)
	time.Sleep(5*time.Second) //? 5 * 1 second = 5 seconds
}
```

at first go runtime - creates the main goroutine (which will call your main() function).
The main goroutine gets its own small stack (on the heap). This stack will start to execute, so main() function
will be executed. First it will see **var x = 10**, x variable will store in main stack. Then **fmt.Println("hello", x)**, it will print: **hello 10**, if x was not in main stack frame, then i would check code segment or data segment, or if closure then it would check on heap memory. Next line is: **go printHello(1)**, here is go keyword, so it should run a go routine, main stackframe(for main go routine) will send this statement to **go runtime**.
Go runtime will create another thread, and will create a stack frame(called printHello stack frame) on heap. another logical processor will run this stack frame in parallel and printHello() function will run now.
main go routine will not stop, after giving **go printHello(1)** to go runtime, it will go to the next step **go printHello(2)** and will give this to **go runtime**, thus main go routine will give these give go routines to go runtime and go runtime will create these five go routines parallel. Main go routine will stole it's work and wil go to step **fmt.Println(a, " ", p)** and will execute it and will print the value of a and p, then it will go to next step which is sleep and it will sleep for 5 seconds as 5\*time.Second.
if there is sleep then after printing, hello 10 and the value of a and p the program will be terminated and thats mean **main go routine** will be deleted and others **go routine** will not be executed as there is no **main go routine**, means the full **process** will be deleted.

-----------------------------------------------------------------------------------------------------------------------------------------

A goroutine is a lightweight, user-level thread managed by the Go runtime(operating system for your Go process.).
It lets you run functions or methods concurrently (at the same time as other code) with very little overhead.

What does the Go runtime do?
The Go runtime manages the “invisible infrastructure” that makes Go’s unique features possible.
It handles things like:

## ⚙️ Go Runtime Responsibilities

| Feature           | What Go Runtime Does                                           |
|-------------------|----------------------------------------------------------------|
| **Goroutines**     | Creates, schedules, and manages goroutines (user threads)     |
| **Scheduler**      | Maps goroutines onto OS threads (M:N scheduling)              |
| **Garbage Collector** | Automatically frees unused memory                         |
| **Stack Management** | Grows/shrinks stacks for goroutines as needed              |
| **Channel Ops**    | Manages communication between goroutines                      |
| **Timers**         | Handles `time.After`, `time.Sleep`, etc.                     |
| **Panic/Recover**  | Manages error handling, panic, and recovery                   |
| **Syscalls & OS Glue** | Provides safe access to OS functions and resources       |


**--------------------------------------------------------------------------------------------------------------------------------------------**
main thread go runtime কে এক্সিকিউট করতেসে। go runtime সব go routine কে এক্সিকিউট করে including main go routine. main go routine পাঁচটা go routine কে go runtime এর কাছে দিয়ে দিসে। তারপর সে ৫ সেকেন্ড স্লিপে গেসে। স্লিপ করবে অনলী main go routine. go runtime মেইন thread এ রান হচ্ছে। সে স্লিপ করবে না। সে বাকি ৫ টা গো রুটিনকে রান করবে। মেইন গো রুটিন ৫ সেকেন্ড স্লিপ করার মধ্যে যেসব গো রুটিন রেজাল্ট শো করতে পারবে, সেগুলা র আউটপুট আমরা দেখবো। পাঁচ সেকেন্ড শেষে মাইন গো রুটিন মরে যাবে। আর মেইন গো রুটিন মরে গেলে main thread ও মরে যাবে। মানে go runtime ও মরে যাবে। আর go runtine মরে গেলে বাকি ৫ টা গো রুটিনও ধ্বংস। কে প্রসেস করল আর না করল, দেখার টাইম নাই|

**Qustion:**
ভাই যদি আমরা go routine এর মধ্যেও ৫ সেকেন্ড স্লিপ করাই মানে printHelloWorld function  এর ভিতরে, তাহলে main go routine  ৫ সেকেন্ড স্লিপ এ আছে আবার এইটা যেই go routine গুলোকে হ্যান্ডেল করার জন্য go runtime কে দিছে তারা ও প্যারালালি ৫ সেকেন্ড স্লিপ করতেছে...আমার প্রশ্ন হচ্ছে তখন ত্ব কোনো go routine তার এক্সিকিউশন শেষ করার কথা না unless তাদের স্লিপ টাইম main go routine এর ছেয়ে কম হয়!

Answer:
দুইটা বা ততোধিক গো রুটিন মেইন গো রুটিনের সাথে রান হবে কনকারেন্সি মেইনটেইন করে। যেহেতু কনকারেন্সি, কোনটা আগে এক্সিকিউট হবে গো রানটাইম দ্বারা আমরা জানি না। মেইন যদি আগে এক্সিকিউট হয়, তাইলে বাকিরা মারা পরবে। কিন্তু যদি বাকিদের মধ্যে কয়েকটা অন্তত আগে এচিকিউট হয়, তাদের আউটপুট অন্তত দেখা যাবে।