package main

import "fmt"

// the fallthrough keyword inside of an switch statement invalidated the condition of the case below and just executes the next block of code
// go automatically adds break statements unlike JS
// fallthrough is rarely used
func main() {
	// we are initiating a variable in the switch statement and need the semi-colon
	// so we can add the following true value condition, it's there as a default
	// switch i := 125; true
	switch i := 125; {
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

// i = 125 OUTPUTS: "big" "positive" "number"
// i = 99 OUTPUTS: "positive" "number"
// i = -1 OUTPUTS: "number"
