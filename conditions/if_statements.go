package main

import "fmt"

func main() {
	score := 5
	// go does not require parentheses
	if score >= 3 {
		fmt.Printf("Good!")
	}
}
