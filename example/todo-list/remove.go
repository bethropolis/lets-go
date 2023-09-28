package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func removeTodo() {
	var i int
	i = getUserChoice("Enter the number of the todo you want to remove: ")

	TODOFile, err := os.OpenFile(todoFile, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer TODOFile.Close()

	todos, err := io.ReadAll(TODOFile)
	if err != nil { 
		panic(err)
	}

	lines := strings.Split(string(todos), "\n")
	if i < 1 || i > len(lines) {
		fmt.Println("Invalid todo number")
		return
	}

	lines = append(lines[:i-1], lines[i:]...)

	err = TODOFile.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = TODOFile.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	_, err = TODOFile.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Todo removed\n")
}

func remove() {
	listTodo()
	removeTodo()
	fmt.Printf("\n back to main menu (enter)")
	fmt.Scanln()
	
}
