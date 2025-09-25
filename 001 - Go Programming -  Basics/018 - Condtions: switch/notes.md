# Conditions: Switch --- in other languages need(switch, case, break(must have after every case if you do not want to excute the next step if the valid excute is executed), default), but in go language needs(switch, case)
> In go we can use falltrough. If we use fall through, then what happen is, that as soon as value one matches the expression, it will still go to the next case and check if expression matches value two. And if there is no fall through in this code block, then it will not move on to the next case, but if there is falltrough then it will check the next case as well.<br>

```go
package main

import "fmt"


func main(){
    fruit := "apple"

	switch fruit{
	case "apple":
		fmt.Println("It's an Apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown Fruit!")
	}

}
```
```bash
It's an Apple
```

------------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
    day := "Monday"

	switch day{
	case "Monday", "Tuesday", "Wednessday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid Day")

	}

}
```