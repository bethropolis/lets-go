package main

import "fmt"

func getUserChoice(prompt string) int {
	var choice int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&choice)
		if err == nil {
			break // Input is valid, exit the loop
		}
		fmt.Println("Invalid input. Please enter a valid number.")
	}
	
	return choice
}