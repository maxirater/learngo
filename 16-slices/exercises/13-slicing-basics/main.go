package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	pf := fmt.Printf
	p := fmt.Println
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	strSlc := strings.Split(data, " ")
	var (
		nums       []int
		evens      []int
		odds       []int
		middle     []int
		first2     []int
		last2      []int
		evensLast2 []int
		oddsLast2  []int
	)
	for _, val := range strSlc {
		n, _ := strconv.Atoi(val)
		nums = append(nums, n)
	}
	pf("%s\t: %v", "nums", nums)
	p()

	//------------------------------
	for _, val := range nums {
		if val%2 == 0 {
			evens = append(evens, val)
		}
	}
	pf("%s\t: %v", "evens", evens)
	p()

	for _, val := range nums {
		if val%2 != 0 {
			odds = append(odds, val)
		}
	}
	pf("%s\t: %v", "odds", odds)
	p()
	middle = nums[2:4]
	first2 = nums[:2]
	last2 = nums[len(nums)-3:]
	evensLast2 = evens[:2]
	oddsLast2 = odds[len(nums)-3:]
	pf("%s\t: %v", "middle", middle)
	p()
	pf("%s\t: %v", "first2", first2)
	p()
	pf("%s\t: %v", "last2", last2)
	p()
	pf("%s\t: %v", "evensLast2", evensLast2)
	p()
	pf("%s\t: %v", "oddsLast2", oddsLast2)
	p()
}
