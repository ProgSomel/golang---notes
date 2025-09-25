package main

import "fmt"


func main(){

	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
    // fmt.Println(sequence2())

	substractor := func() func(int) int{
		countDown := 99
		return func(x int) int{
			countDown -= x
			return countDown
		}
	}()

	//? using the closure substracter
	fmt.Println(substractor(1))
	fmt.Println(substractor(2))
	fmt.Println(substractor(3))
	fmt.Println(substractor(4))
	fmt.Println(substractor(5))
}

func adder() func() int{
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int{
		i++
		fmt.Println("added 1 to i")
		return i
	}
}