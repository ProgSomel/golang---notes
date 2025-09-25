package main

import (
    "fmt"       // Printing messages and getting user input
    "math/rand" // Random number generation
    "time"      // To seed random numbers uniquely
)


func main(){

	source := rand.NewSource(time.Now().UnixNano()) //? time.Now() - Gets the current time as a time.Time value
	                                                //? .UnixNano() - Converts that time to nanoseconds since Unix epoch (January 1, 1970). This returns an int64 representing the exact moment down to nanosecond precision.
													//? rand.NewSource(...) - Creates a new random number source (implements the rand.Source interface) using the nanosecond timestamp as the seed value. The seed determines the starting point for the random number sequence.
	/*
		Why use nanoseconds as seed? Since nanoseconds change very rapidly, 
		this ensures each program run gets a different seed, producing different 
		random sequences. Without seeding (or with a fixed seed), you'd get the 
		same "random" numbers every time.

	*/												
	random := rand.New(source) /*
	rand.New(source) - Creates a new *rand.Rand instance using the source 
	you just created. This gives you a random number generator that you can call methods on.

	Why not use global rand functions? While Go has global functions like rand.Intn(), 
	creating your own instance gives you:

    Thread safety (each instance is independent)
    Control over the seed
    Ability to have multiple generators with different seeds
	*/

	//? Generate random number between 1 and 100
	target := random.Intn(100) /*
	random.Intn(100) - Generates a random integer in the range [0, 100) 
	- that's 0 to 99 inclusive. The n in Intn stands for "integer less than n".
	*/

	/*
		Note about the range: Despite your comment saying "between 1 and 100", this actually generates 0-99. To get 1-100, you'd need:
gotarget := random.Intn(100) + 1  // Now gives 1-100
The overall pattern (seed → source → generator → random number) is the recommended way to generate unpredictable random numbers in Go.
	*/

	//? Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int 
	for{
		fmt.Println("Enter your guessing number: ")
		fmt.Scan(&guess)

		//? check if the guess if correct
		if(guess == target){
			fmt.Println("Congratualions! You guessed the correct number!")
			break;
		}else if(guess < target){
			fmt.Println("To low! Try guessing a higher number. ")
		}else{
			fmt.Println("To high! Try guessing a lower value. ")
		}
	}



}