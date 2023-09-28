package main

import "fmt"

const (
	todoFile = "todo.txt"
)

func main() {
	for {
		fmt.Println("\nWhat do you want to do? ")
		fmt.Println("1. Add a todo")
		fmt.Println("2. List all todos")
		fmt.Println("3. Remove a todo")
		fmt.Println("4. Exit")

		action := getUserChoice("Enter your choice: ")

		switch action {
		case 1:
			addTodo()
		case 2:
			list()
		case 3:
			remove()
		case 4:
			fmt.Println("Goodbye!")
			return // Exit the program gracefully
		default:
			fmt.Printf("Invalid action: %v\n", action)
		}
	}
}
