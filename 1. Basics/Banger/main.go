package main

import "os"

func main() {
	msg := os.Args[1]
	l := len(msg)

	// https://pkg.go.dev/strings#Repeat
	s := msg + strings.Repeat("!, l")

}
