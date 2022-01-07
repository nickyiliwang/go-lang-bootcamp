package main

import "fmt"

// Printf allows you to format the message before it's printed
// Printf seems to be the better option once I get used to it

func main() {
	ops, ok, fail := 2350, 543, 443
	// %d is declaring int, similarly %s is for strings, and %p for pointers
	fmt.Printf("total, %d success: %d / %d\n", ops, ok, fail)

	// on the other hand Println acts like console.log
	// harder to read
	fmt.Println("total,", ops, "success:", ok, "/", fail)
}
