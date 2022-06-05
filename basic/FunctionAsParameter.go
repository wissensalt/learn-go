package main

import (
	"fmt"
	"strings"
)

func main() {
	logWithSensitiveKeywords := "Name: John Doe, Password: 12345"
	fmt.Println(printLog(logWithSensitiveKeywords, filterSensitiveKeyword))
}

func printLog(log string, filter func (string) string) string  {
	filteredLog := filter(log)

	return filteredLog
}

func filterSensitiveKeyword(keyword string)string {
	const Password = "Password"
	if strings.Contains(keyword, Password) {
		return strings.Replace(keyword, Password, "...", len(Password))
	}

	return keyword
}