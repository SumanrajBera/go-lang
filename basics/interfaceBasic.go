package main

import "fmt"

type Message interface {
	sendMessage() string
}

type report struct {
	id      int
	message string
}

func (r report) sendMessage() string {
	return r.message
}

func main() {
	newReport := report{
		id:      123445,
		message: "This is my report on interfaces",
	}
	fmt.Println("Message:", newReport.sendMessage())
}
