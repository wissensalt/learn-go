package main

import "fmt"

func regularFactorial(value int) int {
	result := 1
	for a := value; a > 0; a-- {
		result *= a
	}

	return result
}

func recursiveFactorial(value int) int  {
	if value == 1 {
		return 1
	}

	return value * recursiveFactorial(value - 1)
}

func main() {
	regularResult := regularFactorial(5)
	recursiveResult := recursiveFactorial(5)

	fmt.Println("Regular Result >>> ", regularResult)
	fmt.Println("Recursive Result >>> ", recursiveResult)
}