package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Declare
//
//  Declare and then print four variables using
//  the short declaration statement.
//
// EXPECTED OUTPUT
//  i: 314 f: 3.14 s: Hello b: true
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	//
	i, f, s, b := "i", "f", "s", "b"
	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
}
