// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Do Some Calculations
//
//  1. Print the sum of 50 and 25
//  2. Print the difference of 50 and 15.5
//  3. Print the product of 50 and 0.5
//  4. Print the quotient of 50 and 0.5
//  5. Print the remainder of 25 and 3
//  6. Print the negation of `5 + 2`
//
// EXPECTED OUTPUT
//  75
//  34.5
//  25
//  100
//  1
//  -7
// ---------------------------------------------------------

func main() {
	// cel := os.Args[1]
	cel := "20"
	celF, _ := strconv.ParseFloat(cel, 64)
	far := celF*1.8 + 32
	fmt.Printf("fahrenheit is %.2f\n", far)

	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * .5)
	fmt.Println(50 / .5)
	fmt.Println(25 & 3)
	fmt.Println(-(5 + 2))
}
