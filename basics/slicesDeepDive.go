package main

import "fmt"

// Variadic functions
func sum(nums ...float64) float64 {
	var res float64 = 0
	for _, num := range nums {
		res += num
	}
	return res
}

func main() {
	nums := []float64{1, 28, 9.5, 7.6, 5.2, 7.9, 10.5}
	fmt.Println(sum(nums...))
}
