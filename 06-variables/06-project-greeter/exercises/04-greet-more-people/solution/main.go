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
)

func main() {
	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")
}
