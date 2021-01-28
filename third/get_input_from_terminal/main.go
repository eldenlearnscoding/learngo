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

	greeter()
}

func greeter() {
	for i := 1; i < len(os.Args); i++ {
		// Set the name to be stored in the [1] position of the Args.
		// [0] is used to store the path of the running program
		name := os.Args[i]
		fmt.Println("Hello", name)
	}

	name, age := "Elden", 32
	fmt.Println("I am", name)
	fmt.Println("I am", age, "years old")
}

// Side note: go build -o, the -o arguments allows me to name the compiled program
