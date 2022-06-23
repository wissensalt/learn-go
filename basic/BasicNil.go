package main

import "fmt"

func NewMap(value string) map[string] string  {
	if value == "" {
		return nil
	}

	return map[string]string {
		"Name": value,
	}
}

func main() {
	data := NewMap("")
	fmt.Println(data)

	data2 := NewMap("John")
	fmt.Println(data2)
}
