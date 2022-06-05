package main

import "fmt"

func main() {
	a, b, c := myFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func myFullName() (firstName string, middleName string, lastName string)  {
	firstName = "John"
	middleName = "Doe"
	lastName = "Karebet"

	return
}