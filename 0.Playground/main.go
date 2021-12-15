package main

import "fmt"

// avoiding deadlock
func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
