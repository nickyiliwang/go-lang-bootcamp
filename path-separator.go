// Path Separator
// provides utility functions for working with *url path strings
// ie. path.Split()

// func name (input params, type) (result values)
// func Split (path string) (dir, file string)
package main

import (
	"fmt"
	"path"
)

func Split(pathInput string) (dir, file string) {
	dir, file = path.Split(pathInput)
	return dir, file
}

func main() {
	// Using short declaration replaces var
	dir, file := Split("css/main.css")

	// we can just print one value using underscore
	// _, file = Split("css/main.css")

	fmt.Println("The dir is ", dir)
	fmt.Println("The file is ", file)
}
