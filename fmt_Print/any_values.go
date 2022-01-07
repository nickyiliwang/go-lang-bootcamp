package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// %v is the swiss army knife bc it prints any value
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// %[i]v is selecting Argument index
	fmt.Printf("%v is %v away. In %[2]v kms! %[1]v is very Far !\n", planet, distance)
}
