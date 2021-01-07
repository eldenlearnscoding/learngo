package main

import "fmt"

/* The short declaration below will produce error. It cannot be used at the package level.
   This is because all declarations on package scope start with keyword,
   e.g. package, import, func & etc.

   name := "Elden"
*/

var name = "Elden" // This is OK

func main() {
	var firstname, lastname string
	firstname = "Elden"
	lastname = "Yee"

	decimal1 := 1.12 // := is short declaration, Go will figure out the type.

	fmt.Printf("%q, %q\n", firstname, lastname)
	fmt.Println(decimal1)
	fmt.Println(name)
}
