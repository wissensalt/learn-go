package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error to read the file ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
