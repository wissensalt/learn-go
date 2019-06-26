package main

import "fmt"

func main() {
	var arr [5]int
	for a := 0; a < 5; a++ {
		arr[a] = a
	}
	displayArray(arr)
}

func displayArray(arr [5]int) {
	for a := 0; a < len(arr); a++ {
		fmt.Println(arr[a])
	}
}
