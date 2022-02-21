package main

import "fmt"

func isPalindromeRecurion(input string) bool {

	if len(input) == 0 || len(input) == 1 {
		return true
	}

	if input[0] == input[len(input)-1] {
		return isPalindromeRecurion(input[1 : len(input)-1])
	}

	return false
}

func isPalindrome(input string) bool {
	first := 0
	last := len(input) - 1

	for first < last {
		if input[first] == input[last] {
			first += 1
			last -= 1
		} else {
			return false
		}
	}

	return true
}

func isPalindromeReverse(input string) bool {
	return input == reverse(input)
}

func reverse(input string) string {
	s := ""
	for i := len(input) - 1; i >= 0; i-- {
		s += string(input[i])
	}

	return s
}

func main() {

	s := "akayaka"

	fmt.Println("s is palindrome =", isPalindromeRecurion(s))

	fmt.Println("s is palindrom =", isPalindrome(s))

	fmt.Println("s is palindrom =", isPalindromeReverse(s))
}
