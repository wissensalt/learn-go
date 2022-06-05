package main

import "fmt"

func main() {
	data := "Basic String"
	fmt.Println(data)
	fmt.Println("Length of String : ", len(data))
	fmt.Printf("%c\n", data[0])
	fmt.Printf("%c\n", data[1])
	fmt.Printf("%c\n", data[3])
	fmt.Printf("%c\n", data[4])
}
