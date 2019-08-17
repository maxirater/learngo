package main

import (
	"fmt"
	"os"
)

func main() {
	un, pw := "Ed", "Straker"
	un2, pw2 := "Brad", "Cantrell"

	if len(os.Args) < 3 {
		fmt.Println("Usage: [username][password]")
	} else {
		username, password := os.Args[1], os.Args[2]

		if username == un || un2 == username {
			if password == pw || password == pw2 {
				fmt.Printf("Access granted for %s\n", username)
			} else {
				fmt.Printf("Invalid password for %s\n", username)
			}
		} else {
			fmt.Printf("Acccess denied for %s\n", username)
		}
	}
}
