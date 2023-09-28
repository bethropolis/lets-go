package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false // 1 and any negative number are not prime
	}
	if n == 2 {
		return true // 2 is prime
	}
	if n%2 == 0 {
		return false // Even numbers greater than 2 are not prime
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false // Divisible by a number other than 1 and itself
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(1)) // false
	fmt.Println(isPrime(2)) // true
	fmt.Println(isPrime(3)) // true
	fmt.Println(isPrime(4)) // false
	fmt.Println(isPrime(5)) // true
	fmt.Println(isPrime(6)) // false
	fmt.Println(isPrime(7)) // true
}
