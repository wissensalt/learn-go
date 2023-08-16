package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/fizz-buzz/
func main() {
	fizzBuzz := FizzBuzz(3)
	fmt.Println(fizzBuzz)
}

func FizzBuzz(n int) []string {
	var result = make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
