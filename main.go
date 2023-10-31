package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jeremy-ynov/Hangman-Module/mypackages"
)

func main() {
	// Catch the dictionary name
	args := os.Args

	// Verify if there is a dictionary in arguments
	if len(args) == 2 && mypackages.IsTxt("../"+args[1]) {
		mypackages.RemovePrint(2)
		// print the Menu
		fmt.Println(" --------\033[1m Menu\033[0m --------- ")
		fmt.Println("| [F]  :\033[44m Francais \033[0m      |")
		fmt.Println("| [U]  :\033[7m English  \033[0m      |")
		fmt.Println("| [D]  :\033[101m Deutsch  \033[0m      |")
		fmt.Println("| [E]  :\033[103m Español  \033[0m      |")
		fmt.Println("| [I]  :\033[102m Italiano \033[0m      |")
		fmt.Println("| [Q]  :\033[3m Exit \033[0m          |")
		fmt.Println(" ----------------------- ")

		// Select language
		choseLanguage := mypackages.UserChoice()
		numberOfRound := 0
		switch choseLanguage {
		case "f":
			Game(1, args[1], numberOfRound, []int{0, 0})
		case "u":
			Game(2, args[1], numberOfRound, []int{0, 0})
		case "d":
			Game(3, args[1], numberOfRound, []int{0, 0})
		case "e":
			Game(4, args[1], numberOfRound, []int{0, 0})
		case "i":
			Game(5, args[1], numberOfRound, []int{0, 0})
		case "q":
			// quit case
			mypackages.RemovePrint(2)
			fmt.Println("Thanks for playing")
			time.Sleep(2 * time.Second)
			mypackages.RemovePrint(1)
		default:
			// If invalid revert back to menu
			mypackages.RemovePrint(2)
			main()
		}
	} else {
		// Error case if invalid text document in input
		mypackages.RemovePrint(2)
		fmt.Println("There is no dictionary in arguments. Or the type of the document given in parameters is not an .txt .")
		return
	}

}

func Game(language int, dico string, numberOfRound int, wictoryRate []int) {
	/*
		Function Game
		=============
		Parametters :
		language type(int) : the language id to print the rules
		dico type (string) : the text document with the word list
		numberOfRound type(int) : tracks the number of round already played
		------------
	*/

	// erase previous 13 lines to clean the terminal
	mypackages.RemovePrint(13)

	// open the .txt document
	rules, _ := os.Open("mypackages/rules.txt")

	// Closes the reader once the file has been read
	defer rules.Close()

	// Create a scanner to analize the document
	scanner := bufio.NewScanner(rules)

	// Add each word from the text document into a string list
	rulesList := []string{}
	for scanner.Scan() {
		rulesList = append(rulesList, scanner.Text())
	}

	// Add all the rules in the selected language only into rulesLanguage list
	var rulesLanguage []string
	for i := (language-1)*21 + language - 1; i < (language-1)*21+language+21; i++ {
		rulesLanguage = append(rulesLanguage, rulesList[i])
	}

	// Print the rules in the selected language
	fmt.Println("\n", rulesLanguage[0])
	if numberOfRound > 1 {
		fmt.Println()
		fmt.Println(rulesLanguage[17], numberOfRound, rulesLanguage[18])
	} else {
		fmt.Println()
	}
	// Setting of the hidden Word and the word To Find for the game
	hiddenWord, wordToFind, err := mypackages.Start(language, dico)

	if err != "nil" {
		fmt.Println()
		return
	}

	// Initialisation of Jose our error count
	jose := 0
	usedLetters := []string{}
	// randomly choose the type of jose which will be displayed during the game
	joseType := rand.Intn(8) * 80
	// Assign gender variable to print correct text when the player win or lose
	joseGender := 1 // jose
	if (joseType/80)%2 != 0 {
		joseGender = 2 // josette
	}
	// The game loop
	for jose < 10 {

		// Print the interface for the progresion of the hangman and also the letters used
		mypackages.PrintInterface(language, hiddenWord, jose, joseType, usedLetters)

		// Catch the user imput
		input := mypackages.UserChoice()
		// The case where input is a used letter
		if input != "false" {
			if mypackages.IsUsed(input, usedLetters) {
				mypackages.RemovePrint(3)
				fmt.Println(rulesLanguage[2])
				fmt.Println()

				// The case where the input is in the word To Find
			} else if mypackages.IsIn(input, wordToFind) {
				mypackages.RemovePrint(3)
				fmt.Println()
				fmt.Println()
				hiddenWord = mypackages.Reveal(input, hiddenWord, wordToFind)

				// The case where the input is not in the word To Find
			} else if !mypackages.IsIn(input, wordToFind) {
				mypackages.RemovePrint(3)
				jose++
				usedLetters = append(usedLetters, input)
				fmt.Println(rulesLanguage[3])
				mypackages.TauntError(language)
			}

			// The case where the player dosent have more try "lose case"
			if jose == 10 {
				mypackages.RemovePrint(3)
				mypackages.PrintInterface(language, hiddenWord, jose, joseType, usedLetters)
				if joseGender == 1 {
					fmt.Println(rulesLanguage[4], "\n"+rulesLanguage[6], mypackages.ToUpper(wordToFind))
				} else {
					fmt.Println(rulesLanguage[5], "\n"+rulesLanguage[6], mypackages.ToUpper(wordToFind))
				}
				wictoryRate[1] += 1
				fmt.Println(rulesLanguage[20], (float64(wictoryRate[0])/(float64(wictoryRate[0])+float64(wictoryRate[1])))*100, "%")
			}

			// The case where the hidden word if find "victory case"
			if mypackages.ToUpper(hiddenWord) == mypackages.ToUpper(wordToFind) {
				mypackages.RemovePrint(3)
				if joseGender == 1 {
					fmt.Println(rulesLanguage[7], "\n"+rulesLanguage[8], hiddenWord, ",", rulesLanguage[9])
				} else {
					fmt.Println(rulesLanguage[7], "\n"+rulesLanguage[8], hiddenWord, ",", rulesLanguage[10])
				}
				wictoryRate[0] += 1
				fmt.Println(rulesLanguage[20], (float64(wictoryRate[0])/(float64(wictoryRate[0])+float64(wictoryRate[1])))*100, "%")

				// Leave the loop because without it the program loops indefinitely
				break
			}
			if jose == 9 {
				fmt.Println("\n", 10-jose, rulesLanguage[12])
			} else if jose < 9 {
				fmt.Println("\n", 10-jose, rulesLanguage[11])
			}
		}
	}
	fmt.Println(rulesLanguage[13])
	fmt.Println(rulesLanguage[14])

	// Catch the user answer
	input := mypackages.UserChoice()

	// Case where the user starts another party
	if input == rulesLanguage[15] || input == rulesLanguage[16] {
		numberOfRound++
		Game(language, dico, numberOfRound, wictoryRate)

		// Case where the user leave the game
	} else {
		mypackages.RemovePrint(3)
		numberOfRound++
		fmt.Println(rulesLanguage[17], numberOfRound, rulesLanguage[18])
		fmt.Println(rulesLanguage[19])
		time.Sleep(2 * time.Second)
		main()
	}
}
