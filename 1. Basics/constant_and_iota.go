// iota is a built-in constant generator that generates ever increasing numbers

package main

import (
	"fmt"
)

func main() {
	// verbose and unnecessary
	// const (
	// 	monday    = 0
	// 	tuesday   = 1
	// 	wednesday = 2
	// 	thursday  = 3
	// 	friday    = 4
	// 	saturday  = 5
	// 	sunday    = 6
	// )

	// constants repeats the type of previous const
	// iota is a auto counter & number generator
	// since iota is a value, we can increment the sequence by <iota + 1>
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}
