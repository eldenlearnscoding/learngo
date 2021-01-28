package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	count := len(os.Args) - 1
	fmt.Printf("There are %d people!\n", count)
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Hello great", os.Args[i], "!")
	}
	fmt.Println("Nice to meet you all.")
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	// Search on the web how to solve that.
}
