package main

import (
	"fmt"
	"sync"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for range numDBs {
		<-dbChan
	}
}

// don't touch below this line

func testToken(numDBs int) {
	var wg sync.WaitGroup
	wg.Add(1)
	dbChan := getDatabasesChannel(numDBs, &wg)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	wg.Wait()
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	testToken(3)
	testToken(4)
	testToken(5)
}

func getDatabasesChannel(numDBs int, wg *sync.WaitGroup) chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer wg.Done()
		for i := range numDBs {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()

	return ch
}
