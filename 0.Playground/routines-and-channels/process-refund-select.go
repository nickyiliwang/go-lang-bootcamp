package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			out1 <- "order processed"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "refund processed"
		}
	}()

	for {
		// select is kind of like an switch statement
		// select helps with checking to see if one process is done before the next
		// in this case order is processing faster than refunds
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
	}
}
