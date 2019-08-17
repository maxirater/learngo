package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Correct the lyric
//
//  You have a slice of lyrics data of Beatles' awesome
//  song: Yesterday. Your goal is putting the words into
//  correct positions.
//
//
// STEPS
//
//  1. Prepend "yesterday" to the `lyric` slice.
//
//  2. Put the words to the correct position in the `lyric` slice.
//
//  3. Print the `lyric` slice.
//
//
// EXPECTED OUTPUT
//
//  [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh i believe in yesterday]
//
//
// BONUS
//
//   . Think about when does the append allocates a new backing array.
//   . Then check whether your conclusions are true or not.
//   . You can use the prettyslice package to check the backing array.
//
// ---------------------------------------------------------

// func main() {
// 	// DON'T TOUCH THIS:
// 	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)

// 	// this works correctly, prints: [yesterday all my troubles seemed so far away]
// 	lyric1 := append(lyric[11:12], lyric[:7]...)

// 	// this does not work right, it prints: [yesterday all my troubles seemed so far away all my troubles seemed so far away here to stay] I cant figure out why it repeats "all my troubles seemed so far away"
// 	lyric2 := append(lyric1, lyric[12:]...)

// 	// lyric3 := append(lyric2, lyric[7:13]...)

// 	fmt.Printf("%s\n", lyric2)
// }

// lyric[:7] this overwrites 7 words after "yesterday"
func main() { // 1   2  3     4  5      6    7
	lyric := strings.Fields(`all my troubles seemed so far away oh i believe in yesterday now it looks as though they are here to stay`)
	//                                                                                    ^ this becomes: all my troubles seemed so far away because: follow down the arrow.
	// this works correctly, prints: [yesterday all my troubles seemed so far away]    // |
	fmt.Printf("lyric        %s\n", lyric)        // |
	fmt.Printf("lyric[11:12] %s\n", lyric[11:12]) // |
	fmt.Printf("lyric[:7]    %s\n\n", lyric[:7])  // |
	//                                                                                    |
	//                                                                                    |
	lyric1 := append(lyric[11:12], lyric[:7]...) // THIS OVERWRITES THE BACKING ARRAY OF LYRIC AND LYRIC1
	// lyric[:7] is "all my troubles seemed so far away"
	// lyric[11:12] is "yesterday"
	//
	// SO: after "yesterday", it overwrites the rest with: "all my troubles seemed so far away"

	fmt.Printf("lyric1:      %s\n", lyric1)
	fmt.Printf("lyric:       %s\n\n", lyric[0:cap(lyric)])

	// this does not work right, it prints: [yesterday all my troubles seemed so far away all my troubles seemed so far away here to stay]
	// I cant figure out why it repeats "all my troubles seemed so far away",
	// its supposed to start on index 12 and append: "now it looks as though they are here to stay"
	// lyric2 := append(lyric1, lyric[12:]...)
	// lyric2 := append(lyric1, lyric[12:]...)
	// fmt.Printf("lyric2:       %p\n", lyric2)
	// fmt.Printf("lyric2:       %s\n", lyric2)
}
