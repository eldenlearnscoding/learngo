package main

import (
	"fmt"
	"time"
)

/* The short declaration below will produce error. It cannot be used at the package level.
   This is because all declarations on package scope start with keyword,
   e.g. package, import, func & etc.

   name := "Elden"
*/

var name = "Elden" // This is OK

func main() {
	var firstname, lastname, prevName string
	firstname = "Elden"
	lastname = "Yee"
	prevName = firstname // Saves the firstname "Elden" before the redeclaration to "Karen"
	firstname = "Karen"

	decimal1 := 1.12 // := is short declaration, Go will figure out the type.

	fmt.Printf("%q, %q, %q\n", firstname, lastname, prevName)
	fmt.Println(decimal1)
	fmt.Println(name)

	/* -------------------------------------------- New line ------------------------------------*/

	// Multiple assignments
	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()

	fmt.Println(speed, now)

	// Swapping
	speed, prevSpeed := 100, 50
	fmt.Println("Before swapping:", speed, prevSpeed)

	speed, prevSpeed = prevSpeed, speed // Swap the value between speed and prevSpeed

	fmt.Println("After swapping:", speed, prevSpeed)
}
