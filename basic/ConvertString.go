package main

import "strconv"

func main() {
	fullName := "John Doe"
	firstIndex := fullName[0]
	firstIndexAsString := string(firstIndex)
	integerValue := 100
	integerToString := strconv.Itoa(integerValue)
	stringToInteger, _ := strconv.Atoi(integerToString)

	println(fullName)
	println(firstIndex)
	println(firstIndexAsString)
	println("Original Int : ", integerValue)
	println("Int to String : ", integerToString)
	println("String to Int : ", stringToInteger)
}
