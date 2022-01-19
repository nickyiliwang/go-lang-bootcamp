package main

import (
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`smart|cool|genius`)

// this program guesses the number you input
func main() {

	args := os.Args

	fmt.Println(args[1])
}
