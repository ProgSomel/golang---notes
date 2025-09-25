# 74 - Time
Time Handling in Go, is essential for dealing with dates, times, durations, and time zones.

It provides functionalities to create, manipulate, format and compare times.

In Go, time, values are represented by the time.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? Current Local Time
	fmt.Println(time.Now())
}
```
```bash
2025-09-07 15:57:22.383306 +0600 +06 m=+0.000144126
```

-------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? Current Local Time
	fmt.Println(time.Now())

	//? Specific time
	specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)
}
```
```bash
2025-09-07 16:00:30.128714 +0600 +06 m=+0.000143876
2025-09-07 12:00:00 +0000 UTC
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//? Current Local Time
	fmt.Println(time.Now())

	//? Specific time
	specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	//? Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
}
```
```bash
2025-09-07 16:08:24.084304 +0600 +06 m=+0.000159834
2025-09-07 12:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
```

--------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// //? Current Local Time
	// fmt.Println(time.Now())

	// //? Specific time
	// specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	//? Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
}
```
```bash
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// //? Current Local Time
	// fmt.Println(time.Now())

	// //? Specific time
	// specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	//? Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "20-5-1") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
}
```
```bash
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// //? Current Local Time
	// fmt.Println(time.Now())

	// //? Specific time
	// specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	//? Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "20-5-1") //? Mon Jan 2 15:04:05 MST 2006
	parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") //? Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)
}
```
```bash
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
2020-05-01 18:03:00 +0000 UTC
```

----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// //? Current Local Time
	// fmt.Println(time.Now())

	// //? Specific time
	// specificTime := time.Date(2025, time.September, 07, 12, 0, 0, 0, time.UTC)
	// fmt.Println(specificTime)

	//? Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //? Reference value = Mon Jan 2 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01") //? Reference value = Mon Jan 2 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2", "20-5-1") //? Reference value = Mon Jan 2 15:04:05 MST 2006
	parsedTime3, _ := time.Parse("06-1-2 14-04", "20-5-1 18-03") //? Reference value = Mon Jan 2 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)
	fmt.Println(parsedTime3)
}
```
```bash
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
2020-05-01 00:00:00 +0000 UTC
0001-01-01 00:00:00 +0000 UTC
```

**----------------------------------------------------------------------------------------------------------------------**

## Formatting Time value
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))
}
```
```bash
Formatted time:  25-09-07 19-16
```

**----------------------------------------------------------------------------------------------------------------------**

## manipulating time
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))
	oneDayLater := t.Add(time.Hour*24)
	fmt.Println(oneDayLater)
}
```
```bash
Formatted time:  25-09-07 22-16
2025-09-08 16:22:49.242031 +0600 +06 m=+86400.000134709
```

------------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))
	oneDayLater := t.Add(time.Hour*24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())
}
```
```bash
Formatted time:  25-09-07 24-16
2025-09-08 16:24:11.093374 +0600 +06 m=+86400.000204126
Monday
```

**----------------------------------------------------------------------------------------------------------------------**

## truncate and round time
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))
	oneDayLater := t.Add(time.Hour*24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time", t.Round(time.Hour))
}
```
```bash
Formatted time:  25-09-07 27-16
2025-09-08 16:27:36.678162 +0600 +06 m=+86400.000160126
Monday
Rounded Time 2025-09-07 16:00:00 +0600 +06
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Formatted time: ", t.Format("06-01-02 04-15"))
	oneDayLater := t.Add(time.Hour*24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time", t.Round(time.Hour))
	
	loc, _ := time.LoadLocation("Asia/kolkata")
	t = time.Date(2025, time.September, 8, 14, 16, 40, 00, time.UTC)
	
	//? convert this to specific time zone
	tLocal := t.In(loc)

	//? perform rounding
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)

	fmt.Println("Original Time (UTC): ", t)
	fmt.Println("Original Time (Local): ", tLocal)
	fmt.Println("Rounded Time (UTC): ", roundedTime)
	fmt.Println("Rounded Time (Local): ", roundedTimeLocal)
}
```
```bash
Formatted time:  25-09-07 37-16
2025-09-08 16:37:26.927425 +0600 +06 m=+86400.000146126
Monday
Rounded Time 2025-09-07 17:00:00 +0600 +06
Original Time (UTC):  2025-09-08 14:16:40 +0000 UTC
Original Time (Local):  2025-09-08 19:46:40 +0530 IST
Rounded Time (UTC):  2025-09-08 14:00:00 +0000 UTC
Rounded Time (Local):  2025-09-08 19:30:00 +0530 IST
```

-----------------------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println("Truncated Time: ", t.Truncate(time.Hour))
}
```
```bash
2025-09-07 16:43:17.904341 +0600 +06 m=+0.000165209
Truncated Time:  2025-09-07 16:00:00 +0600 +06
```

**-----------------------------------------------------------------------------------------------------------------------**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	loc, _ := time.LoadLocation("America/New_york")
	fmt.Println("Newyork time: ", t.In(loc))
}
```
```bash
2025-09-07 16:47:56.473196 +0600 +06 m=+0.000185084
Newyork time:  2025-09-07 06:47:56.473196 -0400 EDT
```