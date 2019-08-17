// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
	}
	guess, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())
	var randNum int
	if guess > 10 {
		randNum = rand.Intn(guess)

	} else {

		randNum = rand.Intn(11)
	}

	if guess == randNum {
		fmt.Println("You guessed on the first try!!")
		return
	}
	fmt.Println("Wrong guess")
	for i := 0; i < guess-1; i++ {
		if guess == randNum {
			fmt.Println("You guessed right!!")
			return
		} else {
			fmt.Println("Wrong guess")
			randNum = rand.Intn(11)
		}
	}
}
