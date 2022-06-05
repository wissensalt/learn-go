package main

import "fmt"

func main() {
	determineOddValue := isOddValue
	fmt.Println("Is 10 Odd Value : ", determineOddValue(10))
	fmt.Println("Is 9 Odd Value : ", determineOddValue(9))
}

func isOddValue(value int) bool  {
	return value %2 != 0
}