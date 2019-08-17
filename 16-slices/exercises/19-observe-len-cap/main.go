package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	pf := fmt.Printf
	p := fmt.Println
	// --- #1 ---
	// 1. create a new slice named: games
	//
	// var gamesEmpty []string
	// 2. print the length and capacity of the games slice
	//
	var gamesEmpty []string
	p("empty game slice:")
	pf("len: %d", len(gamesEmpty))
	pf("cap: %d\n", cap(gamesEmpty))
	p()

	// 3. comment out the games slice
	//    then declare it as an empty slice
	//
	var games = []string{}
	// 4. print the length and capacity of the games slice
	//
	p("declared game slice:")
	pf("len: %d\n", len(gamesEmpty))
	pf("cap: %d\n", cap(gamesEmpty))
	p()
	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	//
	games = append(games, "pacman", "mario", "tetris", "doom")
	// 6. print the length and capacity of the games slice
	//
	p("game slice with elements appended:")
	pf("len: %d\n", len(games))
	pf("cap: %d\n", cap(games))
	// 7. comment out everything
	//
	// 8. declare it again using a slice literal
	//    (use the same elements from step 3)
	gamesLiteral := []string{"pacman", "mario", "tetris", "doom"}
	pf("literal len: %d\n", len(gamesLiteral))
	pf("literal cap:%d\n", cap(gamesLiteral))
	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	for i := 0; i <= len(games); i++ {
		s := games[:i] // this is how code is in solution, why assign to s?
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(s), cap(s))
	}
	p()
	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	//
	zero := games[:0]
	// 2. print the games and the new slice's len and cap
	//
	p("games slice:")
	pf("len: %d\n", len(games))
	pf("cap:%d\n", cap(games))
	p("zero slice (slice of game slice that ends in zero):")
	pf("len: %d\n", len(zero))
	pf("cap:%d\n", cap(zero))
	p()
	// 3. append a new element to the new slice
	//
	zero = append(zero, "hey")
	// 4. print the new slice's lens and caps
	//
	p("appended one element to zero slice:")
	pf("len: %d\n", len(zero))
	pf("cap:%d\n", cap(zero))
	p()
	// 5. repeat the last two steps 5 times (use a loop)
	//

	// 6. notice the growth of the capacity after the 5th append
	//
	p("adding 5 more elements")
	for i := 0; i <= len(games); i++ { // INANC HERE <<<<<<<<<<<<<<<<<
		zero = append(zero, "ho") // in your solution code you add to a new slice named s, why?
		fmt.Printf("zero slice: len: %d cap: %d\n", len(zero), cap(zero))

	}
	// Use this slice's elements to append to the new slice:
	fmt.Println()
	aSlice := []string{"ultima", "dagger", "pong", "coldspot", "zetra"}

	// zero := ...
	zero = append(zero, aSlice...)
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	fmt.Println()

	for i := range zero {
		s := zero[:i]
		fmt.Printf("zero's      len: %d cap: %d\n", len(s), cap(s))

	}
	// observe that, the range loop only loops for the length, not the cap.
	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	fmt.Println()

	zero = zero[:cap(zero)]
	for i := range zero {
		s := zero[:i]
		// fmt.Printf("zero[:%d]'s  len: %d cap: %d - \n", i, len(s), cap(s))
		fmt.Printf("%q\n", s)
	}

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	//
	// 2. change the same element of the games slice
	//
	// 3. print the games and the zero slices
	//
	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	zero[0] = "zowie"
	games[0] = "gumba"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
}
