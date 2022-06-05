package main

import "fmt"

func main() {
	person := map[string] string {
		"Name": "John",
		"Address": "Jkt",
	}
	person["Job"] = "Programmer"

	fmt.Println(person)
	fmt.Println("NAME>>>", person["Name"])
	fmt.Println("ADDRESS>>>", person["Address"])
	fmt.Println("JOB>>>", person["Job"])

	//Delete key Job
	delete(person, "Job")
	fmt.Println(person)

	newPerson := make(map[string] string, 2)
	newPerson["Name"] = "Sukar"
	newPerson["Address"] = "Bdg"
	fmt.Println("NEW PERSON >>> ", newPerson)
}
