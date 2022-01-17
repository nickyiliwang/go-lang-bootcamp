package main

import (
	"fmt"
	"strings"
)

func main() {
	// without string.Fields OUTPUTS => "#4: 'n'"
	words := strings.Fields("nick wang is my name, 27 is my age")
	for i, v := range words {
		// OUTPUT [number number]
		// only outputs the bitesize without formating
		fmt.Println(i, v)

		fmt.Printf("#%-2d:  %q\n", i+1, v)

	}
}
