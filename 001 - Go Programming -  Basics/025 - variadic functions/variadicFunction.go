package main

import "fmt"

func sum(sequence int, nums ...int) (int, int){
	total := 0
	for _, v := range nums{
		total += v
		}
	return sequence, total
}

func main(){
	numbers := []int{1, 2, 3, 4, 5}
	sequence, total := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence, "Total: ", total)
}