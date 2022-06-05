package main

import "fmt"

func main() {
	fmt.Println("Counter 1 >>>")
	counter := 1
	for counter <= 5 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("Counter 2 >>>")
	for counter2 := 1; counter2 <= 5; counter2++ {
		fmt.Println(counter2)
	}

	fmt.Println("Iterate Slice >>>")
	slice := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for a := 0; a<len(slice); a++ {
		fmt.Println(slice[a])
	}

	fmt.Println("Iterate Slice via Range>>>")
	for i, value := range slice {
		fmt.Println("Index ", i, " - Value ", value)
	}

	fmt.Println("Iterate Slice via Range Ignore Index>>>")
	for _, value := range slice {
		fmt.Println("Value ", value)
	}

	fmt.Println("Iterate Map via Range>>>")
	person := make(map[string] string)
	person["Name"] = "John"
	person["Address"] = "Bdg"
	for key, value := range person {
		fmt.Println("Key ", key, " Value ", value)
	}
}
