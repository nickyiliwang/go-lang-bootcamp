package main

import (
	"fmt"
	"time"
)

// Channels should not be closed from the receiving end
// Do it from the sending end
func main() {

	out1 := make(chan string)
	go process("order", out1)

	// for every message we want to keep reading out the channel
	// range that works on slices also works here
	// stops reading once the channel is closed
	for msg := range out1 {
		fmt.Println(msg)
	}
}

func process(item string, out chan string) {
	// defer will only run close(out) once the func block is done executing
	defer close(out)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		// puting data into it, out of the channel
		out <- item
	}
}
