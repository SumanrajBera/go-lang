package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// As the parameters are of same type we can just mention it once
func concat2(s1, s2 string) string {
	return s1 + s2
}

func main() {
	fmt.Println("This is the main function")
	fmt.Println(concat("This is", " a function call"))
}
