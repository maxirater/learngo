// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	p := fmt.Println
	pf := fmt.Printf

	if len(os.Args) != 2 {
		p("Pick a number")
		return
	}
	num, err := strconv.Atoi(os.Args[1])

	if err != nil {
		pf("%q is not a number", os.Args[1])
	} else if num%8 == 0 && num%2 == 0 {
		pf("%d is an even number and its divisible by 8", num)
	} else if num%2 == 0 {
		pf("%d is an even number", num)
	} else {
		pf("%d is an odd number", num)
	}
}
