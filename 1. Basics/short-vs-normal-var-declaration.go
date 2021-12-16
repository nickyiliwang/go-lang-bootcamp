// When to use Normal declaration

// 1. If you don't know the initial value
// func main() {
// 	score := 0    // DON'T, doesn't know initial val
// 	var score int // <= do this instead
// }

// 2. When you need a package scoped variable
// package main
// version := 0 // will not work
// var version string // will work

// func main() {
// code
// }

// 3. When you want to group variables together for greater readability, ie. related variables
// func main()  {
// 	var (
// 		// related
// 		video string

// 		// closely related
// 		duration int
// 		current int
// 	)
// }

// When to use Short Declaration
// *Mostly preferred way of declaring

// 1. You know the initial values
// 2. Keep the code concise
// 3. For re-declaration
package main

import "fmt"

func main() {
	width, height := 100, 50

	// DON'T
	// width = 50 // assigns 50 to width
	// color:= "red" // new variable: color

	// DO
	// we are re-declaring and assigning value and declaring a new value
	width, color := 50, "red"

	fmt.Println(width, height, color)
}
