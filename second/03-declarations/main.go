package main

import "fmt"

/* The short declaration below will produce error. It cannot be used at the package level.
   This is because all declarations on package scope start with keyword,
   e.g. package, import, func & etc.

   name := "Elden"
*/

var name = "Elden" // This is OK

func main() {
	/* This is multiple short declaration. The number of variables and the initial values
	   should always match, otherwise there will be error
	*/
	firstname, lastname := "Elden", "Yee"
	myAge, temperature, isAwake := 31, 36.4, true

	decimal1 := 1.12 // := is short declaration, Go will figure out the type.

	fmt.Println(firstname, lastname, myAge, temperature, isAwake)

	firstname, birth := "Karen", 28 // firstname short redeclared to Karen
	fmt.Println(firstname, birth)
	fmt.Println(decimal1)
	fmt.Println(name)
}
