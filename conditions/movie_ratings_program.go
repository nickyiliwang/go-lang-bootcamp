package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func main() {
	args := os.Args
	//  1. Get the age from the command-line.

	if len(args) != 2 {
		fmt.Println("Please enter your age in number")
	} else if n, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("please enter a valid number.", err)

	} else if math.Signbit(float64(n)) {
		fmt.Println("Please enter a none negative number.")

	} else {
		//  2. Return one of the following messages if the age is:
		//     -> Above 17         : "R-Rated"
		//     -> Between 13 and 17: "PG-13"
		//     -> Below 13         : "PG-Rated"
		if n > 17 {
			fmt.Println("R-Rated")
		} else if n >= 13 && n < 17 {
			fmt.Println("PG-13")
		} else {
			fmt.Println("PG-Rated")
		}

	}

}
