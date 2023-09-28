# Prime Number Checker
---

```go
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
```

**Here's the explanation:**

- We start by checking if the number is less than or equal to 1. In that case, it's not prime (return `false`).

- We handle the special case where `n` is 2 because 2 is the only even prime number (return `true`).

- We check if the number is even (greater than 2) because all even numbers (except 2) are not prime (return `false`).

- We use a loop to check divisibility by odd numbers starting from 3 up to the square root of `n`. If `n` is divisible by any of these numbers, it's not prime (return `false`).

- If none of the conditions are met, we conclude that the number is prime (return `true`).
