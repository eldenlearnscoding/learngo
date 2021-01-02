package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello!")
	bye()
	hey()
	fmt.Println(runtime.NumCPU() + runtime.NumCPU())
}
