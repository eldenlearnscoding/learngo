// Type conversion - converts a value to another type
package main

import (
	"fmt"
)

func main() {
	speed := 100
	force := 2.5

	// First convert speed into float64 so it can perform multiplication with force of type float64.
	// Then convert the result to type int so that it matches the type of "speed = ..."
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
