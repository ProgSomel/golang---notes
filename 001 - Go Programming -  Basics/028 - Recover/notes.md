# 028 - Recover - Always use recover with defer function
Recover is a built in function that is used to regain control of a panicking go routine. It is only useful inside defer functions and is used to manage behaviour of a panicking go routine to avoid abrupt termination.
```go
package main

import "fmt"

func process(){
	defer func(){
		if r := recover(); r!=nil{
			fmt.Println("Recovered: ", r)
		} 
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}

func main(){

	process()
	fmt.Println("Returned from process")

}
```
```bash
Start Process
Recovered:  Something went wrong
Returned from process
```

------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func process(){
	// defer func(){
	// 	if r := recover(); r!=nil{
	// 		fmt.Println("Recovered: ", r)
	// 	} 
	// }()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}

func main(){

	process()
	fmt.Println("Returned from process")

}
```
```bash
Start Process
panic: Something went wrong

goroutine 1 [running]:
main.process()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/028 - Recover/recover.go:13 +0x59
main.main()
        /Users/progsomel/Library/Mobile Documents/com~apple~CloudDocs/ProgSomel/Skills/Programming/backend-development/golang/028 - Recover/recover.go:19 +0x13
exit status 2
```

---------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func process(){
	defer func(){
		r := recover()
		if r!=nil{
			fmt.Println("Recovered: ", r)
		} 
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}

func main(){

	process()
	fmt.Println("Returned from process")

}
```