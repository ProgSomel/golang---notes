package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		//? show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the Dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your Choice(1 or 2)")
		var choice int
		_, err := fmt.Scan(&choice)
		if err!=nil || (choice !=1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue;
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break;
		}

		die1 := rand.Intn(6)+1
		die2 := rand.Intn(6)+1

		//? show the results
		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total is: ", die1 + die2)

		//? Ask the user if user want to roll again
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Print("Invalid input, assuming no. ")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break;
		}
	}
}