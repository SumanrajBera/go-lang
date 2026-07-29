package main

import "fmt"

func main() {
	ages := make(map[string]int)
	// Assigning
	ages["John"] = 20
	ages["Max"] = 18
	ages["Sam"] = 21

	fmt.Println(ages["Max"])

	// Deleting
	delete(ages, "Max")

	// Check if a key exists
	_, ok := ages["Max"]
	if !ok {
		fmt.Println("Max's age is deleted")
	}
}
