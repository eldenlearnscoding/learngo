package main

import (
	"fmt"
	"path"
)

func main() {

	color := "green"
	color = "dark " + color

	n := 0.
	n = 3.14 * 2

	fmt.Println(color)
	fmt.Println(n)

	/*-------------------------------------*/
	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Rectangle's Perimeter\n")
	fmt.Printf("--------------------------\n")

	width, height := 5, 6
	perimeter := 2 * (width + height)

	fmt.Println(perimeter)

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Multi Assign\n")
	fmt.Printf("--------------------------\n")
	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	// ADD YOUR CODE BELOW
	lang, version = "go", 2

	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Multi Assign #2\n")
	fmt.Printf("--------------------------\n")
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Multi Short Func\n")
	fmt.Printf("--------------------------\n")

	// ADD YOUR DECLARATIONS HERE
	_, b := multi()

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Swapper\n")
	fmt.Printf("--------------------------\n")

	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"
	color, color2 = "orange", "green"

	fmt.Println(color, color2)

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Swapper #2\n")
	fmt.Printf("--------------------------\n")

	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"
	prevColor, red := red, blue
	blue = prevColor

	fmt.Println(red, blue)

	fmt.Printf("\n--------------------------\n")
	fmt.Printf("Discard The File\n")
	fmt.Printf("--------------------------\n")

	// UNCOMMENT THE CODE BELOW:

	dir, _ := path.Split("secret/file.txt")
	fmt.Println(dir)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
