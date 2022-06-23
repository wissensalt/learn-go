package main

import "fmt"

func getValueString() interface{}  {
	return "A Value"
}

func getValueFloat() interface{}  {
	return 100.50
}

func main() {
	valueString := getValueString()
	var valueAsString string
	valueAsString = valueString.(string)
	fmt.Println("Type Assertions as String : ", valueAsString)

	valueFloat := getValueFloat()
	var valueAsFloat float64
	valueAsFloat = valueFloat.(float64)
	fmt.Println("Type Assertions as Float : ", valueAsFloat)
}
