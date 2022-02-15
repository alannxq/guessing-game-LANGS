package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var guess string
	rand.Seed(time.Now().UnixNano() * 199) // without *199, it was going up in seconds, not fun to play
	num := rand.Intn(100)
	guessed := false

	for guessed == false {
		fmt.Print("Enter Guess from 1-100: ")
		fmt.Scanln(&guess)

		intGuess, err := strconv.Atoi(guess) // converting input into int "Ascii to Integer"
		if err != nil {
			println(err) // uhhh you gotta use the error thingy because it shout at you if not
		}

		if intGuess == num {
			fmt.Println("Correct.\n")
			guessed = true
		} else if intGuess > num {
			fmt.Println("Too high.\n")
		} else {
			fmt.Println("Too low.\n")
		}
	}
}
