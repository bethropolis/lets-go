package main

import (
	"bufio"
	"fmt"
	"os"
)

func listTodo() {
	var i int
	todoFile, err := os.Open(todoFile)
	if err != nil {
		panic(err)
	}
	defer todoFile.Close()

	fmt.Printf("\nLIST OF TODOS: \n")

	scannedTodos := bufio.NewScanner(todoFile)
	for scannedTodos.Scan() {
		fmt.Println(i+1, ">", scannedTodos.Text())
		i++
	}

	if err := scannedTodos.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("\n%d Todo(s) saved \n", i)
}

func list() {
	listTodo()
	fmt.Printf("\n back to main menu (enter)")
	fmt.Scanln()
	
}
