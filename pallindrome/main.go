package main

import "fmt"

func isPallindrome(s string) bool {
	if len(s) < 1 {
		return false
	} else if len(s) == 1 {
		return true
	}

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[(len(s)-1)-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("isPallindrome(\"anna\")  = ", isPallindrome("anna"))
	fmt.Println("isPallindrome(\"anna\")  = ", isPallindrome("anna"))
	fmt.Println("isPallindrome(\"anna\")  = ", isPallindrome("nana"))
}
