//Task: Write a guess-the-number program. 
//Make the computer pick random numbers between 1–100 until it guesses your number, 
//which you declare at the top of the program. 
//Display each guess and whether it was too big or too small.

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var myNumber = 73

	for {
		var guess = rand.Intn(100) + 1
		if guess == myNumber {
			fmt.Printf("%v You got it!", guess)
			break
		} else if guess > myNumber {
			fmt.Printf("%v is high \n", guess)
		} else if guess < myNumber {
			fmt.Printf("%v is low \n", guess)
		} else {
			fmt.Println("You're wrong")

		}
	}

}
