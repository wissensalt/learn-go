package main

import "fmt"

func main() {
	var x int64
	var y int64
	var z int64
	for a := 1; a < 10; a++ {
		if a == 1 {
			x = 0
			y = 1
			z = getResult(x, y)
			fmt.Print(z, " ")
		} else {
			z = getResult(x, y)
			x = y
			y = z
			fmt.Print(z, " ")
		}
	}
}

func getResult(x int64, y int64) int64 {
	return x + y
}
