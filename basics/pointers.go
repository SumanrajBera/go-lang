package main

import "fmt"

func main() {
	var a *int
	x := 5
	y := x // copy of x creates a new reference in memory
	z := &x	// pointer stores address of the in-memory location and can be used with *ptr-name

	fmt.Println("Pointers default Value:", a)
	fmt.Println("--------------------------")

	fmt.Println("Address of x:", &x);
	fmt.Println("Value of x:", x);
	fmt.Println("Address of y:", &y);
	fmt.Println("Value of y:", y);
	fmt.Println("Address of z:", &z);
	fmt.Println("Value of z:", z);
	fmt.Println("--------------------------")
	x = 7;
	fmt.Println("Updating value of x:", x);
	fmt.Println("Doesn't affect vice versa");
	fmt.Println("Value of y:", y);
	fmt.Println("--------------------------")
	fmt.Println("But with pointers we can go to the address and change in memory");
	*z = 10
	fmt.Println("After change in memory value of x:", x);
}