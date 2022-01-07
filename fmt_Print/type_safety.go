package main

import "fmt"

func main() {
	var (
		planet   = "venus" // string  => %s
		distance = 261     // decimal => %d
		orbital  = 224.701 // float	  => %f
		hasLife  = false   // true/f  => %t
	)

	// using proper types will prevent errors

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)

	// adjusting precision
	// %.0f, %.1f, %.2f
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)

}
