package main

import (
	"fmt"
	"strings"
)

type Filter func(value string) string

func main() {
	logWithSensitiveKeywords := "Name: John Doe, Password: 12345"

	fmt.Println(printLog(logWithSensitiveKeywords, filterSensitiveKeyword))

	fmt.Println(printLog(logWithSensitiveKeywords, filterUpperCase))
}

func printLog(log string, filter Filter) string {
	filteredLog := filter(log)

	return filteredLog
}

func filterSensitiveKeyword(keyword string) string {
	const Password = "Password"
	if strings.Contains(keyword, Password) {
		return strings.Replace(keyword, Password, "...", len(Password))
	}

	return keyword
}

func filterUpperCase(words string) string {
	var result string
	for _, word := range strings.Fields(words) {
		if strings.Contains(word, "e") {
			word = strings.ToUpper(word)
		}

		result += word + " "
	}


	return result
}
