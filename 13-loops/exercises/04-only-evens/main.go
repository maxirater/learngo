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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please enter two numbers, second higher than first")
	}
	first, _ := strconv.Atoi(os.Args[1])
	second, _ := strconv.Atoi(os.Args[2])
	res := 0

	for {
		first++
		if first%2 != 0 {
			continue
		}
		res += first

		if first < second {
			fmt.Printf("%d + ", first)
			continue
		}
		fmt.Println(first, "=", res)
		break
	}
}
