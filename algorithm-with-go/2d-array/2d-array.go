package main

import "fmt"

func main() {
	var arr [5][5]int

	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			arr[a][b] = (a + b)
		}
	}

	displayArray(arr)
}

func displayArray(arr [5][5]int) {
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			fmt.Print(arr[a][b], "\t")
		}
		fmt.Println("")
	}
}
