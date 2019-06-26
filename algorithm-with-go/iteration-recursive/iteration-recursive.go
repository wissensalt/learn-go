package main

import "fmt"

func main() {
	iteration(0, 100)
}

func iteration(inc int, val int) {
	if inc <= val {
		fmt.Println(inc)
		inc++
		iteration(inc, val)
	}
}
