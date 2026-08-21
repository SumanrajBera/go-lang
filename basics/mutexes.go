package main

import (
	"fmt"
	"sync"
)

func process1(resource *int, mux *sync.Mutex, wg *sync.WaitGroup) {
	mux.Lock()
	defer mux.Unlock()
	defer wg.Done()
	for range 1_000_000 {
		*resource++
	}
}

func process2(resource *int, mux *sync.Mutex, wg *sync.WaitGroup) {
	mux.Lock()
	defer mux.Unlock()
	defer wg.Done()
	for range 1_000_000 {
		*resource++
	}
}

func process3(resource *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 1_000_000 {
		*resource++
	}
}

func process4(resource *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 1_000_000 {
		*resource++
	}
}
func main() {
	resource1 := 0
	resource2 := 0
	var wg sync.WaitGroup
	var mux sync.Mutex
	wg.Add(4)
	go process1(&resource1, &mux, &wg)
	go process2(&resource1, &mux, &wg)
	go process3(&resource2, &wg)
	go process4(&resource2, &wg)
	wg.Wait()
	fmt.Println("With Mutex--------------------------")
	fmt.Println("Expected resouce1 usage: 2000000")
	fmt.Println("Actual resource1 usage", resource1)

	fmt.Println("Without Mutex--------------------------")
	fmt.Println("Expected resouce2 usage: 2000000")
	fmt.Println("Actual resource2 usage", resource2)
}
