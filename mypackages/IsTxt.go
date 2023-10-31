package mypackages

func IsTxt(fileName string) bool {
	/*
		The function IsTxt return true if the name of file is ended by .txt
		-------------
		input : a string representing the name of a file
		output : bool, true if the end of the name is .txt and false in the other case
	*/
	if len(fileName) >= 4 {
		return fileName[len(fileName)-4:] == ".txt"
	} else {
		return false
	}
}
