package main

import "fmt"

func main() {
	summingIntegers := sum(100, 150, 125, 300, 325)
	fmt.Println(summingIntegers)

	slice := []int {10, 20, 30, 40, 50}
	summingSliceValues := sum(slice...)
	fmt.Println(summingSliceValues)
}

func sum(values ...int) int {
	var result int
	for _, value := range values {
		result += value
	}

	return result
}
