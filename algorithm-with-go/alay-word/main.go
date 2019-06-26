package main

import (
	"fmt"
	"strings"
)

func main() {
	data := [...]string{"kamu", "mau", "tape", "apa", "bebek", "mereka", "itu", "rasanya", "sangat", "enak", "sekali"}
	fmt.Println(convert(data[:]...))
}

func convert(data ...string) string {
	var result string
	for a := 0; a < len(data); a++ {
		word := strings.Split(data[a], "")
		for b := 0; b < len(word); b++ {
			word[b] = convertPerWord(word[b])
		}
		result += strings.Join(word, "") + " "
	}
	return result
}

func convertPerWord(character string) string {
	if strings.ToUpper(character) == "B" {
		return "8"
	} else if strings.ToUpper(character) == "A" {
		return "4"
	} else if strings.ToUpper(character) == "M" {
		return "111"
	} else if strings.ToUpper(character) == "S" {
		return "5"
	} else if strings.ToUpper(character) == "E" {
		return "3"
	} else if strings.ToUpper(character) == "T" {
		return "7"
	} else if strings.ToUpper(character) == "H" {
		return "1-1"
	} else if strings.ToUpper(character) == "G" {
		return "6"
	} else if strings.ToUpper(character) == "R" {
		return "12"
	} else if strings.ToUpper(character) == "I" {
		return "1"
	} else {
		return strings.ToUpper(character)
	}
}
