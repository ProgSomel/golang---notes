# 103 - Goroutines -> are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value.

**Goroutines do not stop the program flow and are non blocking**

Go routines are light-weight threads managed by the go runtime.
They enable concurrent execution of functions, allowing you to perform multiple tasks concurrently within a single Go Program.

Go routines are one of the key features of go, making it easy to write concurrent and parallel programs.

We use go routines to efficiently handle parallel tasks such as input-output operations, calculations and more.

Go routines provide us a way to perform tasks concurrently without manually managing threads.

## How to create go routines?

We use the go keyword to start a new go routine. You add the go keyword preceding the function and execute that function immediately in the main function.

```go
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

func main() {
	sayHello()
}
```

```bash
→ 103 - Goroutines go run goroutines.go
Hello from Go Routine
```

---

```go
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

func main() {
	go sayHello()
}
```

```bash
103 - Goroutines go run goroutines.go
103 - Goroutines
```

Nothing is printed.
So what go keyword does is that it extracts the function from the main thread and it takes it into the background.

Let's consider, a room that is the main room(the main thread, the main function, the main function runs into the main thread), go routine extracts this function, takes it out of this main room. And it will bring this function back to the main thread once it has finished.

```go
func sayHello() {
	//? It will wait 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Go Routine")
}
```

Once the say hello function spends one second

```go
time.Sleep(1 * time.Second)
```

and then prints out this line

```go
fmt.Println("Hello from Go Routine")
```

after that it will be brought back to the main thread.

Now when we ran this program three times, we did not get any output that happend because the main function does not have anything else apart from this go routine. Go routine went back and the main function did not have anything, it reached the next line, next line did not have anything and then it reached the end of the main function. As soon as the main function ends, our program ends. So our go program ended before the go routine could come back to the main thread. The main thread of our program was over, was finished. our program ceased to exist. Our program was not existent, was not alive by the time sayHello finished, but say hello was in the memory, but its reference, its reference point, its source was no longer in the memory, so it could not come back, hence it left in the memory, now it is just out there in the void of our computer's memory, and it will be overwritten by something else now. Our go routine still exists in the memory, because we did not let it complete but our main function is complete.

**Now what do we do?** We wait for this go routine to finish. And there are many ways that we wait for the go routine to finish.
A basic way a basic method to wait for the go routine to finish is **time.Sleep()**

```go
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

func main() {
	go sayHello()
	time.Sleep(2 * time.Second)
}
```

```bash
Hello from Go Routine
```

---

```go
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

func main() {
	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	time.Sleep(2 * time.Second)
}
```

```bash
Beginning Program.
After sayHello function.
Hello from Go Routine
```

**-----------------------------------------------------------------------------------------------------------------------------**

There may be many goroutines that can be run into our program. **all running concurrently**.
It exits, a goroutine eixts when the function's running is complete. Goroutine contains a function and if the function completes then it exits.

**-----------------------------------------------------------------------------------------------------------------------------**

## It is the **go runtime** that manages **goroutine scheduling** and **execution**.

### goroutine scheduling

Go routine scheduling is managed by the go runtime scheduler. It uses **m:n** scheduling model.
**M goroutines run on N OS threads**.
Another thing that the goroutine scheduling does is that it efficiently multiplexes goroutines into available threads.

Go uses a m, n scheduling model where m goroutines are mapped into n operating system threads.
This model allows the go runtime to manage many goroutines with fewer operating system threads improving efficiency and scalability.

We leaned that the go routine scheduler efficiently **multiplexes** go routines into available threads.

**-----------------------------------------------------------------------------------------------------------------------------**

## What is multiplexing?

Multiplexing is like switching, instead of saying the word Multiplexing you can say switching.
The goroutine scheduler multiplexes or switches goroutines into available OS threads.
The scheduler is switching goroutines into the available operating system threads, this means it can run many goroutines on a limited number of threads by dynamically scheduling and resheduling goroutines as neede.
And this efficient use of resources ensures high concurrency and performance.
**Now let's understand it in a general broader terms:**
Imagine having multiple tasks:
You are running a company and you have many tasks that you need done. And you need those tasks to be done by your workers. You have employees. And you have few employees and many tasks. You have many tasks from your clients, and those tasks need to be done to be finished by your employees. Your employees limited number of employees, your few workers. those few workers, those limited number of employees are your operating system's **threads** and your homongous number of tasks are **goroutines**. The scheduler assigns tasks. You are the scheduler, you are the go routine scheduler, you are go runtime. So you assign tasks to workers to employees. And these tasks are dynamically reassigned to them based on workload. It means if one finishes a task it one of the workers, one of the employees, which is a thread. One of the employees finished one task, assign another task, and then another one finished another task, assign another task to that thread, that employee. This ensures that all tasks are completed efficiently without overloading any single worker. We are not assigning tasks on a number basis like we have 50 tasks and we assign those tasks equally between five threads, 10 tasks per thread. But we do not know how much time each task is going to take, So one employee one thread could be overburdened because it has bigger tasks and another employee another thread could get free within half the time because the tasks that he was assigned were smaller tasks, they took less time and now he has Nothing to do. That's not efficiency, that's not efficient for us, that's not efficient for the company. So for that reason we have task scheduler a go routine runtime scheduler that only assigns tasks once a thread has finished processing the previous task.
This is something that we do not have to manually be involved in. Previously there are some older languages where we needed this to be done manually, but in go the go runtime helps us manage this automatically and very efficiently so that we never have to get into scheduling or we never have to bother about go routine scheduling. We have our pease of mind and we can focus on coding our application.

## goroutine execution

goroutine execution is concurrent in nature. goroutine runs independently and concurrently. They are not in the main thread. They are seperate and independent running functions. Our main thread is seperate and go routines are running separately.

**-----------------------------------------------------------------------------------------------------------------------------**

```go
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
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(letter)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello function.")
	printNumbers()
	printLetters()
	time.Sleep(2 * time.Second)
}
```

```bash
Beginning Program.
After sayHello function.
0
1
2
3
4
97
98
99
Hello from Go Routine
100
101
```

---

```go
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
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(letter)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello function.")
	go printNumbers()
	go printLetters()
	time.Sleep(2 * time.Second)
}
```

```bash
Beginning Program.
After sayHello function.
0
97
1
2
98
3
99
4
100
101
Hello from Go Routine
```

---

```go
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

func main() {
	fmt.Println("Beginning Program.")
	go sayHello()
	fmt.Println("After sayHello function.")
	go printNumbers()
	go printLetters()
	time.Sleep(2 * time.Second)
}
```

```bash
Beginning Program.
After sayHello function.
Number:  0 2025-10-06 19:05:19.810448 +0600 +06 m=+0.000211251
a 2025-10-06 19:05:19.810419 +0600 +06 m=+0.000182167
Number:  1 2025-10-06 19:05:19.915531 +0600 +06 m=+0.105294001
b 2025-10-06 19:05:20.015457 +0600 +06 m=+0.205218584
Number:  2 2025-10-06 19:05:20.015799 +0600 +06 m=+0.205560584
Number:  3 2025-10-06 19:05:20.116868 +0600 +06 m=+0.306629417
c 2025-10-06 19:05:20.21581 +0600 +06 m=+0.405570251
Number:  4 2025-10-06 19:05:20.217094 +0600 +06 m=+0.406854459
d 2025-10-06 19:05:20.416714 +0600 +06 m=+0.606472584
e 2025-10-06 19:05:20.617024 +0600 +06 m=+0.806780626
Hello from Go Routine
```

**-----------------------------------------------------------------------------------------------------------------------------**

## Concurrency and Parallelism

**Concurrency** means multiple tasks progress simultaneously and not necessarily at the same time.
But **Parallelism** states that tasks are executed literally at the same time on multiple **processors**.

Goroutines facilitate Concurrency and the go runtime schedules them across available CPUs for Parallelism when possible.

**Key Concepts:**
**Concurrency = "Dealing with multiple things at once"**

- Managing multiple tasks
- Tasks can start, pause, and resume
- One CPU can handle multiple tasks by switching between them

**Parallelism = "Doing multiple things at once"**

- Actually executing multiple tasks simultaneously
- Requires multiple CPUs/cores

### Simple Analogy:

**Concurrency (One Person, Multiple Tasks):**

```bash
You (1 CPU) handling:
- Cooking (start rice, then while it cooks...)
- Answering phone (put cooking on hold)
- Checking email (put phone on hold)
- Back to cooking (rice is done)
```

**Parallelism (Multiple People, Multiple Tasks):**

```bash
You (CPU 1): Cooking
Your friend (CPU 2): Answering phone
Your sibling (CPU 3): Checking email
All happening at the SAME TIME
```

### How Go Works:

**1. Goroutines = Concurrency**

```go
func main() {
    // Start 3 goroutines (concurrent tasks)
    go task1() // "Hey, start this task"
    go task2() // "Hey, start this task too"
    go task3() // "And this one"

    // All 3 can be managed concurrently
    time.Sleep(5 * time.Second)
}

func task1() {
    for i := 0; i < 5; i++ {
        fmt.Println("Task 1:", i)
        time.Sleep(1 * time.Second)
    }
}
```

**2. Go Runtime = Smart Scheduler**
**The Go runtime is like a smart manager:**

```bash
"I have 1000 goroutines to run and 4 CPU cores available.
Let me distribute these goroutines across the cores efficiently."
```

### Visual Example:

#### Scenario: 6 Goroutines, 2 CPUs

**Concurrency (What You Write):**

```go
go task1()  // Goroutine 1
go task2()  // Goroutine 2
go task3()  // Goroutine 3
go task4()  // Goroutine 4
go task5()  // Goroutine 5
go task6()  // Goroutine 6
```

**Parallelism (What Go Runtime Does):**

```bash
CPU 1               CPU 2
-----               -----
task1 ←→ task3      task2 ←→ task5
   ↕                  ↕
task4               task6

Both CPUs work simultaneously (parallel)
Each CPU switches between tasks (concurrent)
```

## Detailed Example:

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    // Check how many CPUs you have
    fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

    // Create multiple goroutines
    var wg sync.WaitGroup

    for i := 1; i <= 6; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers done!")
}

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i < 3; i++ {
        fmt.Printf("Worker %d: task %d (running on CPU thread)\n", id, i)
        time.Sleep(500 * time.Millisecond)
    }
}
```

```bash
Number of CPUs: 10
Worker 6: task 0 (running on CPU thread)
Worker 4: task 0 (running on CPU thread)
Worker 5: task 0 (running on CPU thread)
Worker 2: task 0 (running on CPU thread)
Worker 3: task 0 (running on CPU thread)
Worker 1: task 0 (running on CPU thread)
Worker 4: task 1 (running on CPU thread)
Worker 5: task 1 (running on CPU thread)
Worker 6: task 1 (running on CPU thread)
Worker 3: task 1 (running on CPU thread)
Worker 1: task 1 (running on CPU thread)
Worker 2: task 1 (running on CPU thread)
Worker 3: task 2 (running on CPU thread)
Worker 5: task 2 (running on CPU thread)
Worker 6: task 2 (running on CPU thread)
Worker 4: task 2 (running on CPU thread)
Worker 1: task 2 (running on CPU thread)
Worker 2: task 2 (running on CPU thread)
All workers done!
```

## Key Points:

**1. You Write Concurrent Code:**

```go
go doSomething()  // "I want this to run concurrently"
```

**2. Go Runtime Enables Parallelism:**

```bash
"I'll run your concurrent goroutines on multiple CPUs when possible"
```

**3. It's Automatic:**

- You don't manually assign goroutines to CPUs
- Go runtime handles the scheduling
- More CPUs = more potential for parallelism

## Real-World Analogy:

**You (Programmer)**: "I need these 10 tasks handled concurrently"
**Go Runtime (Manager)**: "I have 4 workers (CPUs) available. I'll distribute these 10 tasks among the 4 workers so they can work in parallel, and each worker can switch between tasks as needed."

## The Magic:

```go
go// You write this (concurrency)
for i := 0; i < 1000; i++ {
    go processData(i)
}

// Go runtime does this (parallelism)
// Spreads 1000 goroutines across your 4/8/16 CPU cores
// Each core runs multiple goroutines concurrently
// All cores work in parallel
```

## Bottom Line:

- **Goroutines** = Your way to write concurrent code
- **Go Runtime** = Smart system that runs your goroutines in parallel across multiple CPUs automatically

You get the benefits of both concurrency (managing multiple tasks) and parallelism (using multiple CPUs) without having to worry about the complex details!

**-----------------------------------------------------------------------------------------------------------------------------**

## Handling errors in Goroutines
When we are handling errors in go routines, we have something, we have a concept called error propagation.

Goroutines execute functions **concurrently**. In that case errors need to be communicated back to the main thread. Use return values or shared error variable if not using channels. If we are not using channels we can use shared error variable. Oherwise when we will start using channels, we will see how we can use return values and handle errors from goroutines.

```go
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
    
    if err != nil {
        fmt.Println("Error: ", err)
    }else{
        fmt.Println("Work completed successfully")
    }
	time.Sleep(2 * time.Second)
}
```
```bash
After sayHello function.
Work completed successfully
a 2025-10-07 17:10:11.256041 +0600 +06 m=+0.001050584
Number:  0 2025-10-07 17:10:11.256025 +0600 +06 m=+0.001034709
Number:  1 2025-10-07 17:10:11.363288 +0600 +06 m=+0.108297834
Number:  2 2025-10-07 17:10:11.463997 +0600 +06 m=+0.209007042
b 2025-10-07 17:10:11.464168 +0600 +06 m=+0.209178542
Number:  3 2025-10-07 17:10:11.565115 +0600 +06 m=+0.310126042
c 2025-10-07 17:10:11.665448 +0600 +06 m=+0.410459459
Number:  4 2025-10-07 17:10:11.665434 +0600 +06 m=+0.410445626
d 2025-10-07 17:10:11.866854 +0600 +06 m=+0.611866667
e 2025-10-07 17:10:12.068709 +0600 +06 m=+0.813722042
Hello from Go Routine
```

-----------------------------------------------------------------------------------------------------------------------------

```go
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
```
```bash
Beginning Program.
After sayHello function.
a 2025-10-07 17:11:57.185094 +0600 +06 m=+0.000207418
Number:  0 2025-10-07 17:11:57.185077 +0600 +06 m=+0.000190126
Number:  1 2025-10-07 17:11:57.289493 +0600 +06 m=+0.104607335
Number:  2 2025-10-07 17:11:57.389887 +0600 +06 m=+0.205001168
b 2025-10-07 17:11:57.390369 +0600 +06 m=+0.205483043
Number:  3 2025-10-07 17:11:57.491032 +0600 +06 m=+0.306147460
Number:  4 2025-10-07 17:11:57.591605 +0600 +06 m=+0.406720876
c 2025-10-07 17:11:57.591624 +0600 +06 m=+0.406739043
d 2025-10-07 17:11:57.793969 +0600 +06 m=+0.609085251
e 2025-10-07 17:11:57.995115 +0600 +06 m=+0.810232418
Hello from Go Routine
Error:  an Error occured in doWork
```

**-----------------------------------------------------------------------------------------------------------------------------**