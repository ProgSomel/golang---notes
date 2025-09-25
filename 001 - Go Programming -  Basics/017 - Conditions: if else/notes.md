# Conditions: if else

```go
package main

import "fmt"


func main(){
	/*
		if condition(){
			block of code
		}
	*/
	age := 18
	if age >= 18{
		fmt.Println("You are an adult")
	}

}
```

----------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	/*
		if condition(){
			block of code
		}else if(){
			block of code
		}else{
			block of code
		}
	*/
	temperatue := 25
	if temperatue >= 30{
		fmt.Println("It's hot outside")
	}else{
		fmt.Println("It's cool outside")
	}

	score := 85

	if score >= 90{
		fmt.Println("Grade A")
	}else if score >= 80{
		fmt.Println("Grade B")
	}else if score >= 70{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Grade D")
	}

}
```
```bash
It's cool outside
Grade B
```

------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	/*
		if condition(){
			block of code
		}else if(){
			block of code
		}else{
			block of code
		}
	*/
	temperatue := 25
	if temperatue >= 30{
		fmt.Println("It's hot outside")
	}else{
		fmt.Println("It's cool outside")
	}

	score := 95

	if score >= 80{
		fmt.Println("Grade B")
	}else if score >= 90{
		fmt.Println("Grade A")
	}else if score >= 70{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Grade D")
	}

}
```
```bash
It's cool outside
Grade B
```

-------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	/*
		if condition(){
			block of code
		}else if(){
			block of code
		}else{
			block of code
		}
	*/
	temperatue := 25
	if temperatue >= 30{
		fmt.Println("It's hot outside")
	}else{
		fmt.Println("It's cool outside")
	}

	score := 85

	if score >= 80{
		fmt.Println("Grade A")
	}
	if score >= 80{
		fmt.Println("Grade B")
	}
	if score >= 70{
		fmt.Println("Grade C")
	}else{
		fmt.Println("Grade D")
	}

}
```
```bash
It's cool outside
Grade A
Grade B
Grade C
```

-----------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	/*
		if condition1(){
			block of code1
			if condition2(){
				code block2
			}
		}
	*/
	num := 18
	if num % 2 == 0{
		if num % 3 == 0{
			fmt.Println("Number is divisible by both 2 and 3")
		}else{
			fmt.Println("Number is divisible by 2 but not 3")
		}
	}else{
		fmt.Println("Number is not divisble by both 2 and 3")
	}
	
}
```
```bash
Number is divisible by both 2 and 3
```

------------------------------------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"


func main(){
	if 10%2 == 0 || 5%2 == 0{
		fmt.Println("Either 10 or 5 are Even")
	}
	if 10%2 == 0 && 5%2 == 0{
		fmt.Println("Both 10 or 5 are Even")
	}
}
```
```bash
Either 10 or 5 are Even
```