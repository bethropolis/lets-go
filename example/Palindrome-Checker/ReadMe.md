## Palindrome Checker

### Program Description

This Go program checks if a given string is a palindrome. A palindrome is a word, phrase, number, or other sequences of characters that reads the same forward and backward (ignoring spaces, capitalization, and non-alphanumeric characters).

### Code Explanation

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// Step 1: Convert the string to lowercase
	s = strings.ToLower(s)
	
	// Step 2: Remove non-alphanumeric characters
	var cleanString []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleanString = append(cleanString, char)
		}
	}
	
	// Step 3: Check if it's a palindrome
	length := len(cleanString)
	for i := 0; i < length/2; i++ {
		if cleanString[i] != cleanString[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("1"))
}
```

**Step 1:** We convert the input string to lowercase using `strings.ToLower()`. This ensures that we ignore capitalization when checking for palindromes.

**Step 2:** We remove non-alphanumeric characters (spaces and special characters) from the string. We use a `for` loop to iterate through each character of the string and keep only letters and digits using `unicode.IsLetter()` and `unicode.IsDigit()`. We store the cleaned characters in the `cleanString` slice.

**Step 3:** We check if the cleaned string is a palindrome. We iterate through the first half of the cleaned string and compare each character to its corresponding character from the end of the string. If any pair of characters doesn't match, we return `false`. If all pairs match, the string is a palindrome, and we return `true`.

### Usage

To use this program, call the `isPalindrome()` function with a string as an argument. It will return `true` if the string is a palindrome and `false` otherwise.

### Extra Tips and Notes

- This palindrome checker handles spaces, capitalization, and non-alphanumeric characters, providing a robust solution.

- Understanding how to manipulate strings, iterate through characters, and use conditional statements is crucial for solving similar text-processing problems.
