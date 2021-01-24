// Type conversion - converts a value to another type
package main

import (
	"fmt"
)

func main() {

	// Convert and fix

	// Expected output : 15.5
	// a, b := 10, 5.5
	// fmt.Println(a + b)
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	// Expected output : 10.5
	// a, b := 10, 5.5
	// a = b
	// fmt.Println(a + b)
	c, d := 10, 5.5
	c = int(d)
	fmt.Println(float64(c) + d)

	// Expected output : 5.5
	// fmt.Println(int(5.5))
	fmt.Println(5.5)

	// Expected output : 9.5
	// age := 2
	// fmt.Println(int(7.5) + int(age))
	age := 2
	fmt.Println(7.5 + float64(age))

	// Expected output : 1127
	// HINTS
	//   maximum of int8  can be 127
	//   maximum of int16 can be 32767
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	//fmt.Println(int8(max) + min)
	fmt.Println(int16(max) + int16(min))
}
