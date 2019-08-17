package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a name")
		return
	}
	name := os.Args[1]

	moods := [...]string{
		"angry", "happy", "sad", "willfull", "joyous", "arfendour",
	}
	rand.Seed(time.Now().UnixNano())
	currentMood := moods[rand.Intn(len(moods))]
	fmt.Printf("%s feels %s", name, currentMood)
}
