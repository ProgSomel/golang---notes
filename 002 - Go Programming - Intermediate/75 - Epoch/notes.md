# 75 - Epoch
Epoch refers to a specific point in time that serves as a reference for timestamps and calculations.

It is often used in computing and programming to represent time as a single number or count of seconds or miliseconds since a defined starting point.

The Epoch time is usuallay defined as the Unix epoch, which is zero hours, zero minutes, zero seconds UTC on January 1st, 1970.

In many programming languages, including Go, time is often represented as the number of seconds or miliseconds elapsed since the Unix epoch.

So this representation is convenient for storing, comparing and manipulating time related data.

Epoch time unit are seconds. Unix time in seconds, which is a unix timestamp.

Milliseconds are used for more precise calculations and to capture smaller time intervals and Epoch.

Epoch time values are Positive values or Negative values.

Positive values represent times after the Unix epoch, which is as mentioned here Jan, 1970, and
Negative values represent times before the Unix epoch, which is the midnight of Jan, 1, 1970.

Epoch time is universal across platforms ans programming languages, facilitating interoperability.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? 00:00:00 UTC on Jan 1, 1970
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time: ", unixTime)
}
```
```bash
Current Unix Time:  1757303026
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? 00:00:00 UTC on Jan 1, 1970
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time: ", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
}
```
```bash
Current Unix Time:  1757303104
2025-09-08 09:45:04 +0600 +06
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? 00:00:00 UTC on Jan 1, 1970
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time: ", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02"))
}
```
```bash
Current Unix Time:  1757303282
2025-09-08 09:48:02 +0600 +06
2025-09-08
```