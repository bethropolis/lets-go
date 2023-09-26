package main

import "fmt"

func main(){
	// declaring variables and types
	var i int
	var s string
	
	// initializing values
	i = 20
	s = "I am a string"

	fmt.Println(i)
	fmt.Println(s)

	// creating and initializing with variables
	var num int = 12
	var isNum bool = true

	fmt.Println(num)
	fmt.Println(isNum)

	// short variable declaration

	l := 100;

	fmt.Println(l)

	// declaring multiple variables
	
	firstname, lastname := "firstName", "LastName"

	fmt.Println(firstname + " "+ lastname)

	// variable declaration block
	var(
		name = "John Doe"
		age = 50
	)

	fmt.Println(name)
	fmt.Println(age)

	join()
	convert()
}