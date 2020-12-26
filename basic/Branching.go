package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Please input a value from 0 to 100")
	value, _ := reader.ReadString('\n')	
	intValue, _ := strconv.Atoi(strings.Replace(value, "\n", "", -1))

	if intValue >= 50 && intValue < 60 {
		fmt.Println("D")
	} else if intValue > 60 && intValue <= 70 {
		fmt.Println("C")
	} else if intValue > 70 && intValue <= 80 {
		fmt.Println("B")
	} else if intValue > 80 {
		fmt.Println("A")
	} else {
		fmt.Println("E")
	}
}

