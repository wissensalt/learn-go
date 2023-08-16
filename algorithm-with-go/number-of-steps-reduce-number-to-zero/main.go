package main

import "fmt"

func main() {
	num, result := GetStepsToZero(8, 0)
	fmt.Println(num, result)
}

func GetStepsToZero(num int, result int) (int, int) {
	if num == 0 {
		return num, result
	}

	result++
	if num%2 == 0 {
		num = num / 2
	} else {
		num--
	}

	return GetStepsToZero(num, result)
}
