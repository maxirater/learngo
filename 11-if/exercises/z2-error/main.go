package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Printf("%q is not a number", input)
		return
	}
	res := num / 3.
	fmt.Printf("%.0f feet is %.2f meters.", num, res)

}
