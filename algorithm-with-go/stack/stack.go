package main

import "fmt"

func main() {
	var currentIndex = 0
	arr := [5]*int{nil, nil, nil, nil, nil}
	display(arr, currentIndex)
	currentIndex, arr = push(arr, currentIndex, 1)
	currentIndex, arr = push(arr, currentIndex, 2)
	currentIndex, arr = push(arr, currentIndex, 3)
	currentIndex, arr = push(arr, currentIndex, 4)
	display(arr, currentIndex)
	// currentIndex, arr = pop(arr, currentIndex)
	// currentIndex, arr = pop(arr, currentIndex)
	// currentIndex, arr = popDefault(arr, currentIndex)
	// currentIndex, arr = popDefault(arr, currentIndex)
	display(arr, currentIndex)
}

func display(arr [5]*int, index int) {
	fmt.Println("Stack Display :", index)
	for a := 0; a < index; a++ {
		fmt.Print(*arr[a])
	}
	fmt.Println("")
}

func push(arr [5]*int, index int, val int) (int, [5]*int) {
	arr[index] = &val
	index++
	return index, arr
}

func pop(arr [5]*int, index int) (int, [5]*int) {
	fmt.Println("POP, arr", index-1, *arr[index-1])
	arr[index] = nil
	index--
	return index, arr
}

func popDefault(arr [5]*int, index int) (int, [5]*int) {
	arr[0] = nil
	index--
	return index, arr
}
