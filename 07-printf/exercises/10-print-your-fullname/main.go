package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	fn, ln := "hey", "ho"
	fmt.Printf("Your name is %s and your last name is %s\n", fn, ln)
	a := 3 / 2.
	fmt.Println(a)
}
