package delivery

import (
	"fmt"
	"os"
)

func BackToMain() {
	var confirmBackToMain string
	fmt.Print("Back to Menu (y/n)? ")
	fmt.Scan(&confirmBackToMain)

	if confirmBackToMain == "y" || confirmBackToMain == "Y" {
		MainMenu()
	}
	return
}

func ExitApp() {
	os.Exit(0)
}
