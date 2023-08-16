package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 123
	input := strconv.Itoa(val)
	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input)-1; j++ {
			valA, _ := strconv.Atoi(string(input[i]))
			valB, _ := strconv.Atoi(string(input[j]))
			fmt.Println("I", valA)
			fmt.Println("J", valB)
			x, _ := strconv.Atoi(string(input[i]) + string(input[j]))
			fmt.Println(x)
		}
	}
}
