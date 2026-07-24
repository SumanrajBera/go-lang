package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	sliceOfArray := array[0:3]
	// Here its just a shallow copy so if we make changes it will show up in original
	fmt.Println(sliceOfArray)

	var newSlice []int
	// A nil slice will print []
	fmt.Println(newSlice)

	thirdSlice := make([]int, 3, 5)
	// 3 is the length and 5 is the upper limit after which Go will have to do reallocation
	fmt.Println(thirdSlice)
}
