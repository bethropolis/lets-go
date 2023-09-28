# Go To-Do List

A simple command-line to-do list manager implemented in Go. This application allows you to add tasks, list tasks, mark tasks as done, and remove tasks from your to-do list.

## Table of Contents

- [Go To-Do List](#go-to-do-list)
  - [Table of Contents](#table-of-contents)
  - [run](#run)
  - [Usage](#usage)
  - [Features](#features)
  - [Input Validation](#input-validation)
  - [File Handling](#file-handling)

## run

1. Run the program using the `go run` command:

   ```shell
   go run .
   ```

## Usage

The Go To-Do List program provides the following options:

1. **Add a Task**: Allows you to add a new task to your to-do list. You'll be prompted to enter a task description and a due date.

2. **List All Tasks**: Displays a list of all tasks in your to-do list, including their descriptions, due dates, and whether they are marked as done or not.

3. **Remove a Task**: Enables you to remove a task from your to-do list by entering its corresponding number.

4. **Exit**: Allows you to exit the program.

## Features

- **Input Validation**: The program includes input validation to ensure that user inputs are valid. It provides user-friendly error messages for invalid inputs.

- **File Handling**: Tasks are stored in a text file (`todo.txt`) for persistence between sessions. The program reads tasks from the file when started and writes them back to the file after modifications.

## Input Validation

The Go To-Do List program ensures that user inputs are valid:

- When removing a task, it prompts the user for a valid task number and provides an error message for invalid inputs.

## File Handling

Tasks are stored in a text file (`todo.txt`) for persistence. The program uses file I/O to:

- Read tasks from the file when the program starts.

- Write tasks back to the file after adding, removing, or marking tasks as done.


happy coding! 