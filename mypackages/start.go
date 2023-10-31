package mypackages

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func Start(language int, dico string) (string, string, string) {

	/*
		Function Start
		===================
		Parametters :
		dico, type(string) : name of the .txt document with the words to be randomly choosen in the game\
		-----------------

		Initialize the game by reading the dico
		Places every word in a string list
		Randomly chooses a word from that list using math/rand import
		Set the blankword as a string of "_" the size of the word
		Set nStartingLetters, the number of letters shown at the start of the game
		Randomly choose the letters to be shown
		Fill the blank word with the letters using the reveal function

		----------------
		Returns :
		blankWord, type(string) : the hidden word that will be showed to the player
		randomWord, type(string) : the answer of the game
		err, type(string) : returns "nil" when dictionnary is valid
	*/

	// Opens the .txt document
	dictionnary, err := os.Open(dico)

	if err != nil {
		fmt.Println("ERROR! : Unable to find " + dico)
		return "", "", ""
	}
	// Closes the reader once the file has been read
	defer dictionnary.Close()

	// Create a scanner to analize the dictionnary
	scanner := bufio.NewScanner(dictionnary)

	// Add each word from the dictionnary into a string list
	wordList := []string{}
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	var finalWordList []string

	for i := (language-1)*84 + language - 1; i < (language-1)*84+language+84; i++ {
		finalWordList = append(finalWordList, wordList[i])
	}

	// Randomly choose a word from the list using math/rand module
	randomWord := finalWordList[rand.Intn(len(finalWordList)-1)]

	// Initialize the blanck word as a string of "_" with the size of the word
	blankWord := ""
	for i := 0; i < len(randomWord); i++ {
		blankWord += "_"
	}

	// Sets the number of letters that will be initialy shown at the start
	nStartingLetters := len(randomWord)/2 - 1

	// List to temporarily hold the inital letters shown
	startingLettersList := []string{}

	i := 0
	for i < nStartingLetters {
		// Randomly choose a letter in the word
		randomLetter := string(randomWord[rand.Intn(len(randomWord)-1)])

		// Check if that letter has not been chosen already using the startingLettersList
		isinlist := false
		for _, letter := range startingLettersList {
			if letter == randomLetter {
				isinlist = true
			}
		}

		// If that letter has not been chosen already
		if !isinlist {
			// Add that letter to the startingLetterslist
			startingLettersList = append(startingLettersList, randomLetter)

			// Show that letter in the blank word using the reveal function
			blankWord = Reveal(randomLetter, blankWord, randomWord)
			i += 1
		}
	}
	return blankWord, randomWord, "nil"
}

func Reveal(letter string, hiddenWord string, wordToFind string) string {
	/*
		Function Reveal
		==============
		Parametters :
		letter, type(string) : letter to be shown in the blank word
		hiddenWord, type(string) : the word filled with blank letters
		wordToFind, type(string) : the answer that need to be found
		------------
		Fills the blank word with the letter taken in the corresponding place using the full word as the reference
		------------
		Returns :
		the hiddenWord with newly added letter, type(string)
	*/

	// Create a rune list with the letters and "_" of hiddenWord
	hiddenList := []rune(hiddenWord)

	// Finds where the letter is in the wordToFind
	for i, char := range wordToFind {
		// If the letter is in the wordToFind add it to the hiddenWord
		if string(char) == letter {
			hiddenList[i] = char - 32
		}
	}

	// return the hiddenWord as a type string
	return string(hiddenList)
}
