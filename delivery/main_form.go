package delivery

import (
	"fmt"
	"strings"
)

func MainMenu() {
	fmt.Println(strings.Repeat("*", 13))
	fmt.Println("Music Library")
	fmt.Println(strings.Repeat("*", 13))
	fmt.Println("1. Add Song")
	fmt.Println("2. Show Song List")
	fmt.Println("3. Search Song By Artist")
	fmt.Println("4. Quit")
	fmt.Print("Select 1-4: ")
}
