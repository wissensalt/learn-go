package main

import "fmt"

var counter int

func main() {
	counter = 1
	displayFibonacci(15, 0, 0, 0)
}

func displayFibonacci(n int, x int, y int, z int) {
	if counter < n {
		if counter == 1 {
			x = 0
			y = 1
			z = getResult(x, y)
			fmt.Print(z, " ")
		} else {
			z = getResult(x, y)
			fmt.Print(z, " ")
			x = y
			y = z
		}
		counter++
		displayFibonacci(n, x, y, z)
	}

}

func getResult(x int, y int) int {
	return x + y
}
