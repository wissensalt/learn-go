package main

import "fmt"

func main() {
	//Read single word
	var name string
	fmt.Println("What is your name ?")
	input, _ := fmt.Scanf("%s", &name)
	switch input {
		case 0:
			fmt.Printf("You must enter argument")
		case 1:
			fmt.Printf("Your name is %v\n", name)
	}

	//Read multiple word
	fmt.Println("What is your name ?")
	secondInput, _ := fmt.Scanf("%q", &name)
	switch secondInput {
	case 0:
		fmt.Printf("You must enter argument include double quote character")
	case 1:
		fmt.Printf("Your name is %v\n", name)
	}
}
