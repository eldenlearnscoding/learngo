package main

import "fmt"

func main() {

	// Multiple short declarations
	i, f, s, b := 314, 3.14, "Hello", true
	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)

	// Short with expression
	sum := (27 + 3.5)
	fmt.Println(sum)

	// Short discard
	a, b := true, true
	_ = b
	fmt.Println(a)

	// Redeclare
	age, yourAge := 30, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)
}
