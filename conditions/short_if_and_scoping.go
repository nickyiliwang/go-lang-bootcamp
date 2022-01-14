package main

import (
	"fmt"
	"os"
	"strconv"
)

// Using the semi-colon between statements we can chain conditions
// but the values are blocked within these function chains
// no need for return statement here when there is nothing else the function needs to do

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Only works with numbers.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		converted := n * 2
		fmt.Printf("%s multiplied by 2 is: %d", a[1], converted)
	}

}
