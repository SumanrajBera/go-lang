package main

import "fmt"

type messageSend struct {
	id      int
	message string
}

func test(m messageSend) {
	fmt.Println("Message:", m.message)
	fmt.Println("This message was sent by", m.id)
}

func basicStruct() {
	test(messageSend{
		id: 123456,
		message: "Hello Brother. How are you?",
	})

	test(messageSend{
		id: 578943,
		message: "I am fine. Whats up?",
	})

	test(messageSend{
		id: 123456,
		message: "Wanna Hangout?",
	})

	test(messageSend{
		id: 578943,
		message: "No. I have a night class today",
	})
}
