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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
	}
	guess, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(11)

	const numGuesses = 6
	if guess == randNum {
		fmt.Println("You guessed on the first try!!")
		return
	}
	fmt.Println("Wrong guess")
	for i := 0; i < numGuesses-1; i++ {
		if guess == randNum {
			fmt.Println("You guessed right!!")
			return
		} else {
			fmt.Println("Wrong guess")
			randNum = rand.Intn(11)
		}
	}
}
