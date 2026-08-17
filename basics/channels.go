package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Order struct {
	ID              string
	Amount          float64
	PaymentVerified bool
}

func main() {
	verifyPayment := make(chan *Order)
	processOrder := make(chan *Order)
	shipOrder := make(chan *Order)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for order := range verifyPayment {
			order.PaymentVerified = true
			fmt.Println("======================Payment Verified============================")
			fmt.Printf("Order Id: %v.\nPayment verification: %t \n", order.ID, order.PaymentVerified)
			fmt.Println("===============================================================")
			processOrder <- order
		}
		close(processOrder)
	}()

	go func() {
		for order := range processOrder {
			fmt.Println("====================Order Processed============================")
			fmt.Printf("Order Id: %v is processed.\n", order.ID)
			fmt.Println("===============================================================")
			shipOrder <- order
		}
		close(shipOrder)
	}()

	go func() {
		defer wg.Done()
		for order := range shipOrder {
			fmt.Println("======================Order Shipped============================")
			fmt.Printf("Order Id: %v is shipped and will be delivered within estimated date. \n", order.ID)
			fmt.Println("===============================================================")
		}
	}()

	orders := []Order{
		{ID: uuid.NewString(), Amount: 4000, PaymentVerified: false},
		{ID: uuid.NewString(), Amount: 499.5, PaymentVerified: false},
		{ID: uuid.NewString(), Amount: 350, PaymentVerified: false},
		{ID: uuid.NewString(), Amount: 105.78, PaymentVerified: false},
	}

	for i := range orders {
		fmt.Println("Order created with ID:", orders[i].ID)
		verifyPayment <- &orders[i]
	}

	close(verifyPayment)
	wg.Wait()
}
