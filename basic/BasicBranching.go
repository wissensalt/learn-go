package main

import "fmt"

func main() {
	condition := false
	if condition {
		fmt.Println("Condition True")
	} else {
		fmt.Println("Condition False")
	}

	//	Short statement
	name := "John Doe"
	if lengthOfName := len(name); lengthOfName <= 5 {
		fmt.Println("Hello ", name)
	} else {
		fmt.Println("Name Not Too long")
	}
}
