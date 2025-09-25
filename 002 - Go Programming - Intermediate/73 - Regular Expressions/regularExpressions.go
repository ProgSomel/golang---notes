package main

import (
	"fmt"
	"regexp"
)

func main(){
	re := regexp.MustCompile(`go`)
	text := "Golang is great"

	fmt.Println("Match:", re.MatchString(text))
}