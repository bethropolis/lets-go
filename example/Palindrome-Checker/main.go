package main

import (
	"fmt"
)

func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(checkPalindrome("race a car"))
	fmt.Println(checkPalindrome(" "))
	fmt.Println(checkPalindrome("0P"))
	fmt.Println(checkPalindrome("a"))
	fmt.Println(checkPalindrome("1"))
}