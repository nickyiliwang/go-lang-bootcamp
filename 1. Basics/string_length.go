package main

import (
	"fmt"
	"unicode/utf8"
)

// len() is not as good because it outputs bites of an string instead of how many chars it contains

func main() {
	name := "Nick Wang"
	fmt.Println(
		utf8.RuneCountInString(name))
}
