## Factorial Calculator

### Program Description

This Go program calculates the factorial of a given number. The factorial of a non-negative integer `n` is the product of all positive integers from `1` to `n`. In mathematical notation, it's often denoted as `n!`.

### Code Explanation

```go
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n < 0 {
		return -1 // Handle negative input gracefully
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}
```

- The `factorial` function takes an integer `n` as input and recursively calculates the factorial. It handles negative input by returning `-1` as an error value.

- In the `main` function, we call `factorial(5)` to calculate the factorial of `5` and print the result.

### Usage

To run the program and calculate the factorial of a different number, modify the argument in the `factorial` function within the `main` function.

### Extra Tips and Notes

- Factorials can grow rapidly, so be cautious with large inputs, as they may lead to integer overflow.

- Handling negative input gracefully with an appropriate error value is a good practice for robust code.

Happy coding! 😊🚀
