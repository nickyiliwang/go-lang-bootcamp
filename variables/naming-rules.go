package main

import "fmt"

// Rules
// 1. Variable names with capital letters are exported
// 2. Variable name can start with underscore, ie. var _speed int
// 3. Can't: start with number (3speed), start with punctuations (!speed, spe!ed)
// *Punctuations have special meanings in go

func main() {
	// variabel declaration, name, static type
	// go is strongly typed
	var speed int

	// speed will output a default value of 0
	fmt.Println(speed)
}
