package mypackages

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func TauntError(language int) {

	/*
		The function tauntEror print a taunt message
		---------------
		input : nothing
		sort : printing of a ramdom line from tauntEror.txt
	*/

	// Open the .txt document
	taunt, _ := os.Open("mypackages/Taunt.txt")

	// Closes the reader once the file has been read
	defer taunt.Close()

	// Create a scanner to analize the txt
	scanner := bufio.NewScanner(taunt)

	// Add each word from the txt into a string list
	wordList := []string{}
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	// randomly choose a word from the list using math/rand module and print it
	fmt.Println(wordList[rand.Intn(29)+30*(language-1)])
}
