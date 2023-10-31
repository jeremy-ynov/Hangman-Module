package mypackages

import (
	"bufio"
	"fmt"
	"os"
)

func PrintInterface(language int, hiddenWord string, jose int, joseType int, usedLetters []string) {
	// Opens the file hangman.txt, initialization of the variable hangman
	hangman, _ := os.Open("mypackages/hangman.txt")
	defer hangman.Close()
	// Reads hangman.txt
	scanner := bufio.NewScanner(hangman)
	var joseList []string
	// Makes a list with what was read in hangman
	for scanner.Scan() {
		joseList = append(joseList, scanner.Text())
	}
	// Prints the advencement of jose, the hangman
	if jose == 1 {
		for i := (jose-1)*7 + joseType; i < jose*7+joseType; i++ {
			fmt.Println(joseList[i])
		}
	} else if jose > 1 {
		for i := (jose-1)*7 + joseType + jose - 1; i < (jose-1)*7+jose+7+joseType; i++ {
			fmt.Println(joseList[i])
		}
	} else {
		// print blank line where jose is supposed to go when starting the game
		for i := 0; i < 7; i++ {
			fmt.Println()
		}

	}
	// Prints the used letters
	first := true
	var word string
	for _, i := range usedLetters {
		if first {
			first = false
			word += ToUpper(i)
		} else {
			word += ", "
			word += ToUpper(i)
		}
	}
	fmt.Print("\n")
	switch language {
	case 1:
		fmt.Println("Vous avez deja utilisé les lettres :", word)
	case 2:
		fmt.Println("You used the letters :", word)
	case 3:
		fmt.Println("Sie haben die Buchstaben bereits verwendet :", word)
	case 4:
		fmt.Println("Ya has usado las letras :", word)
	case 5:
		fmt.Println("Hai già usato le lettere :", word)
	}
	// Prints the hidden word
	fmt.Println(hiddenWord)
}

func ToUpper(s string) string {
	/*
		Function ToUpper
		===============
		Parametters :
		s type(string) : a word or letters to be transformed to uppercase
		---------------
		If character is a lowercase letter
		Transforms that character to it's uppercase equivalent
		--------------
		type(string) : the word or letters in uppercase form
	*/
	list := []rune(s)
	for i := 0; i < len(list); i++ {
		if list[i] >= 97 && list[i] <= 122 {
			list[i] -= 32
		}
	}
	return string(list)
}
