package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		fmt.Println(usage)
		return
	}

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Printf("Error: %q is not a number", arg)
		fmt.Println(usage)
		return
	}

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
