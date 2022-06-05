package main

import "fmt"

func main() {
	firstName, middleName, surName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(surName)
}


func getFullName() (string, string, string) {
	return "John", "Doe", "Karebet"
}