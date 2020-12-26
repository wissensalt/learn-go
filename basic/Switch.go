package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	fmt.Println("Case 1 : Switching Numbers")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input a value from 0 to 4")
	inputValue,_ := reader.ReadString('\n')
	intValue, _ := strconv.Atoi(strings.Replace(inputValue, "\n", "", -1))
	var result string
	switch intValue {
		case 0 :
			result = "Zero"
		case 1 :
			result = "One"
		case 2 :
			result = "Two"
		case 3 : 
			result = "Three"
		case 4 :
			result = "Four"
		default :
			result = "Unknown"
	}
	fmt.Printf("Your input value is %v \n", result)

	fmt.Println("Case 2 : Days in Bahasa")
	result = ""
	switch time.Now().Weekday() {
		case time.Sunday :
			result = "hari Minggu"
		case time.Monday :
			result = "hari Senin"
		case time.Tuesday :
			result = "hari Selasa"
		case time.Wednesday :
			result = "hari Rabu"
		case time.Thursday :
			result = "hari Kamis"
		case time.Friday :
			result = "hari Jumat"
		case time.Saturday :
			result = "hari Sabtu"
	}
	fmt.Printf("Today in Bahasa is %v \n", result)

	fmt.Println("Case 3 : Checking Type of Data")
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
			case bool :
				fmt.Printf("%T is bool\n", t)
			case string :
				fmt.Printf("%T is string\n", t)
			case int :
				fmt.Printf("%T is integer\n", t)
			case *bool :
				fmt.Printf("%T is pointer to boolean\n", t)
			case *int : 
				fmt.Printf("%T is Pointer to integer\n", t)
			default:
				fmt.Printf("Unexpected Type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(20)
	whatAmI("test")
	whatAmI(20.00)
}