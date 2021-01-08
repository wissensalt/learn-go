package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/q" {
			fmt.Println("quit")
			os.Exit(3)
		} else {
			fmt.Println("You type ", scanner.Text())
		}
	}
}
