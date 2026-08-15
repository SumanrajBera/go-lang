package main

import (
	"fmt"
	"sync"
	"time"
)

func emailing(sender string, receiver string) {
	go func() {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("Email received from %s \n", receiver)
	}()

	fmt.Printf("Email sent from %s \n", sender)
}

func emailingWithGroup(sender string, receiver string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("Email received from %s \n", receiver)
	}()

	fmt.Printf("Email sent from %s \n", sender)
}

func main() {
	// Here sleep is a good method but the problem is in real-world we don't have specified time here we are stopping the current routine for 500ms to ensure It is all processed. But if we don't do that what happens is main completes execution and extra routines doesn't even print as main has exited.
	emailing("Samule", "Jackson")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("--------------------------")
	emailing("Atharv", "Bhandungre")
	time.Sleep(500 * time.Millisecond)

	// Hence we will make use of sync package's WaitGroup which allows routine to not end until all the routine's have ended execution
	var wg sync.WaitGroup
	
	fmt.Println("--------------------------")
	fmt.Println("With WaitGroup")
	fmt.Println("--------------------------")
	emailingWithGroup("Samule", "Jackson", &wg)
	emailingWithGroup("Atharv", "Bhandungre", &wg)
	emailingWithGroup("Chetan", "Shetty", &wg)

	wg.Wait()
}
