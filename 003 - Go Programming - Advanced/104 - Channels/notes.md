# 104 - Channels
Channels are a way for go routines to communicate with each other and synchronize their execution.

They provide a means to send and receive values between go routines, facilitating data exchange and coordination.

We use Channels to enable safe and efficient communication between concurrent go routines.

Using Channels helps synchronize and manage the flow of data in concurrent programs.

```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	greeting <- greetString 

	receiver := <- greeting
	fmt.Println(receiver)
}
```
```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/104 - Channels/channels.go:10 +0x49
exit status 2
```
We learned it in the defination of channels, that channels are a way to communicate between goroutines. In our program We do not have any goroutine. We learned that we can not make a channel work inside our main function, or any function, you have to create a go routine in order for the channel to receive or send values.

The issue here with this code is that it tries to send a value to a channel without having a go routine. A go routine should be there to receive from that channel and without a go routine to receive from the channel, it causes a deadlock because **channels** in go are blocking. But goroutines are non-blocking. They are extracted away from the main thread, the main execution thread of our application where the main function is running will continue to run seamlessly in a non-blocking way if we have goroutines. If we have a function here, then that function, if it is not declared with a go keyword
```go
greeting <- greetString
```
it will block the execution of the rest of the statements after that function until the time that function is complete. But if we use a go keyword, that function is extracted out of the main thread, and then the next statements will continue to run before that function is even processed. So that's the non-blocking mechanism and similarly like a function that blocks the execution flow of our main function, a channel will also block the execution of our main function, of our main thread. So that's why we need to receive values into a channel inside a go routine so that it does not block the main execution thread.

```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
	}()

	receiver := <- greeting
	fmt.Println(receiver)
}
```
```bash
Hello
```
We have the output in our terminal. **How is that?**
Right now there is still another concept that is kind of hidden from us in this, because we have declared greeting the channel that is receiveing value into it inside a go function. But channels, they communicate between go routines and **receiver** is not inside a go routine.
```go
receiver := <- greeting
```
The receiver is receiving values from a channel and again receiver is going to receive continous values. **So how is that not blocking?**
Because **receiver** is part of the **main go routine**, the main thread, the main execution thread, **our main function is a go routine**, technically it is a go routine because it is running ocntinously. It is a function and it is the main function of our application, that's why we call it main thread, the main execution thread and the main go routine. And that's how this channel is communicating between this go 
```go
go func() {
	greeting <- greetString
}()
```
routine and the **main go routine**. So the Channel is communicating between different **go routines** here. Receiver is not just an independent receiver, it is a receiver inside another go routine.
**Goroutine = thread, Simple...**
And that makes the **greeting channel** communicate between two go routines.
**So always remember that even if we do not have any go routine we still have a main go routine always.** And that's why we have been using a go routine ever since our first example, which was the main goroutine.

Now you may believe that receiveing from a channel
```go
receiver := <- greeting
```
is non-blocking, but that is not the case. Receiveing from a channel is also blocking and if there is no value to receive, then it will wait for a value to be received and next line
```go
fmt.Println(receiver)
```
will not be executed until the time it receives a value.

Let's demonstrate that:
```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	// greetString := "Hello"

	// go func() {
	// 	greeting <- greetString
	// }()

	receiver := <- greeting
	fmt.Println(receiver)
	fmt.Println("End of Program")
}
```
```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Study/Programming/golang---notes/003 - Go Programming - Advanced/104 - Channels/channels.go:14 +0x35
exit status 2
```
Here we are getting a fatal error, the same error that all go routines are sleep. It did not give us and error that there is no value to receive. It gave us error anout go routine, the same error that we got earlier when we used greeting channel to receive values:
```go
greeting <- greetString
```
And now this time we are using greeting channel to send values to another variable or receiver:
```go
receiver := <- greeting
```

**----------------------------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
	}()

	receiver := <- greeting
	fmt.Println(receiver)

	fmt.Println("End of Program")
}
```
What's happening here, is we have kept our example simple. But in reality this 
```go
greeting <- greetString
```
greeting channel is going to be an open channel constantly open channel, receiving continous values. Let's say we are accessing a weather API. And that weather API is sending us data every half an hour or maybe every hour about wind speed, about temperatures, about clouds, if it is over shadow or if it is cloudy, or if it's clear weather about humidity and a lot of parameters that weather APIs use. We get that data periodically. Maybe not just every hour, maybe every five minutes. So in that case, we use a go routine and we declare a channel
```go
go func() {
	greeting <- greetString
}()
```
and then we continuosly receiving data from that API because that API will also have an open stream, it maybe from a different language. They may have configured their API into a different language. So they may have a stream or if it is a go API from the weather application from the weather API, then they may also be using channels. Let's use a general term, a widely used term which is stream, stream has been established into programming for a very long time and it is a popular term to address a continuous flow. So we will get a continuous stream of weather data from that weather API, and we will use our channels in Go to receive that continuous flow of data from the weather API.

and Receiver
```go
receiver := <- greeting
```
once it is declared it receives value and then moves forward. It receives the value that was sent to greeting and then it moves forward to the next statement. As soon as it receives a value, a piece of value, it can be just one string, it can be just one integer, or it can be a struct. We can make a channel of type a struct that we create maybe person struct or employee struct and then we can pass an instance of that person struct or instance of the employee struct, and it will contain all the values depending on the fields that were declared in that struct. And a channel can be a list also. A channel can carry a list of structs because that's also a type. A list is means a slice of any type or an array of any type. Our receiver will contain the complete array and will process that array according to our requirements.

Now we are sending another value ot our channel
```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"
	}()

	receiver := <- greeting
	fmt.Println(receiver)

	fmt.Println("End of Program")
}
```
```bash
Hello
End of Program
```
We are only getting an output of hello and then the end of the program. The receiver is not receiving the second value. We should use like this:
```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"
	}()

	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)

	fmt.Println("End of Program")
}
```
```bash
Hello
world
End of Program
```

**----------------------------------------------------------------------------------------------------------------------------**

## Now, what will happen if you may be thinking, how about we just declare the receiver inside a go function.
```go
package main

import "fmt"

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
	}()

	go func(){
		receiver := <- greeting
		fmt.Println(receiver)
	}()

	fmt.Println("End of Program")
}
```
Now the channel which is receiving values is in a separate go routine.
```bash
End of Program
```

To see Hello, we need to do this:
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
	}()

	go func(){
		receiver := <- greeting
		fmt.Println(receiver)
	}()

	fmt.Println("End of Program")
	time.Sleep(2 * time.Second)
}
```
```bash
End of Program
Hello
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
	}()

	go func(){
		receiver := <- greeting
		fmt.Println(receiver)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")
}
```
```bash
Hello
End of Program
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"
	}()

	go func(){
		receiver := <- greeting
		fmt.Println(receiver)
		receiver = <- greeting
		fmt.Println(receiver)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")
}
```
```bash
Hello
world
End of Program
```

**----------------------------------------------------------------------------------------------------------------------------**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"

		for _, c := range "abcde"{
			greeting <- "Alphabet: " + string(c)
		}
	}()
	
	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)
	
	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")
}
```
```bash
Hello
world
Alphabet: a
Alphabet: b
Alphabet: c
Alphabet: d
Alphabet: e
End of Program
```

----------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
)

func main() {
	//? variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString
		greeting <- "world"

		for _, c := range "abcde"{
			greeting <- "Alphabet: " + string(c)
		}
	}()
	
	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)
	
	for range 5 {
		rcvr := <- greeting
		fmt.Println(rcvr)
	}

	fmt.Println("End of Program")
}
```
```bash
Hello
world
Alphabet: a
Alphabet: b
Alphabet: c
Alphabet: d
Alphabet: e
End of Program
```