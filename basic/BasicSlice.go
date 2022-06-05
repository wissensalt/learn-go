package main

import "fmt"

func main() {
	var months = [...] string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	var slice1Cap = cap(slice1)
	var slice1Length = len(slice1)

	fmt.Println(slice1)
	fmt.Println(slice1Cap)
	fmt.Println(slice1Length)

	months[5] = "Modified"
	fmt.Println(slice1)

	slice1[0] = "Modified-Slice-0"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "New-Month")
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Slice with Make
	// 1st arg = TYPE
	// 2nd arg = length of slice
	// 3rd arg = capacity of slice
	var slice4 = make([]int, 2, 5)
	slice4[0] = 1
	slice4[1] = 2
	fmt.Println("SLICE4 >>>", slice4)
	fmt.Println("SLICE4 CAP >>>", cap(slice4))
	fmt.Println("SLICE4 LEN >>>", len(slice4))

	copySlice := make([]int, len(slice4), cap(slice4))
	copy(copySlice, slice4)
	fmt.Println("COPY SLICE >>>", copySlice)

	exampleArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("EXAMPLE ARRAY >>>", exampleArray)

	exampleArrayWithoutSize := [...]int{1, 2, 3, 4, 5}
	fmt.Println("ARRAY WITHOUT SIZE", exampleArrayWithoutSize)

	exampleSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("EXAMPLESLICE >>>", exampleSlice)
}
