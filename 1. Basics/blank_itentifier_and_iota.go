package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota)
		// to achieve EST -5, and PST -7
		// we are adding an blank identifier:"_"
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}
