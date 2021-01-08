package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("You must enter argument. Type /help for info")
	} else {
		if args[0] == "/help" {
			fmt.Println("Usage : InteractiveTextConverter 1")
			fmt.Printf("Option : \n" +
				"1. All Upper Case\n" +
				"2. Titled Case\n" +
				"3. All Lower Case")
		} else {
			file, err := os.Open("loremipsum.txt")
			if err != nil {
				fmt.Println("Failed to open file ", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				switch args[0] {
					case "1": fmt.Println(strings.ToUpper(scanner.Text()))
					case "2": fmt.Println(strings.Title(scanner.Text()))
					case "3": fmt.Println(strings.ToLower(scanner.Text()))
				}
			}
		}
	}
}
