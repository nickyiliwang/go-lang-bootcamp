package main

import (
	"fmt"
	"strconv"
)

// nil value means that the value is not initialized yet
// nil itself is a value
// javascript => null

// err := do()
// if err != nil {
// // handle error here or terminate
// }

func main() {
	// some functions are always gonna succeed
	s := strconv.Itoa(42)
	fmt.Println(s)

	text := "123456789"
	i, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(i)
}
