package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	excite := strings.Repeat("!", l)

	// https://pkg.go.dev/strings#Repeat
	s := excite + strings.ToUpper(msg) + excite

	fmt.Println(s)
}
