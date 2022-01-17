package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://go.dev/ref/spec#Short_variable_declarations
// short declaration may only appear inside a function

func main() {
	args := os.Args

	// len starts the count at 1
	if len(args) < 2 {
		fmt.Println("enter an number for the table")
		return
	} else {

		max, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Println("enter a valid number for the table")

		}

		fmt.Printf("%5s", "X")
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()

		for i := 0; i <= max; i++ {

			fmt.Printf("%5d", i)
			for j := 0; j <= max; j++ {
				r := i * j
				fmt.Printf("%5d", r)
			}
			fmt.Println()

		}

	}

}
