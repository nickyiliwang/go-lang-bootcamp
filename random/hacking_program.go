package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var re = regexp.MustCompile(`smart|cool|genius`)

// this program guesses the number you input
func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Enter a number to begin the hacking")
		return
	} else if sen := strings.Join(os.Args[1:], " "); re.MatchString(sen) {
		// os.Args[1:] holds the arguments to the program.
		// joins the words into a sentence
		// regex to check contained words

		fmt.Printf("Thanks human, maybe I'll kill you last! :)")
		return
	}

	if guess, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a number", guess)
	} else if guess < 0 {
		fmt.Printf("No negative numbers like: %d", guess)
	} else {
		t := time.Now()
		s := t.UnixNano()
		rand.Seed(s)

		for n := 0; n != guess; {
			n = rand.Intn(guess + 1)
			if n != guess {
				fmt.Printf("%d is not your number ey \n", n)

			} else {
				fmt.Printf("%d looks like I got it !", n)
			}
		}
	}

}
