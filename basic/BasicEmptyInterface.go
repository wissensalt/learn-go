package main

import "fmt"

func TestEmptyInterface(i int) interface{}  {
	if i == 1 {
		return 1
	}

	if i == 2 {
		return true
	}

	return "Test Empty Interface"
}

func TestParamEmptyInterface(arg interface{}) interface{} {
	if arg == 1 {
		return 1
	}

	if arg == 2 {
		return true
	}

	return "Test param Empty Interface"
}

func main() {
	fmt.Println(TestEmptyInterface(1))
	fmt.Println(TestEmptyInterface(2))
	fmt.Println(TestEmptyInterface(3))

	fmt.Println(">>>>>>>>")

	fmt.Println(TestParamEmptyInterface(1))
	fmt.Println(TestParamEmptyInterface(2))
	fmt.Println(TestParamEmptyInterface(3))
}
