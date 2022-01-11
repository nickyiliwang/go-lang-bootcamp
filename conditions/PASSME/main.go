package main

import (
	"fmt"
	"os"
)

type User struct {
	u string
	p string
}

const (
	useage  = "Usage: [username] [password]"
	errUser = "Access denied for %q. \n"
	errPass = "Invalid password for %q.\n"
	success = "Access granted for %q. \n"
)

func main() {
	user1 := User{"nick", "12345"}
	user2 := User{"jack", "11111"}

	args := os.Args
	if len(args) != 3 {
		fmt.Println(useage)
		return
	}

	u, p := args[1], args[2]

	if (u == user1.u && p == user1.p) || (u == user2.u && p == user2.p) {
		fmt.Printf(success, u)
	} else if u != user1.u || u == user2.u {
		fmt.Printf(errUser, u)
	} else if p != user1.p || p != user2.p {
		fmt.Printf(errPass, u)
	}

}

// func main() {
// 	name := "nick"
// 	username, password := "", ""

// 	if len(os.Args) > 1 {
// 		username, password = os.Args[1], os.Args[2]
// 	}

// 	if username == "nick" && password == "12345" {
// 		fmt.Printf("Access granted for: %s", name)
// 	} else if username == "" || password == "" {
// 		fmt.Printf("Usage: [username], [password]")
// 	} else if username == "nick" && password != "12345" {
// 		fmt.Printf("Invalid password for: %s", username)
// 	} else if username != "nick" && password == "12345" {
// 		fmt.Printf("Access denied for %s", name)
// 	} else {
// 		fmt.Printf("Access denied for %s", username)
// 	}
// }
