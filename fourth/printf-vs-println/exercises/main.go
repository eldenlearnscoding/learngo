package main

import (
	"fmt"
	"os"
)

func main() {

	// EXERCISE: Print Your Age
	age := 30
	fmt.Printf("I am %d years old.\n", age)

	// EXERCISE: Print Your Name and LastName
	name, lastname := "Elden", "Yee"
	fmt.Printf("My name is %s and my lastname is %s.\n", name, lastname)

	//Bonus
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, name, lastname)

	// EXERCISE: False Claims
	tf := false
	fmt.Printf("These are %t claims.\n", tf)

	// EXERCISE: Print the Temperature
	temp := 28.15
	fmt.Printf("Temperature is %0.2f degrees.\n", temp) // Precision at 2 decimal points

	// EXERCISE: Double Quotes
	fmt.Printf("\"hello world\"\n")

	// EXERCISE: Print the Type
	fmt.Printf("Type of %d is %[1]T.\n", 3)
	fmt.Printf("Type of %0.2f is %[1]T.\n", 3.14)
	fmt.Printf("Type of %s is %[1]T.\n", "hello")
	fmt.Printf("Type of %t is %[1]T.\n", true)

	// EXERCISE: Print Your Fullname
	fmt.Printf("Your name is %s and your lastname is %s.\n", os.Args[1], os.Args[2])

	//Bonus
	name1, lastname1 := os.Args[1], os.Args[2]
	msg1 := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg1, name1, lastname1)
}
