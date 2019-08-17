package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into rows by using the newline character ("\n")
//
//     2. For each row, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the rows from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	strSlc := strings.Split(data, "\n")

	dataSlc := make([][]string, len(strings.Split(strSlc[0], ",")))

	for i, val := range strSlc {
		tempSlc := strings.Split(val, ",")
		dataSlc[i] = make([]string, len(tempSlc))
		dataSlc[i] = append(dataSlc[i], tempSlc...)
	}

	locations := make([]string, len(dataSlc[0]))
	sizes := make([]string, len(dataSlc[1]))
	beds := make([]string, len(dataSlc[2]))
	baths := make([]string, len(dataSlc[3]))
	prices := make([]string, len(dataSlc[4]))

	locations = append(locations, dataSlc[0]...)
	sizes = append(sizes, dataSlc[1]...)
	beds = append(beds, dataSlc[2]...)
	baths = append(baths, dataSlc[3]...)
	prices = append(prices, dataSlc[4]...)

	for i := 0; i < len(locations); i++ {
		fmt.Printf("%v\n", locations)
		fmt.Printf("%v\n", sizes)
		fmt.Printf("%v\n", beds)
		fmt.Printf("%v\n", baths)
		fmt.Printf("%v\n", prices)
	}
}

// func main() {
// 	const (
// 		header = "Location,Size,Beds,Baths,Price"
// 		data   = `New York,100,2,1,100000
// New York,150,3,2,200000
// Paris,200,4,3,400000
// Istanbul,500,10,5,1000000`

// 		separator = ","
// 	)
// 	var (
// 		locations []string
// 		// sizes     []int
// 		// beds      []int
// 		// baths     []int
// 		// prices    []int
// 	)

// 	strSlc := strings.Split(data, "\n")

// 	// var dataSlc [][]string

// 	for i, val := range strSlc {
// 		tempSlc := strings.Split(val, ",")
// 		dataSlc[i] = append(dataSlc[i], tempSlc...)
// 	}

// 	for _, val := range dataSlc {
// 		locations = append(locations, val[0])
// 		// append(sizes, val[1])
// 		// append(beds, val[2])
// 		// append(baths, val[3])
// 		// append(prices, val[4])
// 	}
// 	fmt.Printf("%v", locations)
// 	s.Show("hey", locations)
// }
