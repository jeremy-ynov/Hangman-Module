package mypackages

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func UserChoice() string {
	/*
		Foncution UserInput
		==========================
		Takes the letter written into the console by the user
		Checks if it is a valid letter
		Returns that letter
	*/

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() //Enable ability to instantly send any caracters typed into the console

	validLetters := "abcdefghijklmnopqrstuvwxyz"
	fmt.Print("Choice: ")

	buffer := make([]byte, 1)   // Enable a buffer of 1 in order to take only one letter at a time
	os.Stdin.Read(buffer)       // Reads the user input in the console
	userInput := string(buffer) // Transforms that user input into a string form
	fmt.Println()

	// Checks if that letter is valid by comparing it to the alphabet
	for _, letter := range validLetters {

		// If letter is in the alphabet
		if userInput == string(letter) {
			exec.Command("reset").Run() //Reset console to original parameters
			return userInput            // Return the valid letter
		}
	}

	// If letter is not valid, print an error message
	fmt.Print("\033[H\033[2J")
	fmt.Println("Invalid letter! Please enter a valid letter of the alphabet.")
	time.Sleep(2 * time.Second)
	fmt.Print("\033[H\033[2J")
	return "false"
}
