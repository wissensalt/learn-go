package main

import (
	"errors"
	"fmt"
)

func main() {
	div1, err1 := divide(5, 0)
	printError(err1)
	printResult(div1)

	div2, err2 := divide(6, 2)
	printError(err2)
	printResult(div2)
}

func divide(value int, division int) (result int, error error)  {
	if division == 0 {
		return 0, errors.New("division can not be zero")
	}

	return value / division, nil
}

func printResult(result int)  {
	fmt.Println("Result : ", result)
}

func printError(error error) {
	if error != nil {
		fmt.Println("An error occurred : ", error)
	}
}