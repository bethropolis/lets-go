package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func StoreTodo(todo string) {
	todoFile, err := os.OpenFile(todoFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer todoFile.Close()

	todoFile.WriteString(todo + " | " + time.Now().Format("2006-01-02 15:04:05") + "\n")
}

func addTodo() {
	fmt.Printf("\nEnter your todo: ")
	reader := bufio.NewReader(os.Stdin)
	todo, _ := reader.ReadString('\n')
	todo = strings.TrimSpace(todo)
	StoreTodo(todo)
	fmt.Printf("Todo added: %s\n", todo)

	fmt.Printf("\n back to main menu (enter)")
	fmt.Scanln()
}
