# variable

A **variable** is like a **Container** where you can **store data** — numbers, text, etc.
**In Go, you must declare a variable **before** you use it.**

## Declaring a Variable

### \*In Go language variables are created in two different ways:

## 1. **Using var keyword**: In Go language, variables are created using `var` keyword of a particular type, connected with name and provide its initial value. Syntax:

```go
var variable_name type = expression
```

```go
package main

import "fmt"

var name ="John Doe"; //? Variable One

func main() {
    var fruit = "Apple" //? Varibale Two
	fmt.Println(name);
	fmt.Println(fruit);
}
```

```bash
//output
John Doe
Apple

```

### **Important Points**:

- In the above syntax, either type or = expression can be omitted, but not both.
- If the = expression is omitted, then the variable value is determined by its type's default value. The default value is usually 0.
- If the type is removed, then the type of the variable is determined by the value-initialize in the expression. Example:

```go
// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

// Variable declared and
// initialized without the
// explicit type
var myvariable1 = 20
var myvariable2 = "GeeksforGeeks"
var myvariable3 = 34.80

// Display the value and the
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
                                  myvariable1)
fmt.Printf("The type of myvariable1 is : %T\n",
                                  myvariable1)
fmt.Printf("The value of myvariable2 is : %s\n",
                                      myvariable2)
fmt.Printf("The type of myvariable2 is : %T\n",
                                  myvariable2)
fmt.Printf("The value of myvariable3 is : %f\n",
                                      myvariable3)
fmt.Printf("The type of myvariable3 is : %T\n",
                                  myvariable3)
}
```
```bash
The value of myvariable1 is : 20
The type of myvariable1 is : int
The value of myvariable2 is : GeeksforGeeks
The type of myvariable2 is : string
The value of myvariable3 is : 34.800000
The type of myvariable3 is : float64
```
- If the expression is removed, then the variable holds zero-value for the type like zero for the number, false for Booleans, "" for strings, and nil for interface and reference type. So, there is no such concept of an uninitialized variable in Go language. Example:
```go
// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

// Variable declared and
// initialized without expression
var myvariable1 int
var myvariable2 string
var myvariable3 float64

// Display the zero-value of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
                                     myvariable1)

fmt.Printf("The value of myvariable2 is : %s\n",
                                     myvariable2)

fmt.Printf("The value of myvariable3 is : %f",
                                     myvariable3)
}
```
```bash
The value of myvariable1 is : 0
The value of myvariable2 is :
The value of myvariable3 is : 0.000000
```
- If you use type, then you are allowed to declare multiple variables of the same type in the single declaration. Example:
```go
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Multiple variables of the same type
// are declared and initialized
// in the single line
var myvariable1, myvariable2, myvariable3 int = 2, 454, 67

// Display the values of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
myvariable1)

fmt.Printf("The value of myvariable2 is : %d\n",
myvariable2)

fmt.Printf("The value of myvariable3 is : %d",
myvariable3)
}
```
```bash
The value of myvariable1 is : 2
The value of myvariable2 is : 454
The value of myvariable3 is : 67
````

- If you remove type, then you are allowed to declare multiple variables of a different type in the single declaration. The type of variables is determined by the initialized values. Example:

```go
// Go program to illustrate
// concept of variable
package main
import "fmt"

func main() {

// Multiple variables of different types
// are declared and initialized in the single line
var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

// Display the value and
// type of the variables
fmt.Printf("The value of myvariable1 is : %d\n",
                                 myvariable1)

fmt.Printf("The type of myvariable1 is : %T\n",
                                myvariable1)

fmt.Printf("\nThe value of myvariable2 is : %s\n",
                                  myvariable2)

fmt.Printf("The type of myvariable2 is : %T\n",
                                myvariable2)

fmt.Printf("\nThe value of myvariable3 is : %f\n",
                                   myvariable3)

fmt.Printf("The type of myvariable3 is : %T\n",
                                myvariable3)
}
```

```bash
The value of myvariable1 is : 2
The type of myvariable1 is : int

The value of myvariable2 is : GFG
The type of myvariable2 is : string

The value of myvariable3 is : 67.560000
The type of myvariable3 is : float64
```

- You are allowed to initialize a set of variables by the calling function that returns multiple values. Example:


```go
// Here, os.Open function return a
// file in i variable and an error
// in j variable
var i, j = os.Open(name)
```
## 2. **Using short variable declaration**: The local variables which are declared and initialize in the functions are declared by using short variable declaration. Syntax:
```go
variable_name:= expression
```
**Note: Please don't confuse in between := and = as := is a declaration and = is assignment.**
### **Important Points:**
In the above expression, the type of the variable is determined by the type of the expression. Example:
```go
     // Go program to illustrate
     // concept of variable
     package main
     import "fmt"

     func main() {

     // Using short variable declaration
     myvar1 := 39
     myvar2 := "GeeksforGeeks"
     myvar3 := 34.67

     // Display the value and type of the variables
     fmt.Printf("The value of myvar1 is : %d\n", myvar1)
     fmt.Printf("The type of myvar1 is : %T\n", myvar1)

     fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
     fmt.Printf("The type of myvar2 is : %T\n", myvar2)

     fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
     fmt.Printf("The type of myvar3 is : %T\n", myvar3)
     }
```
```bash
     The value of myvar1 is : 39
     The type of myvar1 is : int

     The value of myvar2 is : GeeksforGeeks
     The type of myvar2 is : string

     The value of myvar3 is : 34.670000
     The type of myvar3 is : float64
```
- Most of the local variables are declared and initialized by using short variable declarations due to their brevity and flexibility.
- The var declaration of variables are used for those local variables which need an explicit type that differs from the initializer expression, or for those variables whose values are assigned later and the initialized value is unimportant.
- Using short variable declaration you are allowed to declare multiple variables in the single declaration
Example:
```go
       // Go program to illustrate
       // concept of variable
       package main
       import "fmt"

       func main() {

       // Using short variable declaration
       // Multiple variables of same types
       // are declared and initialized in
       // the single line
       myvar1, myvar2, myvar3 := 800, 34, 56

       // Display the value and
       // type of the variables
       fmt.Printf("The value of myvar1 is : %d\n", myvar1)
       fmt.Printf("The type of myvar1 is : %T\n", myvar1)

       fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
       fmt.Printf("The type of myvar2 is : %T\n", myvar2)

       fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
       fmt.Printf("The type of myvar3 is : %T\n", myvar3)
       }
```
```bash
       The value of myvar1 is : 800
       The type of myvar1 is : int

       The value of myvar2 is : 34
       The type of myvar2 is : int

       The value of myvar3 is : 56
       The type of myvar3 is : int
```
- In a short variable declaration, you are allowed to initialize a set of variables by the calling function that returns multiple values. Example:
```go
       // Here, os.Open function return
       // a file in i variable and an
       // error in j variable
       i, j := os.Open(name)
```

- A short variable declaration acts like an assignment only when for those variables that are already declared in the same lexical block. The variables that are declared in the outer block are ignored. And at least one variable is a new variable out of these two variables as shown in the below example. Example:
```go
        // Go program to illustrate
        // concept of variable
        package main
        import "fmt"

        func main() {

        // Using short variable declaration
        // Here, short variable declaration acts
        // as an assignment for myvar2 variable
        // because same variable present in the same block
        // so the value of myvar2 is changed from 45 to 100
        myvar1, myvar2 := 39, 45
        myvar3, myvar2 := 45, 100

        // If you try to run the commented lines,
        // then compiler will gives error because
        // these variables are already defined
        // myvar1, myvar2 := 43, 47
        // myvar2:= 200

        // Display the values of the variables
        fmt.Printf("The value of myvar1 and myvar2 is : %d %d\n",
                                          myvar1, myvar2)

        fmt.Printf("The value of myvar3 and myvar2 is : %d %d\n",
                                          myvar3, myvar2)
        }
```
```bash
        The value of myvar1 and myvar2 is : 39 100
        The value of myvar3 and myvar2 is : 45 100
```

- Using short variable declaration you are allowed to declare multiple variables of different types in the single declaration. The type of these variables are determined by the expression. Example:
```go
        // Go program to illustrate
        // concept of variable
        package main
        import "fmt"

        func main() {

        // Using short variable declaration
        // Multiple variables of different types
        // are declared and initialized in the single line
        myvar1, myvar2, myvar3 := 800, "Geeks", 47.56

        // Display the value and type of the variables
        fmt.Printf("The value of myvar1 is : %d\n", myvar1)
        fmt.Printf("The type of myvar1 is : %T\n", myvar1)

        fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
        fmt.Printf("The type of myvar2 is : %T\n", myvar2)

        fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
        fmt.Printf("The type of myvar3 is : %T\n", myvar3)
        }
```
```bash
        The value of myvar1 is : 800
        The type of myvar1 is : int

        The value of myvar2 is : Geeks
        The type of myvar2 is : string

        The value of myvar3 is : 47.560000
        The type of myvar3 is : float64
```

**---------------------------------------------------------------------------------------------------------------------**
## In Go, top-level (global) code can only contain declarations, not assignments.
```go
        package main

        import "fmt"

        var name string; //? Variable One
        name = "John Doe";

        func main() {
        var fruit string //? Varibale Two
        fruit = "Apple";
        fmt.Println(name);
        fmt.Println(fruit);
        }
```
**❌ The Problem:**

You wrote this at the **top-level** of the file (outside `func main()`):

```go
        var name string; //? Variable One
        name = "John Doe";
```
**In **Go**, top-level (global) code can **only contain declarations**, not assignments.**

**✅ Solution**:

You must either:

### Option 1: Declare and assign **in one line**:

```go
        var name string = "John Doe"
```

### Option 2: Use short form only **inside** a function:

```go
        name := "John Doe" // ✅ only allowed inside functions
        //
```

### ✅ Corrected Code:

```go
        package main

        import "fmt"

        var name string = "John Doe" // ✅ Declare + assign at top level
        func main() {
        var fruit string //? Variable Two
        fruit = "Apple"  // ✅ Assignment inside function

        fmt.Println(name)
        fmt.Println(fruit)
        }
```
```go
        package main

        import "fmt"

        var name string; //? Declare Here

        func main() {
        name="John Doe"
        var fruit string //? Assign Here
        fruit = "Apple";
        fmt.Println(name);
        fmt.Println(fruit);
        }
```

**⏳ Why This Happens?**

Go is a **compiled language**, and outside of functions, it needs everything to be **declarative and static**.

That means:
✅ **Variable declarations**
❌ **Assignments** (like `name = "John Doe"`) without a declaration

**🎯 Bonus Tip**:
You can even shorten your top-level declaration like this:
```go
        var name = "John Doe" // Go will infer it's a string
```

## Important Links
[nesohq](https://nesohq.github.io/bgce-archive/)
[dev](https://dev.to/diwakarkashyap/variables-in-go-golang-34b4)
[tutorialspoint](https://www.tutorialspoint.com/go/go_variables.htm)