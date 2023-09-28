package main

import (
	"fmt"
)

func main() {
	ops, ok, fail := 2350, 543, 433

	fmt.Println(
		"total:", ops, "- success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)

	var brand string

	// prints the string in quoted-form like this ""
	fmt.Printf("%q\n", brand)

	brand = "Google"
	fmt.Printf("%q\n", brand)


	// without newline
	fmt.Println("hihi")

	// with newline:
	//   \n = escape sequence
	//   \  = escape character
	fmt.Println("hi\nhi")

	// escape characters:
	//   \\ = \
	//   \" = "
	fmt.Println("hi\\n\"hi\"")

		var (
			speed int
			heat  float64
			off   bool
		)
	
		fmt.Printf("%T\n", speed)
		fmt.Printf("%T\n", heat)
		fmt.Printf("%T\n", off)
		fmt.Printf("%T\n", brand)

		var (
			planet   = "venus"
			distance = 261
			orbital  = 224.701
			hasLife  = false
		)
	
		// swiss army knife %v verb
		fmt.Printf("Planet: %v\n", planet)
		fmt.Printf("Distance: %v millions kms\n", distance)
		fmt.Printf("Orbital Period: %v days\n", orbital)
		fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	
		// argument indexing - indexes start from 1
		fmt.Printf(
			"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
			planet, distance,
		)
	
		fmt.Printf("Orbital Period: %f days\n", orbital)
		fmt.Printf("Orbital Period: %.0f days\n", orbital)
		fmt.Printf("Orbital Period: %.1f days\n", orbital)
		fmt.Printf("Orbital Period: %.2f days\n", orbital)
}

