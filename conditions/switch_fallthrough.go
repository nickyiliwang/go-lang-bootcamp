package main

import "fmt"

// the fallthrough keyword inside of an switch statement invalidated the condition of the case below and just executes the next block of code
func main() {
	i := 125
	switch {
	case i > 100:
		fmt.Println("big")
		fallthrough
	case i > 0:
		fmt.Println("positive")
		fallthrough
	default:
		fmt.Println("number")

	}
}
