package mypackages

func IsUsed(letter string, usedLetters []string) bool {
	/* Checks if a letter is already used
	-----------------
	input: a string and a dictionnary, the string is the choice of the player and the dictionnary is the list of all the letter that was already used by the player
	output : bool, false if the letter is used for the first time, true if the letter is already used
	*/
	used := false
	for _, i := range usedLetters {
		if letter == i {
			used = true
		}
	}
	return used
}
