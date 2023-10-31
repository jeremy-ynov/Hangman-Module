package mypackages

func IsIn(letter string, word string) bool {
	/* Checks if a letter is in a word
	-------------
	input: two string, one is the letter chosed by the player the other is the word that the player want to find
	output : bool, true if the letter is in the word, false if it's not in
	*/
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}
