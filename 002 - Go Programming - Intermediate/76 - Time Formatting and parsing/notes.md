# 76 - Time Formatting and Parsing
Formatting and Parsing are crusial for converting time values between human readable formats and machine friendly representations like Unix timestamps.

In Go, The time package provides robust support for these operations, offering a variety of layout patterns to format time and methods to parse time strings into time dot time objects.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? Mon Jan 2 15:04:05 MST 2006 --> Reference time
	// Parsing a time string into a time dot time objects in Go, we need time dot parse function, specifying
	// the layout pattern that matches the format of the input string.

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)

	if err!=nil {
		fmt.Println("Error parsing time: ", err)
		return
	}

	fmt.Println(t)
}
```
```bash
2024-07-04 14:30:18 +0000 UTC
```


------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? Mon Jan 2 15:04:05 MST 2006 --> Reference time
	// Parsing a time string into a time dot time objects in Go, we need time dot parse function, specifying
	// the layout pattern that matches the format of the input string.

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)

	if err!=nil {
		fmt.Println("Error parsing time: ", err)
		return
	}

	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"

	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)
}
```
```bash
2024-07-04 14:30:18 +0000 UTC
2024-07-03 15:18:00 +0000 UTC
```