package main

import "fmt"

func main() {
	n := 5
	// createTop(n)
	createLeftTop(n)
	createLeftBottom(n)
	createRightTop(n)
	createRightBottom(n)

}

func createRightTop(n int) {
	for a := 0; a < n; a++ {
		for b := 0; b < a; b++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func createRightBottom(n int) {
	for a := 0; a < n; a++ {
		for b := n; b > a; b-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func createLeftTop(n int) {
	for a := 0; a < n; a++ {
		for b := n; b >= a; b-- {
			if b == a {
				for c := 0; c < b; c++ {
					fmt.Print("*")
				}
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}
}

func createLeftBottom(n int) {
	for a := 0; a <= n; a++ {
		for b := 0; b < n; b++ {
			if a > b {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}
