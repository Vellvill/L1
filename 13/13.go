package main

import "fmt"

func main() {
	var x, y = 1, 2
	fmt.Printf("Before x: %d, y: %d\n", x, y)
	x, y = y, x
	fmt.Printf("After x: %d, y: %d\n", x, y)
}
