package main

import "fmt"

type Shape struct {

	Rectangle
}

type Rectangle struct {
	length float64
	width float64
}

//? Method with value receiver
func(r Rectangle) Area() float64 {
	return r.length * r.width
}

func main(){
	
	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())

}