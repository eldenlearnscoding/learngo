// Get input from the command line
// Use OS (Operating System) package
// Args variable belongs to the OS package.
//		var Args []string
//		It holds the command-line arguments, starting with the program name.
//		The brackets [] before the string means its a slice; so []string means -
//		Args' type is a "slice of strings", so Args can store multiple string values in it.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Number of items inside os.Args:",
		len(os.Args)) // len function returns the length of the items

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Args[%d] = %v\n", i, os.Args[i])
	}

}

// Side note: go build -o, the -o arguments allows me to name the compiled program
