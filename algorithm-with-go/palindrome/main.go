package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Rotator"
	isPalindrome := IsPalindrome(word)
	fmt.Printf("Word %v Is Palindrome : %v\n", word, isPalindrome)
}

func IsPalindrome(word string) bool {
	if word == "" {
		return false
	}

	word = strings.ToLower(word)
	letters := []rune(word)
	var reversedWord = make([]string, len(letters))
	j := 0
	for i := len(letters); i > 0; i-- {
		reversedWord[j] = string(letters[i-1])
		j++
	}

	if word == strings.Join(reversedWord, "") {
		return true
	}

	return false
}
