package main

import (
	"fmt"
)

func main()  {
	var data [5]int
	fmt.Println("Initial data value ", data)

	for a:=0; a<5; a++ {
		data[a] = (a+1)
	}

	for b:=0; b<5; b++ {
		fmt.Printf("index %v is %v\n", b, data[b])
	}

	fmt.Println("Direct Initial Array at Beginning")
	data2 := [5]int {1, 2, 3, 4, 5}
	for b:=0; b<5; b++ {
		fmt.Printf("index %v is %v\n", b, data2[b])
	}

	fmt.Println("Two Dimensional Array")
	data3 := [3][3]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for c:=0; c<3; c++ {
		for d:=0; d<3; d++ {
			fmt.Printf("index [%v,%v] is %v ", c, d, data3[c][d])
		}
		fmt.Println()
	}
}