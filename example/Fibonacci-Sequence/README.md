
## Fibonacci Sequence

### Program Description

This Go program generates the Fibonacci sequence up to a specified number of terms. The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, usually starting with `0` and `1`. In mathematical notation, it's often defined as: `F(0) = 0, F(1) = 1, and F(n) = F(n-1) + F(n-2)` for `n > 1`.

### Code Explanation

```go
package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(fibonacci(i))
    }
}
```

- The `fibonacci` function calculates the Fibonacci number at position `n` using recursion. It returns `n` if `n` is `0` or `1`, and otherwise, it calculates the sum of `fibonacci(n-1)` and `fibonacci(n-2)`.

- In the `main` function, a `for` loop is used to print the first 10 terms of the Fibonacci sequence.

### Usage

To generate the Fibonacci sequence for a different number of terms, modify the loop condition in the `main` function.

### Extra Tips and Notes

- The Fibonacci sequence can grow rapidly, so be cautious with large values of `n`.

- Recursion is a common way to calculate Fibonacci numbers, but it can be inefficient for large values of `n`. There are more efficient algorithms to calculate Fibonacci numbers if needed.
