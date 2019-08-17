package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter happy or sad")
		return
	}

	name := os.Args[1]
	input := os.Args[2]
	var inputMood int

	if input == "happy" {
		inputMood = 1
	} else if input == "sad" {
		inputMood = 0
	} else {
		fmt.Println("Wrong input, enter happy or sad")
		return
	}

	moods := [...][3]string{
		{"angry", "happy", "arfendour"},
		{"sad", "willfull", "joyous"},
	}

	rand.Seed(time.Now().UnixNano())
	currentMood := moods[inputMood][rand.Intn(len(moods[0]))]

	fmt.Printf("%s feels %s", name, currentMood)
}
