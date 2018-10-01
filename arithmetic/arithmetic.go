package main

import "fmt"

func main() {
	var numberOne int = 20
	var numberTwo int = 3
	var result = numberOne + numberTwo
	fmt.Println("Result Addition", result)

	result = numberOne - numberTwo
	fmt.Println("Result Substraction", result)

	result = numberOne * numberTwo
	fmt.Println("Result Multiplication", result)

	result = numberOne / numberTwo
	fmt.Println("Result Division", result)

	result = numberOne % numberTwo
	fmt.Println("Result Modulus", result)
}
