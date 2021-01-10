package main

import (
	"fmt"
	"path" // Package path implements utility routines for manipulating slash-separated paths.
)

func main() {
	dir, file := path.Split("css/main.css")

	fmt.Printf("dir	: %q\nfile	: %q\n", dir, file)
}
