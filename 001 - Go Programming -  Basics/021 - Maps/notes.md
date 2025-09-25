# 021 - Maps - a built in data structure that associate keys with values. They are like dictionaries in other programming languages, and provide an efficient way to look up data by a key.

Map provide an efficient way to store and retrieve key value pairs. Each key must be unique within the map, and the keys are typical of a comparable type, like strings or integer.

Maps are unordered collections of key value pairs, meaning that there is no guaranteed order when iterating over them. 

```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

}
```
```bash
map[]
```

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	fmt.Println(myMap)

}
```
```bash
map[]
map[key1:9]
```

--------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])

}
```
```bash
map[]
map[code:18 key1:9]
9
```

**--------------------------------------------------------------------------------------------------------**

## what if we mention an incorrect key - it will give zero value of that type
```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

}
```
```bash
map[]
map[code:18 key1:9]
9
0
```

**--------------------------------------------------------------------------------------------------------**

```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 180
	fmt.Println(myMap)
}
```
```bash
map[]
map[code:18 key1:9]
9
0
map[code:180 key1:9]
```

**--------------------------------------------------------------------------------------------------------**

## deleting a key value pair
```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 180
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)
}
```
```bash
map[]
map[code:18 key1:9]
9
0
map[code:180 key1:9]
map[code:180]
```

**--------------------------------------------------------------------------------------------------------**

## completely removing all the key value pairs from a map - we can use the clear method
```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["key4"] = 12

	fmt.Println(myMap)

	clear(myMap)

	fmt.Println(myMap)
}
```
```bash
map[]
map[key1:9 key2:10 key3:11 key4:12]
map[]
```

**--------------------------------------------------------------------------------------------------------**

## when we are using this syntax-> myMap["key1"] = 9 -> we are not just receiving a vslue associated with the key but there is another optional value that we are receiving.
```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["key4"] = 12

	fmt.Println(myMap)

	value, unknownValue := myMap["key1"]
	fmt.Println(value)
	fmt.Println(unknownValue)
}
```
```bash
map[]
map[key1:9 key2:10 key3:11 key4:12]
9
true
```
**Here unknownValue is saying that there is a value associated with the mapKey**

-------------------------------------------------------------------------------------------------------

```go
package main

import "fmt"

func main(){
	//! var mapVariable map[keyType]valueType
	//! mapVariable = make(map[keyType]valueType)
	//! using Map literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["key4"] = 12

	fmt.Println(myMap)

	_, unknownValue := myMap["key1"]
	fmt.Println("Is value associated with key1: ", unknownValue)
}
```
```bash
map[]
map[key1:9 key2:10 key3:11 key4:12]
Is value associated with key1:  true
```

**--------------------------------------------------------------------------------------------------------**

## Initializing a Map in a different way
```go
package main

import "fmt"

func main(){
	myMap := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap)
}
```
```bash
map[a:1 b:2]
```

**--------------------------------------------------------------------------------------------------------**

## Equality check between two Maps
```go
package main

import (
	"fmt"
	"maps"
)

func main(){

	myMap := make(map[string]int)

	myMap["key1"] = 1
	myMap["key2"] = 2

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2){
		fmt.Println("myMap and myMap2 are Equal")
	}
}
```
```bash
map[a:1 b:2]
```

-------------------------------------------------------------------------------------------------------

```go
package main

import (
	"fmt"
	"maps"
)

func main(){

	myMap := make(map[string]int)

	myMap["a"] = 1
	myMap["b"] = 2

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2){
		fmt.Println("myMap and myMap2 are Equal")
	}
}
```
```bash
map[a:1 b:2]
myMap and myMap2 are Equal
```

**--------------------------------------------------------------------------------------------------------**

## iterating over map
```go
package main

import (
	"fmt"
	"maps"
)

func main(){

	myMap := make(map[string]int)

	myMap["a"] = 1
	myMap["b"] = 2

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2){
		fmt.Println("myMap and myMap2 are Equal")
	}

	for k, v := range myMap{
		fmt.Println(k, v)
	}
}
```
```bash
map[a:1 b:2]
myMap and myMap2 are Equal
a 1
```

**--------------------------------------------------------------------------------------------------------**

## zero value of a map - if we have a map that has not initialized but only declared then it is  initialized to nil.
```go
package main

import (
	"fmt"
)

func main(){

	var myMap map[string]string
	if myMap == nil{
		fmt.Println("myMap is initialized to a nil value")
	}else{
		fmt.Println("myMap is not initialized to a nil value")
	}
}
```
```bash
myMap is initialized to a nil value
```

**--------------------------------------------------------------------------------------------------------**

## If we try to initialize the map this way (var myMap map[string]string), if we say that hey, we made a map and we are now assigning (myMap["key"] = "value") some keys and values to this map, it won't work, because when we initialize a map like this, it is initialized to nil value.
```go
package main

import (
	"fmt"
)

func main(){
	var myMap map[string]string
	myMap["key"] = "value" //! nil dereference in map updatenilness
	fmt.Println(myMap)
}
```

### To solve this problem, we have to use make function
```go
package main

import (
	"fmt"
)

func main(){
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["key"] = "value"
	fmt.Println(myMap)
}
```
```bash
map[key:value]
```

**--------------------------------------------------------------------------------------------------------**

## Length of a Map - using len() function
```go
package main

import (
	"fmt"
)

func main(){
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["key"] = "value"
	fmt.Println(myMap)

	fmt.Println("Length of myMap is: ", len(myMap))
}
```
```bash
map[key:value]
Length of myMap is:  1
```

**--------------------------------------------------------------------------------------------------------**

## like arrays and maps, we have multidimensional map/nested maps
```go
package main

import (
	"fmt"
)

func main(){
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["key"] = "value"
	myMap2 := make(map[string]map[string]string)
	myMap2["map1"] = myMap
	fmt.Println(myMap2)
}
```
```bash
map[map1:map[key:value]]
```