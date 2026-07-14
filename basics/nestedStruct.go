package main

import "fmt"

type sendMessage struct {
	message  string
	sender   user
	receiver user
}

type user struct {
	id   int
	name string
}

func testInNestedStruct(m sendMessage) {
	fmt.Println("Sender:", m.sender.name)
	fmt.Println("Receiver:", m.receiver.name)
	fmt.Println("Message:", m.message)
	fmt.Println("===================================")
}

func nestedStruct() {
	var user1 = user{
		id: 12344,
		name: "Sumanraj",
	}

	var user2 = user{
		id: 12344,
		name: "Karan",
	}

	testInNestedStruct(sendMessage{
		message: "Hello Brother",
		sender: user1,
		receiver: user2,
	})

	testInNestedStruct(sendMessage{
		message: "Whats up? How are you?",
		sender: user2,
		receiver: user1,
	})

	testInNestedStruct(sendMessage{
		message: "I am fine.",
		sender: user1,
		receiver: user2,
	})
}
