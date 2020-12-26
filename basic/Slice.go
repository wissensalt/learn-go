package main

import "fmt"

func main() {
	fmt.Println("Basic Slice Operation")
	data := make([]int, 3)
	data[0] = 1
	data[1] = 2
	data[2] = 3
	fmt.Println("Initial Data value  = ", data)
	fmt.Println("Length of the Slice is ", len(data))

	fmt.Println("Appending value 4")
	data = append(data, 4)
	fmt.Println("Current value of data slice is", data)
	fmt.Println("Length of the Slice is ", len(data))

	fmt.Println("Appending value 5")
	data = append(data, 5)
	fmt.Println("Current value of data slice is", data)
	fmt.Println("Length of the Slice is ", len(data))


}
