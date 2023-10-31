package mypackages

import "fmt"

func RemovePrint(nb int) {
	/*
		the function remove the n previosly print
		-----------------------------------------
		input : an int representing the n previosly print to remove
		return : nothing
	*/
	if nb <= 0 {
		return
	}
	for i := 0; i < nb; i++ {
		fmt.Print("\033[H\033[2J")
	}
}
