package main

import (
	"fmt"
	"os"
)

func main() {
    // prints out all the inputed values into an []slice 
	fmt.Printf("%#v\n", os.Args)

    fmt.Println("Path:", os.Args[0])
    fmt.Println("1st argument:", os.Args[1])
    fmt.Println("2nd argument:", os.Args[2])
    fmt.Println("3rd argument:", os.Args[3])
    
    fmt.Println("number of items inside os.Args:", len(os.Args))
    
    
}

// go build -o greeter 