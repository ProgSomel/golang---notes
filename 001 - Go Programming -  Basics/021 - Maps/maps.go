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