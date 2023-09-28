package main

import (
	"fmt"
)

const ROCK, PAPER, SCISSORS = "rock", "paper", "scissors"

func main() {

	choice := pick()
	fmt.Printf("\nuser choice: %s\n", choice)

	user := convertToNumber(choice)
	computer := compute()
	fmt.Printf("computer choice: %s\n", convertToString(computer))
	fmt.Printf("winner: %s\n", winner(user, computer))

	fmt.Printf("\nDo you want to play again? (y/n)")
	fmt.Scanln(&choice)
	if choice == "y" {
		main()
	}
}
