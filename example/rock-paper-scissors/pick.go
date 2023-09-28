package main

import "fmt"

func pick() string {
	var choice string
	fmt.Println("Enter your choice (rock, paper, scissors): ")
	fmt.Scanln(&choice)
	return choice
}
