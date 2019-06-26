package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var arr2 [5]int
	var arr [5]string
	for a := 0; a < len(arr); a++ {
		fmt.Print("Enter Letter: ")
		arr2[a] = a
		text, _ := reader.ReadString('\n')
		arr[a] = text
	}

	fmt.Println("Original", displayWord(arr))
	fmt.Println("Reversed", displayWord(reverseWord(arr)))
}

func reverseWord(arr [5]string) [5]string {
	var result [len(arr)]string
	var inc = 0
	for a := len(arr) - 1; a >= 0; a-- {
		result[inc] = arr[a]
		inc++
	}
	return result
}

func displayWord(arr [5]string) string {
	result := ""
	for a := 0; a < len(arr); a++ {
		result += arr[a]
	}
	return result
}
