package main

import "fmt"

func fixABug() {
	messagesFromDOrris := []string{
		"Whats up?",
		"How are you?",
		"I am fine. Wanna Hangout?",
		"Yes I do",
	}

	numMessages := float64(len(messagesFromDOrris))
	costPerMessage := 0.2

	totalCost := numMessages * costPerMessage

	fmt.Printf("Dorris spent %.2f on text messages today. \n", totalCost)
}
