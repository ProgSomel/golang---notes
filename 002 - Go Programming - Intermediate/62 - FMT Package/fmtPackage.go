package main

import "fmt"


func main(){
	
	err := checkAge(15)
	if err != nil{
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error{
	if age < 18{
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}