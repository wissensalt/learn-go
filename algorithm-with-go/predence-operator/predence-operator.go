package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var arr2 [5]int
	var arr [5]string
	for a := 0; a < len(arr); a++ {
		fmt.Print("Enter Letter: ")
		arr2[a] = a
		text, _ := reader.ReadString('\n')
		arr[a] = text
	}
}
