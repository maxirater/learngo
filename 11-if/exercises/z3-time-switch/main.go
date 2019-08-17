package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	switch t := time.Now().Hour(); {
	case t < 12:
		p("Good Morning")
		p("Woo")
	case t < 18:
		p("Good Afternoon")
	default:
		p("Good Evening")
	}
}
