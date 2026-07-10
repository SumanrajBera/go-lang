package main

import "fmt"

func main()  {
	const age int = 5

	if age >= 18 {
		fmt.Println("You're eligible for voting")
	}else if age >= 10 {
		const remYear = 18 - age
		fmt.Printf("You have to wait for %v more years", remYear)
	}else {
		fmt.Printf("You're too young for this.")
	}
}