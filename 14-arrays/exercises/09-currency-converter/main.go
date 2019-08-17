package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		USD = iota
		GBP
		JYN
		CHY
		GDM
	)
	curr := [...]float64{
		USD: 3.4,
		GBP: 8.2,
		JYN: 5.3,
		CHY: 4.5,
		GDM: 3.0,
	}
	if len(os.Args[1:]) != 1 {
		fmt.Println("please enter a dollar amount")
	}
	input, _ := strconv.ParseFloat(os.Args[1], 64)

	for _, x := range curr {
		fmt.Printf("%f USD is %f XXX\n", input, x)
	}
}
